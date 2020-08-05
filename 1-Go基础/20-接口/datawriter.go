package main

import "fmt"

//定义一个数据写入器
type DataWriter interface {
	WriteData(data interface{}) error
}


//定义文件结构, 用于实现Datawriter
type file struct {

}

//实现DataWriter 接口的WriteData() 方法
func (f *file) WriteData(data interface{}) error {
	//模拟写入数据
	fmt.Println("WriteData:",data)
	return nil
}
