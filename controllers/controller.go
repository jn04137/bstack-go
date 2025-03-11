package controllers

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"wthunder/bstack/repositories"
)

type ControllerStruct struct {
	DBDao *sql.DB
}

func (controller *ControllerStruct) GetController() *chi.Mux {
	userController := UserControllerStruct{
		UserRepo: &repositories.UserRepositoryStruct{DBDao: controller.DBDao},
	}
	teamController := TeamControllerStruct{
		TeamRepo: &repositories.TeamRepositoryStruct{DBBao: controller.DBDao},
	}
	staticController := StaticControllerStruct{}

	r := chi.NewRouter()

	r.Get("/", controller.homePage)
	r.Mount("/user", userController.getController())
	r.Mount("/team", teamController.getController())
	r.Mount("/static", staticController.getController())

	return r
}

func (controller *ControllerStruct) homePage(w http.ResponseWriter, r *http.Request) {
	m := map[string]any {
		"username": r.Context().Value("username"),
		"nanoId": r.Context().Value("nanoId"),
	}

	templ, err := template.ParseFiles("./views/fragments/layout.html", "./views/home.html")
	if err != nil {
		log.Printf("This is the err %v", err.Error())
	}
	templ.ExecuteTemplate(w, "home.html", m)
	return
}
