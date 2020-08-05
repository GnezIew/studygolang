package main

import (
	"fmt"
	"reflect"
)

//定义一个Enum的类型
type Enum int

const (
	Zero Enum = 0
)

func demo(){
	//声明整型变量a并赋初值
	var a int = 1024

	//获取变量a 的反射值对象(a 的地址)
	valueOfA := reflect.ValueOf(&a)

	//取出a地址的元素(a的值)
	valueOfA = valueOfA.Elem()

	//修改a的值为1
	valueOfA.SetInt(1)

	//打印a的值
	fmt.Println(valueOfA.Int())
	fmt.Println(a)
}

func demo2(){
	type dog struct {
		LegCount int
	}

	//获取dog实例地址的反射值对象
	valueOfDog := reflect.ValueOf(&dog{})

	//取出dog实例地址的元素
	valueOfDog = valueOfDog.Elem()

	//获取legCount字段的值
	vLegCount := valueOfDog.FieldByName("LegCount")

	//尝试设置legcount的值
	vLegCount.SetInt(4)

	fmt.Println(vLegCount.Int())
}

func main() {
	//var a int
	//typeOfA := reflect.TypeOf(a)
	//fmt.Println(typeOfA)		//类型对象
	//fmt.Println(typeOfA.Name(),typeOfA.Kind())		//通过类型对象访问任意值的类型信息
	//-------------------------------------------------------------------------------

	//声明一个空结构体
	//type cat struct {
	//
	//}
	//
	////获取结构体实例的反射类型对象
	//typeOfCat :=  reflect.TypeOf(cat{})
	//
	////显示反射类对象的名称和种类
	//fmt.Println(typeOfCat.Name(),typeOfCat.Kind())
	//
	////获取Zero常量的反射类型对象
	//typeOfA := reflect.TypeOf(Zero)
	//
	////显示反射类对象的名称和种类
	//fmt.Println(typeOfA.Name(),typeOfA.Kind())

	//--------------------------------------------------------
	//声明空结构体
	//type cat struct {
	//
	//}
	//
	////创建cat实例
	//ins := &cat{}
	//
	////获取结构体实例的类型对象
	//typeOfCat := reflect.TypeOf(ins)
	//
	////显示反射类型对象的名称和种类
	//fmt.Printf("name:%v kind:%v\n",typeOfCat.Name(),typeOfCat.Kind())
	//
	////获取类型的元素
	//typeOfCat = typeOfCat.Elem()
	//fmt.Printf("element name : %v , element kind : %v\n",typeOfCat.Name(),typeOfCat.Kind())

	//-----------------------------------------------------------------------------------------
	//声明整型变量a 并赋初值
	//var a int = 1024
	//
	////获取变量 a 的反射值对象
	//valueOfA := reflect.ValueOf(a)
	//
	////获取interface{}类型的值，通过类型断言转换
	//var getA int = valueOfA.Interface().(int)
	//
	////获取64位的值，强制类型转换为int类型
	//var getA2 int = int(valueOfA.Int())
	//fmt.Println(getA,getA2)
	//-------------------------------------------------------------------------------
	demo()
}
