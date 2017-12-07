package optreflect

import (
	"reflect"
	"fmt"
	"unsafe"
)

type OptReflect struct{
	structName string
	fieldsMap  map[string]field
}
type field struct {
	offset uintptr
	fieldType string
}
type empty struct {
	etype  *struct{}
	ptr unsafe.Pointer
}
/**
获取表名称
 */
func (t *OptReflect) GetStructName() string{
	return t.structName
}
/**
获取值
 */
func (t *OptReflect) Get(obj interface{}, key string) (interface{},error) {
	if t.fieldsMap == nil{
		return nil,error(fmt.Errorf("对象尚未初始化"))
	}
	v,exist:=t.fieldsMap[key]
	if !exist{
		return nil,error(fmt.Errorf("%s字段不存在",key))
	}
	// 非指针
	on:=reflect.TypeOf(obj).Name()
	if on==""{
		// 如果传入引用类型
		on=reflect.TypeOf(obj).Elem().Name()
	}
	if on!=t.structName{
		return nil,error(fmt.Errorf("给出的类型是%s,要求的类型为%s",on,t.structName))
	}
	ptr := (*empty)(unsafe.Pointer(&obj)).ptr
	ptr = unsafe.Pointer(uintptr(ptr) + v.offset)
	return getInterfaceType(uintptr(ptr),v.fieldType),nil
}
/**
设置值
 */
func (t *OptReflect) Set(obj interface{}, key string, value interface{}) error{
	if t.fieldsMap == nil{
		return error(fmt.Errorf("对象尚未初始化"))
	}
	p:=reflect.TypeOf(obj)
	if p.Kind().String()!="ptr"{
		return error(fmt.Errorf("要求传入指针类型,请检查是否忽略了&引用"))
	}
	n:=p.Name()
	if n!=t.structName{
		if n==""{
			n=p.Elem().String()
		}
		return error(fmt.Errorf("当前传入类型%s,需要传入%s类型结构体的引用",n,t.structName))
	}
	// todo
	return nil
}
/**
使用前初始化,字段别名可以通过tag中的alias设置
如:
type Test struct {
	field1 string
	field2 []string `alias:"oo"`
	field3 int
}
 */
func (t *OptReflect) Init(model interface{}){
	p:=reflect.TypeOf(model)
	if p.Kind().String()!="ptr"{
		panic(fmt.Errorf("要求传入指针类型,请检查是否忽略了&引用"))
	}
	elem:=p.Elem()
	if elem.Kind().String()!="struct"{
		panic(fmt.Errorf("给出的类型是%s,要求的类型为%s",elem.Kind().String(),"struct"))
	}
	if elem.NumField()==0{
		panic(fmt.Errorf("%s不存在可用字段",elem.Kind().String()))
	}
	t.fieldsMap = make(map[string]field,elem.NumField())
	for i:=0;i< elem.NumField();i++{
		f:=elem.Field(i)
		key:=f.Name
		if _,exist:=t.fieldsMap[key];exist{
			t.fieldsMap = nil
			panic(fmt.Errorf("字段名%s被多次定义.请检查结构体%s中tag及field是否存在重复命名。",key,elem.Name()))
		}
		t.fieldsMap[key] = field{f.Offset,f.Type.Kind().String()}
		if alias:=f.Tag.Get("alias");alias!="" && key!=alias{
			if _,exist:=t.fieldsMap[alias];exist{
				t.fieldsMap = nil
				panic(fmt.Errorf("字段名%s被多次定义.请检查结构体%s中tag及field是否存在重复命名。",alias,elem.Name()))
			}
			t.fieldsMap[alias] = field{f.Offset,f.Type.Kind().String()}
		}
		//fmt.Println(fmt.Sprintf("name=%s,tag=%s,type=%s,kind=%s,offset=%d",f.Name,f.Tag.Get("db"),f.Type.Name(),f.Type.Kind().String(),f.Offset))
	}
	t.structName = elem.Name()
}

func getInterfaceType(ptr uintptr, t string) interface{}{
	switch t {
	case "string":
		return *(* string)(unsafe.Pointer(ptr))
	case "int":
		return *(* int)(unsafe.Pointer(ptr))
	case "slice":
		// todo 区分slice类型
		return *(* []string)(unsafe.Pointer(ptr))
	}
	return nil
}