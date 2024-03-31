package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func New() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/persons", GetAllPersons).Methods("GET")
	router.HandleFunc("/persons/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/quest", CreatePerson).Methods("POST")

	return router
}
