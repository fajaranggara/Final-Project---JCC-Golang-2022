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

// Get All Genre godoc
// @Summary Get all genre
// @Description Get list of genre
// @Tags Public
// @Produce json
// @Success 200 {object} []models.Genre
// @Router /genres [get]
func GetAllGenre(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	var genre []models.Genre

	db.Find(&genre)

	c.JSON(http.StatusOK, gin.H{"data": genre})

}

// Get Games by Genre godoc
// @Summary Get list of games in specific genre
// @Description Get all games of spesific genre by id
// @Tags Find Games By
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

// Create Genre godoc
// @Summary Create a new genre
// @Description Only admin have permission to create genre
// @Tags Admin
// @Param Body body GenreInput true "the body to create new genre"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Genre
// @Router /genres [post]
func CreateGenre(c *gin.Context) {
	//check authorization
	cUser, _ := models.GetCurrentUser(c)
	if cUser.Role != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Only for admin level user"})
        return
	}

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

// Update Genre godoc
// @Summary Update existing genre by id
// @Description Only admin have permission to update genre
// @Tags Admin
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Genre Id"
// @Param Body body GenreInput true "the body to create new genre"
// @Success 200 {object} models.Genre
// @Router /genres/{id} [patch]
func UpdateGenre(c *gin.Context) {
	//check authorization
	cUser, _ := models.GetCurrentUser(c)
	if cUser.Role != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Only for admin level user"})
        return
	}

	db := c.MustGet("db").(*gorm.DB)
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
// @Summary Delete existing genre by id
// @Description Only admin have permission to delete genre
// @Tags Admin
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Genre Id"
// @Success 200 {object} map[string]boolean
// @Router /genres/{id} [delete]
func DeleteGenre(c *gin.Context) {
	//check authorization
	cUser, _ := models.GetCurrentUser(c)
	if cUser.Role != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Only for admin level user"})
        return
	}
	
	db := c.MustGet("db").(*gorm.DB)

	var genre models.Genre
	if err := db.Where("id = ?", c.Param("id")).First(&genre).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	db.Delete(&genre)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
