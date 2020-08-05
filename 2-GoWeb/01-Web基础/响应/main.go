package main

import (
	"fmt"
	"log"
	"net/http"
)

func IndexHandler(w http.ResponseWriter,r *http.Request){
	data := "This is Index Page"
	w.Write([]byte(data))
}

//重定向
func moveHandler(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Location","http://baidu.com")
	w.WriteHeader(http.StatusFound)
}

//设置cookie
func setCookie(w http.ResponseWriter,r *http.Request){
	c1 := &http.Cookie{
		Name: "name",
		Value: "lianshi",
		HttpOnly: true,
	}

	c2 := &http.Cookie{
		Name: "age",
		Value: "18",
		HttpOnly: true,
	}
	//将cookie写入响应头
	//w.Header().Set("Set-Cookie",c1.String())
	//w.Header().Add("Set-Cookie",c2.String())
	http.SetCookie(w,c1)
	http.SetCookie(w,c2 )
}

//取Cookie
func getCookie(w http.ResponseWriter,r *http.Request){
	//指定cookie获取
	name,err := r.Cookie("name")
	if err != nil{
		fmt.Fprintln(w,"cannot get cookie of name")
	}

	//获取所有请求携带的cookie
	cookies := r.Cookies()
	fmt.Fprintln(w,name)
	fmt.Fprintln(w,cookies)
	fmt.Fprintln(w,"取Cookie")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/",IndexHandler)
	mux.HandleFunc("/move",moveHandler)
	mux.HandleFunc("/set_cookie",setCookie)
	mux.HandleFunc("/get_cookie",getCookie)

	server := &http.Server{
		Addr: ":8080",
		Handler: mux,
	}

	if err := server.ListenAndServe();err != nil{
		log.Fatal(err)
	}
}
