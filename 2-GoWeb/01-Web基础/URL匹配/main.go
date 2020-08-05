package main

import (
	"fmt"
	"log"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the index page")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is hello Page")
}

func worldHander(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the world Page")
}

func urlHandler(w http.ResponseWriter ,r *http.Request){
	//r *http.Request 所有跟请求相关的都在这里面找
	//w http.ResponseWriter 所有跟响应相关的都在这里面找
	URL := r.URL	// net/url.URL
	//Fprintf根据format参数生成格式化的字符串并写入w。返回写入的字节数和遇到的任何错误
	fmt.Fprintf(w,"Scheme:%s\n",URL.Scheme)
	fmt.Fprintf(w,"Host:%s\n",URL.Host)
	fmt.Fprintf(w,"Path:%s\n",URL.Path)
	fmt.Fprintf(w,"RawPath:%s\n",URL.RawPath)
	fmt.Fprintf(w,"RawQuery:%s\n",URL.RawQuery)
	fmt.Fprintf(w,"Fragment:%s\n",URL.Fragment)
}

//查看header信息的函数
func headerHandler(w http.ResponseWriter,r *http.Request){
	for key,value := range r.Header{
		fmt.Fprintf(w,"%s:%v\n",key,value)
	}
}

//查看请求的内容体
func bodyHandler(w http.ResponseWriter,r *http.Request) {
	data := make([]byte,r.ContentLength)
	//将数据读取到data中
	r.Body.Read(data) 	//忽略错误提示
	defer r.Body.Close()
	//将从请求中读取的内容写入到响应体中
	fmt.Fprintf(w,string(data))
}

//表单提交
func formHandler(w http.ResponseWriter,r *http.Request){
	fmt.Fprint(w,`
<html>
	<head>
		<title>Go Web</title>
	</head>
	<body>
		<form method="post" action="/body">
			<label for="username">用户名:</label>
			<input type="text" id="username" name="username">
			<label for="email">邮箱:</label>
			<input type="text" id="email" name="email">
			<button type="submit">提交</button>
		</form>
	</body>
</html>
`)
}

//当请求form表单数据时，可以使用r.ParseForm解析请求的数据
func form2Handler(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println(r.Form["age"][0])
	fmt.Fprintln(w,"收到了")
}

func index2Handler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w,`
<html>
	<head>
		<title>Go Web</title>
	</head>
	<body>
		<form method="post" action="/form2?lang=cpp&name=ls"
enctype="application/x-www-form-urlencoded">
			<label >Form:</label>
			<input type="text"  name="lang">
			<input type="text"  name="age">
			<button type="submit">提交</button>
		</form>
	</body>
</html>
`)
}

func main() {
	mux := http.NewServeMux()
	//URL如果不是以"/"结尾，是需要进行精确匹配的
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/hello/", helloHandler)
	mux.HandleFunc("/hello/world/", worldHander)
	mux.HandleFunc("/url/",urlHandler)
	mux.HandleFunc("/header",headerHandler)
	mux.HandleFunc("/body",bodyHandler)
	mux.HandleFunc("/form",formHandler)

	mux.HandleFunc("/form2",form2Handler)
	mux.HandleFunc("/index2",index2Handler)
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
