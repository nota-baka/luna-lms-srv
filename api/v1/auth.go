package apiv1

import (
	"Luna.LMS/server/db"
	"github.com/gin-gonic/gin"
)

type AuthWrapper struct {
	dbw db.DB
}

func (w *AuthWrapper) Login(g *gin.Context) {
	// TODO: implement me
	panic("implement me")
}

func (w *AuthWrapper) Logout(g *gin.Context) {
	// TODO: implement me
	panic("implement me")
}

func (w *AuthWrapper) LoginWithTOTP(g *gin.Context) {
	// TODO: implement me
	panic("implement me")
}

func (w *AuthWrapper) SetupTOTPInit(g *gin.Context) {
	// TODO: implement me
	panic("implement me")
}

func (w *AuthWrapper) SetupTOTPConfirm(g *gin.Context) {
	// TODO: implement me
	panic("implement me")
}
