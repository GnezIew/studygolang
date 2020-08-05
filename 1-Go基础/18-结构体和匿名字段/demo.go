package main

import "fmt"

type Student struct {
	//结构体成员在存储时 ， 会有一个内存对齐的问题 ， 会按照数据类型中最大的数据的整数倍进行存储
	//这里最大的为8个字节 所以 1个字节的需要补7个字节 ，总字节数为48
	name string		//16个字节 8 + 8
	sex byte		//1个字节
	age int			//8个字节
	addr string		//16个字节 8 + 8
}

//空结构体
type NULL struct {

}

//结构体标签
/*
	tag是类型的一部分，主要用于通过反射获取字段相关tag设置
*/
type People struct {
	Name string `json:"姓名"`
	Age uint `json:"年龄"`
}

//匿名字段
/*
	通过匿名字段实现“继承关系”
	结构体允许其成员字段在声明时没有字段名而只有类型
*/

type teacher struct {
	People
	Name string
	Score int
}


type Player struct {
	Name string
	HealthPoint int
	MagicPoint int
}

func newPlayer(name string,healthpoint ,magicpoint int) *Player{
	return &Player{				//取结构体地址进行实例化 类似于new()但是却比它效率高
		Name:        "",
		HealthPoint: 0,
		MagicPoint:  0,
	}
}

type Father struct {
	name string
	child *Father	//child的类型为结构体指针
}


//匿名结构体
func printMsgType(msg *struct{
	id int
	data string
}){
	fmt.Printf("%T\n",msg)
}


//构造函数(go语言中没有构造函数，用函数封装结构体的初始化过程)
type Cat struct {
	name string
	color string
}

func newCatName(name string)(*Cat){
	return &Cat{
		name:  name,
	}
}

func newCatColor(color string)(*Cat){
	return &Cat{
		color:color,
	}
}

//继承
type BlackCat struct {
	Cat
}

//----------------

//车轮
type wheel struct {
	Size int
}

//车
type Car struct {
	wheel
	//引擎
	Engine struct{
		Power int
		Type string
	}
}