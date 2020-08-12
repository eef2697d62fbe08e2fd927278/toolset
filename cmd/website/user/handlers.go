package user

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/youngtrashbag/toolset/src/utils"
)

// RenderPage : Renders the page of a user
func (u *User) RenderPage() []byte {

	log.Printf("Rendering User ID=\"%d\"", u.ID)

	hTmpl, err := ioutil.ReadFile("templates/head.html")
	if err != nil {
		log.Panicln(err.Error())
	}
	uTmpl, err := ioutil.ReadFile("templates/user.html")
	if err != nil {
		log.Panicln(err.Error())
	}

	tmpl, err := template.New("user").Parse(string(append(hTmpl[:], uTmpl[:]...)))
	if err != nil {
		log.Panicln(err.Error())
	}

	var templateBytes bytes.Buffer

	err = tmpl.Execute(&templateBytes, u)
	if err != nil {
		log.Println(err.Error())
	}

	return templateBytes.Bytes()
}

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

	utils.LogRequest(req)
}
