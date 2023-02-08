package main

import (
	"PracticalTask/config"
	"PracticalTask/controller"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	config.ConnectDB(config.QueryHistory{})
}

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{config.Client_Domain},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == config.Client_Domain
		},
		MaxAge: 12 * time.Hour,
	}))
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, Welcome to the Gin Backend for my Movie search application !",
		})
	})
	router.GET("/movies/search/:title", controller.MovieSearchHandler)
	router.Run(":8080")
}
