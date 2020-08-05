package main

import (
	"fmt"
	"strings"
	"unsafe"
)

func main() {
	//map映射关系容器，其内部使用散列表(hash)实现
	//map是一种无序的基于 key-value的数据结构
	//map是引用类型
	//var a map[string]int	//只声明
	//fmt.Println(a == nil)
	////map的初始化
	//a = make(map[string]int,8)		//这里是可以不写8的，但是写的好处就是减少内存再分配的次数，增加效率
	//fmt.Println(a)
	//
	////map中添加key-value
	//a["北京"] = 100
	//a["上海"] = 90
	//fmt.Printf("%#v\n",a)
	//fmt.Printf("type:%T\n",a)
	//
	//
	////声明map的同时完成初始化
	//b := map[int]bool{
	//	1:true,
	//	2:false,
	//}
	//fmt.Printf("%#v\n",b)


	//判断map中某个键存不存在
	//var score = map[string]int{
	//	"小明": 100,
	//	"小红": 98,
	//	"韩梅梅":90,
	//}
	//判断"小王" 存不存在 score中
	//value,ok := score["小王"]
	//if ok {
	//	fmt.Println("小王在score中",value)
	//}else{
	//	fmt.Println("小王不在score中",value)
	//}
	//
	////map的遍历(因为map是无序的所以遍历出来的结果可能和你添加的顺序不一致)
	//for k,v :=range score{
	//	fmt.Println(k,v)
	//}
	//
	////只遍历map的key
	//for k :=range score{
	//	fmt.Println(k)
	//}
	//
	////只遍历map中的value
	//for _,value := range score{
	//	fmt.Println(value)
	//}

	//使用delete函数删除map中指定的键值对
	//delete(score,"小明")
	//fmt.Println(score)

	//------------------------------
	//按照某个固定的顺序遍历map
	//var score = make(map[string]int,100)
	////添加50个键值对
	//for i:=0 ;i<50;i++{
	//	key :=fmt.Sprintf("stud%02d",i)	//Sprintf是把格式化字符串，输出到指定的字符串中
	//	value := rand.Intn(100)	//生成0~99的随机整数
	//	score[key] = value
	//}
	//
	////按照key从小到大的顺序去遍历score
	////1.先取出所有的key存放到切片中
	//keys := make([]string,0,100)
	//for k := range score{
	//	keys = append(keys, k)
	//}
	//fmt.Println(keys)
	////2.对取出来的key做排序
	//sort.Strings(keys)
	//fmt.Println(keys)
	////3.按照排序后的key对score排序
	//for _,value := range keys{
	//	fmt.Println(value,score[value])
	//}


	//---------------------------------------
	//元素类型为map的切片
	//var mapSlice = make([]map[string]int,8,8)	//只完成了切片的初始化
	////还需要完成mapSlice中map元素的初始化
	//mapSlice[0] = make(map[string]int,8)		//完成了map的初始化
	//mapSlice[0]["小明"] = 100
	//mapSlice[0]["小李"] = 90
	//mapSlice[1] = make(map[string]int,8)
	//mapSlice[1]["小王"] = 80
	//fmt.Println(mapSlice)

	//---------------------------------------
	//定义一个值为切片的map
	//var sliceMap = make(map[string][]int,8)	//只完成了map的初始化
	//_,ok := sliceMap["中国"]
	//if ok {
	//	fmt.Println(sliceMap["中国"])
	//}else{
	//	//sliceMap中没有“中国”这个键
	//	sliceMap["中国"] = make([]int,8,8)//完成了对切片的初始化,切片长度为8(即切片中默认为8个0值的切片)
	//	sliceMap["中国"][0] = 100
	//	sliceMap["中国"][1] = 200
	//	sliceMap["中国"][2] = 300
	//}
	//sliceMap["巴基斯坦"] = make([]int,0,8)
	//sliceMap["巴基斯坦"] = append(sliceMap["巴基斯坦"], 200)
	//fmt.Println(sliceMap)
	//
	//for k,v := range sliceMap{
	//	fmt.Println(k,v)
	//}

	//--------------------------------------
	//统计一个字符串中每个单词出现的次数
	// how do you do 中每个单词出现的次数
	Str1 := "what are you doing? you just get out here"
	//0.定义一个map[string][0]
	var wordCount = make(map[string]int,10)
	//1.字符串中都有哪些元素
	words := strings.Split(Str1," ")		//ret是切片类型
	//2.遍历单词做统计
	for _,v :=range words{
		_,ok := wordCount[v]
		if ok{ //map中有这个单词的统计记录
			wordCount[v] +=1
		}else{
			wordCount[v] = 1
		}
	}
	fmt.Println(wordCount)


	// map的key的类型 必须支持 == , !=运算符的操作 不可以是函数 map  切片
	// map是无序的,在做查找时速度快
	m := map[int]string{1001:"法师",1002:"贾克斯"}
	fmt.Println(m)
	//删除元素 delete
	//delete(m,1001)
	//fmt.Println(m)

	//计算数据类型在内存中占的字节大小
	//内存大小永远是8个字节 ， 因为存储的是一个地址
	fmt.Println(unsafe.Sizeof(m))
	fmt.Printf("%p\n",m)

	//sync.Map的使用
	demo()
}