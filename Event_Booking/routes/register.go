package routes

import (
	"Event_Booking/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"msg": "couldnt fetch eventid"})
		return
	}
	event,err:=models.GetEventById(eventId)
	if err!=nil{
		context.JSON(http.StatusInternalServerError,gin.H{"msg":"event not found"})
		return
	}

	err=event.Register(userId)
	if err!=nil{
		context.JSON(http.StatusInternalServerError,gin.H{"msg":"event not registerd"})
		return
	}

	context.JSON(http.StatusOK,gin.H{"msg":"Event registerd"})
}


func CancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"msg": "couldnt fetch eventid"})
		return
	}
	var event models.Event
	event.ID=eventId
	err=event.Cancel(userId)
	if err!=nil{
		context.JSON(http.StatusInternalServerError,gin.H{"msg":"event not canceelled"})
		return
	}
	context.JSON(http.StatusOK,gin.H{"msg":"Event cancelled success"})
}