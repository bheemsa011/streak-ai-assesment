package main

import (
	"streak-ai-assesment/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	routes.SetRoutes(router)
	router.Run()

}
