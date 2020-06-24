package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)

	port := ":8000"

	//http.Handle("/", router)
	fmt.Println("Starting Server on Port", port)
	defer fmt.Println("Server shut-down!")
	http.ListenAndServe(port, router)
}

// HomeHandler Handles API call to root (/)
func HomeHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Println("hello home handler")

	if req.Method == http.MethodGet {
		res.Write([]byte("hello get world"))
	} else if req.Method == http.MethodPost {
		res.Write([]byte("hello post world"))
	}
}
