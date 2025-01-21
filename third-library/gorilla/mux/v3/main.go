package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// 匹配域名
func main() {

	router := mux.NewRouter()

	// 只匹配 www.example.com
	router.HandleFunc("/", IndexHandler).Host("www.example.com")
	// http://localhost:8080/
	// 无法访问网站

	// 动态匹配子路由
	router.HandleFunc("/home", HomeHandler).Host("{subdomain:[a-z]+}.example.com")
	// http://localhost:8080/home
	// 无法访问网站

}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "主页面")
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "欢迎回家")
}
