package main

import (
	"fmt"
	"net/http"
)

type police int

func (m police) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Police: Sir, have you been drinking?")
	fmt.Fprintln(w, "Me: THE HANDLER INTERFACE IS SERVEHTTP, RESPONSEWRITER, POINTER TO A REQUEST!")
	fmt.Fprintln(w, "Police: ...Sir, please step out of the vehicle.")
}

func main() {
	var d police
	http.ListenAndServe(":8080", d)
}
