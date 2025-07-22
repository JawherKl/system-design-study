package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"

	handler "github.com/JawherKl/url-shortener/internal/handler"
	repo "github.com/JawherKl/url-shortener/internal/repository"
	service "github.com/JawherKl/url-shortener/internal/service"
)

func main() {
	_ = godotenv.Load()

	redisRepo := repo.NewRedisRepo()
	pgRepo := repo.NewPostgresRepo()
	urlService := service.NewURLService(redisRepo, pgRepo)
	h := handler.NewHandler(urlService)

	r := mux.NewRouter()
	r.HandleFunc("/shorten", h.ShortenURL).Methods("POST")
	r.HandleFunc("/{hash}", h.ResolveURL).Methods("GET")

	log.Println("Server running on :8080")
	http.ListenAndServe(":"+os.Getenv("PORT"), r)
}