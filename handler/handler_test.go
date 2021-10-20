package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOauth(t *testing.T) {
	r := MakeHandler()
	code := 123
	assert := assert.New(t)
	ts := httptest.NewServer(r)

	req, err := http.NewRequest("GET", ts.URL+"/api/v1/auth/google/callback?code=123", nil)
	assert.NoError(err)
	Param := req.URL.Query()
	assert.Equal(Param, code)
}
