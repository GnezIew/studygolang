package main

import "fmt"

func main() {
	//创建动物的名字到实例的映射
	animals := map[string]interface{}{
		"bird": new(bird),
		"pig":  new(pig),
	}

	//遍历映射
	for name,obj := range animals{
		//判断对象是否为飞行动物
		f,isFlyer :=obj.(Flyer)
		//判断对象是否为行走动物
		w,isWalker :=obj.(Walker)
		fmt.Printf("name:%s isFlyer: %v isWalker :%v\n",name,isFlyer,isWalker)
		//如果是飞行动物则调用飞行动物接口
		if isFlyer{
			f.Fly()
		}

		//如果是行走动物则调用行走动物接口
		if isWalker{
			w.Walk()
		}
	}
}
