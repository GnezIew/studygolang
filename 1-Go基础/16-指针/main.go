package main

import "fmt"

//指针和地址有什么区别？
//地址：就是内存地址(用字节来描述的内存地址)
//指针 ：指针是带类型的
// ‘&’表示取地址
// ‘*’根据地址取值
func main() {
	var a int
	fmt.Println(a)
	b := &a	//取变量a的内存地址
	fmt.Println(b)		//存储的a的内存地址
	fmt.Println(&b)		//b的内存地址
	fmt.Printf("%T\n",b)

	d := 100
	b = &d
	fmt.Println(&d)
	fmt.Println(b)

	// '*'取地址的值
	fmt.Println(*b)

	//指针的应用
	a1 := [3]int{1,2,3}
	modifyArray(a1)		//在函数中复制了数组，数组赋值给了内部的a1
	fmt.Println(a1)

	modifyArray2(&a1)
	fmt.Println(a1)


	//指针可以间接修改变量的值
	temp := 20
	fmt.Println("temp修改前的值:",temp)
	var p *int = &temp
	*p = 10
	fmt.Println("temp修改后的值:",temp)

	//内存地址编号为0到255 为系统占用不允许用户读写操作
	/*
		//var p2 *int  // nil
		//*p2 = 123					这里会报错原因在与p2的内存地址编号为0
	*/
	//new创建内存空间
	p2 := new(int)
	*p2 = 123
	fmt.Println(p2)
	fmt.Println(*p2)

}
//定义一个修改数组的函数
func modifyArray(a1 [3]int) {
	a1[0] = 100	//只修改了内部这个a1数组
}
//接受的参数是一个数组的指针
func modifyArray2(a1 *[3]int) {
	//(*a1)[0] = 100	//这里就是通过指针取的原a1的值
	//语法糖 :因为GO语言中指针不支持修改，所以a1[0]默认是(*a1)[0]
	a1[0] =100
}