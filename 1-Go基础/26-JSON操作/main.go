package main

func main() {
	//结构体转成json //结构体的成员一定要首字母大写才能进行转换
	//cl := Class{
	//	Subject: "go语言开发",
	//	Students: []string{
	//		"王宏达",
	//		"老四",
	//		"汉斯",
	//		"安逸君",
	//	},
	//	Price: 8900.88,
	//}
	//slice, err := json.Marshal(cl)		//不带格式的转换
	//if err!=nil{
	//	fmt.Println("json err:",err)
	//}
	//fmt.Println(string(slice))
	//结果:{"Subject":"go语言开发","Students":["王宏达","老四","汉斯","安逸君"],"Price":8900.88}

	//slice, err := json.MarshalIndent(cl,"","   ")		//带格式的转换
	//if err!=nil{
	//	fmt.Println("json err:",err)
	//}
	//fmt.Println(string(slice))
	/*结果:
		{
	   "Subject": "go语言开发",
	   "Students": [
	      "王宏达",
	      "老四",
	      "汉斯",
	      "安逸君"
	   ],
	   "Price": 8900.88
	}
	*/

	//json转换成结构体
	//slice1 := []byte(`{"Subject":"go语言开发","Students":["王宏达","老四","汉斯","安逸君"],"Price":8900.88}`)
	//var cl2 Class
	//err2 := json.Unmarshal(slice1,&cl2)
	//if err2 != nil {
	//	fmt.Println("json err",err2)
	//}else{
	//	fmt.Println(cl2)
	//}

	//map转换成json
	map_json()
	//json转换成map
	json_map()
	
}
