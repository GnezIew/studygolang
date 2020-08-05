package main

import "fmt"

func main() {

	//创建一个无缓冲字符通道
	ch := make(chan string)

	//并发执行服务器逻辑
	go RPCServer(ch)

	//客户端请求数据和结束数据
	recv,err := RPCClient(ch,"hi")
	if err != nil{
		//发生错误打印错误
		fmt.Println(err)
	}else{
		//正常接受到数据
		fmt.Println("Client recived",recv)
	}
}
