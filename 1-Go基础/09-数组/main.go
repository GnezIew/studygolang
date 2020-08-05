package main

import "fmt"


func BubbleSort(arr *[10]int){	//冒泡排序
	for i:=0;i<len(arr)-1;i++{
		for j:=0;j <len(arr)-1-i ;j++  {
			//比较两个相邻元素 满足条件交换数据
			if arr[j] > arr[j+1]{
				arr[j],arr[j+1] = arr[j+1],arr[j]	//多重赋值
			}
		}
	}
}




//数组是同一种数据类型元素的集合
//声明数组时大小是不可变化的，可以修改数组成员
//在数组中长度也数组类型的一部分
func main() {
	//声明
	var a [3]int  //定义一个长度为5，存放int类型的数组
	var b [10]int //定义一个长度为10，存放int类型的数组

	//初始化
	a = [3]int{1, 2, 3}
	b = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(a)
	fmt.Println(b)
	//-------------------------------------------
	//声明+初始化
	var c = [3]string{"北京", "上海", "深圳"}
	c2 := [3]string{"北京", "上海", "深圳"}
	fmt.Println(c)
	fmt.Println(c2)

	//-------------------------------------------
	//可以让编译器根据初始值的个数自行推断数组的长度
	var d = [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(d)
	fmt.Printf("c:%T d:%T\n", c, d)

	//-------------------------------------------
	//根据索引值来初始化
	var e [20]int			//声明
	e = [20]int{19: 1}		//初始化
	fmt.Println(e)


	//数组的基本使用
	fmt.Println(e[19])	//不支持-1

	//遍历数组的方式一
	for i := 0;i<len(a);i++{
		fmt.Println(a[i])
	}

	for k,v:=range c{
		fmt.Println(k)
		fmt.Println(v)
	}


	//求数组所有元素的和 arr = [2,4,6,8]
	arr := [...]int{2,4,6,8}
	sum := 0
	//for i :=0; i< len(arr);i++{
	//	sum +=arr[i]
	//}
	//fmt.Println(sum)

	for _,v := range arr{
		sum +=v
	}
	fmt.Println(sum)

	//求数组【1，3，5，7，8]中两个元素的和为8的两个元素的下标
	arr2 := [...]int{1,3,5,7,8}
	for k1,v1 :=range arr2{
		other:=8-v1
		for j :=k1+1;j < len(arr2);j++{
			if other == arr2[j]{
				fmt.Printf("(%d,%d)\n",k1,j)
			}
		}
	}

	//调用冒泡排序
	arr3 :=[...]int{9,1,5,6,10,8,3,7,2,4}
	//数组作为函数参数为值传递，形参不可以修改实参的值(相当于拷贝了一个新的数组传递)，所以需要一个返回值，或者通过地址传递
	BubbleSort(&arr3)
	fmt.Println(arr3)

}