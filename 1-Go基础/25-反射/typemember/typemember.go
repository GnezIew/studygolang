package main

import (
	"fmt"
	"reflect"
)

//声明一个空结构体
type cat struct {
	Name string

	//带有结构体tag的字段
	Type int `json:"type" id:"100"`
}


//定义一个结构体
type dummy struct {
	a int
	b string

	//嵌入字段
	float32
	bool

	next *dummy
}
func main() {


	//创建cat的实例
	ins := cat{
		Name: "mimi",
		Type: 1,
	}

	//获取结构体实例的反射类型对象
	typeOfCat := reflect.TypeOf(ins)

	//遍历结构体所有成员
	for i := 0; i < typeOfCat.NumField() ;i++  {
		//获取每个成员的结构体实例的反射类型对象
		fieldType := typeOfCat.Field(i)
		//输出成员名和tag
		fmt.Printf("name : %v tag: %v\n",fieldType.Name,fieldType.Tag)
	}

	//通过字段名，找到字段类型信息
	if catType, ok:= typeOfCat.FieldByName("Type");ok{
		//从tag 中取出需要的tag
		fmt.Println(catType.Tag.Get("json"),catType.Tag.Get("id"))
	}

	//-----------------------------------------------------------------------
	//值包装结构体
	d := reflect.ValueOf(dummy{
		a:       1,
		b:       "helloworld",
		float32: 30.0,
		bool: true,
		next:    &dummy{
			a:3,
		},
	})

	//获取字段数量
	fmt.Println("NumField",d.NumField())

	//获取索引为2的字段(float32字段)
	floatField := d.Field(2)
	//输出字段类型
	fmt.Println("Field:",floatField.Type())

	//根据名字找字段
	fmt.Println("FieldByName(\"b\").Type():",d.FieldByName("b").Type())
	fmt.Println("FieldByName(\"b\"):",d.FieldByName("b"))

	//根据索引查找值中,next字段的int字段的值
	fmt.Println("FieldByInex([]int{4,0}).Type():",d.FieldByIndex([]int{4,0}).Type())
	fmt.Println("FieldByInex([]int{4,0}):",d.FieldByIndex([]int{4,0}))

}
