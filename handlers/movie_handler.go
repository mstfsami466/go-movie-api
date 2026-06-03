package handlers

import (
	"fmt"
	"go_movie_api/data"
	"go_movie_api/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func MoviesHandler(c *gin.Context) {
	if len(data.Movies) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Film bulunamadı",
		})
		return
	}

	c.JSON(http.StatusOK, data.Movies)
}


func CategoriesHandler(c *gin.Context){
	if len(data.Categories) == 0{
		c.JSON(http.StatusNotFound, gin.H{
			"message" : "Kategori bulunamadı",
		})

		return
	}
	c.JSON(http.StatusOK, data.Categories)
}

func MovieDetailHandler(c *gin.Context){
	id := c.Param("id")

	for _, movie := range data.Movies{
		if fmt.Sprint(movie.ID) == id{
			c.JSON(http.StatusOK, movie)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"message" : "Film bulunmadi",
	})

}

func MovieDetailQueryHandler(c *gin.Context){
	id := c.Query("id")

	if id == ""{
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : "Film id değeri gönderilmedi",
		})
		return
	}

	for _, movie := range data.Movies{
		if fmt.Sprint(movie.ID) == id{
			c.JSON(http.StatusOK, movie)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"message" : "Film bulunamadi",
	})
}

func MoviesByCategoryHandler(c * gin.Context){
	category := c.Param("category")
	var filteredMovies []models.Movie

	for _, movie := range data.Movies{
		// Büyük/küçük harf duyarsız karşılaştırma yaparak kategoriyi kontrol ediyoruz
		if strings.EqualFold(movie.Category, category){
			filteredMovies = append(filteredMovies, movie)
		}
	}

	if len(filteredMovies) == 0{
		c.JSON(http.StatusNotFound, gin.H{
			"message" : "Bu kategoriye ait film bulunamadi",
		})
		return
	}
	c.JSON(http.StatusOK, filteredMovies)
}

func SearchMoviesHandler(c * gin.Context){
	title := c.Query("title")

	if title == ""{
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : "Arama yapmak için bir başlık değeri gönderilmelidir",
		})
		return
	}

	var filteredMovies []models.Movie

	for _, movie := range data.Movies{
		// Büyük/küçük harf duyarsız karşılaştırma yaparak başlıkta arama yapıyoruz
		if strings.Contains(strings.ToLower(movie.Title), strings.ToLower(title)){
			filteredMovies = append(filteredMovies, movie)
		}
		
	}

	if len(filteredMovies) == 0{
		c.JSON(http.StatusNotFound, gin.H{
			"message" : "Aranan filme ait sonuç bulunmadı",
		})
		return
	}
	c.JSON(http.StatusOK, filteredMovies)

}

