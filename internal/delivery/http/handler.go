package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jameschandra/golang-omdb-api/internal/usecase/movie"
)

type Handler struct {
	movieUsecase movie.Usecase
}

func NewHandler(router *gin.Engine, movieUsecase movie.Usecase) *Handler {
	handler := &Handler{movieUsecase: movieUsecase}

	router.GET("/search", handler.SearchMovies)
	router.GET("/detail/:id", handler.GetMovieDetail)

	return handler
}

func (h *Handler) SearchMovies(c *gin.Context) {
	searchKey := c.Query("s")
	movies, err := h.movieUsecase.SearchMovies(searchKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, movies)
}

func (h *Handler) GetMovieDetail(c *gin.Context) {
	id := c.Param("id")
	movie, err := h.movieUsecase.GetMovieDetail(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, movie)
}
