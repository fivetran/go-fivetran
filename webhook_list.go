package fivetran

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fivetran/go-fivetran/webhooks"
)

// WebhookListService implements the Webhook Management, retrieve List Webhook.
// Ref. https://fivetran.com/docs/rest-api/webhooks#retrievethelistofwebhooks
type WebhookListService struct {
	c      *Client
	limit  *int
	cursor *string
}

func (c *Client) NewWebhookList() *WebhookListService {
	return &WebhookListService{c: c}
}

func (s *WebhookListService) Limit(value int) *WebhookListService {
	s.limit = &value
	return s
}

func (s *WebhookListService) Cursor(value string) *WebhookListService {
	s.cursor = &value
	return s
}

func (s *WebhookListService) Do(ctx context.Context) (webhooks.WebhookListResponse, error) {
	var response webhooks.WebhookListResponse

	url := fmt.Sprintf("%v/webhooks", s.c.baseURL)
	expectedStatus := 200

	headers := s.c.commonHeaders()

	queries := make(map[string]string)
	if s.cursor != nil {
		queries["cursor"] = *s.cursor
	}
	if s.limit != nil {
		queries["limit"] = fmt.Sprint(*s.limit)
	}

	r := request{
		method:           "GET",
		url:              url,
		body:             nil,
		queries:          queries,
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
