package main

import (
	"fmt"
	"sort"
)

func main() {
	//准备英雄列表
	heros := Heros{
		&Hero{
			Name: "吕布",
			Kind: Tank,
		},
		&Hero{
			Name: "李白",
			Kind: Assassin,
		},
		&Hero{
			Name: "妲己",
			Kind: Mage,
		},
		&Hero{
			Name: "貂蝉",
			Kind: Assassin,
		},
		&Hero{
			Name: "关羽",
			Kind: Tank,
		},
		&Hero{
			Name: "诸葛亮",
			Kind: Mage,
		},
	}
	////使用sort包进行排序
	//sort.Sort(heros)
	//
	////遍历英雄列表打印排序结果
	//for _,v := range heros{
	//	fmt.Printf("%+v\n",v)
	//}

	sort.Slice(heros , func(i, j int) bool {
		if heros[i].Kind != heros[j].Kind{
			return heros[i].Kind < heros[j].Kind
		}
		return heros[i].Name < heros[j].Name
	})
	//遍历英雄列表打印排序结果
	for _,v := range heros{
		fmt.Printf("%+v\n",v)
	}
}
