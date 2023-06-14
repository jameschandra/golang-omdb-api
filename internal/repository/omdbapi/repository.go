package omdbapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jameschandra/golang-omdb-api/internal/entities"
)

type Repository struct {
	apiKey  string
	baseURL string
}

func NewOmdbRepository(apiKey, baseURL string) *Repository {
	return &Repository{
		apiKey:  apiKey,
		baseURL: baseURL,
	}
}

func (r *Repository) SearchMovies(searchKey string) ([]entities.Movie, error) {
	resp, err := http.Get(fmt.Sprintf("%s/?apikey=%s&s=%s", r.baseURL, r.apiKey, searchKey))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var searchResponse entities.SearchResponse
	err = json.NewDecoder(resp.Body).Decode(&searchResponse)
	if err != nil {
		return nil, err
	}

	return searchResponse.Search, nil
}

func (r *Repository) GetMovieDetail(id string) (*entities.MovieDetail, error) {
	resp, err := http.Get(fmt.Sprintf("%s/?apikey=%s&i=%s", r.baseURL, r.apiKey, id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var movieDetail entities.MovieDetail
	err = json.NewDecoder(resp.Body).Decode(&movieDetail)
	if err != nil {
		return nil, err
	}

	return &movieDetail, nil
}
