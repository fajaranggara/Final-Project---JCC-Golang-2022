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
// @Tags Find Games By
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

// Create Publisher godoc
// @Summary Create a new publisher
// @Description Only admin have permission to create publisher
// @Tags Admin
// @Param Body body PublisherInput true "the body to create new publisher"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Publisher
// @Router /publishers [post]
func CreatePublisher(c *gin.Context) {
	//check authorization
	cUser, _ := models.GetCurrentUser(c)
	if cUser.Role != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Only for admin level user"})
        return
	}

	var input PublisherInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	publisher := models.Publisher{Name: input.Name}
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	db.Create(&publisher)
	c.JSON(http.StatusOK, gin.H{"data": publisher})
}

// Update Publisher godoc
// @Summary Update existing publisher by id
// @Description Only admin have permission to update publisher
// @Tags Admin
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Publisher Id"
// @Param Body body PublisherInput true "the body to create new publisher"
// @Success 200 {object} models.Publisher
// @Router /publishers/{id} [patch]
func UpdatePublisher(c *gin.Context) {
	//check authorization
	cUser, _ := models.GetCurrentUser(c)
	if cUser.Role != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Only for admin level user"})
        return
	}

	db := c.MustGet("db").(*gorm.DB)
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
	updatedInputPublisher.UpdatedAt = time.Now()

	db.Model(&publisher).Updates(updatedInputPublisher)
	c.JSON(http.StatusOK, gin.H{"data": publisher})
}

// Delete a Publisher godoc
// @Summary Delete existing publisher by id
// @Description Only admin have permission to delete publisher
// @Tags Admin
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Publisher Id"
// @Success 200 {object} map[string]boolean
// @Router /publishers/{id} [delete]
func DeletePublisher(c *gin.Context) {
	//check authorization
	cUser, _ := models.GetCurrentUser(c)
	if cUser.Role != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Only for admin level user"})
        return
	}
	
	db := c.MustGet("db").(*gorm.DB)

	var publisher models.Publisher
	if err := db.Where("id = ?", c.Param("id")).First(&publisher).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	db.Delete(&publisher)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
