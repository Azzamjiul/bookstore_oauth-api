package db

import (
	"github.com/azzamjiul/bookstore_oauth-api/src/domain/access_token"
	"github.com/azzamjiul/bookstore_oauth-api/src/utils/error_utils"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *error_utils.RestErr)
}

type dbRepository struct {
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *error_utils.RestErr) {
	return nil, nil
}
