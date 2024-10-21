package routes

import (
	"streak-ai-assesment/controllers"

	"github.com/gin-gonic/gin"
)

func SetRoutes(router *gin.Engine) {
	router.POST("/find-pairs", controllers.FindPairs)

}
