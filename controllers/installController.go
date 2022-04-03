package controllers

import (
	"final-project/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Get Installed Games from User godoc
// @Summary Get list of installed games
// @Description Get all installed games by user
// @Tags Users
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} []models.InstalledGames
// @Router /users/my-games [get]
func ShowInstalledGames(c *gin.Context) {
	//check authorization
	cUser, _ := models.GetCurrentUser(c)
	if cUser.Role != "user" {
		c.JSON(http.StatusBadRequest, gin.H{"forbidden": "Allowed role: user"})
		return
	}

	db := c.MustGet("db").(*gorm.DB)

	var installed []models.InstalledGames
	if err := db.Where("user_id = ?", cUser.ID).Find(&installed).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": installed})
}

// Install godoc
// @Summary Install a games
// @Description User installing games
// @Tags Users
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Game Id"
// @Success 200 {object} models.InstalledGames
// @Router /games/{id}/install [patch]
func InstallThisGames(c *gin.Context) {
	//check authorization
	cUser, _ := models.GetCurrentUser(c)
	if cUser.Role != "user" {
		c.JSON(http.StatusBadRequest, gin.H{"forbidden": "Allowed role: user"})
		return
	}

	db := c.MustGet("db").(*gorm.DB)

	var game models.Game
	if err := db.Where("id = ?", c.Param("id")).First(&game).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found", "err": c.Param("id")})
		return
	}

	var install models.InstalledGames

	install.GameName = game.Name
	install.IdGame = game.ID
	install.Ratings = game.Ratings
	install.ImageURL = game.ImageURL
	install.UserID = int(cUser.ID)

	db.Create(&install)
	c.JSON(http.StatusOK, gin.H{"data": install})
}

// Uninstall godoc
// @Summary Uninstall a games
// @Description Only user who have permission can uninstall this game
// @Tags Users
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "InstalledGames Id"
// @Success 200 {object} map[string]boolean
// @Router /users/installed/{id} [delete]
func UninstallGame(c *gin.Context) {
	//check authorization
	cUser, _ := models.GetCurrentUser(c)
	if cUser.Role != "user" {
		c.JSON(http.StatusBadRequest, gin.H{"forbidden": "Allowed role: user"})
		return
	}
	
	// get installed games info
	db := c.MustGet("db").(*gorm.DB)
	var install models.InstalledGames
	if err := db.Where("id = ?", c.Param("id")).First(&install).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	// check if current user is installing this games
	if install.UserID != int(cUser.ID) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You don't have permission to uninstall this games"})
		return
	}

	

	db.Delete(&install)

	c.JSON(http.StatusOK, gin.H{"data": true})
}