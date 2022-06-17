package handler

import (
	"net/http"

	"kotak-amal/helper"
	"kotak-amal/transport/formatter"
	"kotak-amal/transport/request"
	"kotak-amal/transport/response"
	"kotak-amal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type userHandler struct {
	userUsecase usecase.UserUsecase
	validator   *validator.Validate
}

func NewUserHandler(userUsecase usecase.UserUsecase, validator *validator.Validate) *userHandler {
	return &userHandler{
		userUsecase: userUsecase,
		validator:   validator,
	}
}

func (u *userHandler) CreateUser(c *gin.Context) {
	var requestUser request.RegisterUserRequest
	if err := c.ShouldBindJSON(&requestUser); err != nil {
		resp := response.GeneralResponse{
			Meta: response.Meta{
				Code:    http.StatusBadRequest,
				Status:  "failed",
				Message: "error while decode request body",
			},
			Data: nil,
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	/*
	 * Checking Validation
	 */
	errorValidation := u.validator.Struct(requestUser)
	if errorValidation != nil {
		errors := helper.FormatValidationError(errorValidation)
		errMsg := gin.H{"errors": errors}

		resp := response.GeneralResponse{
			Meta: response.Meta{
				Code:    http.StatusBadRequest,
				Status:  "failed",
				Message: "register account failed",
			},
			Data: errMsg,
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	res, errorResp := u.userUsecase.CreateUser(c.Request.Context(), &requestUser)
	if errorResp != nil {
		c.JSON(errorResp.Code, errorResp)
		return
	}

	user := formatter.FormatUser(res)
	resp := response.GeneralResponse{
		Meta: response.Meta{
			Code:    http.StatusCreated,
			Status:  "success",
			Message: "account has been registered",
		},
		Data: user,
	}

	c.JSON(http.StatusCreated, resp)
}
