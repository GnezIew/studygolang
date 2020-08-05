package main

import (
	"fmt"
	"log"
	"net/http"
)

//
var userData  = []map[string]string{
	{"name":"admin","password":"111"},
	{"name":"lianshi","password":"123456"},
}

//设置cookie
func setCookie(name,password string,w http.ResponseWriter){
	//设置cookie
	userName := &http.Cookie{
		Name: "name",
		Value: "lianshi",
		HttpOnly: true,
	}
	userPass := &http.Cookie{
		Name: "password",
		Value: "123456",
		HttpOnly: true,
	}
	http.SetCookie(w,userName)
	http.SetCookie(w,userPass)
}

func LoginHandler(w http.ResponseWriter,r *http.Request){
	if r.Method == "GET"{
		fmt.Fprintln(w,`
<html>
	<head>
		<title>Go Web</title>
	</head>
	<body>
		<form method="post" action="/login">
			<label >登录:</label>
			<input type="text"  name="loginName"/>
			<input type="text"  name="password"/>
			<button type="submit">提交</button>
		</form>
	</body>
</html>

`)
	}else{
		r.ParseForm()
		if r.FormValue("loginName") == "lianshi"{
			if r.FormValue("password") == "123456"{
				setCookie("lianshi","123456",w)
				w.Header().Set("Location","http://127.0.0.1:8080/index")
				w.WriteHeader(http.StatusFound)
			}else{
				fmt.Fprintln(w,"密码错误")
			}
		}else{
			fmt.Fprintln(w,"账号错误")
		}
	}

}

func indeHandler(w http.ResponseWriter,r *http.Request){
	//判断cookie
	name,err := r.Cookie("name")
	password,err2 := r.Cookie("password")
	if err != nil || err2 != nil {
		w.Header().Set("Location","http://127.0.0.1:8080/login")
		w.WriteHeader(http.StatusFound)
		return
	}
	if name.Value == "lianshi" && password.Value == "123456"{
		fmt.Fprintln(w,"欢迎",name.Value,"同学")
	}


}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/index",indeHandler)
	mux.HandleFunc("/login",LoginHandler)
	server := &http.Server{
		Addr: ":8080",
		Handler: mux,
	}
	if err := server.ListenAndServe();err != nil{
		log.Fatal(err)
	}
}
