package main

import "fmt"

func main() {

	//创建字典实例
	dict :=NewDictionary()

	//添加游戏数据
	dict.Set("My Factory",60)
	dict.Set("Terra Craft",36)
	dict.Set("Don't Hungry",24)

	//获取值及打印值
	favorite := dict.Get("Terra Craft")
	fmt.Println("favorite:",favorite)

	//遍历所有的字典元素
	dict.Visit(func(key, value interface{}) bool {

		//将值转为int类型，并判断是否大于40
		if value.(int) > 40{
			//输出很贵
			fmt.Println(key,"很贵")
			return true
		}
		//默认输出“很便宜”
		fmt.Println(key,"is cheap")
		return true
	})
}
