package handler

import (
	"air-line-reservation-backend/internal/application/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(
	userService services.UserService,
) *UserHandler {
	return &UserHandler{
		userService,
	}
}

func (handler *UserHandler) GetUser(c *gin.Context) {
	ctx := c.Request.Context()
	fmt.Println(c.Param("id"))

	userId := c.Param("id")
	result, err := handler.userService.GetUser(ctx, userId)
	fmt.Println(result)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": result})
}
