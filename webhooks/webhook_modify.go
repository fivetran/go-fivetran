package webhooks

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// WebhookModifyService implements the Webhook Management, Modify a Webhook.
// Ref. https://fivetran.com/docs/rest-api/webhooks#updatewebhook
type WebhookModifyService struct {
	httputils.HttpService
	webhookId *string
	url       *string
	events    *[]string
	active    *bool
	secret    *string
	runTests  *bool
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

func (s *WebhookModifyService) Do(ctx context.Context) (WebhookResponse, error) {
	var response WebhookResponse
	if s.webhookId == nil {
		return response, fmt.Errorf("missing required webhookId")
	}
	url := fmt.Sprintf("/webhooks/%v", *s.webhookId)
	err := s.HttpService.Do(ctx, "PATCH", url, s.request(), nil, 200, &response)
	return response, err
}
