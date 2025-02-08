package controllers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type UserControllerStruct struct {}

func (userController *UserControllerStruct) getController() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/auth_page", userController.authPage)
	r.Post("/user_signup", userController.createUser)
	r.Post("/user_login", userController.userLogin)

	return r
}

func (userController *UserControllerStruct) authPage(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles("./views/fragments/layout.html", "./views/auth_page.html")
	if err != nil {
		log.Printf("This is the err %v", err.Error())
	}
	templ.ExecuteTemplate(w, "auth_page.html", nil)
	return
}

func (userController *UserControllerStruct) userLogin(w http.ResponseWriter, r *http.Request) {
	userName := r.FormValue("username")
	userPass := r.FormValue("password")

	log.Printf("Username: %s, Password: %s", userName, userPass)
}

func (userController *UserControllerStruct) createUser(w http.ResponseWriter, r *http.Request) {
	userName := r.FormValue("username")
	userPass := r.FormValue("password")

	log.Printf("Username: %s, Password: %s", userName, userPass)
}

