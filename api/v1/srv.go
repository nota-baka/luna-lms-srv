package apiv1

import (
	"Luna.LMS/server/db"
	"github.com/gin-gonic/gin"
)

type SrvWrapper struct {
	dbw db.DB
}

func (w *SrvWrapper) Ping(g *gin.Context) {
	// TODO: implement me
	panic("implement me")
}
