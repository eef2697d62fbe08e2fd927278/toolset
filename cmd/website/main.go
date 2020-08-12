package main

import (
	"log"
	"github.com/gorilla/mux"
	"github.com/youngtrashbag/toolset/cmd/backend/note"
	"github.com/youngtrashbag/toolset/src/tag"
	"github.com/youngtrashbag/toolset/src/user"
)

func main() {
	router := mux.NewRouter()

	/* user handlers */
	router.HandleFunc("/user", user.Handler)
	router.HandleFunc("/user/{id}", user.HandleByID)

	/* note handlers */
	router.HandleFunc("/note", note.Handler)

	/* tag handlers */

	port := ":8000"

	log.Printf("Starting Server on Port \"%s\"\n", port)
	defer log.Println("Server shut-down!")
	log.Fatal(http.ListenAndServe(port, router))
}
