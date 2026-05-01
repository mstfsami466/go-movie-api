package main

import (
	"fmt"
	"go_movie_api/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/filmler", handlers.MoviesHandler)
	http.HandleFunc("/kategoriler", handlers.CategoriesHandler)
	http.HandleFunc("/film", handlers.MovieDetailHandler)

	fmt.Println("Film api çalışıyor...")
	fmt.Println("http://localhost:8080 adresine giderek api'ye erişebilirsiniz.")
	http.ListenAndServe(":8080", nil)
	

}