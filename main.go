package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	
	"wthunder/bstack/controllers"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	controller := controllers.ControllerStruct{}

	r.Mount("/", controller.GetController())

	http.ListenAndServe(":8080", r)
}
