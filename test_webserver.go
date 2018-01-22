package main

import (
	"io"
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w,
     "Hi, This is an example of http service in golang!")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, r.URL.Path)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/", echoHandler)

   // http.HandleFunc("/", handler)
   // http.ListenAndServe(":8080", nil)
    http.ListenAndServe(":8080", mux)
    //http.ListenAndServeTLS(":8080", "server.crt","server.key", nil)
}
