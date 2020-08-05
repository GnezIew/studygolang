package main

import "fmt"

/*
	方法定义(可以理解为类中的方法)
	func (方法接受者) 方法名 (参数列表) (返回值列表) {
		代码值
		返回值
  	}
	分为指针对象 和值对象
*/
type Student struct {
	Name string
	Age int
	Score int
}

//值对象接受者
func (s Student) SayHello () {
	//fmt.Println(s.Name,":hello")
	s.Name = "为为"
}

//指针对象接受者
func (s *Student) SayHello2 () {
	//fmt.Println(s.Name,":hello")
	s.Name = "为为"
}

/*
	两者区别:值对象为接受者，修改结构体的属性，只会在方法内改变
	指针对象的话能够永久改变结构体的属性
*/

//方法的继承和重写-------------------------------------------
type People struct {
	Name string
	Age int
}
//建议使用指针对象接受者
func (p *People) GetName(){
	fmt.Println("父类方法:",p.Name)
}

type Student2 struct {
	People
	Score int
}

//同一个对象的方法名不能重名
func (s *Student2) GetName(){
	fmt.Println("子类重写父类方法:",s.Name)
}


func main() {
	stu := Student{
		Name:  "曾为",
		Age:   18,
		Score: 100,
	}
	//对象.方法()
	stu.SayHello()
	fmt.Println("值对象为接受者，方法修改结构体属性后:",stu)
	stu.SayHello2()
	fmt.Println("指针对象为接受者，方法修改结构体属性后:",stu)


	//方法的继承和重写
	//子类继承父类 子类可以使用父类结构体成员属性 可以使用父类方法

	stu2 := Student2{
		People: People{
			"董聃",
			18,
		},
		Score:  100,
	}
	//子类调用父类方法
	stu2.People.GetName()
	//子类重写父类方法
	stu2.GetName()


	//-------------------------------------------
	//实例化属性
	p := new(Property)
	//设置属性
	p.SetValue(100)
	//返回属性
	fmt.Println(p.GetValue())


	//-------二维矢量模拟玩家移动------------------------
	//实例化玩家对象，并且设置速度为0.5
	player := NewPlayer(0.5)
	//让玩家移动到3,1点
	player.MoveTo(Vec2{3,1})

	//如果没有到达就一直循环
	for !player.IsArrived() {
		//更新玩家位置
		player.Update()

		//打印每次移动后玩家的位置
		fmt.Println(player.Pos())
	}

	//-------------为类型添加方法-----------------------
	var b MyInt
	fmt.Println(b.Iszero())
	b = 1
	fmt.Println(b.Add(2))
}
