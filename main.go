package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World v.2023.06-29_21.05 From " + os.Getenv("MODE")))
	})
	log.Println("Starting server at " + os.Getenv("PORT"))
	http.ListenAndServe(":"+os.Getenv("PORT"), r)
}
