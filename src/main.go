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
	router.HandleFunc("/", Handler)

	/* user frontend handlers */
	//router.HandleFunc("/user", user.Handler)
	//router.HandleFunc("/user/{id}", user.HandleByID)

	/* note frontend handlers */
	//router.HandleFunc("/note", note.Handler)

	/* tag frontend handlers */

	apiRouter := router.PathPrefix("/api").Subrouter()

	/* user api handlers */
	apiRouter.HandleFunc("/user", user.APIHandleCreate)
	apiRouter.HandleFunc("/user/id/{id}", user.Handle)
	//apiRouter.HandleFunc("/user/{username}", user.Handle)	//TODO: do i want to add this, becuase it would disrupt api design

	/* note api handlers */
	//apiRouter.HandleFunc("/note", note.APIHandleCreate)
	apiRouter.HandleFunc("/note/id/{id}", note.Handle)

	/* tag api handlers */
	//apiRouter.HandleFunc("/tag", tag.APIHandleCreate)
	//apiRouter.HandleFunc("/tag/{id}", tag.APIHandleByID)

	port := ":8000"

	log.Printf("Starting Server on Port \"%s\"\n", port)
	defer log.Println("Server shut-down!")
	log.Fatal(http.ListenAndServe(port, router))
}

// Handler : Handles API call to root (/)
func Handler(res http.ResponseWriter, req *http.Request) {
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

	log.Printf("%s Method on \"%s\", StatusCode:%d, Message:\"%s\"\n", req.Method, req.URL.Path, statusCode, message)

	res.WriteHeader(statusCode)
	res.Write([]byte(message))
}
