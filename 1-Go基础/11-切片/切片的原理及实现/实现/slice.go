package main

/*
#include<stdlib.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

//Slice的实现
type Slice struct {
	Data unsafe.Pointer //指向数据的指针
	len  int            //有效数据的长度
	cap  int            //有效数据的容量
}

//数据类型在内存中占的字节
const TAG = 8

//Create()创建Slice报错
func CreateSliceError() {
	panic("CreateSlice: create slice error!")
}

//Print()切片为空报错
func SliceEmptyError() {
	panic("Slice is nil")
}

//下标越界
func RangeoutError() {
	panic("out of range slice")
}

//Create(长度 , 容量 ,数据)
func (s *Slice) Create(l int, c int, Data ...int) {
	//对传递参数判断
	if l < 0 || c < 0 || l > c || len(Data) > l {
		CreateSliceError()
	}
	//创建堆内存空间 赋值给slice的Data数据
	// C语言开辟空间空间 如果成功为堆区内存地址 如果失败返回值为 NULL 相当于nil
	s.Data = C.malloc(C.ulonglong(c) * TAG) //基于容量开辟
	s.len = l
	s.cap = c

	//将s.Data转成Go语言中可以计算的指针 uintptr
	/*
		uintptr 指针类型(可以计算) , 普通指针是不能进行计算的(*int)
		unsafe.Pointer 万能指针
	*/
	p := uintptr(s.Data)
	for i := 0; i < s.len; i++ {
		*(*int)(unsafe.Pointer(p)) = 0
		p += TAG
	}
	p = uintptr(s.Data)
	for _, v := range Data {
		//数据存储
		*(*int)(unsafe.Pointer(p)) = v //不同类型的指针不能相互转化,只能通过unsafe.Pointer
		//指针偏移 存储下一个数据
		p += TAG //一个整型是8个字节
	}
}

//Print打印切片
func (s *Slice) Print() { // (s *Slice)以指针作为接受者,引用传递 //在方法内能够修改对象
	if s == nil {
		SliceEmptyError()
	}
	//将s.Data转成可计算的指针 再通过指针找出数据的值
	p := uintptr(s.Data)
	for i := 0; i < s.len; i++ {
		//打印数据
		fmt.Print(*(*int)(unsafe.Pointer(p)), " ")
		//指针偏移
		p += TAG
	}
	//设置打印后换行
	fmt.Println()
}

//Append()函数 切片追加
/*
	如果数据未超过1024，cap扩容为上一次的一倍
	如果数据超过1024，cap扩容为上一次的1/4
*/
func (s *Slice) Append(Data ...int) {
	if s == nil {
		SliceEmptyError()
	}
	if len(Data) == 0 {
		return
	}
	//如果添加的数据超出容量 需要扩容
	if s.len+len(Data) > s.cap {
		//扩充容量
		// C.realloc(指针, 字节大小)
		s.Data = C.realloc(s.Data, C.ulonglong(s.cap*2*TAG))
		s.cap *= 2
	}
	//计算指针的偏移,指向插入数据的位置
	p := uintptr(s.Data)
	for i := 0; i < s.len; i++ {
		p += TAG
	}
	//添加数据
	for _, v := range Data {
		*(*int)(unsafe.Pointer(p)) = v
		p += TAG
	}

	//更新长度
	s.len += len(Data)
}

//GetData 获取元素 GetData(下标)  返回值为int元素
func (s *Slice) GetData(index int) int {
	if s == nil {
		SliceEmptyError()
	}
	if index < 0 || index > s.len-1 {
		RangeoutError()
	}
	p := uintptr(s.Data)
	//将指针指向下标的内存位置
	for i := 0; i < index; i++ {
		p += TAG
	}
	res := *(*int)(unsafe.Pointer(p))
	return res
}

//Search 查找元素 Search(元素) 返回值为int下标
func (s *Slice) Search(Data int) int {
	//更具元素查找下标 返回第一次出现的位置
	if s == nil {
		SliceEmptyError()
	}
	p := uintptr(s.Data)
	for i := 0; i < s.len; i++ {
		if *(*int)(unsafe.Pointer(p)) == Data {
			//返回下标位置
			return i
		}
		//指针偏移
		p += TAG
	}
	return -1
}

//Delete 删除元素 Delete(下标)
func (s *Slice) Delete(index int) {
	if s == nil {
		SliceEmptyError()
	}

	if index < 0 || index > s.len {
		RangeoutError()
	}
	//将指针指向要删除元素的位置
	p := uintptr(s.Data)
	for i := 0; i < index; i++ {
		p += TAG
	}
	//用下一个指针为当前的指针赋值
	temp := p
	for i := index; i < s.len; i++ {
		temp += TAG
		*(*int)(unsafe.Pointer(p)) = *(*int)(unsafe.Pointer(temp))
		p += TAG
	}
	s.len -= 1
}

//Insert 插入元素 Insert (下标  元素)
func (s *Slice) Insert(index int, Data int) {
	if s == nil {
		SliceEmptyError()
	}
	if index < 0 || index > s.len  {
		RangeoutError()
	}
	//如果下标和 s.len-1相同
	if index == s.len-1{
		//调用追加函数
		s.Append(Data)
	}
	//扩容(写或不写都行)
	if s.len == s.cap{
		s.Data = C.realloc(s.Data,C.ulonglong(s.cap*2*TAG))
		s.cap *=2
	}
	//获取需要插入数据位置
	p := uintptr(s.Data)
	for i := 0; i < index; i++ {
		p += TAG
	}

	//获取末尾的指针
	temp := uintptr(s.Data)
	temp += uintptr(s.len) * TAG

	for k := s.len; k > index; k-- {
		*(*int)(unsafe.Pointer(temp)) = *(*int)(unsafe.Pointer(temp - TAG))
		temp -= TAG
	}
	*(*int)(unsafe.Pointer(p)) = Data
	s.len+=1
}

//Destroy 销毁切片
func (s *Slice) Destroy(){
	//释放开辟的堆空间
	C.free(s.Data)
	//将slice结构体重置
	s.Data = nil
	s.len = 0
	s.cap = 0
}