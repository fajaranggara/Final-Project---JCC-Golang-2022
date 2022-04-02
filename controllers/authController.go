package controllers

import (
	"final-project/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Get User Profile godoc
// @Summary Get info of current login user
// @Description Get logged in user info
// @Tags User
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /users/profile [get]
func GetUserProfile(c *gin.Context){
	// get current user
	usr, err := models.GetCurrentUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You need to sign in to add review"})
		return
	}

	usr.HidePassword()

	c.JSON(http.StatusOK, gin.H{"data": usr})
}


type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type ChangePasswordInput struct {
    CurrentPassword string `json:"current_password" binding:"required"`
    NewPassword string `json:"new_password" binding:"required"`
}

// LoginUser godoc
// @Summary Login as an user.
// @Description Logging in to get jwt token to access admin or user api by roles.
// @Tags Auth
// @Param Body body LoginInput true "the body to login a user"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /login [post]
func Login(c *gin.Context){
	db := c.MustGet("db").(*gorm.DB)
    var input LoginInput

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    usr := models.User{}

    usr.Username = input.Username
    usr.Password = input.Password

    token, err := models.LoginCheck(usr.Username, usr.Password, db)

    if err != nil {
        fmt.Println(err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
        return
    }

    user := map[string]string{
        "username": usr.Username,
        "email":    usr.Email,
    }

    c.JSON(http.StatusOK, gin.H{"message": "login success", "user": user, "token": token})
}

// Register godoc
// @Summary Register a user.
// @Description registering a user from public access.
// @Tags Auth
// @Param Body body RegisterInput true "the body to register a user"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /register [post]
func Register(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
    var input RegisterInput

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    usr := models.User{}

    usr.Username = input.Username
    usr.Email = input.Email
    usr.Password = input.Password
    usr.Role = "user"

    _, err := usr.SaveUser(db)

    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user := map[string]string{
        "username": input.Username,
        "email":    input.Email,
    }

    c.JSON(http.StatusOK, gin.H{"message": "registration success", "user": user})
}


// Change password godoc
// @Summary Change users password.
// @Description renew users password.
// @Tags User
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param Body body ChangePasswordInput true "the body to change user password"
// @Success 200 {object} map[string]interface{}
// @Router /users/change-password [patch]
func ChangePassword(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
    
    // get current user
	usr, err := models.GetCurrentUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You need to sign in to add review"})
		return
	}

    var input ChangePasswordInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    // check current password if true
    if err := models.VerifyPassword(input.CurrentPassword, usr.Password); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong password"})
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
