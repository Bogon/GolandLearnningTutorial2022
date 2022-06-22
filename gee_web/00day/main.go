package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	// 处理根节点请求
	http.HandleFunc("/", indexHandle)
	// 处理 hello
	http.HandleFunc("/hello", helloHandle)

	// 开始请求监听
	log.Fatal(http.ListenAndServe(":9999", nil))
}

// indexHandle
func indexHandle(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}

// helloHandle
func helloHandle(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}
