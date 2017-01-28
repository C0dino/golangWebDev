package main

import (
	"io"
	"net/http"
)

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

func me(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Stephen David Pendino")
}

func i(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Index")
}

func main() {
	http.Handle("/dog", http.HandlerFunc(d))
	http.Handle("/cat", http.HandlerFunc(c))
	http.Handle("/me", http.HandlerFunc(me))
	http.Handle("/", http.HandlerFunc(i))
	http.ListenAndServe(":8080", nil)
}
