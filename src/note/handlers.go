package note

import (
	"encoding/json"
	"fmt"
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

// GetByID : api request for a note of a certain id passed in url
func GetByID(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		panic(err.Error())
	}
	note := GetNoteByID(int64(id))

	if note.id == -1 {
		//note not found
		res.WriteHeader(404)
	}

	json.NewEncoder(res).Encode(note)
}
