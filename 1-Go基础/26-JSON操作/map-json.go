package main

import (
	"encoding/json"
	"fmt"
)

func map_json(){
	m := make(map[string]interface{})
	m["subject"] = "go语言开发"
	m["students"] = []string{"王宏达","老四","汉斯","安逸君","法师"}
	m["price "] = 8988.88

	//slice,err:= json.Marshal(m)
	//结果：{"price ":8988.88,"students":["王宏达","老四","汉斯","安逸君","法师"],"subject":"go语言开发"}
	slice,err := json.MarshalIndent(m,"","   ")
	/*
		结果:
			{
	   "price ": 8988.88,
	   "students": [
	      "王宏达",
	      "老四",
	      "汉斯",
	      "安逸君",
	      "法师"
	   ],
	   "subject": "go语言开发"
	}
	*/
	if err != nil{
		fmt.Println("json err: ",err)
	}else{
		fmt.Println(string(slice))
	}
}

//json解析成map
func json_map(){
	slice := []byte(`{"price ":8988.88,"students":["王宏达","老四","汉斯","安逸君","法师"],"subject":"go语言开发"}`)
	//创建一个map
	m := make(map[string]interface{})
	//var temp interface{}
	err := json.Unmarshal(slice,&m)
	if err != nil{
		fmt.Println("json err:",err)
	}
	fmt.Println(m)
	//fmt.Printf("%T\n",temp)

	//提取数据
	//map的value为interface 需要进行类型断言 获取数据内容
	for k,v :=range m{
		switch tv := v.(type) {
		//case interface{}:
		//	fmt.Println(k,v)
		case string:
			fmt.Println("string类型数据：",k,tv)
		case float64:
			fmt.Println("float64类型数据",k,tv)
		//case []string:			//err
		//	fmt.Println("[]string类型数据:",k,v)
		case []interface{}:
			fmt.Println("复合类型数据:",k,tv)
			for i, val := range tv {
				fmt.Println(i,val)
			}
		}
	}
}