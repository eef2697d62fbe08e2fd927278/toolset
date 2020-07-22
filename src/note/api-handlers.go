package note

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Handler : supposed to handle the /note resource
func Handler(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		fmt.Println("GetRequest")
	} else if req.Method == http.MethodPost {
		fmt.Println("PostRequest")
	}
}

// TODO: this function is flawed. use code from userhandler and implement it here

// APIHandleByID : handles the requests to notes with a certain id
func APIHandleByID(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		// requesting a note
		res.Header().Set("Content-Type", "application/json")

		params := mux.Vars(req)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			log.Panicln(err.Error())
		}
		note := GetByID(int64(id))

		if note.ID == -1 {
			//note not found
			res.WriteHeader(http.StatusNotFound)
		}

		json.NewEncoder(res).Encode(note)

	} else if req.Method == http.MethodPatch {
		// updating a note
	}
}