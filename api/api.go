package api

import (
	"github.com/gin-gonic/gin"
)

type Srv interface {
	Ping(g *gin.Context)
}

type Auth interface {
	Login(g *gin.Context)
	Logout(g *gin.Context)
	LoginWithTOTP(g *gin.Context)
	SetupTOTPInit(g *gin.Context)
	SetupTOTPConfirm(g *gin.Context)
}

type API interface {
	Srv() Srv
	Auth() Auth
}
