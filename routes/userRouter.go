package routes

import (
	"github.com/Micah-Shallom/modules/controllers"
	"github.com/Micah-Shallom/modules/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate()) //authenticates users before allowing access to the below routes
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/users/:user_id", controllers.GetUser())
}