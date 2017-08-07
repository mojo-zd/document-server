package tools

import (
	"encoding/base64"
	"errors"
	"strings"
)

const (
	prefix = "Basic "
	secret = "my_secret"
)

type BasicAuth struct {
	UserName string
	Password string
}

func ParseBasicAuth(basicAuth string) (basic BasicAuth, err error) {
	if !strings.HasPrefix(basicAuth, prefix) {
		return
	}
	authBytes, err := base64.StdEncoding.DecodeString(basicAuth[len(prefix):])
	if err != nil {
		return
	}
	auth := string(authBytes)
	s := strings.IndexByte(auth, ':')
	if s < 0 {
		err = errors.New("basic auth invalid")
		return
	}
	basic.UserName = auth[:s]
	basic.Password = auth[s+1:]
	return
}
