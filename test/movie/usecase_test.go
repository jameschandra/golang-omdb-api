package test

import (
	"testing"

	"github.com/jameschandra/golang-omdb-api/internal/usecase/movie"
	"github.com/jameschandra/golang-omdb-api/test/mock/omdbapi"
)

func TestSearchMovies(t *testing.T) {
	repo := omdbapi.NewMockRepository()

	usecase := movie.NewMovieUsecase(repo)

	movies, err := usecase.SearchMovies("Batman")

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if movies == nil {
		t.Errorf("Expected non-nil movies")
	}
}

func TestGetMovieDetail(t *testing.T) {
	repo := omdbapi.NewMockRepository()

	usecase := movie.NewMovieUsecase(repo)

	movie, err := usecase.GetMovieDetail("tt0374526")

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if movie == nil {
		t.Errorf("Expected non-nil movie")
	}
}
