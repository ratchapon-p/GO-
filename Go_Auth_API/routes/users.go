package routes

import (
	"net/http"

	"example.com/go_auth_api/models"
	"example.com/go_auth_api/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
		
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user."})
		return
		
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created"})

}

func login(context *gin.Context)  {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
		
	}

	err = user.ValidCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user."})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not authenticate user."})
	}

	context.JSON(http.StatusOK, gin.H{"nessage": "Login Successfully!", "token": token})


}
