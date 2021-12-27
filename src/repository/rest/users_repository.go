package rest

import (
	"time"

	"github.com/azzamjiul/bookstore_oauth-api/src/domain/users"
	"github.com/azzamjiul/bookstore_oauth-api/src/utils/error_utils"
	"github.com/azzamjiul/bookstore_oauth-api/src/utils/http_utils"
	"github.com/mercadolibre/golang-restclient/rest"
)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: "localhost:8081",
		Timeout: 100 * time.Millisecond,
	}
)

type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, *error_utils.RestErr)
}

type usersRepository struct{}

func NewRepository() RestUsersRepository {
	return &usersRepository{}
}

func (r *usersRepository) LoginUser(email string, password string) (*users.User, *error_utils.RestErr) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}

	var user users.User
	headers := map[string]string{"Content-Type": "application/json"}

	_, err := http_utils.New().Post("http://localhost:8081/users/login", request, headers, &user)
	if err != nil {
		return nil, error_utils.NewInternalServerError(err.Error())
	}

	return &user, nil
}
