package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/renatodaltiba/rest-api-go/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		books := main.Group("/books")
		{
			books.GET("/:id", controllers.ShowBook)
			books.POST("/", controllers.CreateBook)
			books.GET("/", controllers.ShowBooks)
			books.PUT("/:id", controllers.UpdateBook)
			books.DELETE("/:id", controllers.DeleteBook)
		}
	}

	return router
}
