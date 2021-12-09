package access_token

import (
	"strings"

	"github.com/azzamjiul/bookstore_oauth-api/src/utils/error_utils"
)

type Repository interface {
	GetById(string) (*AccessToken, *error_utils.RestErr)
	Create(AccessToken) *error_utils.RestErr
	UpdateExpirationTime(AccessToken) *error_utils.RestErr
}

type Service interface {
	GetById(string) (*AccessToken, *error_utils.RestErr)
	Create(AccessToken) *error_utils.RestErr
	UpdateExpirationTime(AccessToken) *error_utils.RestErr
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(accessTokenId string) (*AccessToken, *error_utils.RestErr) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, error_utils.NewBadRequestError("invalid acces token id")
	}

	accessToken, err := s.repository.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}

	return accessToken, nil
}

func (s *service) Create(at AccessToken) *error_utils.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}

	return s.repository.Create(at)
}

func (s *service) UpdateExpirationTime(at AccessToken) *error_utils.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}

	return s.repository.UpdateExpirationTime(at)
}
