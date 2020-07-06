package main

import (
	"log"
	"net/http"
)

// Handler : Handles API call to root (/)
func Handler(res http.ResponseWriter, req *http.Request) {
	var statusCode int
	var message string

	if req.Method == http.MethodGet {
		statusCode = 200
		message = "hello get world"
	} else if req.Method == http.MethodPost {
		statusCode = 201
		message = "hello post world"
	} else if req.Method == http.MethodPut {
		statusCode = 202
		message = "hello put world"
	} else if req.Method == http.MethodDelete {
		statusCode = 204
		message = "hello delete world"
	} else {
		statusCode = 404
	}

	log.Printf("%s Method on \"%s\", StatusCode:%d, Message:\"%s\"", req.Method, req.URL.Path, statusCode, message)

	res.WriteHeader(statusCode)
	res.Write([]byte(message))
}
