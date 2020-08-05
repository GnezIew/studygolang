package main

import (
	"errors"
	"fmt"
	"time"
)

//模拟RPC客户端的请求和接受消息封装
func RPCClient (ch chan string, req string) (string,error){
	//向服务器发送请求
	ch <- req

	//等待服务器返回
	select {
	case ack := <-ch:							//接受服务器返回数据
		return ack,nil
	case <- time.After(time.Second):			//超时
		return "",errors.New("Time Out")

	}
}


//模拟RPC服务器端接受客户端请求和回应
func RPCServer(ch chan string){
	for {
		//接受客服端请求
		data := <-ch
		//打印接受到的数据
		fmt.Println("server received:",data)

		//通过睡眠函数让程序执行阻塞2秒的任务
		time.Sleep(time.Second*2)

		//向客户端反馈已收到
		ch <- "roger"
	}
}
