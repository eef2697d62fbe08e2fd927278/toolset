package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/youngtrashbag/toolset/src/note"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/note", note.Handler)

	port := ":8000"

	//http.Handle("/", router)
	fmt.Println("Starting Server on Port", port)
	defer fmt.Println("Server shut-down!")
	http.ListenAndServe(port, router)
}

// HomeHandler : Handles API call to root (/)
func HomeHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Println("hello home handler")

	if req.Method == http.MethodGet {
		res.Write([]byte("hello get world"))
	} else if req.Method == http.MethodPost {
		res.Write([]byte("hello post world"))
	}
}
