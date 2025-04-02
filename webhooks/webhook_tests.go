package webhooks

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type WebhookTestService struct {
	httputils.HttpService
	webhookId *string
	event     *string
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

	url := fmt.Sprintf("/webhooks/%v/test", *s.webhookId)
	err := s.HttpService.Do(ctx, "POST", url, s.request(), nil, 200, &response)
	return response, err
}