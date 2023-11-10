package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/migueelzz/crud-go/src/configuration/logger"
	"github.com/migueelzz/crud-go/src/configuration/validation"
	"github.com/migueelzz/crud-go/src/controller/model/request"
	"github.com/migueelzz/crud-go/src/controller/model/response"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zap.String("journey", "createUser"),
	)

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "createUser"))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	response := response.UserResponse{
		ID:    "test",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}

	logger.Info("User created successfully",
		zap.String("journey", "createUser"))

	c.JSON(http.StatusOK, response)
}
