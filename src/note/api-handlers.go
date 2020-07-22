package note

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/youngtrashbag/toolset/src/utils"
)

type jNote struct {
	ID           int64  `json:"id"`
	Title        string `json:"title"`
	Content      string `json:"content"`
	CreationDate string `json:"creation_date"`
	AuthorID     int64  `json:"author_id"`
}

// Handle : handles api requests for notes
func Handle(res http.ResponseWriter, req *http.Request) {
	for _, i := range req.Header["Accept"] {
		if i == "application/json" {
			if req.Method == http.MethodGet {

				res.Header().Set("Content-Type", "application/json")

				params := mux.Vars(req)

				id, err := strconv.Atoi(params["id"])
				if err != nil {
					log.Panicln(err.Error())
				}

				n := GetByID(int64(id))

				if n.ID != -1 {

					var t string
					utils.ConvertTime(&n.CreationDate, &t)
					j := jNote{
						ID:           n.ID,
						Title:        n.Title,
						Content:      n.Content,
						CreationDate: t,
						AuthorID:     n.AuthorID,
					}

					json.NewEncoder(res).Encode(j)
					res.WriteHeader(http.StatusOK)
				} else {
					//user not in database
					message := "Note not found"
					res.WriteHeader(http.StatusNotFound)
					json.NewEncoder(res).Encode(utils.NewResponse(message))
					log.Printf(message)
				}
			} else {
				res.WriteHeader(http.StatusMethodNotAllowed)
			}
		} else {
			res.WriteHeader(http.StatusBadRequest)
		}
	}

	utils.LogRequest(res, req)
}

// HandleTags : handles api requests for tagIDs for a specific note
func HandleTags(res http.ResponseWriter, req *http.Request) {
	for _, i := range req.Header["Accept"] {
		if i == "application/json" {
			if req.Method == http.MethodGet {

				res.Header().Set("Content-Type", "application/json")

				params := mux.Vars(req)

				id, err := strconv.Atoi(params["id"])
				if err != nil {
					log.Panicln(err.Error())
				}

				n := GetByID(int64(id))

				if n.ID != -1 {
				} else {
					//user not in database
					message := "Note not found"
					res.WriteHeader(http.StatusNotFound)
					json.NewEncoder(res).Encode(utils.NewResponse(message))
					log.Printf(message)
				}
			} else {
				res.WriteHeader(http.StatusMethodNotAllowed)
			}
		} else {
			res.WriteHeader(http.StatusBadRequest)
		}
	}

	utils.LogRequest(res, req)
}
