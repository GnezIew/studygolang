package main

import "fmt"

func main() {
	var a = 100
	var b = "沙河哪吒"
	var c =false
	fmt.Println(a,b,c)

	//%v俗称占位符
	fmt.Printf("a = %v\n",a)	//在不知道变量的类型的时候，可以用%v来进行输出
	//%T查看类型
	fmt.Printf("a的类型：%T\n",a)
	//百分号转义 %%
	fmt.Printf("%d%%\n",a)

	//%+d带符号的整型,会输出整数的正负号
	fmt.Printf("%+d\n",a)

	//%q打印单引号/双引号
	fmt.Printf("%q\n",b)

	//指定长度的整型
	fmt.Printf("|%5d|\n",a)		//长度为5右对齐，左边留白
	fmt.Printf("|%-5d|\n",a)		//左对齐右边留白
	fmt.Printf("|%05d|\n",a)			//右对齐，

	//浮点数%f默认保留小数点后两位
	f1 := 3.1415926
	fmt.Printf("%f\n",f1)
	fmt.Printf("%.2f\n",f1)
	//g表示总共多少数字
	fmt.Printf("%.2g\n",f1)

	//指定子串长度
	s1 := "这是一个字符串"
	//最小宽度
	fmt.Printf("|%20s|\n",s1)	//右对齐，最小宽度20
	fmt.Printf("|%-20s|\n",s1)		//左对齐，最小宽度20

	//最大宽度
	fmt.Printf("|%.5s|\n",s1)	//最大宽度为5

	//范围宽度
	fmt.Printf("|%10.20s|\n",s1)	//最小宽度为10，最大宽度为20	右对齐

	//布尔值
	fmt.Printf("%t\n",c)
}