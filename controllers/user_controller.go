package controllers

import (
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

}

func (userController *UserControllerStruct) userLogin(w http.ResponseWriter, r *http.Request) {

}

func (userController *UserControllerStruct) createUser(w http.ResponseWriter, r *http.Request) {

}

