package controllers

import (
	"html/template"
	"log"
	"net/http"
	
	"github.com/go-chi/chi/v5"

	"wthunder/bstack/repositories"
)

type TeamControllerStruct struct {
	TeamRepo *repositories.TeamRepositoryStruct
}

func (controller *TeamControllerStruct) getController() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/teams_page", controller.getTeamsPage)
	r.Get("/team_page/{teamNanoId}", controller.viewTeam)

	r.Post("/create", controller.createTeam)

	return r
}

func (controller *TeamControllerStruct) getTeamsPage(w http.ResponseWriter, r *http.Request) {

	templ, err := template.ParseFiles("./views/teams_page.html", "./views/fragments/layout.html")
	if err != nil {
		log.Printf("This is the err %v", err.Error())
	}

	teams, err := controller.TeamRepo.GetTeams()
	log.Printf("These are the teams: %v", teams)
	
	m := map[string]any {
		"username": r.Context().Value("username"),
		"nanoId": r.Context().Value("nanoId"),
		"teams": teams,
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

func (controller *TeamControllerStruct) viewTeam(w http.ResponseWriter, r *http.Request) {
	nanoIdParam := chi.URLParam(r, "teamNanoId")
	repo := controller.TeamRepo
	
	templ, err := template.ParseFiles("./views/team_page.html", "./views/fragments/layout.html")
	if err != nil {
		log.Printf("This is the err %v", err.Error())
	}

	log.Printf("this is the nanoIdParam: %v", nanoIdParam)
	team, err := repo.GetTeam(nanoIdParam)
	if err != nil {
		log.Printf("This is the err %v", err.Error())
	}

	players, err := repo.GetPlayersOnTeam(nanoIdParam)
	
	m := map[string]any {
		"username": r.Context().Value("username"),
		"nanoId": r.Context().Value("nanoId"),
		"team": team,
		"players": players,
	}
	
	templ.ExecuteTemplate(w, "team_page.html", m)
	return
}

