package controllers

import (
	"final-project/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GenreInput struct {
	Name        string `json:"name"`
}

// Get all Genre godoc
// @Summary Get all Genre
// @Description Get list of Genre
// @Tags Genre
// @Produce json
// @Success 200 {object} []models.Genre
// @Router /genres [get]
func GetAllGenre(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	var genres []models.Genre

	db.Find(&genres)

	c.JSON(http.StatusOK, gin.H{"data": genres})

}

// Create Genre godoc
// @Summary Create a Genre
// @Description Create new Genre
// @Tags Genre
// @Param Body body GenreInput true "the body to create new genre"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Genre
// @Router /genres [post]
func CreateGenre(c *gin.Context) {
	var input GenreInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	genre := models.Genre{Name: input.Name}
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	db.Create(&genre)
	c.JSON(http.StatusOK, gin.H{"data": genre})
}

// Get Genre godoc
// @Summary Get Genre by id
// @Description Get one Genre by id
// @Tags Genre
// @Produce json
// @Param id path string true "Genre Id"
// @Success 200 {object} models.Genre
// @Router /genres/{id} [get]
func GetGenreById(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	// check if exist and get data
	var genre models.Genre
	if err := db.Where("id = ?", c.Param("id")).First(&genre).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": genre})
}

// Get games from one Genre godoc
// @Summary Get games by Genre by id
// @Description Get all games of spesific Genre by id
// @Tags Genre
// @Produce json
// @Param id path string true "Genre Id"
// @Success 200 {object} []models.Game
// @Router /genres/{id}/games [get]
func GetGamesByGenreId(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	var games []models.Game
	if err := db.Where("genre_id = ?", c.Param("id")).Find(&games).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": games})

}

// Update Genre godoc
// @Summary update a Genre by id
// @Description update one Genre by id
// @Tags Genre
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Genre Id"
// @Param Body body GenreInput true "the body to create new genre"
// @Success 200 {object} models.Genre
// @Router /genres/{id} [patch]
func UpdateGenre(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	// get rating if exist
	var genre models.Genre
	if err := db.Where("id = ?", c.Param("id")).First(&genre).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	var input GenreInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInputGenre models.Genre

	updatedInputGenre.Name = input.Name
	updatedInputGenre.UpdatedAt = time.Now()

	db.Model(&genre).Updates(updatedInputGenre)
	c.JSON(http.StatusOK, gin.H{"data": genre})
}

// Delete a Genre godoc
// @Summary delete a Genre by id
// @Description delete one Genre by id
// @Tags Genre
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Genre Id"
// @Success 200 {object} map[string]boolean
// @Router /genres/{id} [delete]
func DeleteGenre(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	var genre models.Genre
	if err := db.Where("id = ?", c.Param("id")).First(&genre).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	db.Delete(&genre)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
