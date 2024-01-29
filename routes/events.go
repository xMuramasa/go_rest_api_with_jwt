package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/go_rest_api/models"
	"github.com/gin-gonic/gin"
)

func getEvents(ctx *gin.Context) {

	events, err := models.GetAllEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			gin.H{
				"message": "Could not fetch events",
			},
		)
		return
	}
	ctx.JSON(
		http.StatusOK,
		events,
	)
}

func getEvent(ctx *gin.Context) {
	eId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"message": "Invalid id"},
		)
		return
	}
	event, err := models.GetEventById(eId)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "Could not fetch event"},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		event,
	)

}

func createEvent(ctx *gin.Context) {

	id := ctx.GetInt64("id")

	var event models.Event
	err := ctx.ShouldBindJSON(&event)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "Could not parse data",
			},
		)
		return
	}

	event.UserId = id

	err = event.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			gin.H{
				"message": "Could not create events",
			},
		)
		return
	}

	ctx.JSON(
		http.StatusCreated,
		gin.H{
			"message": "Event created",
		},
	)

}

func updateEvent(ctx *gin.Context) {
	eId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"message": "Invalid id"},
		)
		return
	}

	id := ctx.GetInt64("id")

	e, err := models.GetEventById(eId)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "Event does not exist"},
		)
		return
	}

	if e.UserId != id {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized delete"})
		return
	}

	var updatedEvent models.Event
	err = ctx.ShouldBindJSON(&updatedEvent)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "Could not parse data",
			},
		)
		return
	}

	updatedEvent.Id = eId
	err = updatedEvent.Update()
	if err != nil {
		fmt.Println(err)
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "Could not update event",
			},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"message": "Event updated",
		},
	)
}

func deleteEvent(ctx *gin.Context) {
	id := ctx.GetInt64("id")

	eId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"message": "Invalid id"},
		)
		return
	}

	event, err := models.GetEventById(eId)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "Event does not exist"},
		)
		return
	}

	if event.UserId != id {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized update"})
		return
	}

	err = event.Delete()
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "Event could not be deleted"},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{"message": "Event deleted"},
	)

}
