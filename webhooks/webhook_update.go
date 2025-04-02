package webhooks

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type WebhookUpdateService struct {
	httputils.HttpService
	webhookId *string
	url       *string
	events    *[]string
	active    *bool
	secret    *string
	runTests  *bool
}

func (s *WebhookUpdateService) request() *webhookUpdateRequest {
	return &webhookUpdateRequest{
		Url:    s.url,
		Events: s.events,
		Active: s.active,
		Secret: s.secret,
	}
}

func (s *WebhookUpdateService) Url(value string) *WebhookUpdateService {
	s.url = &value
	return s
}

func (s *WebhookUpdateService) Secret(value string) *WebhookUpdateService {
	s.secret = &value
	return s
}

func (s *WebhookUpdateService) Active(value bool) *WebhookUpdateService {
	s.active = &value
	return s
}

func (s *WebhookUpdateService) Events(value []string) *WebhookUpdateService {
	s.events = &value
	return s
}

func (s *WebhookUpdateService) WebhookId(value string) *WebhookUpdateService {
	s.webhookId = &value
	return s
}

func (s *WebhookUpdateService) RunTests(value bool) *WebhookUpdateService {
	s.runTests = &value
	return s
}

func (s *WebhookUpdateService) Do(ctx context.Context) (WebhookResponse, error) {
	var response WebhookResponse
	if s.webhookId == nil {
		return response, fmt.Errorf("missing required webhookId")
	}
	url := fmt.Sprintf("/webhooks/%v", *s.webhookId)
	err := s.HttpService.Do(ctx, "PATCH", url, s.request(), nil, 200, &response)
	return response, err
}
