package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// 参数路由
func main() {

	router := mux.NewRouter()
	// 路由参数类型未限制
	router.HandleFunc("/product/{id}", ProductHandler)
	// 使用正则表达式路由参数类型 限制为数字
	router.HandleFunc("/book/{bid:[0-9]+}", BookHandler)
	http.ListenAndServe(":8080", router)
}

// 商品/编号
func ProductHandler(w http.ResponseWriter, r *http.Request) {
	pid := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Product ID: %v\n", pid)
}

// http://localhost:8080/product/123
// 网页输出：Product ID: map[id:123]

func BookHandler(w http.ResponseWriter, r *http.Request) {
	bid := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Book ID: %v\n", bid)
}

// http://localhost:8080/book/100101
// Book ID: map[bid:100101]

// http://localhost:8080/book/wiwi
// 404 page not found
