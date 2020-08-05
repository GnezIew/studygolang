package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Name string `json:"name"`
}

func JsonHandler(w http.ResponseWriter,r *http.Request){
	data := make([]byte,r.ContentLength)
	r.Body.Read(data)
	defer r.Body.Close()
	//json反序列化
	user := new(User)
	json.Unmarshal(data,user)
	fmt.Printf("%#v\n",user)
	fmt.Println(user.Name)
	fmt.Fprintf(w,string(data))

}



func main() {
	//创建多路复用器
	mux := http.NewServeMux()
	mux.HandleFunc("/json",JsonHandler)
	server := &http.Server{
		Addr: ":8080",
		Handler: mux,
	}

	if err := server.ListenAndServe();err != nil{
		log.Fatal(err)
	}
}
