package apiv1

import (
	"encoding/base64"
	"net/http"
	"time"
	"uuid"

	"Luna.LMS/server/cryptohelpers"
	"Luna.LMS/server/db"
	"Luna.LMS/server/entities"
	"github.com/gin-gonic/gin"
)

type AuthWrapper struct {
	dbw       db.DB
	debugMode bool
}

type RequestLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResponseLogin struct {
	AccountId string               `json:"account_id"`
	Name      string               `json:"name"`
	Role      entities.AccountRole `json:"role"`
}

func (w *AuthWrapper) Login(g *gin.Context) {
	var body RequestLogin
	if err := g.ShouldBindJSON(&body); err != nil || body.Email == "" || body.Password == "" {
		g.AbortWithStatus(http.StatusBadRequest)
		return
	}
	account, err := w.dbw.Accounts().GetByEmail(g.Request.Context(), body.Email)
	if err != nil {
		g.AbortWithError(http.StatusUnauthorized, err)
		return
	}
	if !cryptohelpers.VerifyPassword(body.Password, account.PasswordHash) {
		g.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	if _, err := w.dbw.TOTPs().GetByAccountId(g.Request.Context(), account.Id); err != nil {
		token, err := cryptohelpers.GenerateToken()
		if err != nil {
			g.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		curTime := time.Now()
		device := entities.Device{
			Id:             uuid.NewV7().String(),
			AccountId:      account.Id,
			LastOnlineAt:   curTime,
			TokenHash:      cryptohelpers.HashToken(token),
			TokenExpiresAt: curTime.Add(time.Minute * 10),
		}
		if err := w.dbw.Devices().Add(g.Request.Context(), &device); err != nil {
			g.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		// TODO: Make domain configurable
		g.SetCookie("token", base64.StdEncoding.EncodeToString(token), 60*5, "/", "localhost", !w.debugMode, true)
		g.SetCookie("device_id", device.Id, 60*60*30, "/", "localhost", !w.debugMode, true)
		g.JSON(http.StatusOK, ResponseLogin{
			AccountId: account.Id,
			Name:      account.Name,
			Role:      account.Role,
		})
		return
	}

	totpAttempt := entities.TOTPAttempt{
		Id:        uuid.NewV7().String(),
		AccountId: account.Id,
		ExpiresAt: time.Now().Add(time.Minute * 5),
	}
	err = w.dbw.TOTPsAttempts().Add(g.Request.Context(), &totpAttempt)
	if err != nil {
		g.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	g.SetCookie("totp_attempt_id", totpAttempt.Id, int(time.Until(totpAttempt.ExpiresAt)), "/", "localhost", !w.debugMode, true)
	g.Status(http.StatusOK)
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
