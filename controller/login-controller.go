package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rfaelmarini/call-cat-api/dto"
	"github.com/rfaelmarini/call-cat-api/service"
)

type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginService service.LoginService
	jWtService   service.JWTService
}

func LoginHandler(loginService service.LoginService,
	jWtService service.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jWtService:   jWtService,
	}
}

func (controller *loginController) Login(ctx *gin.Context) string {
	var credential dto.LoginCredentials
	err := ctx.ShouldBind(&credential)
	if err != nil {
		return "no data found"
	}

	isUserAuthenticated := controller.loginService.Login(credential.Username, credential.Password)
	if isUserAuthenticated {
		return controller.jWtService.GenerateToken(credential.Username, true)
	}

	return ""
}
