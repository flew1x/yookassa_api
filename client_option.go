package yookassa

import (
	"net/http"
)

func WithBaseURL(url string) func(*YooKassaClient) {
	return func(c *YooKassaClient) {
		c.baseURL = url
	}
}

func WithHTTPClient(client *http.Client) func(*YooKassaClient) {
	return func(c *YooKassaClient) {
		c.httpClient = client
	}
}