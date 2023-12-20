package webhooks

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// WebhookDetailsService implements the Webhook Management, retrieve Webhook Details.
// Ref. https://fivetran.com/docs/rest-api/webhooks#retrievewebhookdetails
type WebhookDetailsService struct {
	httputils.HttpService
	webhookId *string
}

func (s *WebhookDetailsService) WebhookId(value string) *WebhookDetailsService {
	s.webhookId = &value
	return s
}

func (s *WebhookDetailsService) Do(ctx context.Context) (WebhookResponse, error) {
	var response WebhookResponse
	if s.webhookId == nil {
		return response, fmt.Errorf("missing required webhookId")
	}
	url := fmt.Sprintf("/webhooks/%v", *s.webhookId)
	err := s.HttpService.Do(ctx, "GET", url, nil, nil, 200, &response)
	return response, err
}