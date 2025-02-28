package controllers

import (
	"os"
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v5"
	"github.com/matoous/go-nanoid/v2"
	"golang.org/x/crypto/bcrypt"

	"wthunder/bstack/models"
	"wthunder/bstack/repositories"
	"wthunder/bstack/middlewares"
)

type UserControllerStruct struct {
	UserRepo *repositories.UserRepositoryStruct
}

func (controller *UserControllerStruct) getController() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/auth_page", controller.authPage)
	r.Get("/players_page", controller.playersPage)
	r.Get("/player_dashboard", controller.userDashboard)

	r.Post("/user_signup", controller.createUser)
	r.Post("/user_login", controller.userLogin)

	return r
}

func (controller *UserControllerStruct) authPage(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles("./views/fragments/layout.html", "./views/auth_page.html")
	
	m := map[string]interface{} {
		"username": r.Context().Value("username"),
		"nanoId": r.Context().Value("nanoId"),
	}

	if err != nil {
		log.Printf("This is the err %v", err.Error())
	}
	templ.ExecuteTemplate(w, "auth_page.html", m)
	return
}

func (controller *UserControllerStruct) userLogin(w http.ResponseWriter, r *http.Request) {
	userRepo := controller.UserRepo
	userName := r.FormValue("username")
	userPass := r.FormValue("password")

	player, err := userRepo.GetUserWithPassHash(userName)
	if err != nil {
		log.Printf("Error finding player: %s; Error from repo: %v", player, err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(player.Password), []byte(userPass))
	if err != nil {
		log.Printf("Password didn't match: %v", err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	claims := middlewares.BstackCustomClaims {
		Username: player.Username,
		NanoId: player.NanoId,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtSecret := os.Getenv("JWT_SECRET")
	signedString, err := token.SignedString([]byte(jwtSecret))

	if err != nil {
		log.Printf("Something happened with jwtToken creation: %v", err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	cookie := http.Cookie{
		Name: "bstack_token",
		Value: signedString,
		HttpOnly: true,
		Path: "/",
	}

	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/", http.StatusSeeOther)
	return 
}

func (controller *UserControllerStruct) createUser(w http.ResponseWriter, r *http.Request) {
	userName := r.FormValue("username")
	userPass := r.FormValue("password")
	hashedPassword, bcryptErr := bcrypt.GenerateFromPassword([]byte(userPass), 10)
	if bcryptErr != nil {
		log.Printf("Error occurred hashing password\n%v", bcryptErr.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	nanoId, err := gonanoid.New(); if err != nil {
		log.Printf("Error occurred generating nanoid\n%v", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	newPlayer := models.Player{
		Username: userName,
		Password: string(hashedPassword),
		NanoId: nanoId,
	}

	err = controller.UserRepo.CreateUser(newPlayer)
	if err != nil {
		log.Printf("Failed to create new player\n%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}

func (controller *UserControllerStruct) playersPage(w http.ResponseWriter, r *http.Request) {
	repo := controller.UserRepo
	// serve page that shows players on site
	templ, err := template.ParseFiles("./views/fragments/layout.html", "./views/players_page.html")
	if err != nil {
		log.Printf("This is the err %v", err.Error())
	}
	
	players, err := repo.GetUsers()
	if err != nil {
		log.Printf("This is the err %v", err.Error())
	}

	log.Printf("These are the players: %v", players)
	
	m := map[string]interface{} {
		"username": r.Context().Value("username"),
		"nanoId": r.Context().Value("nanoId"),
		"players": players,
	}
	templ.ExecuteTemplate(w, "players_page.html", m)
	return
}

func (controller *UserControllerStruct) userDashboard(w http.ResponseWriter, r *http.Request) {
	m := map[string]interface{} {
		"username": r.Context().Value("username"),
		"nanoId": r.Context().Value("nanoId"),
	}
	
	templ, err := template.ParseFiles("./views/fragments/layout.html", "./views/user_dashboard.html")
	if err != nil {
		log.Printf("This is the err %v", err.Error())
	}

	templ.ExecuteTemplate(w, "user_dashboard.html", m)
	return
}


