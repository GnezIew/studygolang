package main

import (
	"fmt"
	"os"
)

//命令行写入器
type consoleWriter struct {

}

//实现LogWriter 的Writer()方法
func (f *consoleWriter) Write(data interface{}) error{
	//将数据序列化为字符串
	str := fmt.Sprintf("%v\n",data)

	//数据以字节数组写入命令行
	_,err := os.Stdout.Write([]byte(str))
	return err
}

//创建命令喊写入器实例
func newConsoleWriter() *consoleWriter{
	return &consoleWriter{}
}