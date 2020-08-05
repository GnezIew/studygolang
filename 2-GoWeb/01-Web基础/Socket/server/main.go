package main

import (
	"bufio"
	"fmt"
	"net"
)

//处理连接函数
func process(conn net.Conn){
	defer conn.Close()	//关闭连接
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n,err := reader.Read(buf[:]) //读取数据
		if err != nil {
			fmt.Println("read from client failed,err:",err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client端发来的数据:",recvStr)
		sendStr := "我收到了"
		conn.Write([]byte(sendStr))	//发送数据
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("read from client failed,err:",err)
		return
	}
	for {
		conn,err := listen.Accept()		//建立链接
		if err !=nil{
			fmt.Println("accept failed,err:",err)
			continue
		}
		go process(conn)
	}
}
