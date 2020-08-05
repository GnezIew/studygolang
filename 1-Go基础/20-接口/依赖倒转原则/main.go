package main

import (
	"fmt"
)

//抽象层
type Car interface {
	Run()
}

type Driver interface {
	Driver(car Car)
}

//实现层
type Benz struct {

}
func (benz *Benz)Run(){
	fmt.Println("Benz is running")
}

type Bmw struct {

}
func (bmw *Bmw)Run(){
	fmt.Println("Benz is running")
}

type Zhang_3 struct {

}

func(zhang3 *Zhang_3) Driver(car Car){
	fmt.Println("zhang3 drive car")
	car.Run()
}

type li_4 struct {

}
func (li4 *li_4) Driver(car Car){
	fmt.Println("li4 drive car")
	car.Run()
}


func main() {
	//业务逻辑
	//张三开宝马
	//var bmw Car
	//bmw = &Bmw{}
	//
	//var zhang3 Driver
	//zhang3 = &Zhang_3{}
	//
	//zhang3.Driver(bmw)
	//
	////李四开奔驰
	//var benz Car
	//benz = &Benz{}
	//
	//var li4 Driver
	//li4 = &li_4{}
	//
	//li4.Driver(benz)

	//intel系列电脑
	com1 := NewComputer(&IntelCPU{},&IntelMemory{},&IntelCard{})
	com1.Dowork()

	//杂牌子
	com2 := NewComputer(&IntelCPU{},&KingstonMemory{},&NvidiaCard{})
	com2.Dowork()
}
