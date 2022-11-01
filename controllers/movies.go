package controllers

import (
	"REST_Postgres/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateMovieInput struct {
	Title       string `json:"title" binding:"required"`
	ReleaseYear uint16 `json:"release_year" binding:"required"`
}

type UpdateMovieInput struct {
	Title       string `json:"title"`
	ReleaseYear uint16 `json:"release_year"`
}

func GetAllMovies(c *gin.Context) {
	var movies []models.Movie
	models.DB.Find(&movies)

	c.IndentedJSON(http.StatusOK, gin.H{"movies": movies})
}

func GetMovie(c *gin.Context) {
	var movie models.Movie
	err := models.DB.Where("id = ?", c.Param("id")).First(&movie).Error
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Movie doesn't exist"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"movie": movie})
}

func CreateMovie(c *gin.Context) {
	var input CreateMovieInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	movie := models.Movie{Title: input.Title, ReleaseYear: input.ReleaseYear}
	models.DB.Create(&movie)
	c.IndentedJSON(http.StatusOK, gin.H{"movie": movie})
}

func EditMovie(c *gin.Context) {
	var movie models.Movie
	err := models.DB.Where("id = ?", c.Param("id")).First(&movie).Error
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Movie doesn't exists"})
		return
	}

	var input UpdateMovieInput
	err = c.ShouldBind(&input)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&movie).Update(input)

	c.IndentedJSON(http.StatusOK, gin.H{"movie": movie})

}

func DeleteMovie(c *gin.Context) {
	var movie models.Movie
	err := models.DB.Where("id = ?", c.Param("id")).First(&movie).Error
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Movie doesn't exist"})
		return
	}

	models.DB.Delete(&movie)
	c.IndentedJSON(http.StatusOK, gin.H{"deleted": true})
}
