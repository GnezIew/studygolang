package main

import (
	"fmt"
	"unsafe"

)

// Go语言规定所有数据作为函数参数为值传递
func Demo(slice []int) {
	//追加元素
	fmt.Printf("append()增加元素前 slice中数据集合的地址:%p\n",slice)
	slice = append(slice, 6)
	fmt.Println("函数中:", slice)
	fmt.Printf("append()增加元素后 slice中数据集合的地址:%p\n",slice)

	fmt.Println("slice在内存中的大小:", unsafe.Sizeof(slice))
}
func main() {
	slice := []int{1,2,3,4}
	fmt.Printf("函数调用之前slice中数据集合的地址:%p\n",slice)
	//将切片作为函数参数传递(值传递)
	Demo(slice)
	fmt.Println("定义中:", slice)
	fmt.Printf("函数调用之后 主函数中 slice 中数据集合的地址:%p\n",slice)
	fmt.Println("slice在内存中的大小:", unsafe.Sizeof(slice))

	/*
		切片的结构体如下 :
			type slice struct {
				array unsafe.Pointer		//8个字节	//指针指向数据集的地址 万能指针类型 对应C语言中的void*
				len   int				    //8个字节
				cap   int				    //8个字节
		     }
	*/
	
}
