package controllers

import (
	"final-project/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ReviewInput struct {
	Rate  		uint    	  `json:"rate"`
	Content		string    `json:"content"`
}

// Get Reviews from Games godoc
// @Summary Get reviews of specific games
// @Description Get all reviews of spesific games
// @Tags Public
// @Produce json
// @Param id path string true "Game Id"
// @Success 200 {object} []models.Review
// @Router /games/{id}/reviews [get]
func GetGamesReview(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	var reviews []models.Review
	if err := db.Where("game_id = ?", c.Param("id")).Find(&reviews).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": reviews})
}

// Create Review godoc
// @Summary Create a review
// @Description Create new review and rate(1-5)
// @Tags Users
// @Param Body body ReviewInput true "the body to create new review"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Game Id"
// @Success 200 {object} models.Review
// @Router /users/games/{id}/add-reviews [post]
func AddReview(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	//check authorization
	cUser, _ := models.GetCurrentUser(c)
	if cUser.Role != "user" {
		c.JSON(http.StatusBadRequest, gin.H{"forbidden": "Allowed role: user"})
		return
	}

	var input ReviewInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Rate > 5 || input.Rate < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Rate must in range 1 - 5"})
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	review := models.Review{
		Rate: int(input.Rate), 
		Content: input.Content, 
		GameID: id, 
		UserID: int(cUser.ID),
	}
	db.Create(&review)

	var game models.Game
	if err := db.Where("id = ?", c.Param("id")).First(&game).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	var updateGameRating models.Game
	updateGameRating.Ratings = CalculateRating(&game, int(input.Rate))
	updateGameRating.RatingsCounter = game.RatingsCounter + 1
	updateGameRating.UpdatedAt = time.Now()

	db.Model(&game).Updates(updateGameRating)

	c.JSON(http.StatusOK, gin.H{"add": review})
}

// Update Review godoc
// @Summary Update existing review by id
// @Description Only user who create this review have permission to update
// @Tags Users
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Review Id"
// @Param Body body ReviewInput true "the body to create new review"
// @Success 200 {object} models.Review
// @Router /users/games/reviews/{id} [patch]
func UpdateReview(c *gin.Context) {	
	db := c.MustGet("db").(*gorm.DB)
	//check authorization
	cUser, _ := models.GetCurrentUser(c)
	if cUser.Role != "user" {
		c.JSON(http.StatusBadRequest, gin.H{"forbidden": "Allowed role: user"})
		return
	}

	var review models.Review
	if err := db.Where("id = ?", c.Param("id")).First(&review).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	// check if current user is owner of this review
	if review.UserID != int(cUser.ID) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You don't have permission to edit this review"})
		return
	}

	var input ReviewInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInputReview models.Review

	updatedInputReview.Rate = int(input.Rate)
	updatedInputReview.Content = input.Content
	updatedInputReview.UpdatedAt = time.Now()

	db.Model(&review).Updates(updatedInputReview)
	c.JSON(http.StatusOK, gin.H{"data": review})
}

// Delete Review godoc
// @Summary Delete existing review by id
// @Description Only user who create this review have permission to update
// @Tags Users
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Review Id"
// @Success 200 {object} map[string]boolean
// @Router /users/games/reviews/{id} [delete]
func DeleteReview(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	//check authorization
	cUser, _ := models.GetCurrentUser(c)
	if cUser.Role != "user" {
		c.JSON(http.StatusBadRequest, gin.H{"forbidden": "Allowed role: user"})
		return
	}

	var review models.Review
	if err := db.Where("id = ?", c.Param("id")).First(&review).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	// check if current user is owner of this review
	if review.UserID != int(cUser.ID) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You don't have permission to delete this review"})
		return
	}

	db.Delete(&review)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
