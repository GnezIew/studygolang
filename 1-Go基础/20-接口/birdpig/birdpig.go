package main

import "fmt"

//定义飞行动物接口
type Flyer interface {
	Fly()
}

//定义行走动物接口
type Walker interface {
	Walk()
}

//定义鸟类
type bird struct {

}

//实现飞行动物接口
func (b *bird) Fly(){
	fmt.Println("bird: fly")
}

//为鸟类添加walk()方法,实现行走动物接口
func (b *bird) Walk() {
	fmt.Println("bird: walk")
}

//定义猪
type pig struct {

}

//为猪添加Walk()方法，实现行走动物接口
func (p *pig) Walk(){
	fmt.Println("pig : walk")
}