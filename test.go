package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       // parse arguments, you have to call this by yourself
	fmt.Println(r.Form) // print form information in server side
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	s := make([]string, 3)
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
		s = append(s, strings.Join(v, ""))
	}
	for v := range s {
		fmt.Fprintf(w, strconv.Itoa(v)) // send data to client side

	}

}

func main() {
	http.HandleFunc("/dave", sayhelloName)   // set router
	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
