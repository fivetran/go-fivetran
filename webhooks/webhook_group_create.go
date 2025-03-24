package webhooks

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type WebhookGroupCreateService struct {
	httputils.HttpService
	groupId *string
	url     *string
	events  *[]string
	active  *bool
	secret  *string
}

func (s *WebhookGroupCreateService) request() *webhookGroupCreateRequest {
	return &webhookGroupCreateRequest{
		Url:    s.url,
		Events: s.events,
		Active: s.active,
		Secret: s.secret,
	}
}

func (s *WebhookGroupCreateService) Url(value string) *WebhookGroupCreateService {
	s.url = &value
	return s
}

func (s *WebhookGroupCreateService) Secret(value string) *WebhookGroupCreateService {
	s.secret = &value
	return s
}

func (s *WebhookGroupCreateService) Active(value bool) *WebhookGroupCreateService {
	s.active = &value
	return s
}

func (s *WebhookGroupCreateService) Events(value []string) *WebhookGroupCreateService {
	s.events = &value
	return s
}

func (s *WebhookGroupCreateService) GroupId(value string) *WebhookGroupCreateService {
	s.groupId = &value
	return s
}

func (s *WebhookGroupCreateService) Do(ctx context.Context) (WebhookResponse, error) {
	var response WebhookResponse
	if s.groupId == nil {
		return response, fmt.Errorf("missing required groupId")
	}
	url := fmt.Sprintf("/webhooks/group/%v", *s.groupId)
	err := s.HttpService.Do(ctx, "POST", url, s.request(), nil, 200, &response)
	return response, err
}
