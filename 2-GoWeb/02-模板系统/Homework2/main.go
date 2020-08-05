package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	//登录页面
	if r.Method == "GET" {
		t, err := template.ParseFiles("2-GoWeb/02-模板系统/Homework2/login.html")
		if err != nil {
			fmt.Println(err)
		}

		data := map[string]interface{}{
			"now": time.Now().Format("2006/01/02/ 15:04:05"),
		}
		//渲染模板
		t.Execute(w, data)
	} else {
		//解析form表单
		r.ParseForm()
		username := r.FormValue("username")
		password := r.FormValue("password")
		if username == "lianshi" && password == "123456" {
			msg := "登录成功"
			//下发cookie
			c1 := http.Cookie{
				Name:     "username",
				Value:    username,
				MaxAge:   0,
				HttpOnly: true,
			}
			http.SetCookie(w,&c1)
			fmt.Fprintln(w, msg)
		} else {
			//返回登录失败
			msg := "登录失败"
			fmt.Fprintln(w, msg)
		}
	}

}

func userInfoHandler(w http.ResponseWriter, r *http.Request) {
	//返回当前请求的用户的个人信息
	username,_ := r.Cookie("username")
	t,_ := template.ParseFiles("2-GoWeb/02-模板系统/Homework2/userInfo.html")
	t.Execute(w,username.Value)

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", loginHandler)
	mux.HandleFunc("/userinfo", userInfoHandler)
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
