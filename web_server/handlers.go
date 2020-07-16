package main

import (
	"fmt"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hola mundo")
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This the API Endpoint")
}