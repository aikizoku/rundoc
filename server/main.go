package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/unrolled/render"
)

func handler(w http.ResponseWriter, r *http.Request) {
	src := map[string]interface{}{
		"hoge": "aaaaaa",
		"fuga": 1,
	}
	ren := render.New()
	ren.JSON(w, http.StatusOK, src)
}

func main() {
	r := chi.NewRouter()
	r.Get("/sample", handler)
	r.Post("/sample", handler)
	r.Put("/sample", handler)
	r.Delete("/sample", handler)
	http.ListenAndServe(":8080", r)
}
