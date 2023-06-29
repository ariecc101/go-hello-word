package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World v.2023.06-29_19.32 From DEV"))
	})
	log.Println("Starting server at 5000...")
	http.ListenAndServe(":5000", r)
}
