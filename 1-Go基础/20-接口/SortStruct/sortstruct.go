package main

//声明英雄的分类
type Herokind int

//定义HeroKind常量，类似于枚举
const (
	None Herokind = iota
	Tank
	Assassin
	Mage
)

//定义英雄名单的结构
type Hero struct {
	Name string   //英雄名字
	Kind Herokind //英雄的种类
}

//将英雄指针的切片定义为Heros类型
type Heros []*Hero

//实现sort.Interface接口去元素数量方法
func (s Heros) Len() int{
	return len(s)
}


//实现sort.Interface接口比较元素方法
func (s Heros) Less (i,j int) bool {
	//如果英雄的分类不一致时，优先对分类进行排序
	if s[i].Kind != s[j].Kind{
		return s[i].Kind < s[j].Kind
	}
	//默认按英雄名字字符进行升序排序
	return s[i].Name < s[j].Name
}

//实现sort.Interface接口交换元素方法
func (s Heros) Swap(i, j int) {
	s[i] , s[j] = s[j],s[i]
}
