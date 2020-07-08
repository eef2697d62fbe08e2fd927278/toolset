package user

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/youngtrashbag/toolset/src/database"
)

type jUser struct {
	ID           int64 `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	password     string
	CreationDate time.Time `json:"creation_date"`
}

// APIHandleCreate : handles the creation a user
func APIHandleCreate(res http.ResponseWriter, req *http.Request) {
	for _, i := range req.Header["Content-Type"] {
		if i == "application/json" {
			if req.Method == http.MethodPost {

				b, err := ioutil.ReadAll(req.Body)
				if err != nil {
					log.Panicln(err.Error())
				}

				r := bytes.NewReader(b)
				jDecoder := json.NewDecoder(r)

				var u User
				err = jDecoder.Decode(&u)
				if err != nil {
					log.Panicln(err.Error())
				}

				id := u.Insert()

				if id != -1 {
					// if id == -1 then the user could not be created
					message := "Inserted User with ID " + string(id) + " into Database\n"
					log.Println(message)
					json.NewEncoder(res).Encode(database.NewResponse(message))

					res.WriteHeader(http.StatusCreated)
				} else {
					message := "Could not Insert User into Database"
					log.Panicln(message)
					json.NewEncoder(res).Encode(database.NewResponse(message))
					res.WriteHeader(http.StatusBadRequest)
				}
			} else {
				res.WriteHeader(http.StatusMethodNotAllowed)
			}
		} else {
			res.WriteHeader(http.StatusNotAcceptable)
		}
	}
}

// APIHandleByID : handles requests for users with a specified id
func APIHandleByID(res http.ResponseWriter, req *http.Request) {
	for _, i := range req.Header["Accept"] {
		if i == "application/json" {
			if req.Method == http.MethodGet {

				res.Header().Set("Content-Type", "application/json")

				params := mux.Vars(req)
				id, err := strconv.Atoi(params["id"])
				if err != nil {
					log.Panicln(err.Error())
				}

				u := GetByID(int64(id))

				if u.ID != -1 {
					json.NewEncoder(res).Encode(u)
					res.WriteHeader(http.StatusOK)
				} else {
					//user not in database
					message := "User not found"
					res.WriteHeader(http.StatusNotFound)
					json.NewEncoder(res).Encode(database.NewResponse(message))
					log.Printf(message)
				}
			} else {
				res.WriteHeader(http.StatusBadRequest)
			}
		}
	}
}
