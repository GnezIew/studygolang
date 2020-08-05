package main

import "fmt"

func main() {
	//var a [3]int
	//a = [3]int{1,2,3}
	////声明二维数组
	//var b [3][2]int
	//b = [3][2]int{
	//	{1,2},
	//	{3,4},
	//	{5,6},
	//}
	//fmt.Println(a)
	//fmt.Println(b)
	//
	////声明的同时完成赋值
	//var c = [3][2]int{
	//	{1,2},
	//	{3,4},
	//}
	//fmt.Println(c)
	//
	//
	////注意事项：多维数组除了第一层其他层都不能用[...]
	//var d = [...][2]int{
	//	{2,3},
	//	{4,5},
	//}
	//fmt.Println(d)
	//
	////多维数组的使用
	//fmt.Println(d[1][1])

	//for i:=0; i<len(d);i++{
	//	for j:=0;j<len(d[i]);j++{
	//		fmt.Println(d[i][j])
	//	}
	//}

	//for _,v:=range d{
	//	for _,v2:=range v{
	//		fmt.Println(v2)
	//	}
	//}

	//数组是值类型
	a := [2]int{1,2}
	b := a		//相当于把a整个拷贝一份赋值给了b
	b[0] = 100
	fmt.Println(a)
	fmt.Println(b)

	c := [2][2]int{
		{1,2},
		{3,4},
	}
	d := c
	d[0][0] =100
	fmt.Println(c)
	fmt.Println(d)

}
