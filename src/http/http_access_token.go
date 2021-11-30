package http

import (
	"github.com/azzamjiul/bookstore_oauth-api/src/domain/access_token"
)

type AccessTokenHandler interface {
	GetById()
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewHandler(service access_token.Service) AccessTokenHandler {

}
