package main

import (
	"fmt"
	"sort"
)

func main() {
	//准备一个内容被打乱顺序的字符串切片
	names := MyStringList{
		"3. Triple Kill",
		"5. Penta Kill",
		"2. Double Kill",
		"4. Quadra Kill",
		"1. First Blood",
	}

	//使用sort包进行排序
	sort.Sort(names)

	//遍历打印结果
	for _,v := range names{
		fmt.Printf("%s\n",v)
	}


	names2 := float32Slice{
		1.2,
		5.6,
		0.1,
	}
	float32s(names2)

	for _,v := range names2{
		fmt.Printf("%f\n",v)
	}
}
