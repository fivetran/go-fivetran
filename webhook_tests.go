package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
)

// WebhookTestService implements the test method for Webhook Management API.
// Ref. https://fivetran.com/docs/rest-api/webhooks#testwebhook
type WebhookTestService struct {
	c         *Client
	webhookId *string
	event     *string
}

type webhookTestRequest struct {
	Event *string `json:"event,omitempty"`
}

type WebhookTestResponse struct {
	Code string `json:"code"`
	Data struct {
		Succeed bool   `json:"succeed"`
		Status  int    `json:"status"`
		Message string `json:"message"`
	} `json:"data"`
}

func (c *Client) NewWebhookTest() *WebhookTestService {
	return &WebhookTestService{c: c}
}

func (s *WebhookTestService) request() *webhookTestRequest {
	return &webhookTestRequest{
		Event: s.event,
	}
}

func (s *WebhookTestService) WebhookId(value string) *WebhookTestService {
	s.webhookId = &value
	return s
}

func (s *WebhookTestService) Event(value string) *WebhookTestService {
	s.event = &value
	return s
}

func (s *WebhookTestService) Do(ctx context.Context) (WebhookTestResponse, error) {
	var response WebhookTestResponse

	if s.webhookId == nil {
		return response, fmt.Errorf("missing required webhookId")
	}

	url := fmt.Sprintf("%v/webhooks/%v/test", s.c.baseURL, *s.webhookId)
	expectedStatus := 200

	headers := s.c.commonHeaders()
	headers["Content-Type"] = "application/json"
	headers["Accept"] = restAPIv2

	reqBody, err := json.Marshal(s.request())
	if err != nil {
		return response, err
	}

	r := request{
		method:           "POST",
		url:              url,
		queries:          nil,
		body:             reqBody,
		headers:          headers,
		client:           s.c.httpClient,
		handleRateLimits: s.c.handleRateLimits,
		maxRetryAttempts: s.c.maxRetryAttempts,
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
