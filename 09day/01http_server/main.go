package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// net/http server

func f1(w http.ResponseWriter, r *http.Request) {
	n, err := ioutil.ReadFile("./xx.txt")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%v", err)))
	}
	w.Write(n)
}

func f2(w http.ResponseWriter, r *http.Request) {
	// 对于 get 请求，参数都放在 URL上，请求体中没有数据
	fmt.Println(r.URL.Query()) // 自动识别URL中的参数
	paramas := r.URL.Query()
	name := paramas.Get("name")
	age := paramas.Get("age")
	fmt.Println(name, age)
	fmt.Println(r.Method)
	fmt.Println(ioutil.ReadAll(r.Body)) // 在服务端打印客户端发来的请求body
	fmt.Println(r.Header)
	w.Write([]byte("OK"))
}

func main() {
	http.HandleFunc("/hello", f1)
	http.HandleFunc("/xxx/", f2)
	http.ListenAndServe("127.0.0.1:9090", nil)
}
