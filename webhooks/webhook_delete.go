package webhooks

import (
	"context"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// WebhookDeleteService implements the Webhook Management, Delete a Webhook.
// Ref. https://fivetran.com/docs/rest-api/webhooks#deletewebhook
type WebhookDeleteService struct {
	httputils.HttpService
	webhookId *string
}

func (s *WebhookDeleteService) WebhookId(value string) *WebhookDeleteService {
	s.webhookId = &value
	return s
}

func (s *WebhookDeleteService) Do(ctx context.Context) (common.CommonResponse, error) {
	var response common.CommonResponse
	
	if s.webhookId == nil {
		return response, fmt.Errorf("missing required webhookId")
	}

	url := fmt.Sprintf("/webhooks/%v", *s.webhookId)
	err := s.HttpService.Do(ctx, "DELETE", url, nil, nil, 200, &response)
	return response, err
}