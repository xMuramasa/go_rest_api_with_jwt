package routes

import (
	"net/http"
	"strconv"

	"example.com/go_rest_api/models"
	"github.com/gin-gonic/gin"
)

func RegisterForEvent(ctx *gin.Context) {
	userId := ctx.GetInt64("id")

	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"message": "Invalid id"},
		)
		return
	}
	e, err := models.GetEventById(eventId)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "could not fetch event"},
		)
		return
	}

	err = e.Register(userId)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "could not make registration"},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{"message": "registration complete"},
	)

}

func DeleteRegistration(ctx *gin.Context) {
	userId := ctx.GetInt64("id")

	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"message": "Invalid id"},
		)
		return
	}

	var event models.Event
	event.Id = eventId
	err = event.CancelRegistration(userId)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "could not cancel registration"},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{"message": "registration canceled"},
	)
}
