package account

import (
	"ecommerce/global"
	"ecommerce/internal/model"
	"ecommerce/internal/service"
	"ecommerce/pkg/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// management controller Login User
var Login = new(cUserLogin)

type cUserLogin struct{}

func (c *cUserLogin) Login(ctx *gin.Context) {
	err := service.UserLogin().Login(ctx)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrorCodeParamInvalid, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.ErrorCodeSuccess, nil)
}

// User Register document
// @Summary      User Register
// @Description  When user is registered send otp to email
// @Tags         account management
// @Accept       json
// @Produce      json
// @Param        payload body model.RegisterInput true "payload"
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /user/register [post]
func (c *cUserLogin) Register(ctx *gin.Context) {
	var params model.RegisterInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrorCodeParamInvalid, err.Error())
		return
	}

	codeStatus, err := service.UserLogin().Register(ctx, &params)
	if err != nil {
		global.Logger.Error("Error registering user OTP", zap.Error(err))
		response.ErrorResponse(ctx, codeStatus, err.Error())
		return
	}

	response.SuccessResponse(ctx, codeStatus, nil)
}

// Verify OTP Login By User
// @Summary      Verify OTP Login By User
// @Description  Verify OTP Login By User
// @Tags         account management
// @Accept       json
// @Produce      json
// @Param        payload body model.VerifyInput true "payload"
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /user/verify_account [post]
func (c *cUserLogin) VerifyOTP(ctx *gin.Context) {
	var params model.VerifyInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrorCodeParamInvalid, err.Error())
		return
	}

	result, err := service.UserLogin().VerifyOTP(ctx, &params)
	if err != nil {
		global.Logger.Error("Error verifying OTP", zap.Error(err))
		response.ErrorResponse(ctx, response.ErrorInvalidOTP, err.Error())
		return
	}

	response.SuccessResponse(ctx, response.ErrorCodeSuccess, result)
}

// Update Password Register
// @Summary      Update Password Register
// @Description  Update Password Register
// @Tags         account management
// @Accept       json
// @Produce      json
// @Param        payload body model.UpdatePasswordRegisterInput true "payload"
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /user/update_pass_register [post]
func (c *cUserLogin) UpdatePasswordRegister(ctx *gin.Context) {
	var params model.UpdatePasswordRegisterInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrorCodeParamInvalid, err.Error())
		return
	}

	result, err := service.UserLogin().UpdatePasswordRegister(ctx, params.UserToken, params.UserPassword)
	if err != nil {
		response.ErrorResponse(ctx, result, err.Error())
		return
	}

	response.SuccessResponse(ctx, response.ErrorCodeSuccess, result)
}
