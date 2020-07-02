package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/youngtrashbag/toolset/src/database"
	"github.com/youngtrashbag/toolset/src/note"
	"github.com/youngtrashbag/toolset/src/user"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/note", note.Handler)

	// test api handlers
	router.HandleFunc("/api/user", user.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/api/note/{id}", note.GetByID).Methods(http.MethodGet)

	//test convert
	t := time.Now()
	var s string

	database.ConvertTime(&t, &s)
	fmt.Println(s)

	port := ":8000"

	//http.Handle("/", router)
	fmt.Println("Starting Server on Port: ", port)
	defer fmt.Println("Server shut-down!")
	log.Fatal(http.ListenAndServe(port, router))
}

// HomeHandler : Handles API call to root (/)
func HomeHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Println("called HomeHandler")

	if req.Method == http.MethodGet {
		res.WriteHeader(200)
		res.Write([]byte("hello get world"))
	} else if req.Method == http.MethodPost {
		res.WriteHeader(201)
		res.Write([]byte("hello post world"))
	} else if req.Method == http.MethodPut {
		res.WriteHeader(202)
		res.Write([]byte("hello put world"))
	} else if req.Method == http.MethodDelete {
		res.WriteHeader(204)
		res.Write([]byte("hello delete world"))
	}
}
