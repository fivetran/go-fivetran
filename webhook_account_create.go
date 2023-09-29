package fivetran

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fivetran/go-fivetran/webhooks"
)

// WebhookAccountCreateService implements the Webhook Management, Create a Account Webhook.
// Ref. https://fivetran.com/docs/rest-api/webhooks#createaccountwebhook
type WebhookAccountCreateService struct {
	c      *Client
	url    *string
	events *[]string
	active *bool
	secret *string
}

type webhookAccountCreateRequest struct {
	Url    *string   `json:"url,omitempty"`
	Events *[]string `json:"events,omitempty"`
	Active *bool     `json:"active,omitempty"`
	Secret *string   `json:"secret,omitempty"`
}

func (c *Client) NewWebhookAccountCreate() *WebhookAccountCreateService {
	return &WebhookAccountCreateService{c: c}
}

func (s *WebhookAccountCreateService) request() *webhookAccountCreateRequest {
	return &webhookAccountCreateRequest{
		Url:    s.url,
		Events: s.events,
		Active: s.active,
		Secret: s.secret,
	}
}

func (s *WebhookAccountCreateService) Url(value string) *WebhookAccountCreateService {
	s.url = &value
	return s
}

func (s *WebhookAccountCreateService) Secret(value string) *WebhookAccountCreateService {
	s.secret = &value
	return s
}

func (s *WebhookAccountCreateService) Active(value bool) *WebhookAccountCreateService {
	s.active = &value
	return s
}

func (s *WebhookAccountCreateService) Events(value []string) *WebhookAccountCreateService {
	s.events = &value
	return s
}

func (s *WebhookAccountCreateService) Do(ctx context.Context) (webhooks.WebhookResponse, error) {
	var response webhooks.WebhookResponse
	url := fmt.Sprintf("%v/webhooks/account", s.c.baseURL)
	expectedStatus := 200

	headers := s.c.commonHeaders()
	headers["Content-Type"] = "application/json"
	headers["Accept"] = restAPIv2

	reqBody, err := json.Marshal(s.request())
	if err != nil {
		return response, err
	}

	r := request{
		method:  "POST",
		url:     url,
		body:    reqBody,
		queries: nil,
		headers: headers,
		client:  s.c.httpClient,
	}

	respBody, respStatus, err := r.httpRequest(ctx)
	if err != nil {
		return response, err
	}

	if err := json.Unmarshal(respBody, &response); err != nil {
		return response, err
	}

	if respStatus != expectedStatus {
		err := fmt.Errorf("status code: %v; expected: %v", respStatus, expectedStatus)
		return response, err
	}

	return response, nil
}
