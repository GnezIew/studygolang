package main

//四则运算实现

//定义一个接口
type Opter interface {
	Result() int
}

//定义父类
type Operta struct {
	n1 int
	n2 int
}

//加法子类
type Add struct {
	Operta
}

//减法子类
type Sub struct {
	Operta
}

//乘法子类
type Muli struct {
	Operta
}

//除法子类
type Div struct {
	Operta
}

// 工厂类
type factory struct {
}

//加法方法
func (a *Add) Result() int {
	return a.n1 + a.n2
}

//减法方法
func (s *Sub) Result() int {
	return s.n1 - s.n2
}

//乘法方法
func (m *Muli) Result() int {
	return m.n1 * m.n2
}

//除法方法
func (d *Div) Result() int {
	return d.n1 / d.n2
}

//多态
func Result(opter Opter) int {
	value := opter.Result()
	return value
}

//工厂方法
func (f *factory) CreateFactory(n1 int, n2 int, ch string) int {
	var o Opter
	switch ch {
	case "+":
		a := Add{Operta{n1, n2}}
		o = &a
	case "-":
		s := Sub{Operta{n1, n2}}
		o = &s
	case "*":
		m := Muli{Operta{n1, n2}}
		o = &m
	case "/":
		d := Div{Operta{n1, n2}}
		o = &d
	}
	value := Result(o)
	return value
}
