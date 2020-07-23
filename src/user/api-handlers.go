package user

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/youngtrashbag/toolset/src/utils"
)

type jUser struct {
	ID           int64  `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	CreationDate string `json:"creation_date"`
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

				var usr struct {
					email    string // `json:"email"`
					username string // `json:"username"`
					password string // `json:"password"`
				}

				err = jDecoder.Decode(&usr)
				if err != nil {
					log.Panicln(err.Error())
				}

				u := NewUser(usr.email, usr.username, usr.password)

				id := u.Insert()

				if id == -1 {
					message := "Could not Insert User into Database"
					log.Panicln(message)
					json.NewEncoder(res).Encode(utils.NewResponse(message))
					res.WriteHeader(http.StatusBadRequest)
				} else if id == -2 {

					message := "Username or Email already taken"
					log.Panicln(message)
					json.NewEncoder(res).Encode(utils.NewResponse(message))
					res.WriteHeader(http.StatusBadRequest)
				} else {
					message := "Inserted User with ID " + string(id) + " into database\n"
					log.Println(message)
					json.NewEncoder(res).Encode(utils.NewResponse(message))

					res.WriteHeader(http.StatusCreated)
				}
			} else {
				res.WriteHeader(http.StatusMethodNotAllowed)
			}
		} else {
			res.WriteHeader(http.StatusNotAcceptable)
		}
	}
}

// HandleByID : handles api requests for users by id
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

				u := GetByID(int64(id))

				if u.ID == -1 {
					//user not in database
					message := "User not found"
					res.WriteHeader(http.StatusNotFound)
					json.NewEncoder(res).Encode(utils.NewResponse(message))
					log.Printf(message)
				} else {
					var tm string
					utils.ConvertTime(&u.CreationDate, &tm)
					j := jUser{
						ID:           u.ID,
						Username:     u.Username,
						Email:        u.Email,
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

	utils.LogRequest(res, req)
}
