package note

import (
	"fmt"
	"net/http"
)

// Handler : supposed to handle the /note resource
func Handler(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		fmt.Println("GetRequest")
	} else if req.Method == http.MethodPost {
		fmt.Println("PostRequest")
	}
}
