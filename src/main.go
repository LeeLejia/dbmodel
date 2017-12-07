package main

import (
	"./optreflect"
	"fmt"
)
type TestObj struct {
	field1 string `TA`
	field2 []string `db:"nihao"`
	field3 int	  `db:65`
}
type TestObj2 struct {
	ss string `TA`
	g []string `db:"nihao"`
	wg string `TA`
	qw int	  `db:65`
	l []string `db:"nihao"`
	fh int	  `db:65`
}
func (t *TestObj)T()  {
	fmt.Println("fuck!")
}
func main(){
	//sss:=TestObj{
	//	field1:"abcdefghijklmnopyq",
	//	field2:[]string{"a"},
	//	field3:3,
	//}
	//vvv:=TestObj{
	//	field1:"abcdlmnopyq",
	//	field2:[]string{"ass"},
	//	field3:4553,
	//}
	//fmt.Println(*(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&sss))+unsafe.Offsetof(vvv.field3))))
	//fmt.Println(unsafe.Pointer(&vvv))

	a:=&optreflect.OptReflect{}
	//s:="xsss"
	//_,err:=a.Get(&TestObj{},"xxx")
	//if err!=nil{
	//	fmt.Println(err.Error())
	//}
	//err=a.Set("xxxxxx","kiy","xxx")
	//if err!=nil{
	//	fmt.Println(err.Error())
	//}
	a.Init(&TestObj{})
	//err:=a.Set(&aa,"kiy","xxx")
	//if err!=nil{
	//	fmt.Println(err.Error())
	//}
	ooo:=TestObj{field1:"abcdesklmnopyq",field2:[]string{"a","b","c"}, field3:4523}
	v,err:=a.Get(&ooo,"field1")
	if err==nil{
		fmt.Println(v)
	}
	v,err=a.Get(&ooo,"field3")
	if err==nil{
		fmt.Println(v)
	}
	v,err=a.Get(&ooo,"field2")
	if err==nil{
		fmt.Println(v)
	}
}