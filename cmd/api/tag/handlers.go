package tag

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/youngtrashbag/toolset/pgk/utils"
)

type jTag struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	CreationDate string `json:"creation_date"`
}

// HandleByID : handles api requests for tags by id
func HandleByID(res http.ResponseWriter, req *http.Request) {
	for _, i := range req.Header["Accept"] {
		if i == "application/json" {
			if req.Method == http.MethodGet {

				res.Header().Set("Content-Type", "application/json")

				params := mux.Vars(req)

				id, err := strconv.Atoi(params["id"])
				if err != nil {
					log.Panicln(err.Error())
				}

				t := GetByID(int64(id))

				if t.ID == -1 {
					//user not in database
					message := "Note not found"
					res.WriteHeader(http.StatusNotFound)
					json.NewEncoder(res).Encode(utils.NewResponse(message))
					log.Printf(message)
				} else {
					var tm string
					utils.ConvertTime(&t.CreationDate, &tm)
					j := jTag{
						ID:           t.ID,
						Name:         t.Name,
						CreationDate: tm,
					}

					json.NewEncoder(res).Encode(j)
					res.WriteHeader(http.StatusOK)
				}
			} else {
				res.WriteHeader(http.StatusMethodNotAllowed)
			}
		} else {
			res.WriteHeader(http.StatusBadRequest)
		}
	}

	utils.LogRequest(req)
}

// HandleNotes : handles api requests for noteIDs for a specific tag
func HandleNotes(res http.ResponseWriter, req *http.Request) {
	for _, i := range req.Header["Accept"] {
		if i == "application/json" {
			if req.Method == http.MethodGet {

				res.Header().Set("Content-Type", "application/json")

				params := mux.Vars(req)

				id, err := strconv.Atoi(params["id"])
				if err != nil {
					log.Panicln(err.Error())
				}

				t := GetByID(int64(id))

				var tArr []int64

				if t.ID == -1 {
					//user not in database
					message := "Tag not found"
					res.WriteHeader(http.StatusNotFound)
					json.NewEncoder(res).Encode(utils.NewResponse(message))
					log.Printf(message)
				} else {
					tArr = t.GetNoteIDs()

					type tagArr struct {
						TagID int64   `json:"tag_id"`
						Notes []int64 `json:"note_id"`
					}

					j := tagArr{
						TagID: t.ID,
						Notes: tArr,
					}

					json.NewEncoder(res).Encode(j)
					res.WriteHeader(http.StatusOK)
				}
			} else {
				res.WriteHeader(http.StatusMethodNotAllowed)
			}
		} else {
			res.WriteHeader(http.StatusBadRequest)
		}
	}

	utils.LogRequest(req)
}
