package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joshua468/church_event_management/internal/model"
	"github.com/joshua468/church_event_management/internal/service"
	"github.com/joshua468/church_event_management/internal/util"
)

type UserController struct {
	service service.UserService
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{service}
}

func (c *UserController) RegisterRoutes(router *gin.Engine) {
	router.POST("/api/register", c.RegisterUser)
	router.POST("/api/login", c.LoginUser)
}

func (c *UserController) RegisterUser(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.service.RegisterUser(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, user)
}

func (c *UserController) LoginUser(ctx *gin.Context) {
	var loginData struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&loginData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := c.service.AuthenticateUser(loginData.Email, loginData.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}
	token, err := util.GenerateToken(user.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
