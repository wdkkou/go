package main

import (
	"fmt"
	"net/http"
)

// MyHandler is struct
type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "handle!")
}

// HelloHandler is struct
type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

// Wdk is struct
type Wdk struct{}

func (h *Wdk) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "wdkkou!")
}

func main() {
	handler := MyHandler{} // handlerはハンドラ。
	hello := HelloHandler{}
	wdkkou := Wdk{}
	server := http.Server{
		Addr: "127.0.0.1:8080",
		//		Addr:    ":8080", // これでもOK
	}

	http.Handle("/", &handler)
	http.Handle("/hello", &hello)
	http.Handle("/wdkkou", &wdkkou)
	server.ListenAndServe()
}
