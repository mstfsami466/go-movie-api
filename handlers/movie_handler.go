package handlers

import (
	"encoding/json"
	"fmt"
	"go_movie_api/data"
	"net/http"
)

func MoviesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if len(data.Movies) == 0 {
		http.Error(w, "No movies found", http.StatusNotFound)
		return
	}
	err := json.NewEncoder(w).Encode(data.Movies)
	if err != nil {
		http.Error(w, "Error encoding movies data", http.StatusInternalServerError)
		return
	}
}

func CategoriesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if len(data.Categories) == 0 {
		http.Error(w, "No categories found", http.StatusNotFound)
		return
	}

	err := json.NewEncoder(w).Encode(data.Categories)
	if err != nil {
		http.Error(w, "Error encoding categories data", http.StatusInternalServerError)
		return
	}
}

func MovieDetailHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	id := r.URL.Query().Get("id")

	for _, movie := range data.Movies{
		if fmt.Sprint(movie.ID) == id{
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	http.Error(w, "Film bulunamadı", http.StatusNotFound)
	
}

