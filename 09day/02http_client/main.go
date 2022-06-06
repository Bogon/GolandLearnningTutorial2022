package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// net/http client

// 共用一个链接，适用于请求比较频繁
var (
	client = http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: false,
		},
	}
)

func main() {

	// response, err := http.Get("http://127.0.0.1:9090/xxx/?name=sb&age=18")
	// if err != nil {
	// 	fmt.Println("http client get path failed, error: ", err)
	// 	return
	// }

	data := url.Values{}
	urlObj, _ := url.Parse("http://127.0.0.1:9090/xxx/")
	data.Set("name", "萧山")
	data.Set("age", "120")
	queryStr := data.Encode()
	fmt.Println(queryStr)
	urlObj.RawQuery = queryStr
	request, err := http.NewRequest("GET", urlObj.String(), nil)
	if err != nil {
		fmt.Println("http client get path failed, error: ", err)
		return
	}
	// 发起请求
	// response, err := http.DefaultClient.Do(request)
	// if err != nil {
	// 	fmt.Println("http client get path failed, error: ", err)
	// 	return
	// }

	// 适用于请求不是特别频繁的请求
	// 禁用 KeepAlives 的client
	// tr := &http.Transport{
	// 	DisableKeepAlives: true,
	// }
	// client := http.Client{
	// 	Transport: tr,
	// }
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("http client get path failed, error: ", err)
		return
	}
	defer response.Body.Close() // 一定要记得关闭 resp.Body

	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("read response body failed, error:", err)
		return
	}
	fmt.Println(string(b))
	fmt.Println(response.Status)
	fmt.Println(response.StatusCode)
	fmt.Println(response.Header)
	fmt.Println(ioutil.ReadAll(response.Body))
}
