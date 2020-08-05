package main

import (
	"code.zengwei.com/studygolang/demoTest"
	"fmt"
	"reflect"
	"unsafe"
)

/*
	一、自定义类型和类型别名
	byte : uint8的别名
	rune : int32的别名
*/
//自定义类型
// NewInt 是一个新类型
type NewInt int //基于Go语言int类型自定义的一种新类型

//类型别名:只存在代码编写过程中，代码编译之后根本不存在 "MyInt" 只是增加代码可读性
type MyInt = int
type FUNCTYPE func(int, bool, string)

func demo(a int, b bool, c string) {
	fmt.Println("函数调用")
}
func demo2() {
	fmt.Println("我是demo2")
}
func demo3(f func()) {
	f()
}

/*
	二、结构体
	结构体名称或成员变量 首字母大写表示能够对外暴露，小写只能在内部使用
*/
//创建新的类型要使用type关键字

type student struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	/*
		var a NewInt
		fmt.Println(a)
		fmt.Printf("%T\n",a)

		var b MyInt
		fmt.Println(b)
		fmt.Printf("%T\n",b)
	*/
	var f FUNCTYPE
	//将函数赋值给函数类型的变量
	f = demo
	f(10, true, "123")
	demo3(demo2)
	//----------------------------
	//结构实例化 ---- 第一种 // 实例化 = 声明
	var stu1 = student{
		name:   "曾为",
		age:    19,
		gender: "男",
		hobby:  []string{"篮球", "足球", "台球"},
	}
	fmt.Println(stu1)
	//结构体支持访问属性
	fmt.Println(stu1.name)
	fmt.Println(stu1.age)
	fmt.Println(stu1.gender)
	fmt.Println(stu1.hobby)

	//结构实例化 ---- 第二种
	//struct是值类型
	//如果初始化时没有给属性(字段) 设置对应的初始值 , 那么这些属性都是默认的空置
	var stu2 = student{}
	fmt.Println(stu2.name)
	fmt.Println(stu2.age)
	fmt.Println(stu2.gender)
	fmt.Println(stu2.hobby)

	////结构实例化 ---- 第三种
	var stu3 = new(student)
	//(*stu3).name = "董聃"
	stu3.name = "董聃"
	stu3.age = 18
	fmt.Println(*stu3)

	// 实例化方法4 :var stu4 = &student{} 和上面方法3一样

	//初始化方法1
	var stu4 = student{
		"曾为",
		18,
		"男",
		[]string{"画画"},
	}
	fmt.Println(stu4.hobby)

	//初始化方法2
	var stu5 = &student{
		name:   "董聃",
		age:    18,
		gender: "女",
		hobby:  []string{"曾为"},
	}
	fmt.Println(*stu5)

	var stu6 student
	stu6.name = "曾为"
	stu6.hobby = []string{"篮球"}
	stu6.age = 18
	stu6.gender = "男"

	var ret interface{}
	ret = 123
	//t1 := reflect.TypeOf(ret)
	//类型断言
	ret2, ok := ret.(int)
	fmt.Println(ret2+1, ok)

	//结构体名称或成员变量 首字母大写表示能够对外暴露，小写只能在内部使用
	var Stu demoTest.Student
	Stu.Id = 100
	Stu.Name = "曾为"

	//go func() {
	//	fmt.Println("hello world")
	//}()
	//time.Sleep(time.Second * 10)

	//------------------------------------------------------------
	var str string
	str = "hello"
	fmt.Println("str在内存中的大小:", unsafe.Sizeof(str)) // 8 + 8  一个存储的是地址 一个存储的是字符串长度
	var s Student
	fmt.Printf("%p\n", &s)
	fmt.Printf("%p\n", &s.name)
	fmt.Printf("%p\n", &s.sex)
	fmt.Printf("%p\n", &s.age)
	fmt.Printf("%p\n", &s.addr)
	fmt.Println(unsafe.Sizeof(s))

	//空结构体
	/*
		作用:
			在特殊的channel中使用 不可以写入数据 只有close关闭才能进行输出操作
	*/
	//a := struct {
	//
	//}{}
	a := NULL{}
	a2 := struct {
	}{}
	fmt.Println("空结构体的内存大小:", unsafe.Sizeof(a))
	//空结构体指向同一片地址，避免内存滥用
	fmt.Printf("a空结构体在内存中的内存地址:%p\n", &a)
	fmt.Printf("a2空结构体在内存中的内存地址:%p\n", &a2)

	//结构体标签
	typeofCat := reflect.TypeOf(People{})
	fmt.Printf("%T\n", typeofCat)
	//通过反射获取Tag设置
	for i := 0; i < typeofCat.NumField(); i++ {
		field := typeofCat.Field(i)
		tag := field.Tag.Get("json")
		fmt.Println(tag)
	}

	//匿名字段 ("继承")
	peo := People{
		Name: "曾为",
		Age:  18,
	}
	teach := teacher{
		People: peo,
		Name:   "为为",
		Score:  100,
	}
	fmt.Println(teach.Name)		//就近原则
	fmt.Println(teach.People.Name)


	//创建指针类型的结构体
	tank := new(Player)
	tank.Name ="Canonn"	//这里的本来应该是(*tank).Name 由于语法糖技术，所以 tank.Name 等价于 (*tank).Name
	tank.HealthPoint = 300
	fmt.Printf("tank的地址:%p\n",tank)


	//初始化
	//father := &Father{
	//	name:"爷爷",
	//	child:&Father{
	//		name : "我",
	//	},
	//}
	//fmt.Println(*father)

	//实例化一个匿名结构体
	msg := &struct {		//定义部分
		id int
		data string
	}{					//初始化
		1,
		"hello",
	}
	printMsgType(msg)

	//构造函数(go语言中没有构造函数，用函数封装结构体的初始化过程)
	cat :=newCatName("咪咪")
	fmt.Println(*cat)
	cat = newCatColor("橘猫")
	fmt.Println(*cat)

	//继承
	blackcat := BlackCat{Cat{color:"黑色",name:"黑猫警长"}}
	fmt.Println("继承:",blackcat)



	c := Car{
		//初始化轮子
		wheel :wheel{Size:18},
		Engine: struct {
			Power int
			Type  string
		}{Power:143 , Type:"1.4T" },
	}
	fmt.Println(c)

}

