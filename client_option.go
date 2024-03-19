package yookassa

import (
	"net/http"
)

func WithBaseURL(url string) func(*yooKassaClient) {
	return func(c *yooKassaClient) {
		c.baseURL = url
	}
}

func WithHTTPClient(client *http.Client) func(*yooKassaClient) {
	return func(c *yooKassaClient) {
		c.httpClient = client
	}
}