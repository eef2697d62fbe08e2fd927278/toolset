package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/youngtrashbag/toolset/cmd/api/note"
	"github.com/youngtrashbag/toolset/cmd/api/tag"
	"github.com/youngtrashbag/toolset/cmd/api/user"
	"github.com/youngtrashbag/toolset/pkg/utils"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", Handler)

	apiRouter := router.PathPrefix("/api").Subrouter()

	/* user api handlers */
	apiRouter.HandleFunc("/user", user.APIHandleCreate)
	apiRouter.HandleFunc("/user/{id}", user.HandleByID)
	//apiRouter.HandleFunc("/user/{username}", user.Handle)	//TODO: do i want to add this, becuase it would disrupt api design

	/* note api handlers */
	apiRouter.HandleFunc("/note/{id}", note.HandleByID)
	apiRouter.HandleFunc("/note/{id}/tags", note.HandleTags)

	/* tag api handlers */
	apiRouter.HandleFunc("/tag/{id}", tag.HandleByID)
	apiRouter.HandleFunc("/tag/{id}/notes", tag.HandleNotes)

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

	utils.LogRequest(req)

	res.WriteHeader(statusCode)
	res.Write([]byte(message))
}
