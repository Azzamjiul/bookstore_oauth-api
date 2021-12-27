package http_utils

import (
	"encoding/json"
	"fmt"

	"github.com/azzamjiul/bookstore_oauth-api/src/utils/error_utils"
	"github.com/parnurzeal/gorequest"
)

type client struct{}

type Client interface {
	Get(string, map[string]string, interface{}) *error_utils.RestErr
}

func New() Client {
	return &client{}
}

func (c client) Get(url string, headers map[string]string, resp interface{}) *error_utils.RestErr {
	request := gorequest.New()
	_, bytes, err := request.Get(url).EndBytes()

	if err != nil {
		fmt.Println(err)
		return error_utils.NewInternalServerError("error external api")
	}

	if err := json.Unmarshal(bytes, &resp); err != nil {
		return error_utils.NewInternalServerError("error while reading response content")
	}

	return nil
}
