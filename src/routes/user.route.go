package routes

import (
	controllers "gin-mongo/src/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(routes *gin.Engine) {
	routes.POST("/", controllers.CreateUser)
	routes.GET("/", controllers.GetAllUsers)
	routes.GET("/:id", controllers.GetUserById)
	routes.PUT("/:id", controllers.UpdateUser)
	routes.DELETE("/:id", controllers.DeleteUserById)
}
