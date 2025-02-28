package controllers

import (
	"log"
	"net/http"
	"html/template"
	
	"github.com/go-chi/chi/v5"
)

type TeamControllerStruct struct {

}

func (controller *TeamControllerStruct) getController() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/teams_page", controller.getTeamsPage)
	r.Post("/create", controller.createTeam)

	return r
}

func (controller *TeamControllerStruct) getTeamsPage(w http.ResponseWriter, r *http.Request) {
	m := map[string]interface{} {
		"username": r.Context().Value("username"),
		"nanoId": r.Context().Value("nanoId"),
	}

	templ, err := template.ParseFiles("./views/teams_page.html", "./views/fragments/layout.html")
	if err != nil {
		log.Printf("This is the err %v", err.Error())
	}

	templ.ExecuteTemplate(w, "teams_page.html", m)
	return
}

func (controller *TeamControllerStruct) createTeam(w http.ResponseWriter, r *http.Request) {
	teamName := r.FormValue("teamName")
	teamDesc := r.FormValue("teamDesc")

	//m := map[string]interface{} {
	//	"username": r.Context().Value("username"),
	//	"nanoId": r.Context().Value("nanoId"),
	//}

	/*
		Check if user is logged in (should be handled by middleware)
		Check if user has 5 teams already
		Use user nano_id to create a team and set as owner
	*/

	log.Printf("create team endpoint was hit")
	log.Printf("teamName: %s, teamDesc: %s", teamName, teamDesc)

	return
}

