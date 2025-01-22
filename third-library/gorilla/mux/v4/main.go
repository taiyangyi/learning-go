package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// 子路由匹配
func main() {

	router := mux.NewRouter()

	// 创建一个子路由器，处理 /products/ 开头的所有路径
	productRouter := router.PathPrefix("/products/").Subrouter()
	productRouter.HandleFunc("/{key}", ProductHandler).Methods("GET")

	// 创建另一个子路由器，处理 /users/ 开头的所有路径
	usersRouter := router.PathPrefix("/users/").Subrouter()
	usersRouter.HandleFunc("/login", UserLoginHandler).Methods("POST")
	usersRouter.HandleFunc("/register", UserRegisterHandler).Methods("GET")

	http.ListenAndServe(":8080", router)

}

// ProductHandler 处理产品详情页
func ProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	fmt.Fprintf(w, "Product details for key: %s", key)
}

// UserLoginHandler 用户登录处理
func UserLoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "User login page")
}

// UserRegisterHandler 用户注册处理
func UserRegisterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "User register page")
}
