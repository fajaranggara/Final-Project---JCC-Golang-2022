package controllers

import (
	"final-project/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CategoryInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Get All Category godoc
// @Summary Get all category
// @Description Get list of category
// @Tags Public
// @Produce json
// @Success 200 {object} []models.Category
// @Router /categories [get]
func GetAllCategory(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	var categories []models.Category

	db.Find(&categories)

	c.JSON(http.StatusOK, gin.H{"data": categories})

}

// Get Games by Category godoc
// @Summary Get list of games in specific category
// @Description Get all games of spesific category by id
// @Tags Public
// @Produce json
// @Param id path string true "Category Id"
// @Success 200 {object} []models.Game
// @Router /categories/{id}/games [get]
func GetGamesByCategoryId(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	var games []models.Game
	if err := db.Where("category_id = ?", c.Param("id")).Find(&games).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": games})
}

// Create Category godoc
// @Summary Create a new category
// @Description Only admin have permission to create category
// @Tags Admin
// @Param Body body CategoryInput true "the body to create new category"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Category
// @Router /admin/add-categories [post]
func CreateCategory(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	//check authorization
	cUser, _ := models.GetCurrentUser(c)
	if cUser.Role != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"forbidden": "Allowed role: admin"})
        return
	}

	var input CategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category := models.Category{Name: input.Name, Description: input.Description}

	db.Create(&category)
	c.JSON(http.StatusOK, gin.H{"data": category})
}

// Update Category godoc
// @Summary Update existing category by id
// @Description Only admin have permission to update category
// @Tags Admin
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Category Id"
// @Param Body body CategoryInput true "the body to create new category"
// @Success 200 {object} models.Category
// @Router /admin/categories/{id} [patch]
func UpdateCategory(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	//check authorization
	cUser, _ := models.GetCurrentUser(c)
	if cUser.Role != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"forbidden": "Allowed role: admin"})
        return
	}

	var category models.Category
	if err := db.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	var input CategoryInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInputCategory models.Category

	updatedInputCategory.Name = input.Name
	updatedInputCategory.Description = input.Description
	updatedInputCategory.UpdatedAt = time.Now()

	db.Model(&category).Updates(updatedInputCategory)
	c.JSON(http.StatusOK, gin.H{"data": category})
}

// Delete a Category godoc
// @Summary Delete existing category by id
// @Description Only admin have permission to delete category
// @Tags Admin
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Category Id"
// @Success 200 {object} map[string]boolean
// @Router /admin/categories/{id} [delete]
func DeleteCategory(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	//check authorization
	cUser, _ := models.GetCurrentUser(c)
	if cUser.Role != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"forbidden": "Allowed role: admin"})
        return
	}

	var category models.Category
	if err := db.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	db.Delete(&category)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
