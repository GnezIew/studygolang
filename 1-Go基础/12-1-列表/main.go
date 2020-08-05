package main

import (
	"container/list"
	"fmt"
)

func main() {
	/*
		列表 内部的实现原理是双链表
	*/
	//列表的声明 两种声明方式
	lst := list.New()
	//var lst  list.List

	/*
		列表的插入：
			1.头部插入	PushFront
			2.尾部插入	PushBack
		两种方法都会返回一个*list.Element结构(可以理解为一个节点),可以配合Remove()方法进行删除
			3.InsertAfter 任意节点之后插入
	        4.InsertBefore 任意节点之前插入
	*/
	//头部插入
	lst.PushFront(1)
	//尾部插入
	element1 :=lst.PushBack(2)
	//节点之后插入
	lst.InsertAfter("后",element1)
	//节点前插入
	lst.InsertBefore("前",element1)
	//节点删除
	lst.Remove(element1.Prev())

	//访问列表中的每一个元素
	for i:=lst.Front();i != nil ;i = i.Next() {
		fmt.Println(i.Value)
	}

}
