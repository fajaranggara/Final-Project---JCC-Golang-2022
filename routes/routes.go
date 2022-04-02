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
		r.POST("/login", controllers.Login)
		r.POST("/register", controllers.Register) // default role is "user"

		r.GET("/categories", controllers.GetAllCategory)
		r.GET("/categories/:id", controllers.GetCategoryById)
		r.GET("/categories/:id/games", controllers.GetGamesByCategoryId)

		r.GET("/genres", controllers.GetAllGenre)
		r.GET("/genres/:id", controllers.GetGenreById)
		r.GET("/genres/:id/games", controllers.GetGamesByGenreId)

		r.GET("/publishers", controllers.GetAllPublisher)
		r.GET("/publishers/:id", controllers.GetPublisherById)
		r.GET("/publishers/:id/games", controllers.GetGamesByPublisherId)

		r.GET("/games", controllers.GetAllGame)
		r.GET("/games/:id", controllers.GetGameById)
		r.GET("/games/:id/reviews", controllers.GetGamesReview)
		
	}

	// USER LEVEL: can access by user with role{"user", "admin"}
	{
		r.PATCH("/users/:id/change-password", controllers.ChangePassword, middlewares.JwtAuthMiddleware())
	
		r.POST("/:id/reviews", controllers.AddReview, middlewares.JwtAuthMiddleware())

		reviewMiddleware := r.Group("/reviews")
		reviewMiddleware.Use(middlewares.JwtAuthMiddleware())
		// Can access by user who create this review
		{
			reviewMiddleware.PATCH("/:id", controllers.UpdateReview)
			reviewMiddleware.DELETE("/:id", controllers.DeleteReview)
		}
		
	}
	
	// ADMIN LEVEL: can access by user with role{"admin"}
	{
		categoryMiddleware := r.Group("/categories")
		categoryMiddleware.Use(middlewares.JwtAuthMiddleware())
		categoryMiddleware.POST("/", controllers.CreateCategory)
		categoryMiddleware.PATCH("/:id", controllers.UpdateCategory)
		categoryMiddleware.DELETE("/:id", controllers.DeleteCategory)

		genreMiddleware := r.Group("/genres")
		genreMiddleware.Use(middlewares.JwtAuthMiddleware())
		genreMiddleware.POST("/", controllers.CreateGenre)
		genreMiddleware.PATCH("/:id", controllers.UpdateGenre)
		genreMiddleware.DELETE("/:id", controllers.DeleteGenre)

		publisherMiddleware := r.Group("/publishers")
		publisherMiddleware.Use(middlewares.JwtAuthMiddleware())
		publisherMiddleware.POST("/", controllers.CreatePublisher)
		publisherMiddleware.PATCH("/:id", controllers.UpdatePublisher)
		publisherMiddleware.DELETE("/:id", controllers.DeletePublisher)

		gameMiddleware := r.Group("/games")
		gameMiddleware.Use(middlewares.JwtAuthMiddleware())
		gameMiddleware.POST("/", controllers.CreateGame)
		gameMiddleware.PATCH("/:id", controllers.UpdateGame)
		gameMiddleware.DELETE("/:id", controllers.DeleteGame)
	}


	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
