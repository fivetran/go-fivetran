package fivetran

import (
	"context"
	"encoding/json"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
	"github.com/fivetran/go-fivetran/webhooks"
)

// WebhookModifyService implements the Webhook Management, Modify a Webhook.
// Ref. https://fivetran.com/docs/rest-api/webhooks#updatewebhook
type WebhookModifyService struct {
	c         *Client
	webhookId *string
	url       *string
	events    *[]string
	active    *bool
	secret    *string
	runTests  *bool
}

type webhookModifyRequest struct {
	Url      *string   `json:"url,omitempty"`
	Events   *[]string `json:"events,omitempty"`
	Active   *bool     `json:"active,omitempty"`
	Secret   *string   `json:"secret,omitempty"`
	RunTests *bool     `json:"run_tests,omitempty"`
}

func (c *Client) NewWebhookModify() *WebhookModifyService {
	return &WebhookModifyService{c: c}
}

func (s *WebhookModifyService) request() *webhookModifyRequest {
	return &webhookModifyRequest{
		Url:    s.url,
		Events: s.events,
		Active: s.active,
		Secret: s.secret,
	}
}

func (s *WebhookModifyService) Url(value string) *WebhookModifyService {
	s.url = &value
	return s
}

func (s *WebhookModifyService) Secret(value string) *WebhookModifyService {
	s.secret = &value
	return s
}

func (s *WebhookModifyService) Active(value bool) *WebhookModifyService {
	s.active = &value
	return s
}

func (s *WebhookModifyService) Events(value []string) *WebhookModifyService {
	s.events = &value
	return s
}

func (s *WebhookModifyService) WebhookId(value string) *WebhookModifyService {
	s.webhookId = &value
	return s
}

func (s *WebhookModifyService) RunTests(value bool) *WebhookModifyService {
	s.runTests = &value
	return s
}

func (s *WebhookModifyService) Do(ctx context.Context) (webhooks.WebhookResponse, error) {
	var response webhooks.WebhookResponse

	if s.webhookId == nil {
		return response, fmt.Errorf("missing required webhookId")
	}

	url := fmt.Sprintf("%v/webhooks/%v", s.c.baseURL, *s.webhookId)
	expectedStatus := 200

	headers := s.c.commonHeaders()
	headers["Content-Type"] = "application/json"
	headers["Accept"] = restAPIv2

	reqBody, err := json.Marshal(s.request())
	if err != nil {
		return response, err
	}

	r := httputils.Request{
		Method:           "PATCH",
		Url:              url,
		Body:             reqBody,
		Queries:          nil,
		Headers:          headers,
		Client:           s.c.httpClient,
		HandleRateLimits: s.c.handleRateLimits,
		MaxRetryAttempts: s.c.maxRetryAttempts,
	}

	respBody, respStatus, err := r.Do(ctx)
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
