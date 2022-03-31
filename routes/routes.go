package routes

import (
	"final-project/controllers"
	"final-project/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"gorm.io/gorm"

	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// set db to gin context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.POST("/login", controllers.Login)
	r.POST("/register", controllers.Register)


	r.GET("/categories", controllers.GetAllCategory)
	r.GET("/categories/:id", controllers.GetCategoryById)
	r.GET("/categories/:id/games", controllers.GetGamesByCategoryId)
	
	// middleware for category
	categoryMiddleware := r.Group("/categories")
	categoryMiddleware.Use(middlewares.JwtAuthMiddleware())
	categoryMiddleware.POST("/", controllers.CreateCategory)
	categoryMiddleware.PATCH("/:id", controllers.UpdateCategory)
	categoryMiddleware.DELETE("/:id", controllers.DeleteCategory)
	// end

	r.GET("/games", controllers.GetAllGame)
	r.GET("/games/:id", controllers.GetGameById)

	// middleware for game
	gameMiddleware := r.Group("/games")
	gameMiddleware.Use(middlewares.JwtAuthMiddleware())
	gameMiddleware.POST("/", controllers.CreateGame)
	gameMiddleware.PATCH("/:id", controllers.UpdateGame)
	gameMiddleware.DELETE("/:id", controllers.DeleteGame)
	// end


	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
