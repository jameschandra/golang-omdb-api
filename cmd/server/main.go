package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jameschandra/golang-omdb-api/internal/delivery/http"
	"github.com/jameschandra/golang-omdb-api/internal/repository/omdbapi"
	"github.com/jameschandra/golang-omdb-api/internal/usecase/movie"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	apiKey := os.Getenv("API_KEY")
	baseURL := os.Getenv("BASE_URL")

	router := gin.Default()

	omdbRepo := omdbapi.NewOmdbRepository(apiKey, baseURL)
	movieUsecase := movie.NewMovieUsecase(omdbRepo)
	http.NewHandler(router, *movieUsecase)

	router.Run()
}
