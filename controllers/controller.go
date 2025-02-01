package controllers

import (
	"log"
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type ControllerStruct struct {

}

func (controller *ControllerStruct) GetController() *chi.Mux {
	userController := UserControllerStruct{}
	staticController := StaticControllerStruct{}

	r := chi.NewRouter()

	r.Get("/", controller.homePage)
	r.Mount("/user", userController.getController())
	r.Mount("/static", staticController.getController())

	return r
}

func (controller *ControllerStruct) homePage(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles("./views/fragments/layout.html", "./views/home.html")
	if err != nil {
		log.Printf("This is the err %v", err.Error())
	}
	templ.ExecuteTemplate(w, "home.html", nil)
	return
}
