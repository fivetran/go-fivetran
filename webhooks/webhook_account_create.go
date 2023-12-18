package webhooks

import (
	"context"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// WebhookAccountCreateService implements the Webhook Management, Create a Account Webhook.
// Ref. https://fivetran.com/docs/rest-api/webhooks#createaccountwebhook
type WebhookAccountCreateService struct {
	httputils.HttpService
	url    *string
	events *[]string
	active *bool
	secret *string
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

func (s *WebhookAccountCreateService) Do(ctx context.Context) (WebhookResponse, error) {
	var response WebhookResponse
	err := s.HttpService.Do(ctx, "POST", "/webhooks/account", s.request(), nil, 200, &response)
	return response, err
}
