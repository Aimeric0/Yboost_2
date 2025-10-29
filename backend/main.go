package main

import (
	"Yboost_2/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8080")
}
