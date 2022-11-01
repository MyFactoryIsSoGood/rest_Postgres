package main

import (
	"REST_Postgres/controllers"
	"REST_Postgres/models"
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	models.ConnectDB()

	route.GET("/movies", controllers.GetAllMovies)
	route.GET("/movies/:id", controllers.GetMovie)
	route.PATCH("/movies/:id", controllers.EditMovie)
	route.POST("/movies", controllers.CreateMovie)
	route.DELETE("/movies/:id", controllers.DeleteMovie)

	route.Run(":8080")
}
