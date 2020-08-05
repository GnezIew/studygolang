package main

import (

	"fmt"
)

func testSwitch3() {
	switch n := 7; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	default:
		fmt.Println(n)
	}
}

func switchDemo4() {
	age := 30
	switch {
	case age%2 == 0:
		fmt.Println("偶数")
	case age%2 != 0:
		fmt.Println("奇数")
	}
}

func switchDemo5() {
	tmpe := "a"

	switch tmpe {
	case "a":
		fmt.Println("a")
		fallthrough //会下穿，无条件执行下一个case语句
	case "b":
		fmt.Println("b")
	}
}
func main() {
	//age :=17
	//if age > 18{
	//	fmt.Println("澳门首家线上赌场开业了")
	//}else if age < 18 {
	//	fmt.Println("Warning")
	//}else{
	//	fmt.Println("成年了！")
	//}
	//
	//if age2 := 8;age2 > 18{			//这里限制了age2的作用域范围
	//	fmt.Println("成年了")
	//}else if age2 < 18{
	//	fmt.Println("未成年")
	//}

	//for循环
	//age := 18
	//for age > 0{			//相当于其他语言的while
	//	fmt.Println(age)
	//	age--
	//}

	//switch循环
	//finger :=4
	//
	//switch finger {
	//case 1:
	//	fmt.Println("大拇指")
	//case 2:
	//	fmt.Println("食指")
	//case 3:
	//	fmt.Println("中指")
	//case 4:
	//	fmt.Println("无名指")
	//}
	testSwitch3()
	switchDemo4()
	switchDemo5()

	//goto
	//for i:=0; i < 5;i++{
	//	for j:=0;j<3;j++{
	//		//当i = 2，j = 2时跳出两层for循环
	//		if i == 2 && j==2{
	//			goto breaklab
	//		}
	//		fmt.Printf("%d--%d\n",i,j)
	//	}
	//}
	//breaklab:		//定义一个标签
	//fmt.Println("两层for循环结束")

	//break语句可以结束for,switch,select的代码块
	//break语句后面可以添加标签，表示退出某个标签对应的代码块
	//BreakDemo1:
	//	for i :=0;i<5;i++{
	//		for j:=0;j<10;j++{
	//			if j==2{
	//				break BreakDemo1
	//			}
	//			fmt.Printf("%d--%d\n",i,j)
	//		}
	//
	//	}

	//continue只能在for循环中使用

	//for i :=0;i<5;i++{
	//	for j:=0;j<10;j++{
	//		if i == 2 && j==2{
	//			continue
	//		}
	//		fmt.Printf("%v--%v\n",i,j)
	//	}
	//}

	//用for循环写一个九九乘法表
	//for i :=1;i<10;i++{
	//	for j:=1;j<=i ;j++  {
	//		fmt.Printf("%d * %d = %d\t",j,i,i*j)
	//	}
	//	fmt.Println()
	//}

	//求200-1000以内的所有的素数

	for i := 200; i <= 1000; i++ {
		Flag := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				Flag = false
				break
			}
		}
		if Flag {
			fmt.Printf("%d\t\n", i)
		}
	}

	var m int
	fmt.Print("请输入月份:")
	fmt.Scan(&m)
	switch m {
	case 1, 3, 5, 7, 8, 10, 12:
		fmt.Println("31")
	case 4, 6, 9, 11:
		fmt.Println("30")
	case 2:
		fmt.Println("28 or 29")
	}

	//古代数学家 张丘建 “算经”
	// 百钱百鸡 ： 公鸡5钱一只 母鸡3钱一只  小鸡1钱三只 现在需要用百钱买百鸡
	// 100 买公鸡 20只
	// 100 买母鸡 33只
	// 100 买小鸡 100只
	demo()

}
