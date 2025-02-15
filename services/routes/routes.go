package routes

import (
	"net/http"

	"github.com/codinomello/webjetos-go/services/handlers"
)

func HandleAllRoutes(router *http.ServeMux) {
	// Rota principal (index.templ)
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := handlers.HandleTemplIndex(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Rota projetos (projects.templ)
	router.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
		if err := handlers.HandleTemplProjects(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Rota para postar um projeto
	router.HandleFunc("/post-projects", handlers.HandlePostProjects)
}
