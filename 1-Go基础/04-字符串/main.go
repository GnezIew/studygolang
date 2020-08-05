package main

import (
	"fmt"
	"strings"
)

//字符串操作

func main() {
	s1 := "alexdsb"
	fmt.Println(len(s1))

	//拼接字符串
	s2 := "python"
	fmt.Println(s1+s2)
	s3 := fmt.Sprintf("%s--%s",s1,s2)
	fmt.Println(s3)

	//字符串分割
	ret := strings.Split(s1,"x")
	fmt.Println(ret)
	//判断是否包含
	res := strings.Contains(s1,"dsb")
	fmt.Println(res)

	//判断前缀和后缀
	res2 := strings.HasPrefix(s1,"als")
	res3 := strings.HasSuffix(s1,"sb")
	fmt.Println(res2,res3)

	//判断字串出现的位置
	s4 := "Python and Java"
	res4 := strings.Index(s4,"a")
	res5 := strings.LastIndex(s4,"a")
	fmt.Println(res4,res5)

	//合并操作 join
	a1 := []string{"Python","PHP","JavaScript","Ruby","Golang"}		//切片
	fmt.Println(strings.Join(a1,"-"))

}
