package models

type Movie struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Category    string   `json:"category"`
	Year        int      `json:"year"`
	Description string   `json:"description"`
	Director    string   `json:"director"`
	Actors      []string `json:"actors"`
	Duration    int      `json:"duration"`
	IMDBScore   float64  `json:"imdb_score"`
	PosterURL   string   `json:"poster_url"`
	TrailerURL  string   `json:"trailer_url"`
}