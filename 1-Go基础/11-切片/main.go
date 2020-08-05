package main

import "fmt"

//定义一个删除切片元素函数
func remove(Slc []string,loc int) []string{
	Slc = append(Slc[:loc],Slc[loc+1:]...) //a[2：]...自动将a[2:]中元素一个个提取出来
	return Slc
}


func main() {
	//切片(slice)是一个拥有相同类型元素的可变长度的序列，他是基于数组类型做的一层封装，
	//非常灵活，支持自动扩容
	//切片是引用类型的
	//切片由三部分组成:地址，大小，容量

	//var a = []int{1,2,3}		//切片
	//var b  = [3]int{4,5,6}		//数组
	//fmt.Println(a,b)
	//fmt.Printf("a:%T b:%T\n",a,b)
	//
	//
	////声明切片的方式2：从数组得到切片
	////通过引用的方式
	//var c []int
	//c = b[:]
	//b[0] =100
	//fmt.Println(c)
	//c[0] = 50
	//fmt.Println(b)
	////
	//////这样通过直接赋值的方式，得到的d的类型也是切片
	//d := a[0:2]				//(也是引用a[x1:x2]中x1的地址)
	//a[0] = 100
	//fmt.Printf("d:%v\n,%p\n",d,d)
	//fmt.Printf("a:%v\n,%p\n",a,a)
	//fmt.Println(d)
	////-------------------------------
	//
	////切片的大小(目前元素的数量)
	//fmt.Println(len(a))
	////容量(底层数组最大能放多少)
	//x := [...]string{"北京","上海","深圳","广州","成都","杭州","重庆"}
	//y := x[1:4]		//这里y记录的是“上海”的地址，所以从“上海”开始往后算容量
	//fmt.Println(y)
	//fmt.Printf("切片y的长度是:%d\n",len(y))	//切片的大小
	//fmt.Printf("切片y的容量是:%d\n",cap(y))	//切片的容量


	//append()内置方法
	//var a  = []int{}			//每次都是以2的倍数进行扩容(重新申请地址空间),如果是一次性追加很多元素，且超过他最大容量>1，则容量扩充为总长度加1
	//fmt.Printf("a:%v 长度:%d 容量:%d 地址:%p\n",a,len(a),cap(a),a)
	//a = append(a,1)
	//fmt.Printf("a:%v 长度:%d 容量:%d 地址:%p\n",a,len(a),cap(a),a)
	//a = append(a,2)
	//fmt.Printf("a:%v 长度:%d 容量:%d 地址:%p\n",a,len(a),cap(a),a)
	//a = append(a,3)
	//fmt.Printf("a:%v 长度:%d 容量:%d 地址:%p\n",a,len(a),cap(a),a)
	//a = append(a,4)
	//fmt.Printf("a:%v 长度:%d 容量:%d 地址:%p\n",a,len(a),cap(a),a)
	//a = append(a,5,6,7,8,9)
	//fmt.Printf("a:%v 长度:%d 容量:%d 地址:%p\n",a,len(a),cap(a),a)
	//a = append(a,10)
	//fmt.Printf("a:%v 长度:%d 容量:%d 地址:%p\n",a,len(a),cap(a),a)
	//a = append(a,11)
	//fmt.Printf("a:%v 长度:%d 容量:%d 地址:%p\n",a,len(a),cap(a),a)


	//copy()复制切片
	//a := []int{1,2,3}
	//b := a
	////var c []int  //仅仅只是声明了，还没申请内存空间
	////c =make([]int , 3,3)	//申请了为int类型，长度为3，容量为3
	////上面两行等价于：
	//var c = []int{0,0,0}		//这里之所以不用var c = []int{}，是因为这里容量为0拷贝不进去
	//copy(c,a)		//深拷贝，返回值为拷贝元素个数
	//b[0] =100
	//fmt.Printf("a:%v,%p\n",a,a)		//不难发现a,b指向的是同一片地址空间(切片是引用类型)
	//fmt.Printf("b:%v,%p\n",b,b)
	//fmt.Printf("c:%v,%p\n",c,c)


	//切片中元素删除(没有高级的用法)
	//a := []string{"北京","上海","深圳","广州"}
	//a = append(a[:1],a[2:]...) //a[2：]...自动将a[2:]中元素一个个提取出来
	//fmt.Println(a)

	a := []string{"北京","上海","深圳","广州"}
	a = remove(a,3)
	fmt.Println(a)

	//练习
	//定义一个数组
	//a := [...]int{1,3,5,7,9,11,13}	//数组
	////基于数组得到一个切片
	//b := a[:]			//切片
	////修改切片中的第一个元素为100
	//b[0] = 100
	////打印数组中第一个元素的值
	//fmt.Println(a[0])
	//
	//c := a[2:5]
	//fmt.Printf("c:%v,长度：%d ,容量:%d,c地址空间:%p\n",c,len(c),cap(c),c)
	//
	//d := c[:5]		//由于切片c的长度是不够的所以需要引用底层数组，又因为c的底层数组为[5,7,9,11,13],所以可以切c[:5]
	//fmt.Println(cap(d))
	//fmt.Printf("d的地址:%p\n",d)
	//
	//d = append(d,200,300)	//由于添加两个元素后，超过了最大容量，所以会重新开辟一个内存空间，地址会发生变化
	//fmt.Println(cap(d))
	//fmt.Printf("d的地址:%p\n",d)
	//
	////由于上面d重新开辟了空间，所以d和c之间没有引用关系，所以d发生的改变不会影响到c
	//d[0] = 1000
	//fmt.Println(d[0])
	//fmt.Println(c[0])



	//a := [...]int{1,2,3}		//值类型
	//b := a[:]					//引用类型
	//fmt.Printf("%p\n",&a)		//默认取第一个元素的地址
	//fmt.Printf("%p\n",b)
	//
	//fmt.Printf("b的容量:%d,b的地址:%p\n",cap(b),b)
	//b = append(b, 4,5,6,7,8,9)
	//fmt.Printf("b的容量:%d,b的地址:%p\n",cap(b),b)


	//nil值的切片的长度和容量，相当于只声明了的切片(没有地址空间的)
	//var a []int		//仅仅声明
	//var b = []int{}		//声明并且初始化
	//c := make([]int,0)
	//if a == nil {
	//	fmt.Println("a是一个nil")
	//}
	//fmt.Println(b)
	//if b == nil{
	//	fmt.Println("b是一个nil")
	//}else{
	//	fmt.Println("b不是一个nil")
	//}
	//
	//if c == nil{
	//	fmt.Println("c是一个nil")
	//}else{
	//	fmt.Println("c不是一个nil")
	//}


	//------------------------
	//切片的遍历
	//a := []int{1,2,3,4,5}
	//
	//for i:=0;i<len(a);i++{
	//	fmt.Println(a[i])
	//}
	//
	//for index,value :=range a{
	//	fmt.Println(index,value)
	//}

	//----------------------------
	//练习题
	//a := make([]string,5,10)//开辟了空间，因为长度为5那么会有5个” “([     ])
	//for i:=0;i<10;i++{
	//	a = append(a, fmt.Sprintf("%v",i))
	//}
	//fmt.Println(a)

	//使用内置的sort包对数组var a = [...]int{3,7,8,9,1]进行排序
	//var a = [...]int{3,7,8,9,1}
	//b := a[:]
	//b[0] =100
	//sort.Ints(b)		//sort()里面接收的是切片,因为这个切片的底层就是数组a,所以切片排序，数组也会被排序
	//fmt.Println(a)
	//fmt.Println(b)



	/*
		切片的截取与拷贝
		1.切片截取是在源切片基础上进行操作的 修改一个会影响另外一个
	*/
	s1 := []int{9,1,5,6,10,8,3,7,2,4}
	//s2 := s1[起始位置,结束位置,容量] //切片的截取是左闭右开的
	s2 := s1[2:5]
	fmt.Println(s2)
	//修改切片s2
	s2[2] = 123
	fmt.Println(s2)
	fmt.Println(s1)
	/*
		s1对应的数据集合中第一个数据的地址 	例：0xc000064050
		s2对应的数据集合中第三个数据的地址	例：0xc000064060
		因为地址是十六进制见十六进一
		所以 第一个数据地址为 ：0xc000064050
			第二个数据地址为 ：0xc000064058
			第三个数据地址为 ：0xc000064060
	*/
	fmt.Printf("s1中指向数据的地址: %p\n",s1)
	fmt.Printf("s2中指向数据的地址: %p\n",s2)

	//切片的拷贝
	s3 := make([]int ,10,20)
	copy(s3,s1)
	s3[2] = 666
	fmt.Println("切片s1中的数据:",s1)
	fmt.Println("切片s3中的数据:",s3)
	fmt.Printf("s1中指向数据的地址: %p\n",s1)
	fmt.Printf("s3中指向数据的地址: %p\n",s3)
}
