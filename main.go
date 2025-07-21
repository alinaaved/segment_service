package main

import (
	"segment_service/db"
	"segment_service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8080")
}
