package fivetran

import (
	"context"
	"encoding/json"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
	"github.com/fivetran/go-fivetran/webhooks"
)

// WebhookDetailsService implements the Webhook Management, retrieve Webhook Details.
// Ref. https://fivetran.com/docs/rest-api/webhooks#retrievewebhookdetails
type WebhookDetailsService struct {
	c         *Client
	webhookId *string
}

func (c *Client) NewWebhookDetails() *WebhookDetailsService {
	return &WebhookDetailsService{c: c}
}

func (s *WebhookDetailsService) WebhookId(value string) *WebhookDetailsService {
	s.webhookId = &value
	return s
}

func (s *WebhookDetailsService) Do(ctx context.Context) (webhooks.WebhookResponse, error) {
	var response webhooks.WebhookResponse

	if s.webhookId == nil {
		return response, fmt.Errorf("missing required webhookId")
	}

	url := fmt.Sprintf("%v/webhooks/%v", s.c.baseURL, *s.webhookId)
	expectedStatus := 200

	headers := s.c.commonHeaders()
	headers["Content-Type"] = "application/json"
	headers["Accept"] = restAPIv2

	r := httputils.Request{
		Method:           "GET",
		Url:              url,
		Body:             nil,
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
