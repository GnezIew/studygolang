package main

import "fmt"

//抽象的银行业务员
type AbstractBanker interface {
	DoBusi()	//抽象的处理业务的接口
}

//存款的业务员
type SaveBanker struct {}

func (sb *SaveBanker) DoBusi(){
	fmt.Println("进行了存款的操作")
}

//转账的业务员
type TransferBanker struct {

}
func (tb *TransferBanker) DoBusi(){
	fmt.Println("进行了转账")
}

//支付的业务员
type PayBanker struct {

}
func (pb *PayBanker) DoBusi(){
	fmt.Println("进行了存款")
}

//实现架构层(基于抽象层进行业务封装-针对interface接口进行封装)
func BankerBusiness(banker AbstractBanker){
	//通过接口向下调用(多态)
	banker.DoBusi()
}

func main() {
	/*
		开闭原则：
			一个软件实体如类、模块和函数应该对扩展开放，对修改关闭
			简单来说就是子啊修改需求的时候，应该尽量通过扩展来实现
			而不是通过修改代码来实现变化
	*/


	//实现业务调用
	//进行存款
	BankerBusiness(&SaveBanker{})

	//进行转账
	BankerBusiness(&TransferBanker{})

	//进行支付
	BankerBusiness(&PayBanker{})
}
