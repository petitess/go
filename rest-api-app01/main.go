package main

import (
	"github.com/gin-gonic/gin"
	"site.org/abc/db"
	"site.org/abc/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
