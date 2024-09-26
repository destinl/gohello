package main

import (
	"fmt"
	"net/http"
	"strings"
)

// func main() {
// 	http.HandleFunc("/", hello)
// 	err := http.ListenAndServe(":8080", nil)
// 	if err != nil {
// 		log.Fatal("ListenAndServe: ", err)
// 	}
// }

func hello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
