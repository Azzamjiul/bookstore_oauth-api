package db

import (
	"github.com/azzamjiul/bookstore_oauth-api/src/clients/cassandra"
	"github.com/azzamjiul/bookstore_oauth-api/src/domain/access_token"
	"github.com/azzamjiul/bookstore_oauth-api/src/utils/error_utils"
	"github.com/gocql/gocql"
)

const (
	queryGetAccessToken    = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token = ?;"
	queryCreateAccessToken = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES (?,?,?,?);"
	queryUpdateExpires     = "UPDATE access_tokens SET expires = ? WHERE access_token = ?;"
)

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *error_utils.RestErr)
	Create(access_token.AccessToken) *error_utils.RestErr
	UpdateExpirationTime(access_token.AccessToken) *error_utils.RestErr
}

type dbRepository struct {
}

func NewRepository() DbRepository {
	return &dbRepository{}
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *error_utils.RestErr) {
	var result access_token.AccessToken
	if err := cassandra.GetSession().Query(queryGetAccessToken, id).Scan(
		&result.AccessToken,
		&result.UserId,
		&result.ClientId,
		&result.Expires,
	); err != nil {
		if err == gocql.ErrNotFound {
			return nil, error_utils.NewNotFoundError("no access token found with given id")
		}
		return nil, error_utils.NewInternalServerError("error when trying to get current id")
	}
	return &result, nil
}

func (r *dbRepository) Create(at access_token.AccessToken) *error_utils.RestErr {
	err := cassandra.GetSession().Query(queryCreateAccessToken, at.AccessToken, at.UserId, at.ClientId, at.Expires).Exec()
	if err != nil {
		return error_utils.NewInternalServerError(err.Error())
	}
	return nil
}

func (r *dbRepository) UpdateExpirationTime(at access_token.AccessToken) *error_utils.RestErr {
	err := cassandra.GetSession().Query(queryUpdateExpires, at.Expires, at.AccessToken).Exec()
	if err != nil {
		return error_utils.NewInternalServerError(err.Error())
	}
	return nil
}
