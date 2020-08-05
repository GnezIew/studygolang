package main

import "fmt"

func demo() {
	//古代数学家 张丘建 “算经”
	// 百钱百鸡 ： 公鸡5钱一只 母鸡3钱一只  小鸡1钱三只 现在需要用百钱买百鸡
	// 100 买公鸡 20只
	// 100 买母鸡 33只
	// 100 买小鸡 100只
	//count := 0
	//for i := 0; i <= 20; i++ {
	//	for j := 0; j <= 33; j++ {
	//		for k := 0; k <= 100; k++ {
	//			count++
	//			if k%3 == 0 && i+j+k == 100 && i*5+j*3+k/3 == 100 {
	//				fmt.Println("公鸡:", i, "母鸡:", j, "小鸡:", k)
	//			}
	//		}
	//	}
	//}
	//fmt.Println(count)

	//优化
	count := 0
	for i := 0; i <= 20; i++ {
		for j := 0; j <= 33; j++ {
			count ++
			// 小鸡的个数 100 - 公鸡 + 母鸡
			k := 100 - i - j
			if k%3==0 && i*5 + j*3 + k/3 == 100 {
				fmt.Println("公鸡:", i, "母鸡:", j, "小鸡:", k)
			}
		}
	}
	fmt.Println(count)
}
