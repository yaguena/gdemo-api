package main

import (
	"net/http"

	"github.com/gaguena/gdemo-api/controllers"
	"github.com/gaguena/gdemo-api/models"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	handler := controllers.New()

	server := &http.Server{
		Addr:    "0.0.0.0:8008",
		Handler: handler,
	}

	models.ConnectDatabase()

	server.ListenAndServe()
}
