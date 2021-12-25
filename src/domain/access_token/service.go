package access_token

import (
	"strings"

	"github.com/azzamjiul/bookstore_oauth-api/src/repository/rest"
	"github.com/azzamjiul/bookstore_oauth-api/src/utils/error_utils"
)

type Repository interface {
	GetById(string) (*AccessToken, *error_utils.RestErr)
	Create(AccessToken) *error_utils.RestErr
	UpdateExpirationTime(AccessToken) *error_utils.RestErr
}

type Service interface {
	GetById(string) (*AccessToken, *error_utils.RestErr)
	Create(AccessTokenRequest) (*AccessToken, *error_utils.RestErr)
	UpdateExpirationTime(AccessToken) *error_utils.RestErr
}

type service struct {
	restUsersRepo rest.RestUsersRepository
	repository    Repository
}

func NewService(usersRepo rest.RestUsersRepository, repo Repository) Service {
	return &service{
		repository:    repo,
		restUsersRepo: usersRepo,
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

func (s *service) Create(request AccessTokenRequest) (*AccessToken, *error_utils.RestErr) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	user, err := s.restUsersRepo.LoginUser(request.Username, request.Password)
	if err != nil {
		return nil, err
	}

	// Generate a new access token:
	at := GetNewAccessToken(user.Id)
	at.Generate()

	// Save the new access token in Cassandra:
	if err := s.repository.Create(at); err != nil {
		return nil, err
	}

	return &at, nil
}

func (s *service) UpdateExpirationTime(at AccessToken) *error_utils.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}

	return s.repository.UpdateExpirationTime(at)
}
