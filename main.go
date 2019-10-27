package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/",myMiddleware(myHandler))

	http.ListenAndServe(":8080", nil)
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("hi")
	fmt.Println(r)
	fmt.Println(w)
	w.WriteHeader(http.StatusOK)
}

func myMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("middleware!!")
		next(w,r)
	}
}