package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("provisionally")

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
}

// HomeHandler Handles API call to root (/)
func HomeHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Println("hello home handler")
}
