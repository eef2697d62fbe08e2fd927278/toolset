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

				if id != -1 {
					// if id == -1 then the user could not be created
					message := "Inserted User with ID " + string(id) + " into database\n"
					log.Println(message)
					json.NewEncoder(res).Encode(utils.NewResponse(message))

					res.WriteHeader(http.StatusCreated)
				} else {
					message := "Could not Insert User into Database"
					log.Panicln(message)
					json.NewEncoder(res).Encode(utils.NewResponse(message))
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

					var t string
					utils.ConvertTime(&u.CreationDate, &t)
					j := jUser{
						ID:           u.ID,
						Username:     u.Username,
						Email:        u.Email,
						CreationDate: t,
					}

					json.NewEncoder(res).Encode(j)
					res.WriteHeader(http.StatusOK)
				} else {
					//user not in database
					message := "User not found"
					res.WriteHeader(http.StatusNotFound)
					json.NewEncoder(res).Encode(utils.NewResponse(message))
					log.Printf(message)
				}
			} else {
				res.WriteHeader(http.StatusBadRequest)
			}
		}
	}
}

/*
	frontend handlers
*/

// HandleByID : handles frontend requests
func HandleByID(res http.ResponseWriter, req *http.Request) {
	for _, i := range req.Header["Accept"] {
		if i == "text/*" {
			if req.Method == http.MethodGet {

				res.Header().Set("Content-Type", "text/html")

				params := mux.Vars(req)
				id, err := strconv.Atoi(params["id"])
				if err != nil {
					log.Panicln(err.Error())
				}

				user := GetByID(int64(id))

				if user.ID != -1 {
					log.Print("user found")
					res.Write(user.RenderPage())
				} else {
					// user not found in database
					res.Write([]byte("User Not Found"))
					res.WriteHeader(http.StatusNotFound)
				}

			}
		}
	}

	utils.LogRequest(res, req)
}
