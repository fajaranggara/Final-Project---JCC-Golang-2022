package controllers

import (
	"final-project/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AddReviewInput struct {
	Rate  		int    	  `json:"rate"`
	Content		string    `json:"content"`
	UserID		int		  `json:"user_id"`
}
type UpdateReviewInput struct {
	Rate  		int    	  `json:"rate"`
	Content		string    `json:"content"`
}

// Get Reviews from Game godoc
// @Summary Get reviews of a game by gameID
// @Description Get all reviews of spesific game
// @Tags Game
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
// @Description Create new review
// @Tags Review
// @Param Body body AddReviewInput true "the body to create new review"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Game Id"
// @Success 200 {object} models.Review
// @Router /games/:id/reviews [post]
func AddReview(c *gin.Context) {
	var input AddReviewInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// get current user
	usr, err := models.GetCurrentUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You need to sign in to add review"})
		return
	}

	gameId,_ := strconv.Atoi(c.Param("id"))
	review := models.Review{Rate: input.Rate, Content: input.Content, GameID: gameId, UserID: int(usr.ID)}
	
	db := c.MustGet("db").(*gorm.DB)

	var game models.Game
	if err := db.Where("id = ?", c.Param("id")).First(&game).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	var updateGameRating models.Game
	updateGameRating.Ratings = CalculateRating(&game, input.Rate)
	updateGameRating.RatingsCounter = game.RatingsCounter + 1
	updateGameRating.UpdatedAt = time.Now()

	db.Model(&game).Updates(updateGameRating)
	db.Create(&review)
	c.JSON(http.StatusOK, gin.H{"add": review, "game": game})

}

// Update Review godoc
// @Summary Update existing review by id
// @Description Only user who create this review have permission to update
// @Tags Review
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Review Id"
// @Param Body body UpdateReviewInput true "the body to create new review"
// @Success 200 {object} models.Review
// @Router /reviews/{id} [patch]
func UpdateReview(c *gin.Context) {	
	db := c.MustGet("db").(*gorm.DB)

	var review models.Review
	if err := db.Where("id = ?", c.Param("id")).First(&review).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	// check if current user is the same user who create this review
	usr, _ := models.GetCurrentUser(c)
	if review.UserID != int(usr.ID) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You don't have permission to edit this review"})
		return
	}

	var input UpdateReviewInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInputReview models.Review

	updatedInputReview.Rate = input.Rate
	updatedInputReview.Content = input.Content
	updatedInputReview.UpdatedAt = time.Now()

	db.Model(&review).Updates(updatedInputReview)
	c.JSON(http.StatusOK, gin.H{"data": review})
}

// Delete Review godoc
// @Summary Delete existing review by id
// @Description Only user who create this review have permission to update
// @Tags Review
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Review Id"
// @Success 200 {object} map[string]boolean
// @Router /reviews/{id} [delete]
func DeleteReview(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var review models.Review
	if err := db.Where("id = ?", c.Param("id")).First(&review).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	// check if current user is the same user who create this review
	usr, _ := models.GetCurrentUser(c)
	if review.UserID != int(usr.ID) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You don't have permission to delete this review"})
		return
	}

	db.Delete(&review)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
