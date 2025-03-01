package routes

import (
	//"Event_Booking/db"
	"Event_Booking/models"


	//"Context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// update an event(PUT)
func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"msg": "couldnt fetch eventid"})
		return
	}
	userId:=context.GetInt64("userId")
	event,err:= models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"msg": "not a valid id"})
		return
	}

	if(event.UserID!=userId){
		context.JSON(http.StatusUnauthorized,gin.H{"msg":"user not authorized to update"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "missing fields"})
		return
	}
	updatedEvent.ID = eventId

	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"msg": "could not update event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"msg": "updated event successfully"})
}

// remove an event(DELETE)
func deleteEventById(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"msg": "couldnt fetch eventid"})
		return
	}
	userId:=context.GetInt64("userId")
	event,err:= models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"msg": "not a valid id"})
		return
	}

	if(event.UserID!=userId){
		context.JSON(http.StatusUnauthorized,gin.H{"msg":"user not authorized to delete"})
		return
	}
	
	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"msg": "couldnt dlt"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"msg": "delete success"})

}

// get all events
func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"msg": "couldnt fetch"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Hello"})
	context.JSON(http.StatusOK, events)
}

// get event by id
func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"msg": "couldnt fetch eventid"})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"msg": "not a valid id"})
		return
	}
	context.JSON(http.StatusOK, event)
}

// post to create a new event
func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "missing fields"})
		return
	}

	userId:=context.GetInt64("userId")
	event.UserID = userId
	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"msg": "couldnt create event"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "event created", "event": event})

}
