package main

import (
	"fmt"
	"runtime"
	"time"
)

type P *int
type Q *int

type T []int

func F(t T) {}

//为结构体其别名
type stu student
type student struct {
	name string
	age  int
}

//常量和全局变量优于init()执行
var a = test()

func init() {
	fmt.Println(1)
}
func init() {
	fmt.Println(2)
}

func test() int {
	fmt.Println(3)
	return 3
}

func init() {
	var pcs [1]uintptr
	runtime.Callers(1, pcs[:])
	fn := runtime.FuncForPC(pcs[0])
	fmt.Println(fn.Name())
}

func init() {
	var pcs [1]uintptr
	runtime.Callers(1, pcs[:])
	fn := runtime.FuncForPC(pcs[0])
	fmt.Println(fn.Name())
}

var x int

func init() {
	x++
}

//12.
func f(a int, b uint) {
	/*
		copy(dst, src []Type) int
		返回被复制的元素数量，它会是 len(src) 和 len(dst) 中较小的那个
	*/
	val := copy(make([]struct{}, a), make([]struct{}, b))
	fmt.Printf("%d\n", val)
}

//13.
type Hash [32]byte

//func MustNotBeZero(h Hash){
//	if h == Hash{
//		fmt.Println(1)
//	}
//}

func MustNotBeZero2(h Hash) {
	if h == [32]byte{} {
		fmt.Println(1)
	}
}

//20.
//func A ( string string) string{
//	return string + string
//}
//
//func B(len int) int {
//	return len +len
//}
//
//func C(val ,default string) string{
//	if val == "" {
//		return default
//	}
//	return val
//}

/*
	命名规则:有数字，字母下划线组成，不能以数字和下划线开头,不能以25个关键字做变量名，
	但是36个预设字是可以重新定义的，但是一般不去重新定义
*/

//21.select
func A() int {
	fmt.Println("调用A函数")
	time.Sleep(100 * time.Millisecond)
	return 1
}

func B() int {
	fmt.Println("调用B函数")
	time.Sleep(1000 * time.Millisecond)
	return 2
}
//23.map的值
type student2 struct {
	Name string
	Age  int
}


//25、
//type People struct{
//
//}
//
//func (p *People) ShowA(){
//	fmt.Println("showA")
//	p.ShowB()
//}
//
//func (p *People) ShowB(){
//	fmt.Println("showB")
//}
//
//type Teacher struct {
//	People
//}
//
//func (t *Teacher) ShowB(){
//	fmt.Println("teacher showB")
//}

//26.
func calc(index string,a , b int) int {
	ret := a + b
	fmt.Println(index,a,b,ret)
	return ret
}


//30.
type People interface {
	Speak(string) string
}
type Student struct {
}
func (stu *Student)Speak(think string) (talk string){
	if think == "法师"{
		talk = "法师,我爱你哟~"
	}else {
		talk = "hi"
	}
	return talk
}

func main() {
	//1.死锁
	//t := time.NewTicker(100)
	//for range t.C{
	//	t.Stop()
	//}
	/*
		time.NewTicker 会返回一个新的Ticker ,每隔参数时间，就会向通道发送当前时间
		t.Stop()只能关闭信号 不关闭通道通信
		t.Stop()关闭了信号不再向通道发送数据了， 而t.C还在接受数据，导致通道阻塞，死锁
	*/

	//2.硬编码格式
	//var x uint32 = 100.5 + 100.5 + 2147483447.0
	//fmt.Println(x)
	///*
	//	uint32 取值范围是0~ 2^32次幂
	//	int32  取值范围是-2^31 ~ 2^31
	//*/

	//3.字符串处理
	//s := strings.TrimRight("abcdefedcba","fedabc")
	////s := strings.TrimRight("abcdefedmcba","fedabc")
	////s := strings.TrimSuffix("abcdefedmcba","abc")
	//fmt.Printf("%q\n",s)
	/*
		TrimRight(s string, cutset string) 返回右往左匹配，去除s中含cutset的字符/字符串，如果遇到不匹配，则停止扫描
		TrimSuffix(s, suffix string) string 返回去除s可能的后缀suffix的字符串。
	*/

	//4.位运算
	//a := 2 ^ 15		// 相当于 10 ^ 1111 = 1101  = 13
	//b := 4 ^ 15		// 相当于 100 ^ 1111 = 1011 = 11
	////fmt.Println(a &^ b)
	//if a > b{
	//	fmt.Println("a")
	//}else{
	//	fmt.Println("b")
	//}
	/*
		&(AND)		相同位都1时结果才为1
		|(OR) 		相同位只要一个为1 即为1
		^(XOR)		相同位相同为0 不同为1
		&^(AND NOT) 相同位后数为0 ，则用前数对应位替代， 后数为1则取0
	*/

	// 5.切片
	//v := []int{1,2,3}
	//fmt.Printf("%p\n",v)
	//for i,_ :=range v {
	//	v = append(v,i)
	//	fmt.Printf("%p\n",v)
	//}
	//fmt.Println(v)
	/*
		并不会出现死循环
		range中的v 相当于对上面的v的拷贝
	*/

	//6.字符类型
	//r:= []rune("☃")
	//fmt.Println(r)
	//r:= []rune("你好")
	//for _,v := range r{
	//	fmt.Printf("%c\n",v)
	//}
	/*
		rune 数据类型 ，支持中文编码集
		是否报错
	*/

	//7.编码
	//可以支持图片，是因为他属于中文编码的一种
	//fmt.Println("Seattle, WA☔ ")

	//8.类型别名
	//var p P = new(int)
	//*p +=8
	////当操作为引用类型 别名和原类型是可以进行赋值
	//var x *int = p
	//fmt.Printf("%T\n",p)
	//fmt.Printf("%T\n",x)
	//var q Q = x
	//fmt.Printf("%T\n",q)
	//
	//*q++
	//fmt.Println(*p,*q)
	/*
		var  p P = new(int)
		var  q Q = p
		不能这样会报错类型不一致，p,q不同类型

		当操作为引用类型 别名和原类型是可以进行赋值
		如果p为 int类型(或其他非引用类型) 是不能赋值的
	*/

	//9.map的值操作
	//m := make(map[string]int)
	//m["foo"]++
	//fmt.Println(m["foo"])
	/*
		m["foo"]++ 相当于 为map添加元素，默认值为零值
		即 m["foo"] = 0  m["foo"] +=1
	*/
	//拓展
	//"asdadwqfqvefgbrhsf"统计字符出现的个数
	//str := "asdadwqfqvefgbrhsf"
	//m2 := make(map[string]int)
	//for _,v := range str{
	//	m2[string(v)]+=1
	//}
	//fmt.Println(m2)

	//拓展2
	//mStu := make(map[int]student)
	//mStu[1001] = student{
	//	name: "王宏达",
	//	age:  18,
	//}
	//fmt.Println(mStu[1001].name)
	//mStu[1001].name = "老四"
	/*
			map中的value类型如果是结构体的话，是只能读取不能修改
			原因在于map中的value不支持寻址的操作
			也就是时候value是值类型的，相当于拷贝了，如果想有写操作，
			那么在声明的时候，可以创建指针类型的结构体
		 	或者整体的修改
			mStu[1001] = student{
				name: "老曾",
				age:  18,
			}
	*/
	//mStu2 := make(map[int]*student)
	//mStu2[1001] = &student{
	//	name: "王宏达",
	//	age:  18,
	//}
	//mStu2[1001].name = "老曾"
	//fmt.Println(mStu2[1001].name)

	//10.函数名
	/*
		go语言中函数名不允许重名
	*/

	//11.函数名init
	/*
		一个包下不可以有两个相同的函数名，init()除外
	*/

	//12.函数名init()扩展
	//主函数中不能调用init(),一个init函数只能执行一次，
	//fmt.Println(x)

	//13.切片的拷贝操作
	//f(90,314)
	/*
		Sizeof返回类型v本身数据所占用的字节数。
		返回值是“顶层”的数据占有的字节数。
		例如，若v是一个切片，它会返回该切片描述符的大小，而非该切片底层引用的内存的大小。
	*/
	//slice := make([]struct{},10)
	//fmt.Println(unsafe.Sizeof(slice))		//一个占8字节(64为操作系统)

	//14.数组的比较运算
	//MustNotBeZero(Hash{})	//不能用类型别名进行比较操作
	//MustNotBeZero2(Hash{})

	//18.协程
	//go func() {
	//	panic("Boom")
	//}()
	//for i:=0;i< 10000;i++{
	//	fmt.Print(i,",")
	//}
	//fmt.Println("Done")

	//go func() {
	//	defer func() {
	//		err := recover()
	//		fmt.Println(err)
	//	}()
	//	panic("Boom")
	//}()
	//for i := 0; i < 10000; i++ {
	//	fmt.Print(i, ",")
	//}
	//fmt.Println("Done")

	//19.匿名类型
	//var x []int
	//var y T
	//y = x
	//F(x)
	//_ = y
	///*
	//	x 和 y 可以直接赋值吗
	//	是可以的，切片是引用类型
	//*/
	//
	//var s = student{
	//	name: "z",
	//	age:  0,
	//}
	////这里不能赋值的原因是结构体不是引用类型
	////var s1 stu
	////s1 = s
	//fmt.Println(s)

	//20. 数据类型

	//21.select
	//ch := make(chan int, 1)
	//go func() {
	//	select {
	//	case ch <- A():
	//	case ch <- B():
	//	default:
	//		ch <- 3
	//	}
	//}()
	//fmt.Println(<-ch)
	/*
		结果是：1或2	(函数先执行在进行选择)
		select 多路复用 case分支满足了，就不会执行default
		如果两个case都满足条件，是伪随机选择一个执行的
	*/

	//22.defer延迟调用(入栈)
	//defer func() {
	//	fmt.Println("1")
	//}()
	//defer func() {
	//	fmt.Println("2")
	//}()
	//defer func() {
	//	fmt.Println("3")
	//}()
	//panic("触发异常")
	///*
	//	panic 函数调用时，都需要将程序所有的内容进行处理后,再进行错误提示(清空栈区等等)，defer会被调用
	//*/

	//23.map的值
	//m := make(map[string]*student2)
	//stus := []student2{
	//	{Name:"zhou",Age:24},
	//	{Name:"li",Age:23},
	//	{Name:"wang",Age:22},
	//}
	//for _,stu2 := range stus{	//
	//	m[stu2.Name] = &stu2
	//}
	//fmt.Println(m)
	/*
		1.map是无序的
		2.for循环中的vaule的地址是固定的
		上面的不足指出在于得到的map中的值都是同一个
		改进版本:
		m := make(map[string]*student2)
		stus := []*student2{
			{Name:"zhou",Age:24},
			{Name:"li",Age:23},
			{Name:"wang",Age:22},
		}
		for _,stu2 := range stus{	//

			m[stu2.Name] = stu2
		}
		fmt.Println(m)
		for _,v := range m{
			fmt.Println(v)
		}
	*/

	//24.等待组
	//runtime.GOMAXPROCS(1)
	//wg := sync.WaitGroup{}
	//wg.Add(10)
	//for i:=0;i<10 ;i++  {
	//	go func(i int) {
	//		fmt.Println(i)
	//		wg.Done()
	//	}(i)
	//}
	//wg.Wait()
	//fmt.Println("等待结束")
	/*
		结果：单核且不超过256,则先输出最后一个，然后从头顺序输出
		多核则是无序输出
	*/

	//25.方法继承
	//t := Teacher{}
	//t.ShowA()
	/*
		结果为：showA showB，
	向下继承，子类能够找到父类，父类不能找到子类
	*/

	//26.defer
	//a1 := 1
	//b1 := 2
	//defer calc("1",a1,calc("3",a1,b1))
	//a1 = 0
	//defer calc("2",a1,calc("4",a1,b1))
	//b1 = 1

	//27.切片的截取
	//a1 := []int{0,1,2,3,4,5,6,7}
	//b1 := a1[:3]
	//fmt.Println(len(b1))
	//fmt.Println(cap(b1))
	//b1 = append(b1,4)
	//fmt.Println(a1)
	//fmt.Println(b1)
	/*
		b1的长度为切片截取的长度
		b1的容量 = 原切片的容量 - 截取起始位置的下标
		如果添加元素未超过容量大小,那么b1和a1指向的是同一片内存空间，b1的修改同时会改变a1
		如果添加元素超过容量大小,那么b1和a1 指向的就是不同的内存空间，b1的修改不会影响a1
	*/

	//28.数据类型
	//i := GetValue()
	////类型断言 ， 基于接口类型
	////接口变量.(type) 结果为数据类型
	//switch i.(type) {
	//case int:
	//	fmt.Println("int")
	//case string:
	//	fmt.Println("string")
	//case interface{}:
	//	fmt.Println("interface")
	//default:
	//	fmt.Println("interface")
	//}

	//29.结构体比较
	//sn1 := struct {
	//	age int
	//	name string
	//}{age:11,name:"qq"}
	//sn2 := struct {
	//	age int
	//	name string
	//}{age:11,name:"qq"}
	//if sn1 == sn2 {
	//	fmt.Println("sn1 == sn2")
	//}else{
	//	fmt.Println(sn1 != sn2)
	//}
	/*
		结构体比较：
		比较的内容 是比较的是具体的数值内容,还有成员的位置顺序
	*/

	//30.结构体比较2
	//sn1:= struct {
	//	age int
	//	m map[string]string
	//}{age:11,m: map[string]string{"a":"1"}}
	//sn2 := struct {
	//	age int
	//	m map[string]string
	//}{age:11,m: map[string]string{"a":"1"}}
	///*
	//	不能比较
	//	原因:map,slice,函数 不支持比较
	//*/
	////if sn1 == sn2{
	////	fmt.Println("sn1 == sn2")
	////}
	//if sn1.m["a"] == sn2.m["a"]{
	//	fmt.Println("值相同")
	//}

	//31.方法集
	var peo People = &Student{}
	think := "法师"
	fmt.Println(peo.Speak(think))
}
//func GetValue() interface{} {
//	return 1
//}
