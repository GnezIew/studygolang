package main

import "fmt"

func main() {
	//算数运算符
	//n1 := 5
	//n2 :=3
	//fmt.Println(n1+n2)
	//fmt.Println(n1-n2)
	//fmt.Println(n1*n2)
	//fmt.Println(n1 /n2)
	//fmt.Println(n1 % n2)
	//n2 ++
	//fmt.Println(n2)
	//n1--
	//fmt.Println(n1)

	//关系运算符
	//fmt.Println(n1==n2)
	//fmt.Println(n1!=n2)
	//fmt.Println(n1 >= n2)
	//fmt.Println(n1 <= n2)

	//逻辑运算符
	a:=true
	b :=false
	//两个条件都成立才为真
	fmt.Println(a && b)
	//两个条件有一个成立就为真
	fmt.Println(a || b)

	//原来为真取非就为假，原来为假取非后就为真
	fmt.Println(!a)

	//位运算
	//1101 13 /  11 3
	fmt.Printf("13的二进制%b\n",13)
	fmt.Printf("3的二进制%b\n",3)
	// &(两个对应的二进制位都为1才为1)
	fmt.Println(13 & 3)
	// | (两个对应的二进制位只要有一个为1 就为1)
	fmt.Println(13|3)
	// ^ (异或，两个对应的二进制位相异时，结果为1)
	fmt.Println(13^3)

	// << 左移	(比如3左移2位，3的二进制为11左移后为 1100 => 12)
	fmt.Println(3<<2)
	// >> 右移 (比如3右移1位，3的二进制为11右移后为1 )
	fmt.Println(3>>1)

	//赋值运算符
	num := 10
	x := 2
	num /= x //等价于num = num / x
	fmt.Println(num)
}
