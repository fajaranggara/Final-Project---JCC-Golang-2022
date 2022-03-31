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

// Get all Category godoc
// @Summary Get all Category
// @Description Get list of Category
// @Tags Category
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

// Create Category godoc
// @Summary Create a Category
// @Description Create new Category
// @Tags Category
// @Param Body body CategoryInput true "the body to create new category"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Category
// @Router /categories [post]
func CreateCategory(c *gin.Context) {
	var input CategoryInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category := models.Category{Name: input.Name, Description: input.Description}
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	db.Create(&category)
	c.JSON(http.StatusOK, gin.H{"data": category})
}

// Get Category godoc
// @Summary Get Category by id
// @Description Get one Category by id
// @Tags Category
// @Produce json
// @Param id path string true "Category Id"
// @Success 200 {object} models.Category
// @Router /categories/{id} [get]
func GetCategoryById(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	// check if exist and get data
	var category models.Category
	if err := db.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": category})
}

// Get games from one Category godoc
// @Summary Get games by Category by id
// @Description Get all games of spesific Category by id
// @Tags Category
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

// Update Category godoc
// @Summary update a Category by id
// @Description update one Category by id
// @Tags Category
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Category Id"
// @Param Body body CategoryInput true "the body to create new category"
// @Success 200 {object} models.Category
// @Router /categories/{id} [patch]
func UpdateCategory(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	// get rating if exist
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
// @Summary delete a Category by id
// @Description delete one Category by id
// @Tags Category
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Category Id"
// @Success 200 {object} map[string]boolean
// @Router /categories/{id} [delete]
func DeleteCategory(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	var category models.Category
	if err := db.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	db.Delete(&category)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
