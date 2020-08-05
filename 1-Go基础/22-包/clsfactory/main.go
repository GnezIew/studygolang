package main

import (
	"code.zengwei.com/studygolang/22-包/clsfactory/base"
	_ "code.zengwei.com/studygolang/22-包/clsfactory/cls1"
	_ "code.zengwei.com/studygolang/22-包/clsfactory/cls2"
)
func main() {
	//根据字符串动态创建一个Class1实例
	c1 := base.Create("Class1")		// 相当于var c1 Class1或 c1 := new(Class1)
	c1.Do()

	c2 := base.Create("Class2")
	c2.Do()
}
