package routes

import (
	"net/http"

	"example.com/go_rest_api/models"
	"example.com/go_rest_api/utils"
	"github.com/gin-gonic/gin"
)

func signup(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "Could not parse user data",
			},
		)
		return
	}

	err = user.Save()
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "Could not create user",
			},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"message": "User created",
		},
	)

}

func login(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"message": "Could not parse data"},
		)
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		ctx.JSON(
			http.StatusUnauthorized,
			gin.H{"mesasge": err.Error()},
		)
		return
	}

	tok, err := utils.GenerateToken(user.Email, user.Id)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"mesasge": err.Error()},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{"message": "login succesful", "token": tok},
	)

}
