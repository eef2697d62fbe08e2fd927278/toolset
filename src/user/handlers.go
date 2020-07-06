package user

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/youngtrashbag/toolset/src/database"
)

// HandleCreate : handles the creation a user
func HandleCreate(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		for _, i := range req.Header["Content-Type"] {
			if i == "application/json" {
				u := NewUser("mail@mail.com", "username", "password")
				id := u.Insert()

				if id != -1 {
					message := "Inserted User with ID " + string(id) + " into Database\n"
					log.Println(message)
					json.NewEncoder(res).Encode(database.NewResponse(message))

					res.WriteHeader(201)
				} else {
					message := "Could not Insert User into Database"
					log.Panicln(message)
					json.NewEncoder(res).Encode(database.NewResponse(message))
					res.WriteHeader(400)
				}
			}
		}
	} else {
		res.WriteHeader(405)
	}
}

// HandleByID : handles requests for users with a specified id
func HandleByID(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		for _, i := range req.Header["Accept"] {
			if i == "application/json" {
				res.Header().Set("Content-Type", "application/json")

				params := mux.Vars(req)
				id, err := strconv.Atoi(params["id"])
				if err != nil {
					log.Panicln(err.Error())
				}

				u := GetByID(int64(id))

				if u.ID != -1 {
					json.NewEncoder(res).Encode(u)
					res.WriteHeader(201)
				} else {
					//user not in database
					message := "User not found"
					res.WriteHeader(404)
					json.NewEncoder(res).Encode(database.NewResponse(message))
					log.Printf(message)
				}
			} else if i == "text/*" {
				// render html
			}
		}
	} else {
		res.WriteHeader(400)
	}
}
