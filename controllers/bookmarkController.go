package controllers

import (
	"final-project/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Get Bookmarked Games from User godoc
// @Summary Get list of bookmarked games
// @Description Get all games in users bookmark
// @Tags Users
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} []models.Bookmark
// @Router /users/bookmarks [get]
func ShowUserBookmark(c *gin.Context) {
	cUser, _ := models.GetCurrentUser(c)
	if cUser.Role != "user" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You need to login"})
        return
	}

	db := c.MustGet("db").(*gorm.DB)

	var bookmarks []models.Bookmark
	if err := db.Where("user_id = ?", cUser.ID).Find(&bookmarks).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bookmarks})
}

// Bookmark godoc
// @Summary Bookmarked games
// @Description User add games to bookmark
// @Tags Games
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Game Id"
// @Success 200 {object} models.Bookmark
// @Router /games/{id}/add-to-bookmark [patch]
func AddGameToBookmark(c *gin.Context) {
	//check authorization
	cUser, err := models.GetCurrentUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You need to login"})
        return
	}

	db := c.MustGet("db").(*gorm.DB)

	var game models.Game
	if err := db.Where("id = ?", c.Param("id")).First(&game).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found", "err": c.Param("id")})
		return
	}

	var bookmark models.Bookmark

	bookmark.GameName = game.Name
	bookmark.IdGame = game.ID
	bookmark.Ratings = game.Ratings
	bookmark.ImageURL = game.ImageURL
	bookmark.UserID = int(cUser.ID)

	db.Create(&bookmark)
	c.JSON(http.StatusOK, gin.H{"data": bookmark})
}

// Delete Bookmark godoc
// @Summary Delete games in users bookmark
// @Description Only user who have permission can delete this bookmark
// @Tags Users
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Bookmark Id"
// @Success 200 {object} map[string]boolean
// @Router /users/bookmarks/{id} [delete]
func DeleteBookmarkedGame(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var bookmark models.Bookmark
	if err := db.Where("id = ?", c.Param("id")).First(&bookmark).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	// check if current user is the same user who create this review
	usr, _ := models.GetCurrentUser(c)
	if bookmark.UserID != int(usr.ID) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You don't have permission to delete this review"})
		return
	}

	db.Delete(&bookmark)

	c.JSON(http.StatusOK, gin.H{"data": true})
}