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
		// auth
		r.POST("/login", controllers.Login)
		r.POST("/register", controllers.Register) // default role is "user"

		// public
		r.GET("/categories", controllers.GetAllCategory)
		r.GET("/genres", controllers.GetAllGenre)
		r.GET("/publishers", controllers.GetAllPublisher)
		r.GET("/games", controllers.GetAllGame)
		r.GET("/games/:id", controllers.GetGameById)
		r.GET("/games/:id/reviews", controllers.GetGamesReview)

		// find games by
		r.GET("/categories/:id/games", controllers.GetGamesByCategoryId)
		r.GET("/genres/:id/games", controllers.GetGamesByGenreId)
		r.GET("/publishers/:id/games", controllers.GetGamesByPublisherId)
	
	}

	userMiddleware := r.Group("/users").Use(middlewares.JwtAuthMiddleware())
	reviewMiddleware := r.Group("/reviews").Use(middlewares.JwtAuthMiddleware())
	categoryMiddleware := r.Group("/categories").Use(middlewares.JwtAuthMiddleware())
	genreMiddleware := r.Group("/genres").Use(middlewares.JwtAuthMiddleware())
	publisherMiddleware := r.Group("/publishers").Use(middlewares.JwtAuthMiddleware())
	gamesMiddleware := r.Group("/games").Use(middlewares.JwtAuthMiddleware())

	// USER LEVEL: can be access by user with role{"user", "admin"}
	{
		// users
		userMiddleware.PATCH("/change-password", controllers.ChangePassword)
		userMiddleware.GET("/profiles", controllers.GetUserProfile)
		userMiddleware.GET("/bookmarks", controllers.ShowUserBookmark)
		userMiddleware.DELETE("/bookmarks/:id", controllers.DeleteBookmarkedGame)
		userMiddleware.PATCH("/regist-publisher", controllers.RegisPublisher)
		
		
		// games
		gamesMiddleware.POST("/:id/add-reviews", controllers.AddReview)
		reviewMiddleware.PATCH("/:id", controllers.UpdateReview)
		reviewMiddleware.DELETE("/:id", controllers.DeleteReview)
		gamesMiddleware.PATCH("/:id/add-to-bookmark", controllers.AddGameToBookmark)
	}

	// PUBLISHER LEVEL: can be access by userr with role{"publisher", "admin"}
	{
		gamesMiddleware.POST("/", controllers.CreateGame)
		gamesMiddleware.PATCH("/:id", controllers.UpdateGame)
		gamesMiddleware.DELETE("/:id", controllers.DeleteGame)

	}
	
	// ADMIN LEVEL: can be access by user with role{"admin"}
	{// admin
		categoryMiddleware.POST("/", controllers.CreateCategory)
		categoryMiddleware.PATCH("/:id", controllers.UpdateCategory)
		categoryMiddleware.DELETE("/:id", controllers.DeleteCategory)
		
		genreMiddleware.POST("/", controllers.CreateGenre)
		genreMiddleware.PATCH("/:id", controllers.UpdateGenre)
		genreMiddleware.DELETE("/:id", controllers.DeleteGenre)

		publisherMiddleware.POST("/", controllers.CreatePublisher)
		publisherMiddleware.PATCH("/:id", controllers.UpdatePublisher)
		publisherMiddleware.DELETE("/:id", controllers.DeletePublisher)
	}

	
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
