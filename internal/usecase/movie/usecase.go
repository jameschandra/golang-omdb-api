package movie

import (
	"github.com/jameschandra/golang-omdb-api/internal/entities"
)

type Usecase struct {
	movieRepo entities.MovieRepository
}

func NewMovieUsecase(movieRepo entities.MovieRepository) *Usecase {
	return &Usecase{
		movieRepo: movieRepo,
	}
}

func (u *Usecase) SearchMovies(searchKey string) ([]entities.Movie, error) {
	return u.movieRepo.SearchMovies(searchKey)
}

func (u *Usecase) GetMovieDetail(id string) (*entities.MovieDetail, error) {
	return u.movieRepo.GetMovieDetail(id)
}
