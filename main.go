package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type JsonResponse struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type User struct {
	ID    int    `json:"id"`
	EMail string `json:"email"`
	Name  string `json:"name"`
}

type Token struct {
	Token string `json:"token"`
	ID    int    `json:"id"`
}

func main() {
	r := gin.Default()

	r.POST("/auth", func(c *gin.Context) {

		var requestBody JsonResponse

		if err := c.BindJSON(&requestBody); err != nil {

		}

		if requestBody.UserName == "jakub" && requestBody.Password == "jakub" {
			c.JSON(http.StatusOK, Token{
				Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjoiMTIzNCIsIm5hbWUiOiJKb2huIERvZSIsImlhdCI6MTUxNjIzOTAyMn0.6sqXAR68gEsFhodi8q3vmMeNRDrsyrHeWQ45RxCQPac",
				ID:    123,
			})

		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Bad credentials"})
		}

	})

	r.GET("/user-info/:id", func(c *gin.Context) {

		c.JSON(http.StatusOK, User{ID: 123, EMail: "jakub@jakub.cz", Name: "Jakub Englický"})
	})

	r.Run()
}
