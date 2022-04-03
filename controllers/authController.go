package controllers

import (
	"final-project/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

// LoginUser godoc
// @Summary Login for user, publisher, or admin.
// @Description Logging in to get jwt token to access admin or user api by roles.
// @Tags Authentication & Authorization
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
    }

    c.JSON(http.StatusOK, gin.H{"message": "login success", "user": user, "token": token})
}

// Register godoc
// @Summary Create an account.
// @Description Create account, role: user.
// @Tags Authentication & Authorization
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
