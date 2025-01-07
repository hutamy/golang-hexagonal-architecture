package api

import (
	"fmt"

	"github.com/hutamy/golang-hexagonal-architecture/shared/enum"
)

type RestClientRepository interface {
	Get(endpoint string, headers map[string]string, body interface{}) (interface{}, error)
	GetPathRequest(endpoint string, headers map[string]string, body interface{}) (interface{}, error)
	Post(endpoint string, headers map[string]string, body interface{}) (interface{}, error)
	Put(endpoint string, headers map[string]string, body interface{}) (interface{}, error)
}

func NewRestClient(restType enum.RestClient) (RestClientRepository, error) {
	switch restType {
	case enum.Resty:
		return NewRestyClient(), nil
	default:
		return nil, fmt.Errorf("invalid rest client type")
	}
}
