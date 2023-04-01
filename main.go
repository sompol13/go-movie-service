package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/api/v1/movies", getMovies)
	router.GET("/api/v1/movies/:id", getMovieById)
	router.Run("localhost:3001")
}

type movie struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var movies = []movie{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getMovies(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, movies)
}

func getMovieById(context *gin.Context) {
	id := context.Param("id")
	for _, a := range movies {
		if a.ID == id {
			context.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{
		"message": "album not found",
	})
}
