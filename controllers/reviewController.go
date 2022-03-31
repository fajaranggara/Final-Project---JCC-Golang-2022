package controllers

import (
	"final-project/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ReviewInput struct {
	Rate  		int    	  `json:"rate"`
	Content		string    `json:"content"`
	GameID  	int		  `json:"game_id"`
	UserID		int		  `json:"user_id"`
}

// Get all Review godoc
// @Summary Get all Review
// @Description Get list of Review
// @Tags Review
// @Produce json
// @Success 200 {object} []models.Review
// @Router /reviews [get]
func GetAllReview(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	var reviews []models.Review

	db.Find(&reviews)

	c.JSON(http.StatusOK, gin.H{"data": reviews})

}


// Create Review godoc
// @Summary Create a Review
// @Description Create new Review
// @Tags Review
// @Param Body body ReviewInput true "the body to create new review"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Review
// @Router /reviews [post]
func AddReview(c *gin.Context) {
	var input ReviewInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	review := models.Review{Rate: input.Rate, Content: input.Content, GameID: input.GameID, UserID: input.UserID}
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	db.Create(&review)
	c.JSON(http.StatusOK, gin.H{"data": review})

}

// Update Review godoc
// @Summary update a Review by id
// @Description update one Review by id
// @Tags Review
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Review Id"
// @Param Body body ReviewInput true "the body to create new review"
// @Success 200 {object} models.Review
// @Router /reviews/{id} [patch]
func UpdateReview(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	// get rating if exist
	var review models.Review
	if err := db.Where("id = ?", c.Param("id")).First(&review).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	var input ReviewInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInputReview models.Review

	updatedInputReview.Rate = input.Rate
	updatedInputReview.Content = input.Content
	updatedInputReview.GameID = input.GameID
	updatedInputReview.UserID = input.UserID
	updatedInputReview.UpdatedAt = time.Now()

	db.Model(&review).Updates(updatedInputReview)
	c.JSON(http.StatusOK, gin.H{"data": review})
}

// Delete a Review godoc
// @Summary delete a Review by id
// @Description delete one Review by id
// @Tags Review
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Review Id"
// @Success 200 {object} map[string]boolean
// @Router /reviews/{id} [delete]
func DeleteReview(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	var review models.Review
	if err := db.Where("id = ?", c.Param("id")).First(&review).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	db.Delete(&review)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
