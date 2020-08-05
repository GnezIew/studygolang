package main

import "fmt"

func main() {
	//new是用来初始化值类型指针的
	//make是用来初始化slice, map,chan
	var a = new(int)	//得到一个int类型的指针
	fmt.Println(a)
	*a = 10
	fmt.Println(*a)

	var b = new([3]int)
	fmt.Println(b)
	//(*b)[0] = 10
	b[0] = 10
	fmt.Println(*b)


}
