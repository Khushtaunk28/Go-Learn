package routes

import (
	"Event_Booking/models"
	"Event_Booking/utils"
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

func Login(context *gin.Context){
	var user models.User

	err:=context.ShouldBindJSON(&user)
	if err!=nil{
		context.JSON(http.StatusBadRequest, gin.H{"message": "missing fields"})
		return
	}

	err=user.ValidateCred()
	if err!=nil{
		context.JSON(http.StatusUnauthorized,gin.H{"msg":"invalid password"})
		return
	}
	token,err:=utils.GenerateToken(user.Email,user.ID)
	if err!=nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message": "jwt token not generated"})
		return
	}


	context.JSON(http.StatusOK,gin.H{"msg":"Login success","token":token})
	
}
