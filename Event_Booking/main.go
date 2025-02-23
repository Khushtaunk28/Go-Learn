package main

import (
	"Event_Booking/db"
	"Event_Booking/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisteredRoutes(server)
	server.Run(":8080")
}
