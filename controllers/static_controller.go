package controllers

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi/v5"
)

type StaticControllerStruct struct {}

func (staticController *StaticControllerStruct) getController() *chi.Mux {
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir,"static"))
	
	r := chi.NewRouter()
	fileServer(r, "/", filesDir)

	r.Get("/", staticController.fileServer)

	return r
}

func (staticController *StaticControllerStruct) fileServer(w http.ResponseWriter, r *http.Request) {

}

func fileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
