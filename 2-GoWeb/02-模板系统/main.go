package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"text/template"
	"time"
)

type User struct {
	Name string
	Age  int
}

type AgeInfo struct {
	Age           int
	GreaterThan60 bool
	GreaterThan40 bool
}

func stringLiteralTemplate() {
	s := "My name is {{ .Name }}. I am {{ .Age }} years old.\n"
	//创建一个模板并解析s这个字符串
	t, err := template.New("test").Parse(s)
	if err != nil {
		log.Fatal("Parse string literal template error:", err)
	}

	u := User{
		Name: "lianshi ",
		Age:  18,
	}
	err = t.Execute(os.Stdout, u) //去执行渲染动作
	if err != nil {
		log.Fatal("Execute string literal template error:", err)
	}
}

func fileTemplate(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatal("Parse file template error:", err)
	}

	u := User{
		Name: "ls",
		Age:  18,
	}
	err = t.Execute(w, u)

	if err != nil {
		log.Fatal("Execute file template error:", err)
	}
}

//条件动作
func agePage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("agePage.html")
	if err != nil {
		log.Fatal("Parse error", err)
	}

	rand.Seed(time.Now().Unix())
	age := rand.Intn(100)
	info := AgeInfo{
		Age:           age,
		GreaterThan60: age > 60,
		GreaterThan40: age > 40,
	}
	err = t.Execute(w,info)
	if err != nil {
		log.Fatal("Execute error",err)
	}
}

type Item struct {
	Name string
	Price int
}
//迭代动作
func itemRange(w http.ResponseWriter,r *http.Request){
	t,err := template.ParseFiles("itemRange.html")
	if err != nil{
		log.Fatal("Parse error :",err)
	}
	items := []Item{
		{"iphone",699},
		{"ipad",799},
		{"iWatch",199},
		{"MacBook",999},
	}
	err = t.Execute(w,items)
	if err != nil{
		log.Fatal("Execute error:",err)
	}
}

func main() {
	//stringLiteralTemplate()
	//
	//fileTemplate()

	mux := http.NewServeMux()
	mux.HandleFunc("/index", fileTemplate)
	mux.HandleFunc("/agePage",agePage)
	mux.HandleFunc("/itemRange",itemRange)
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
