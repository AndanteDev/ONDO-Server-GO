package httputil

import (
	"errors"
)

var (
	ErrFailedExchangeToken error = errors.New("failed to exchange")
	ErrFaildGetUserInfo    error = errors.New("failed to get userInfo")
	ErrRespBodyNil         error = errors.New("err respBody Nil")
	ErrInvalidOauthState   error = errors.New("invalid google oauth state")
)
