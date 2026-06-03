package main

import (
	"fmt"
	"go_movie_api/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	routes.RegisterRoutes(r)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	fmt.Println("Film API çalışıyor. Port : " + port)

	if err := r.Run(":" + port); err != nil {
		fmt.Println("Sunucu başlatılırken hata oluştu : ", err)
	}

}
