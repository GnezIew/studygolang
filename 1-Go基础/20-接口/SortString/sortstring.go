package main

import "sort"

//将[]string 定义为MyStringList类型
type MyStringList []string

//实现sort.Interface接口的获取元素数量方法
func (m MyStringList) Len() int {
	return len(m)
}

//实现sort.Interface接口的比较元素方法
func (m MyStringList) Less(i, j int) bool {
	return m[i] < m[j]
}

//实现sort.Interface接口的交换元素方法
func (m MyStringList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

type float32Slice []float32

func (f float32Slice) Len() int {
	return len(f)
}

func (f float32Slice) Less(i, j int) bool {
	return f[i] < f[j]
}

func (f float32Slice) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func float32s (a []float32) {
	sort.Sort(float32Slice(a))
}
