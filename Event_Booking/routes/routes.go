package routes

import (
	"Event_Booking/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisteredRoutes(server *gin.Engine){
	//events routes
	server.GET("/events", getEvents)
	server.GET("/events/:id",getEvent)


	authenticated:= server.Group("/")
	authenticated.Use(middlewares.Authenticaton)
	authenticated.POST("/events",createEvent)
	authenticated.PUT("/events/:id",updateEvent)
	authenticated.DELETE("/events/:id",deleteEventById)


	// server.POST("/events",middlewares.Authenticaton ,createEvent)
	// server.PUT("/events/:id",updateEvent)
	// server.DELETE("/events/:id",deleteEventById)

	//user handling routes
	server.POST("/signup",SaveUser)
	server.POST("/login",Login)

}