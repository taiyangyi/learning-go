package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	// 创建路由器实例
	router := mux.NewRouter()

	// 域名匹配
	router.Host("www.example.com")
	router.Host("{subdomain:[a-z]+}.example.com}")

	// 前缀匹配
	router.PathPrefix("/books/")
	// 方法匹配
	router.Methods("GET", "POST")
	// schemes 匹配
	router.Schemes("http", "https")
	// Header 头信息匹配
	router.Headers("Content-Type", "application/json")
	// 参数匹配
	router.Queries("key", "value")
	// 自定义匹配
	router.MatcherFunc(func(request *http.Request, match *mux.RouteMatch) bool {
		return request.ProtoMajor == 0
	})

	// 链式调用
	router.HandleFunc("/home", HomeHandler).
		Host("www.example.com").
		Methods("GET").
		Schemes("http")

}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "主页面")
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "欢迎回家")
}
