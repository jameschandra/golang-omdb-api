package omdbapi

import (
	"github.com/jameschandra/golang-omdb-api/internal/entities"
)

type Repository struct{}

func NewMockRepository() *Repository {
	return &Repository{}
}

func (r *Repository) SearchMovies(searchKey string) ([]entities.Movie, error) {
	return []entities.Movie{}, nil
}

func (r *Repository) GetMovieDetail(id string) (*entities.MovieDetail, error) {
	return &entities.MovieDetail{}, nil
}
