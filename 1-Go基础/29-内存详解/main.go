package main

var obj1 *Object
var obj2 *Object

type Object struct {
	data interface{}
}

func (obj *Object) Demo(){
	//初始化
	obj1 = nil
	obj2 = obj
	//垃圾回收机制开始工作
	//obj1被标记

	//重新赋值
	obj1 = obj
	obj2 = nil
}


func main() {

}