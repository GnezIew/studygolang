package main

import (
	"encoding/json"
	"fmt"
)

//声明角色的结构体
type Actor struct {

}

//为角色添加一个事件处理函数
func (a *Actor) OnEvent (param interface{}){
	fmt.Println("actor event",param)
}

//全局事件
func GlobalEvent(param interface{}){
	fmt.Println("global event:",param)
}

func main() {
	//实例化一个角色
	a := new(Actor)

	//注册名为OnSkill的回调
	RegisterEvent("OnSkill",a.OnEvent)

	//再次在OnSkill上注册全局事件
	RegisterEvent("OnSkill",GlobalEvent)

	//调用时间，所有注册的同名函数都会被调用
	CallEvent("OnSkill",100)

	//--------------------------------
	//生成一段Json数据
	jsonData := genJsonData()
	fmt.Println(string(jsonData))

	//只需要屏幕和指纹识别信息的结构和实例
	screenAndTouch := struct {
		Screen
		HasTouchID bool
	}{}
	//反序列化到screenAndTouch
	json.Unmarshal(jsonData,&screenAndTouch)

	//输出screenAndTouch的详细结构
	fmt.Printf("%+v\n",screenAndTouch)

	//只需要电池和指纹识别信息的结构和实例
	batteryAndTouch := struct {
		Battery
		HasTouchID bool
	}{}

	//反序列化到batteryAndTouch
	json.Unmarshal(jsonData,&batteryAndTouch)
	fmt.Printf("%+v\n",batteryAndTouch)
}
