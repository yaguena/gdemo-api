package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gaguena/gdemo-api/models"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

type PersonInput struct {
	Name   string `json:"name" validate:"required"`
	Email  string `json:"email" validate:"required"`
	Reward int    `json:"reward" validate:"required"`
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var input PersonInput

	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &input)

	err := validator.New().Struct(input)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Validation Error")
		return
	}

	person := &models.Person{
		Name:   input.Name,
		Email:  input.Email,
		Reward: input.Reward,
	}

	models.DB.Create(person)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(person)

}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	var person models.Person

	if err := models.DB.Where("id = ?", id).First(&person).Error; err != nil {
		RespondWithError(w, http.StatusNotFound, "Person not found")
		return
	}

	json.NewEncoder(w).Encode(person)
}

func GetAllPersons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var persons []models.Person
	models.DB.Find(&persons)

	json.NewEncoder(w).Encode(persons)
}
