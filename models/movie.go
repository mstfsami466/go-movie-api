package models


type Movie struct {
	ID       	int   `json:"id"`
	Title    	string   `json:"title"`
	Category 	string   `json:"category"`
	Year    	int      `json:"year"` 
}
