package controllers

import (
	"final-project/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PublisherInput struct {
	Name        string `json:"name"`
	ImageURL	string `json:"image_url"`
}

// Get All Publisher godoc
// @Summary Get all publisher
// @Description Get list of publisher
// @Tags Public
// @Produce json
// @Success 200 {object} []models.Publisher
// @Router /publishers [get]
func GetAllPublisher(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	var publisher []models.Publisher

	db.Find(&publisher)

	c.JSON(http.StatusOK, gin.H{"data": publisher})
}

// Get Games by Publisher godoc
// @Summary Get list of games in specific publisher
// @Description Get all games of spesific publisher by id
// @Tags Public
// @Produce json
// @Param id path string true "Publisher Id"
// @Success 200 {object} []models.Game
// @Router /publishers/{id}/games [get]
func GetGamesByPublisherId(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	var games []models.Game
	if err := db.Where("publisher_id = ?", c.Param("id")).Find(&games).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": games})
}
