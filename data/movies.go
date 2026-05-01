package data

import "go_movie_api/models"

var Movies = []models.Movie{
	{ID: 1, Title: "Inception", Category: "Bilim Kurgu", Year: 2010},
	{ID: 2, Title: "The Dark Knight", Category: "Aksiyon", Year: 2008},
	{ID: 3, Title: "Interstellar", Category: "Bilim Kurgu", Year: 2014},
	{ID: 4, Title: "The Godfather", Category: "Suç", Year: 1972},
	{ID: 5, Title: "Pulp Fiction", Category: "Suç", Year: 1994},
	{ID: 6, Title: "The Shawshank Redemption", Category: "Drama", Year: 1994},
	{ID: 7, Title: "The Matrix", Category: "Bilim Kurgu", Year: 1999},
	{ID: 8, Title: "Forrest Gump", Category: "Drama", Year: 1994},
	{ID: 9, Title: "The Avengers", Category: "Aksiyon", Year: 2012},
	{ID: 10, Title: "The Hangover", Category: "Komedi", Year: 2009},
	{ID: 11, Title: "The Lion King", Category: "Animasyon", Year: 1994},
	{ID: 12, Title: "The Conjuring", Category: "Korku", Year: 2013},
	{ID: 13, Title: "The Lord of the Rings: The Fellowship of the Ring", Category: "Macera", Year: 2001},
	

}


var Categories = []string{
	"Bilim Kurgu",
	"Aksiyon",
	"Suç",
	"Drama",
	"Komedi",
	"Romantik",
	"Animasyon",
	"Korku",
	"Macera",
}
