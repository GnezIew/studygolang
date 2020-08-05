package main

import (
	"fmt"
)

func goRoutineA(a <-chan int){
	val := <-a
	fmt.Println("goRoutineA:",val)
}

func goRoutineB(b chan int){
	val := <-b
	fmt.Println("goRoutineB:",val)
}


func main() {
	//ch := make(chan int,10)
	//fmt.Println(unsafe.Sizeof(ch))		//8个字节
	//ch := make(chan int ,3)
	//go goRoutineA(ch)
	//go goRoutineB(ch)
	//ch <- 3
	//ch <- 4
	//time.Sleep(time.Second)

}
