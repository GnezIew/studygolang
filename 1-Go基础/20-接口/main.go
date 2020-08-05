package main

import (
	"fmt"
)

/*
	接口 包括动态值 和 动态类型
*/
//接口接受参数类型的判断
func E(x interface{}){
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("non-empty interface")
}

//接口定义 (人为的加一个 ”er“ 表示接口)
type Humaner interface {
	//方法的声明
	SayHello()
}

type Student struct {
	name string
	age int
	score int
}
func (s *Student) SayHello(){
	fmt.Println("大家好我叫:",s.name)
}

type teacher struct {
	name string
	class string
}

func (t *teacher) SayHello(){
	fmt.Println("我是老师:",t.name)
}

//函数 -多态的实现 将接口作为函数参数
func SayHello(humaner Humaner){
	humaner.SayHello()
}

func printType(v interface{}){
	switch v.(type) {
	case int:
		fmt.Println(v,"is int")
	case string :
		fmt.Println(v,"is string")
	case bool:
		fmt.Println(v,"is bool")
	}
}
func main() {
	var x *int = nil		//这里仅仅动态值为nil 而动态类型不是
	E(x)
	var x2 interface{} = nil
	E(x2)
	//
	stu := Student{
		name:  "曾为",
		age:   24,
		score: 100,
	}
	teach := teacher{
		name:  "董聃",
		class: "WEB",
	}
	////定义接口类型变量
	//var h Humaner
	////必须讲对象的地址赋值给接口类型变量
	//h = &stu
	//h.SayHello()
	//
	//h = &teach
	//h.SayHello()
	SayHello(&stu)
	SayHello(&teach)

	//四则运算
	f := factory{}
	res := f.CreateFactory(1,2,"/")
	fmt.Println(res)

	//--------------datawriter.go-----------
	//实例化file
	F := new(file)

	//声明一个Datawrite接口
	var writer DataWriter

	//将接口赋值f,也就是*file类型
	writer = F

	//使用Datawriter接口进行数据写入
	writer.WriteData("data")

	printType(1024)
	printType("pig")
	printType(true)


	var temp map[string]string
	temp = map[string]string{"sad":"asd"}
	v,ok := temp["aa"]
	fmt.Println(v,ok)
}
