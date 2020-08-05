package main

//定义属性结构
type Property struct {
	value int
}
//Property的方法 设置属性值
func (p *Property) SetValue (v int){
	//修改p的值
	p.value = v
}

//返回属性值
func (p *Property) GetValue () int{
	//返回属性值
	return p.value
}


//-------------为类型添加方法-----------------------
//将int定义为MyInt类型
type MyInt int

//为MyInt添加Iszero()方法
func (m MyInt) Iszero() bool{
	return m == 0
}

//为MyInt添加Add()方法
func (m MyInt) Add(other int) int {
	return other +int(m)
}


