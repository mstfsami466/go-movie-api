package routes

import (
	"go_movie_api/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/", handlers.HomeHandler)

	r.GET("/movies", handlers.MoviesHandler)
	r.GET("/movies/:id", handlers.MovieDetailHandler)

	r.GET("/movie", handlers.MovieDetailQueryHandler)

	r.GET("/categories", handlers.CategoriesHandler)
	r.GET("/categories/:category/movies", handlers.MoviesByCategoryHandler)

	r.GET("/search/movies", handlers.SearchMoviesHandler)
}
