package routes

import (
	"github.com/kon2300/go-echo-vue-nuxt/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/users", controllers.GetAllUsers)
	e.POST("/users", controllers.CreateUser)
	e.GET("/users/:id", controllers.GetUser)
	e.PUT("/users/:id", controllers.UpdateUser)
	e.DELETE("/users/:id", controllers.DeleteUser)

	e.GET("/articles", controllers.GetAllArticles)
	e.POST("/articles", controllers.CreateArticle)
	e.GET("/articles/:id", controllers.GetArticle)
	e.PUT("/articles/:id", controllers.UpdateArticle)
	e.DELETE("/articles/:id", controllers.DeleteArticle)

	e.POST("/article-like", controllers.CreateArticleLike)
	e.DELETE("/article-like/:id", controllers.DeleteArticleLike)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
