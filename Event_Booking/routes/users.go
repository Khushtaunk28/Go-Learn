package routes

import (
	"Event_Booking/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SaveUser(context *gin.Context){
	var user models.User
	err:=context.ShouldBindJSON(&user)
	if err!=nil{
		context.JSON(http.StatusBadRequest, gin.H{"message": "missing fields"})
		return
	}
	// user.ID=1;
	err=user.Save()
	if err!=nil{
		context.JSON(http.StatusInternalServerError,gin.H{"msg":"couldnt save user"})
		return
	}
	context.JSON(http.StatusOK,gin.H{"msg":"user created and saved"})

}
