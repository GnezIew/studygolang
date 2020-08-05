package main

import "reflect"

type State interface {
	//获取状态名称
	Name()	string

	//状态是否允许在同状态间转移
	EnableSameTransit() bool

	//能否从当前状态转移到指定的某种状态
	CanTransitTo(name string) bool

	//响应状态开始时
	OnBegin()

	//响应状态开始后
	OnEnd()
}

//从状态实例获取状态名
func StateName(s State) string{
	if s == nil{
		return "None"
	}
	//使用反射获取状态的名称
	return reflect.TypeOf(s).Elem().Name()
}