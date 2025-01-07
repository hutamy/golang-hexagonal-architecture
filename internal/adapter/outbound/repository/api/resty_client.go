package api

import (
	"gopkg.in/resty.v1"
)

type RestyClient struct {
	client *resty.Client
}

// NewRestyClient new object resty client
func NewRestyClient() *RestyClient {
	return &RestyClient{
		client: resty.New(),
	}
}

// Post send using post method
func (c *RestyClient) Post(endpoint string, headers map[string]string, body interface{}) (interface{}, error) {
	return c.client.R().SetHeaders(headers).SetBody(body).Post(endpoint)
}

// Put send using put method
func (c *RestyClient) Put(endpoint string, headers map[string]string, body interface{}) (interface{}, error) {
	return c.client.R().SetHeaders(headers).SetBody(body).Put(endpoint)
}

// GetClient get client of resty
func (c *RestyClient) Get(endpoint string, headers map[string]string, body interface{}) (interface{}, error) {
	return c.client.R().SetHeaders(headers).SetBody(body).Post(endpoint)
}

// GetClient get client of resty
func (c *RestyClient) GetPathRequest(endpoint string, headers map[string]string, body interface{}) (interface{}, error) {
	return c.client.R().SetHeaders(headers).SetBody(body).Get(endpoint)
}
