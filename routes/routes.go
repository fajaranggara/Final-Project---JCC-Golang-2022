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

	// PUBLIC LEVEL: no login needed
	{
		r.GET("/games", controllers.GetAllGame)
		r.GET("/games/:id", controllers.GetGameById)
		r.GET("/games/:id/reviews", controllers.GetGamesReview)

		r.GET("/categories/:id/games", controllers.GetGamesByCategoryId)
		r.GET("/categories", controllers.GetAllCategory)
		r.GET("/genres/:id/games", controllers.GetGamesByGenreId)
		r.GET("/genres", controllers.GetAllGenre)
		r.GET("/publishers/:id/games", controllers.GetGamesByPublisherId)
		r.GET("/publishers", controllers.GetAllPublisher)
	
	}


	r.POST("/login", controllers.Login)
	r.POST("/register", controllers.Register) // default role: user
	
	r.GET("/profiles", controllers.GetUserProfile, middlewares.JwtAuthMiddleware())
	r.PATCH("/change-password", controllers.ChangePassword, middlewares.JwtAuthMiddleware())
	r.PATCH("/regist-publisher", controllers.RegisPublisher, middlewares.JwtAuthMiddleware())
	r.PATCH("/regist-admin", controllers.RegisAdmin, middlewares.JwtAuthMiddleware())

	// USER LEVEL: can be access by user with role{"user"}
	user := r.Group("/users").Use(middlewares.JwtAuthMiddleware())
	{
		user.PATCH("/games/:id/add-to-bookmark", controllers.AddGameToBookmark)
		user.GET("/bookmarks", controllers.ShowUserBookmark)
		user.DELETE("/bookmarks/:id", controllers.DeleteBookmarkedGame)

		user.PATCH("/games/:id/install", controllers.InstallThisGames)
		user.GET("/my-games", controllers.ShowInstalledGames)
		user.DELETE("/installed/:id", controllers.UninstallGame)

		user.POST("/games/:id/add-reviews", controllers.AddReview)
		user.PATCH("/games/reviews/:id", controllers.UpdateReview)
		user.DELETE("/games/reviews/:id", controllers.DeleteReview)

	}


	// PUBLISHER LEVEL: can be access by userr with role{"publisher"}
	publisher := r.Group("/publisher").Use(middlewares.JwtAuthMiddleware())
	{
		publisher.POST("/add-games", controllers.CreateGame)
		publisher.PATCH("/games/:id", controllers.UpdateGame)
		publisher.DELETE("/games/:id", controllers.DeleteGame)
	}
	
	// ADMIN LEVEL: can be access by user with role{"admin"}
	admin := r.Group("/admin").Use(middlewares.JwtAuthMiddleware())
	{
		admin.POST("/add-categories", controllers.CreateCategory)
		admin.PATCH("/categories/:id", controllers.UpdateCategory)
		admin.DELETE("/categories/:id", controllers.DeleteCategory)
		
		admin.POST("/add-genres", controllers.CreateGenre)
		admin.PATCH("/genres/:id", controllers.UpdateGenre)
		admin.DELETE("/genres/:id", controllers.DeleteGenre)
	}

	
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
