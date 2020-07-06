package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/youngtrashbag/toolset/src/note"
	"github.com/youngtrashbag/toolset/src/user"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)

	router.HandleFunc("/note", note.Handler)

	apiRouter := router.PathPrefix("/api").Subrouter()

	// user api handlers
	apiRouter.HandleFunc("/user", user.APIHandleCreate)
	apiRouter.HandleFunc("/user/{id}", user.APIHandleByID)

	//note api handlers
	//apiRouter.HandleFunc("/note", note.APIHandleCreate)
	apiRouter.HandleFunc("/note/{id}", note.APIHandleByID)

	port := ":8000"

	log.Printf("Starting Server on Port \"%s\"", port)
	defer log.Println("Server shut-down!")
	log.Fatal(http.ListenAndServe(port, router))
}

// HomeHandler : Handles API call to root (/)
func HomeHandler(res http.ResponseWriter, req *http.Request) {
	var statusCode int
	var message string

	if req.Method == http.MethodGet {
		statusCode = 200
		message = "hello get world"
	} else if req.Method == http.MethodPost {
		statusCode = 201
		message = "hello post world"
	} else if req.Method == http.MethodPut {
		statusCode = 202
		message = "hello put world"
	} else if req.Method == http.MethodDelete {
		statusCode = 204
		message = "hello delete world"
	} else {
		statusCode = 404
	}

	log.Printf("%s Method on \"%s\", StatusCode:%d, Message:\"%s\"", req.Method, req.URL.Path, statusCode, message)

	res.WriteHeader(statusCode)
	res.Write([]byte(message))
}
