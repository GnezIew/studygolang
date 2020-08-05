package main

import "fmt"

func main() {
	//用单引号包裹字符

	s1 := "Golang"
	c1 := 'G' //输出ASCII码，占一个字节(8位 ,8 bit)
	fmt.Println(s1, c1)

	s2 := "中国"
	c2 := '中' //UTF-8下中文占3个字节
	fmt.Println(s2, c2)

	s3 := "hello沙河" //11
	fmt.Println(len(s3))

	//rune类型在处理复合类型时要用到
	//遍历字符串
	for i := 0; i < len(s3); i++ {
		fmt.Printf("%c\n", s3[i])
	}
	fmt.Println()

	//for range 循环是按照rune类型去遍历的
	for k, v := range s3 {
		fmt.Printf("%d%c\n", k, v)
	}

	//修改字符串
	s4 := "big"
	//强制类型转换;将字符串强制转换成字节数组类型
	byteS4 := []byte(s4)
	byteS4[0] = 'p'
	fmt.Println(string(byteS4))

	s5 := "白萝卜rab"
	runeS5 := []rune(s5)
	runeS5[0] = '红'
	fmt.Println(string(runeS5))

	//作业题
	//1.把hello --> "olleh"
	s6 := "hello"
	byteArray := []byte(s6)
	//s7 := ""
	//for i:=len(byteArray)-1;i>=0;i--{
	//	s7 =s7 + string(byteArray[i])
	//}
	//fmt.Println(s7)

	for i := 0; i < len(s6)/2; i++ {
		byteArray[i], byteArray[len(s6)-1-i] = byteArray[len(s6)-1-i], byteArray[i]
	}
	fmt.Println(string(byteArray))

	//反引号能够将字符串中的内容原样输出，不会转义
	str := `瞅你咋地\n不服就干`
	fmt.Println(str) //结果 : 瞅你咋地\n不服就干

	//类型转换是不会进行四舍五入的
	price := 3.25
	weight := 5
	fmt.Println(price * float64(weight)) //不同类型是不能进行运算的，所以需要转化类型

	fmt.Println(int(price * float64(weight)))

}
