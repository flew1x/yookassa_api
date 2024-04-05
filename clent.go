package yookassa

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

type yooKassaClient struct {
	baseURL string
	shopID  string
	apiKEY  string
	httpClient *http.Client
}



// NewConfig creates a newConfig yooKassaClient with the given shop ID, API key, and options.
// It returns a pointer to the yooKassaClient.
func NewConfig(shopId, apiKey string, opts ...func(c *yooKassaClient)) *yooKassaClient {
	c := &yooKassaClient{
		baseURL: "https://api.yookassa.ru/v3/payments",
		shopID:  shopId,
		apiKEY:  apiKey,
		httpClient:    &http.Client{},
	}

	for _, o := range opts {
		o(c)
	}

	return c
}

func (c *yooKassaClient) doRequest(method string, payload io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, c.baseURL, payload)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Idempotence-Key", uuid.NewV4().String())
	req.SetBasicAuth(c.shopID, c.apiKEY)

	return c.httpClient.Do(req)
}

func (c *yooKassaClient) sendRequest(method string, payload interface{}) (*http.Response, error) {
	var buf bytes.Buffer
	if payload != nil {
		if err := json.NewEncoder(&buf).Encode(payload); err != nil {
			return nil, fmt.Errorf("failed to marshal payload: %w", err)
		}
	}

	return c.doRequest(method, &buf)
}

// SendPaymentRequest sends a payment request to YooKassa and returns the payment response and any error encountered.
//
// It takes a paymentRequest of type PaymentRequest and returns a PaymentResponse and an error.
func (c *yooKassaClient) SendPaymentRequest(paymentRequest PaymentRequest) (PaymentResponse, error) {
	resp, err := c.sendRequest(http.MethodPost, paymentRequest)
	if err != nil {
		return PaymentResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return PaymentResponse{}, fmt.Errorf("request failed with status %s", resp.Status)
	}

	var paymentResponse PaymentResponse
	if err = json.NewDecoder(resp.Body).Decode(&paymentResponse); err != nil {
		return PaymentResponse{}, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return paymentResponse, nil
}

// GetPayment retrieves a payment with the given ID from the YooKassa API.
func (c *yooKassaClient) GetPayment(ctx context.Context, paymentID string) (PaymentResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+"/"+paymentID, nil)
	if err != nil {
		return PaymentResponse{}, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(c.shopID, c.apiKEY)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PaymentResponse{}, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return PaymentResponse{}, fmt.Errorf("request failed with status %s", resp.Status)
	}

	var paymentResponse PaymentResponse
	if err = json.NewDecoder(resp.Body).Decode(&paymentResponse); err != nil {
		return PaymentResponse{}, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return paymentResponse, nil
}

