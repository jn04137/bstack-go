package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"wthunder/bstack/controllers"
	"wthunder/bstack/middlewares"
	"wthunder/bstack/services"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	dbConn, dbConnErr := services.CreateDBConnection()
	if dbConnErr != nil {
		log.Fatalf("Couldn't make connection to database, %v", dbConnErr)
	}
	controller := controllers.ControllerStruct{
		DBDao: dbConn,
	}

	r.Use(middlewares.CheckUserAuth)
	r.Mount("/", controller.GetController())

	http.ListenAndServe(":8080", r)
}
