package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", IndexHandler)
	router.HandleFunc("/home", HomeHandler)
	http.ListenAndServe(":8080", router)

}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "主页面")
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "欢迎回家")
}
