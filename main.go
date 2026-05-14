package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shaneoh10/events-backend/db"
	"github.com/shaneoh10/events-backend/routes"
)

func main() {
	db.Init()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
