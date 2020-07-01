package note

import (
	"encoding/json"
	"fmt"
	"net/http"

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

func ByID(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)
	
	note := GetNoteByID(int64(params["id"]))

	if note.id = -1 {
		//post not found
		res.WriteHeader(404)
	}
	json.NewEncoder(res).Encode(note)
}
