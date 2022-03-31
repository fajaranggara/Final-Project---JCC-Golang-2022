package controllers

import (
	"final-project/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GameInput struct {
	Name        string 		`json:"name"`
	ReleaseDate string    	`json:"release_date"`
	Price 		int 		`json:"price"`
	Description string 		`json:"description"`
	ImageURL 	string 		`json:"image_url"`
	CategoryID	int			`json:"category_id"`
}

// Get all Game godoc
// @Summary Get all Game
// @Description Get list of Game
// @Tags Game
// @Produce json
// @Success 200 {object} []models.Game
// @Router /games [get]
func GetAllGame(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	var games []models.Game

	db.Find(&games)

	c.JSON(http.StatusOK, gin.H{"data": games})

}

// Create Game godoc
// @Summary Create a Game
// @Description Create new Game
// @Tags Game
// @Param Body body GameInput true "the body to create new game"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Game
// @Router /games [post]
func CreateGame(c *gin.Context) {
	var input GameInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	game := models.Game{Name: input.Name, 
		ReleaseDate: input.ReleaseDate, 
		Price: input.Price,
		Description: input.Description, 
		ImageURL: input.ImageURL, 
		CategoryID: input.CategoryID,
	}
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	db.Create(&game)
	c.JSON(http.StatusOK, gin.H{"data": game})
}

// Get Game godoc
// @Summary Get Game by id
// @Description Get one Game by id
// @Tags Game
// @Produce json
// @Param id path string true "Game Id"
// @Success 200 {object} models.Game
// @Router /games/{id} [get]
func GetGameById(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	// check if exist and get data
	var game models.Game
	if err := db.Where("id = ?", c.Param("id")).First(&game).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": game})
}

// Update Game godoc
// @Summary update a Game by id
// @Description update one Game by id
// @Tags Game
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Game Id"
// @Param Body body GameInput true "the body to create new game"
// @Success 200 {object} models.Game
// @Router /games/{id} [patch]
func UpdateGame(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	// get rating if exist
	var game models.Game
	if err := db.Where("id = ?", c.Param("id")).First(&game).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	var input GameInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInputGame models.Game

	updatedInputGame.Name = input.Name
	updatedInputGame.ReleaseDate = input.ReleaseDate
	updatedInputGame.Price = input.Price
	updatedInputGame.Description = input.Description
	updatedInputGame.ImageURL = input.ImageURL
	updatedInputGame.CategoryID = input.CategoryID
	updatedInputGame.UpdatedAt = time.Now()

	db.Model(&game).Updates(updatedInputGame)
	c.JSON(http.StatusOK, gin.H{"data": game})
}

// Delete a Game godoc
// @Summary delete a Game by id
// @Description delete one Game by id
// @Tags Game
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Game Id"
// @Success 200 {object} map[string]boolean
// @Router /games/{id} [delete]
func DeleteGame(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	var game models.Game
	if err := db.Where("id = ?", c.Param("id")).First(&game).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	db.Delete(&game)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
