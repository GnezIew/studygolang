package main

import "fmt"

//测试文件
func main() {
	var s Slice
	//创建切片
	s.Create(5,10,1,2,3,4,5)
	//追加数据
	s.Append(6,7)
	//打印切片
	s.Print()
	//根据下标获取数据元素的值
	res := s.GetData(1)
	fmt.Println("下标为1的元素值:",res)
	//根据元素值查找下标
	index := s.Search(5)
	fmt.Println("元素值为5的下标:",index)

	//删除对应下标的元素
	s.Delete(4)
	s.Print()

	//插入切片
	s.Insert(2,8)
	s.Print()

	//销毁切片
	s.Destroy()

}
