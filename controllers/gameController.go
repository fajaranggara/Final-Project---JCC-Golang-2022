package controllers

import (
	"final-project/models"
	"math"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AddGameInput struct {
	Name        string 		`json:"name"`
	Description string 		`json:"description"`
	ImageURL 	string 		`json:"image_url"`
	GenreID		int			`json:"genre_id"`
	CategoryID	int			`json:"category_id"`
}
type UpdateGameInput struct {
	Name        string 		`json:"name"`
	Description string 		`json:"description"`
	ImageURL 	string 		`json:"image_url"`
	GenreID		int			`json:"genre_id"`
	CategoryID	int			`json:"category_id"`
}

// Get all Game godoc
// @Summary Get all Game
// @Description Get list of Game
// @Tags Public
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

// Get Game by ID godoc
// @Summary Get Game by id
// @Description Get one game by id
// @Tags Public
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

// Create Games godoc
// @Summary Create a new games
// @Description Only publisher and admin have permission to create games
// @Tags Publisher
// @Param Body body AddGameInput true "the body to create new games"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Game
// @Router /publisher/add-games [post]
func CreateGame(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	//check authorization
	cUser, _ := models.GetCurrentUser(c)
	if cUser.Role != "publisher" {
		c.JSON(http.StatusBadRequest, gin.H{"forbidden": "Allowed role: publisher"})
        return
	}

	var input AddGameInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var publisher models.Publisher
	if err := db.Where("user_id = ?", cUser.ID).First(&publisher).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	game := models.Game{Name: input.Name,
		Ratings: 0,
		RatingsCounter: 0, 
		ReleaseDate: time.Now(),
		Description: input.Description, 
		ImageURL: input.ImageURL,
		GenreID: input.GenreID, 
		CategoryID: input.CategoryID,
		PublisherID: publisher.ID,
	}

	db.Create(&game)
	c.JSON(http.StatusOK, gin.H{"data": game})
}

// Update Game godoc
// @Summary Update existing game by id
// @Description Only admin have permission to update game
// @Tags Publisher
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Game Id"
// @Param Body body UpdateGameInput true "the body to create new game"
// @Success 200 {object} models.Game
// @Router /publisher/games/{id} [patch]
func UpdateGame(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	//check authorization
	cUser, _ := models.GetCurrentUser(c)
	if cUser.Role != "publisher" {
		c.JSON(http.StatusBadRequest, gin.H{"forbidden": "Allowed role: publisher"})
        return
	}

	var game models.Game
	if err := db.Where("id = ?", c.Param("id")).First(&game).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	// check if current user is publisher of this games
	uidPublisher, err := getUserIdByPublisherId(game.PublisherID, db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if uidPublisher != int(cUser.ID) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You don't have permission to edit this games"})
		return
	}

	var input UpdateGameInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInputGame models.Game

	updatedInputGame.Name = input.Name
	updatedInputGame.Description = input.Description
	updatedInputGame.ImageURL = input.ImageURL
	updatedInputGame.GenreID = input.GenreID
	updatedInputGame.CategoryID = input.CategoryID
	updatedInputGame.UpdatedAt = time.Now()

	db.Model(&game).Updates(updatedInputGame)
	c.JSON(http.StatusOK, gin.H{"data": game})
}

// Delete a Game godoc
// @Summary Delete existing game by id
// @Description Only admin have permission to delete game
// @Tags Publisher
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Game Id"
// @Success 200 {object} map[string]boolean
// @Router /publisher/games/{id} [delete]
func DeleteGame(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	//check authorization
	cUser, _ := models.GetCurrentUser(c)
	if cUser.Role != "publisher" {
		c.JSON(http.StatusBadRequest, gin.H{"forbidden": "Allowed role: publisher"})
        return
	}
	
	var game models.Game
	if err := db.Where("id = ?", c.Param("id")).First(&game).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	// check if current user is publisher of this games
	uidPublisher, err := getUserIdByPublisherId(game.PublisherID, db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if uidPublisher != int(cUser.ID) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You don't have permission to delete this games"})
		return
	}

	db.Delete(&game)

	c.JSON(http.StatusOK, gin.H{"data": true})
}


func CalculateRating(game *models.Game, newRate int) int {
	counter := game.RatingsCounter + 1

	rating := ((float64(game.Ratings) * float64(game.RatingsCounter)) + float64(newRate)) / float64(counter)
	return int(math.Round(rating))
}

func getUserIdByPublisherId(publisherId int, db *gorm.DB) (int, error) {
	var publisher models.Publisher
	if err := db.Where("user_id = ?", publisherId).First(&publisher).Error; err != nil {
		return 0, err
	}

	return publisher.UserID, nil
}