package fivetran

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
)

// WebhookDeleteService implements the Webhook Management, Delete a Webhook.
// Ref. https://fivetran.com/docs/rest-api/webhooks#deletewebhook
type WebhookDeleteService struct {
	c         *Client
	webhookId *string
}

func (c *Client) NewWebhookDelete() *WebhookDeleteService {
	return &WebhookDeleteService{c: c}
}

func (s *WebhookDeleteService) WebhookId(value string) *WebhookDeleteService {
	s.webhookId = &value
	return s
}

func (s *WebhookDeleteService) Do(ctx context.Context) (common.CommonResponse, error) {
	var response common.CommonResponse

	if s.webhookId == nil {
		return response, fmt.Errorf("missing required WebhookId")
	}

	url := fmt.Sprintf("%v/webhooks/%v", s.c.baseURL, *s.webhookId)
	expectedStatus := 200

	headers := s.c.commonHeaders()

	r := request{
		method:  "DELETE",
		url:     url,
		body:    nil,
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
