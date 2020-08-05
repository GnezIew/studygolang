package main

//字典结构
type Dictionary struct {
	data map[interface{}]interface{}
}

//根据键获取值
func (d *Dictionary) Get(key interface{}) interface{}{
	return d.data[key]
}

//设置键值
func (d *Dictionary) Set(key interface{},value interface{}) {
	d.data[key] = value
}

//遍历所有的键值，如果回调返回值为false，停止遍历
func (d *Dictionary) Visit(callback func(k,v interface{}) bool) {
	if callback == nil {
		return
	}
	for k, v :=range d.data{
		if !callback(k,v){
			return
		}
	}
}


//清空所有的数据
func (d *Dictionary) Clear(){
	d.data = make(map[interface{}]interface{})
}

//创建一个字典
func NewDictionary() *Dictionary{
	d := &Dictionary{}

	//初始化map
	d.Clear()
	return d
}