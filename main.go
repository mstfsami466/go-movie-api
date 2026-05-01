package main

import (
	"fmt"
	"go_movie_api/handlers"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/movies", handlers.MoviesHandler)
	http.HandleFunc("/categories", handlers.CategoriesHandler)
	http.HandleFunc("/movie", handlers.MovieDetailHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Film API çalışıyor. Port:", port)
	http.ListenAndServe(":"+port, nil)
}
