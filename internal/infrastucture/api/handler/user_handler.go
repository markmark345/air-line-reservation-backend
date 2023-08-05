package handler

import (
	"air-line-reservation-backend/internal/application/services"
	"air-line-reservation-backend/internal/domain/entities"
	"air-line-reservation-backend/internal/infrastucture/api/schema"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator"

	"github.com/markmark345/air-line-v1-common/api/responses"
	validatorscustom "github.com/markmark345/air-line-v1-common/validators-custom"
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
	userId := c.Param("id")

	result, err := handler.userService.GetUser(ctx, userId)
	if err != nil {
		responses.Failure(c, "user_not_found", nil)
		return
	}

	responses.Success(c, "success", result)
}

func (handler *UserHandler) CreateUser(c *gin.Context) {
	ctx := c.Request.Context()
	user := entities.User{}
	body := schema.CreateUser{}

	if err := c.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		responses.Error(c, err)
		return
	}
	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		fmt.Println(err)
		validatorscustom.ValidateCustomError(c, err)
		return
	}

	if err := c.ShouldBindBodyWith(&user, binding.JSON); err != nil {
		responses.Error(c, err)
		return
	}

	fmt.Println(user)

	err := handler.userService.CreateUser(ctx, &user)
	if err != nil {
		responses.Error(c, err)
		return
	}

	responses.Success(c, "created", nil)
}
