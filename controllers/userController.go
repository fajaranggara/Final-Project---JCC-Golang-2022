package controllers

import (
	"final-project/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RegisPublisherInput struct {
	LogoUrl			string `json:"logo_url"`
}

type ChangePasswordInput struct {
    CurrentPassword string `json:"current_password" binding:"required"`
    NewPassword string `json:"new_password" binding:"required"`
}


// Get User Profile godoc
// @Summary Get info of current login account
// @Description Get user, publisher, or admin info
// @Tags Authentication & Authorization
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /profiles [get]
func GetUserProfile(c *gin.Context){
	// get current user
	usr, err := models.GetCurrentUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You need to sign in"})
		return
	}

	usr.HidePassword()

	c.JSON(http.StatusOK, gin.H{"data": usr})
}

// Change password godoc
// @Summary Change password account.
// @Description Used by User, Publisher, or Admin for change password account.
// @Tags Authentication & Authorization
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param Body body ChangePasswordInput true "the body to change user password"
// @Success 200 {object} map[string]interface{}
// @Router /change-password [patch]
func ChangePassword(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
    
    // get current user
	usr, err := models.GetCurrentUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You need to sign in"})
		return
	}

    var input ChangePasswordInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    // check current password if true
    if err := models.VerifyPassword(input.CurrentPassword, usr.Password); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Current password is wrong"})
		return
    }

    hashedNewPassword, errPassword := bcrypt.GenerateFromPassword([]byte(input.NewPassword), bcrypt.DefaultCost)
	if errPassword != nil {
		return
	}
    newPassword := string(hashedNewPassword)
    newUser := models.User{}

    newUser.Username = usr.Username
    newUser.Email    = usr.Email
    newUser.Password = newPassword
    newUser.UpdatedAt= time.Now()

    db.Model(&usr).Updates(newUser)
    c.JSON(http.StatusOK, gin.H{"message": "change password success"})
}


// Become Publisher godoc
// @Summary Change role into publisher.
// @Description Create an account before change your role into publisher. User or Admin can change into this role
// @Tags Authentication & Authorization
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param Body body RegisPublisherInput true "the body to become publisher"
// @Success 200 {object} map[string]interface{}
// @Router /regist-publisher [patch]
func RegisPublisher(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)

	//check authorization
	cUser, err := models.GetCurrentUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You need to sign in"})
		return
	}
	if cUser.Role == "publisher" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You're already publisher"})
		return
	}

	var input RegisPublisherInput
	if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	// update current user role
	updatedUser := models.User{}
	updatedUser.Role = "publisher"
	updatedUser.UpdatedAt = time.Now()
	db.Model(&cUser).Updates(updatedUser)

	// create publisher
	publisher := models.Publisher{
		Name: cUser.Username,
		LogoURL: input.LogoUrl,
		UserID: int(cUser.ID),
	}
	db.Create(&publisher)
	c.JSON(http.StatusOK, gin.H{"message": "You're now a publisher"})
}

// Become Admin godoc
// @Summary Change role into admin.
// @Description Create an account before chhange your role. User or Publisher can change into this role (temporary func)
// @Tags Authentication & Authorization
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param Body body RegisPublisherInput true "the body to become admin"
// @Success 200 {object} map[string]interface{}
// @Router /regist-admin [patch]
func RegisAdmin(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)

	//check authorization
	cUser, err := models.GetCurrentUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You need to sign in"})
		return
	}
	if cUser.Role == "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You're already admin"})
		return
	}

	// update current user role
	updatedUser := models.User{}
	updatedUser.Role = "admin"
	updatedUser.UpdatedAt = time.Now()
	db.Model(&cUser).Updates(updatedUser)

	c.JSON(http.StatusOK, gin.H{"message": "You're now an admin"})
}
