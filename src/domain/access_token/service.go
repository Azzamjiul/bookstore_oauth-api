package access_token

import "github.com/azzamjiul/bookstore_oauth-api/src/utils/error_utils"

type Repository interface {
	GetById(string) (*AccessToken, *error_utils.RestErr)
}

type Service interface {
	GetById(string) (*AccessToken, *error_utils.RestErr)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(string) (*AccessToken, *error_utils.RestErr) {
	return nil, nil
}
