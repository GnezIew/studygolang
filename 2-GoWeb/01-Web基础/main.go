package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello,web!")
}

func main() {
	http.HandleFunc("/", hello) //注册一个函数hello 专门用来处理 / 这个path
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal()
	}
}
