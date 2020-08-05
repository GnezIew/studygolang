package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func fileHandler(w http.ResponseWriter,r *http.Request){
	fmt.Fprint(w,`
<html>
	<head>
		<title>Go Web</title>
	</head>
	<body>
		<form method="post" action="/multipartform?lang=cpp&name=dj"
enctype="multipart/form-data">
			<label >MultipartForm:</label>
			<input type="text"  name="lang"/>
			<input type="text"  name="age"/>
			<input type="file"  name="uploaded"/>
			<button type="submit">提交</button>
		</form>
	</body>
</html>
`)
}

func multipartFormHandler(w http.ResponseWriter,r *http.Request){
	r.ParseMultipartForm(4096)
	fmt.Fprintln(w,r.MultipartForm)

	fileHeader := r.MultipartForm.File["uploaded"][0]
	file,err := fileHeader.Open()
	if err != nil {
		fmt.Println("Open failed",err)
		return
	}
	data,err := ioutil.ReadAll(file)
	if err == nil{
		ioutil.WriteFile(fileHeader.Filename,data,0777)
	}
	fmt.Fprintln(w,"收到了")
}

func main() {
	//创建多路复用器
	mux := http.NewServeMux()
	mux.HandleFunc("/file",fileHandler)
	mux.HandleFunc("/multipartform",multipartFormHandler)

	server := &http.Server{
		Addr: ":8080",
		Handler: mux,
	}

	if err := server.ListenAndServe();err != nil{
		log.Fatal(err)
	}
}
