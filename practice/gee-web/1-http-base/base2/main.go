package main

import (
	"fmt"
	"log"
	"net/http"
)

type Engine struct{}

// 这里将结构体指针定为函数方法的接收器，在函数方法中的操作可以改变Engine的值，参考https://www.runoob.com/go/go-method.html
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}

func main() {
	engine := new(Engine)
	log.Fatal(http.ListenAndServe(":7999", engine))
}
