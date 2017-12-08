package main

import (
	"./optreflect"
	"fmt"
)
type TestObj struct {
	field1 string
	field2 []string
	field3 int
	field4 []int	`alias:"f4"`
	field5 byte 	`alias:"f5"`
	field6 []byte	`alias:"f6"`
}
func (t *TestObj)T()  {
	fmt.Println("fuck!")
}
func main(){
	a:=&optreflect.OptReflect{}
	a.Init(&TestObj{})
	ooo:=TestObj{
		field1:"test string",
		field2:[]string{"a","b","c"},
		field3:4523,
		field4:[]int{1,2,3,5},
		field5:byte(4),
		field6:[]byte{2,2,2},
		}
	err:=a.Set(&ooo,"field1","change test string to test this sentence!")
	if err!=nil{
		fmt.Println(err)
	}
	err=a.Set(&ooo,"field1",0)
	if err!=nil{
		fmt.Println(err)
	}
	err=a.Set(&ooo,"field1",0)
	if err!=nil{
		fmt.Println(err)
	}
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
	v,err=a.Get(&ooo,"field4")
	if err==nil{
		fmt.Println(v)
	}
	v,err=a.Get(&ooo,"f5")
	if err==nil{
		fmt.Println(v)
	}
	v,err=a.Get(&ooo,"f6")
	if err==nil{
		fmt.Println(v)
	}
}