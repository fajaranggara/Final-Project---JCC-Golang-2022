package controllers

import (
	"final-project/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PublisherInput struct {
	Name        string `json:"name"`
	ImageURL	string `json:"image_url"`
}

// Get all Publisher godoc
// @Summary Get all Publisher
// @Description Get list of Publisher
// @Tags Publisher
// @Produce json
// @Success 200 {object} []models.Publisher
// @Router /publishers [get]
func GetAllPublisher(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	var publishers []models.Publisher

	db.Find(&publishers)

	c.JSON(http.StatusOK, gin.H{"data": publishers})

}

// Create Publisher godoc
// @Summary Create a Publisher
// @Description Create new Publisher
// @Tags Publisher
// @Param Body body PublisherInput true "the body to create new publisher"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Publisher
// @Router /publishers [post]
func CreatePublisher(c *gin.Context) {
	var input PublisherInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	publisher := models.Publisher{Name: input.Name,  ImageURL: input.ImageURL}
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	db.Create(&publisher)
	c.JSON(http.StatusOK, gin.H{"data": publisher})
}

// Get Publisher godoc
// @Summary Get Publisher by id
// @Description Get one Publisher by id
// @Tags Publisher
// @Produce json
// @Param id path string true "Publisher Id"
// @Success 200 {object} models.Publisher
// @Router /publishers/{id} [get]
func GetPublisherById(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	// check if exist and get data
	var publisher models.Publisher
	if err := db.Where("id = ?", c.Param("id")).First(&publisher).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": publisher})
}

// Get games from one Publisher godoc
// @Summary Get games by Publisher by id
// @Description Get all games of spesific Publisher by id
// @Tags Publisher
// @Produce json
// @Param id path string true "Publisher Id"
// @Success 200 {object} []models.Game
// @Router /publishers/{id}/games [get]
func GetGamesByPublisherId(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	var publishers []models.Game
	if err := db.Where("publisher_id = ?", c.Param("id")).Find(&publishers).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": publishers})

}

// Update Publisher godoc
// @Summary update a Publisher by id
// @Description update one Publisher by id
// @Tags Publisher
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Publisher Id"
// @Param Body body PublisherInput true "the body to create new publisher"
// @Success 200 {object} models.Publisher
// @Router /publishers/{id} [patch]
func UpdatePublisher(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	// get rating if exist
	var publisher models.Publisher
	if err := db.Where("id = ?", c.Param("id")).First(&publisher).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	var input PublisherInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInputPublisher models.Publisher

	updatedInputPublisher.Name = input.Name
	updatedInputPublisher.ImageURL = input.ImageURL
	updatedInputPublisher.UpdatedAt = time.Now()

	db.Model(&publisher).Updates(updatedInputPublisher)
	c.JSON(http.StatusOK, gin.H{"data": publisher})
}

// Delete a Publisher godoc
// @Summary delete a Publisher by id
// @Description delete one Publisher by id
// @Tags Publisher
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Publisher Id"
// @Success 200 {object} map[string]boolean
// @Router /publishers/{id} [delete]
func DeletePublisher(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	var publisher models.Publisher
	if err := db.Where("id = ?", c.Param("id")).First(&publisher).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	db.Delete(&publisher)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
