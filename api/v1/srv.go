package apiv1

import (
	"net/http"

	"Luna.LMS/server/db"
	"github.com/gin-gonic/gin"
)

type SrvWrapper struct {
	dbw db.DB
}

type PingResponse struct {
	Status string `json:"status"`
}

func (w *SrvWrapper) Ping(g *gin.Context) {
	g.JSON(http.StatusOK, PingResponse{Status: "pong"})
}
