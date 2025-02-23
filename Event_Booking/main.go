package main

import (
	"Event_Booking/db"
	"Event_Booking/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	server.GET("/events", getEvents)

	server.POST("/events", createEvent)

	server.Run(":8080")

}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"msg": "couldnt fetch"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Hello"})
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "missing fields"})
		return
	}
	event.ID = 1
	event.UserID = 1
	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"msg": "couldnt create event"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "event created", "event": event})

}
