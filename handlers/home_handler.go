package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message" : "Film Api'ye Hoş Geldiniz",
		"status" : "success",
	})
}