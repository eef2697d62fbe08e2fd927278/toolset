package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/youngtrashbag/toolset/src/note"
	"github.com/youngtrashbag/toolset/src/user"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/note", note.Handler)
	router.HandleFunc("/createUser", user.CreateUser).Methods(http.MethodPost)

	// test api handlers
	router.HandleFunc("/api/note/{id}", note.ByID).Methods(http.MethodGet)

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
