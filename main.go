package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vismaygamit/eventbooking/db"
	"github.com/vismaygamit/eventbooking/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") // localhost:8080
}
