package yookassa

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

type YooKassaClient struct {
	baseURL string
	shopID  string
	apiKEY  string
	httpClient *http.Client
}

// NewConfig creates a newConfig YooKassaClient with the given shop ID, API key, and options.
// It returns a pointer to the YooKassaClient.
func NewConfig(shopId, apiKey string, opts ...func(c *YooKassaClient)) *YooKassaClient {
	c := &YooKassaClient{
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
// SendPaymentRequest sends a payment request to YooKassa and returns the payment response and any error encountered.
//
// It takes a paymentRequest of type PaymentRequest and returns a PaymentResponse and an error.
func (c *YooKassaClient) SendPaymentRequest(paymentRequest PaymentRequest) (PaymentResponse, error) {
	payload, err := json.Marshal(paymentRequest)
	if err != nil {
		return PaymentResponse{}, err
	}

	req, err := http.NewRequest("POST", c.baseURL, bytes.NewBuffer(payload))
	if err != nil {
		return PaymentResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Idempotence-Key", uuid.NewV4().String())
	req.SetBasicAuth(c.shopID, c.apiKEY)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PaymentResponse{}, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return PaymentResponse{}, err
	}

	var paymentResponse PaymentResponse
	err = json.Unmarshal(respBody, &paymentResponse)
	if err != nil {
		return PaymentResponse{}, err
	}

	return paymentResponse, nil
}

// GetPayment retrieves a payment with the given ID from the YooKassa API.
//
// paymentID string - the ID of the payment to retrieve.
// PaymentResponse - the payment response object.
// error - an error if the request fails.
func (c *YooKassaClient) GetPayment(paymentID string) (PaymentResponse, error) {
	req, err := http.NewRequest("GET", c.baseURL + "/" + paymentID, nil)
	if err != nil {
		return PaymentResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(c.shopID, c.apiKEY)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PaymentResponse{}, err
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return PaymentResponse{}, err
	}

	var paymentResponse PaymentResponse
	err = json.Unmarshal(respBody, &paymentResponse)

	return paymentResponse, err
}