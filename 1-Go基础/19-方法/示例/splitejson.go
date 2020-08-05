package main

import "encoding/json"

//定义手机屏幕
type Screen struct {
	Size       float32 //屏幕尺寸
	ResX, ResY int     //屏幕水平和垂直分辨率
}

//定义电池
type Battery struct {
	Capacity int		//容量
}

func genJsonData() 	[]byte{
	//完整数据结构
	raw := &struct {
		Screen
		Battery
		HasTouchID bool
	}{
		//屏幕参数
		Screen:Screen{
			Size: 5.5,
			ResX: 1920,
			ResY: 1820,
		},
		Battery:Battery{Capacity:2910},
		HasTouchID :true,
	}

	//将数据序列化为JSON
	jsonData,_ := json.Marshal(raw)
	return jsonData
}