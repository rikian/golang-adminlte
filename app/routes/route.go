package routes

import (
	"golte/app/controllers"

	"github.com/gin-gonic/gin"
)

func Route(route *gin.Engine) {
	route.GET("", controllers.Get)
}
