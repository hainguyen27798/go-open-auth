package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-open-auth/global"
	"github.com/go-open-auth/internal/dto"
	"github.com/go-open-auth/internal/services"
	"github.com/go-open-auth/pkg/response"
)

type AuthController struct {
	authService services.IAuthService
}

func NewAuthController(authService services.IAuthService) *AuthController {
	return &AuthController{
		authService,
	}
}

func (ac *AuthController) Register(c *gin.Context) {
	var params dto.UserRegistrationRequestDTO
	if err := c.ShouldBindBodyWithJSON(&params); err != nil {
		response.ValidateErrorResponse(c, err)
		return
	}

	response.MessageResponse(c, ac.authService.Register(params))
}

func (ac *AuthController) Login(c *gin.Context) {
	var params dto.UserLoginRequestDTO
	if err := c.ShouldBindBodyWithJSON(&params); err != nil {
		response.ValidateErrorResponse(c, err)
		return
	}

	if res, errCode := ac.authService.Login(params); errCode != nil {
		response.MessageResponse(c, *errCode)
	} else {
		response.OkResponse(c, response.LoginSuccess, *res)
	}
}

func (ac *AuthController) RefreshToken(c *gin.Context) {
	refreshToken := c.GetHeader(global.RefreshTokenKey)

	newToken, errCode := ac.authService.RefreshToken(refreshToken)
	if errCode != nil {
		response.MessageResponse(c, *errCode)
	} else {
		response.OkResponse(c, response.CodeSuccess, newToken)
	}
}

func (ac *AuthController) Logout(c *gin.Context) {}
