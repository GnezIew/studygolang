package main

import (
	"fmt"
	"math"
)


func main() {
	var a int = 10
	var b int = 077
	var c int = 0xff
	fmt.Println(a,b)
	fmt.Printf("%b\n",a)
	fmt.Printf("%o\n",b)
	fmt.Printf("%x\n",c)
	//求C变量的内存地址
	fmt.Printf("%p\n",&c)

	//浮点数常量
	//建议使用float64代替float32
	fmt.Println(math.MaxFloat32)
	var a2 float64 = 3.14			//精确到小数点后15位
	var b2 float32 = 3.14
	fmt.Printf("%.20f\n",a2)		// ‘.20f’保留20位，默认保留六位，会对第七位四舍五入
	fmt.Printf("%.20f\n",b2)

	//Go语言中不允许将整型强制转换为布尔型。


	//字符串转义
	fmt.Println("\"c:\\go\"")
	var s1 = "单行文本"
	var s2 = `这是
              "不需要转义"
              多行字符`

	fmt.Println(s1)
	fmt.Println(s2)
}
