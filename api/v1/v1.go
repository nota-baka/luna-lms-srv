package apiv1

import (
	"errors"

	"Luna.LMS/server/api"
	"Luna.LMS/server/config"
	"Luna.LMS/server/db"
)

type APIv1Wrapper struct {
	dbw         db.DB
	srvWrapper  *SrvWrapper
	authWrapper *AuthWrapper
}

func (w *APIv1Wrapper) Srv() api.Srv {
	return w.srvWrapper
}

func (w *APIv1Wrapper) Auth() api.Auth {
	return w.authWrapper
}

// NewAPIv1Wrapper creates new APIv1 handler with provided database wrapper
func NewAPIv1Wrapper(dbw db.DB, serverMode config.ServerMode) (*APIv1Wrapper, error) {
	if dbw == nil {
		return nil, errors.New("failed to create APIv1 wrapper: provided with nil")
	}

	debugMode := serverMode == config.ServerModeDebug

	return &APIv1Wrapper{
		dbw:         dbw,
		srvWrapper:  &SrvWrapper{dbw: dbw, debugMode: debugMode},
		authWrapper: &AuthWrapper{dbw: dbw, debugMode: debugMode}}, nil
}
