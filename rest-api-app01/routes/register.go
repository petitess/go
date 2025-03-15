package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"site.org/abc/models"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id. " + err.Error()})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not find event. " + err.Error()})
		return
	}

	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not register. " + err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Added user to event"})
}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id. " + err.Error()})
		return
	}
	var event models.Event
	event.ID = eventId
	event.CancelRegistration(userId)

	err = event.CancelRegistration(userId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not cancel. " + err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Canceled"})

}
