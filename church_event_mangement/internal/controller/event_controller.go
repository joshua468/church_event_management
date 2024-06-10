package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joshua468/church_event_management/internal/model"
	"github.com/joshua468/church_event_management/internal/service"
)

type EventController struct {
	service service.EventService
}

func NewEventController(service service.EventService) *EventController {
	return &EventController{service}
}

func (c *EventController) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/events", c.CreateEvent)
	router.GET("/events/:id", c.GetEvent)
}

func (c *EventController) CreateEvent(ctx *gin.Context) {
	var event model.Event
	if err := ctx.ShouldBindJSON(&event); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.service.CreateEvent(&event); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, event)
}

func (c *EventController) GetEvent(ctx *gin.Context) {
	idParam := ctx.Param("id")
	idUint64, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	event, err := c.service.GetEventByID(uint(idUint64))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}
	ctx.JSON(http.StatusOK, event)
}
