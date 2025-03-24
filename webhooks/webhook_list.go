package webhooks

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type WebhookListService struct {
	httputils.HttpService
	limit  *int
	cursor *string
}

func (s *WebhookListService) Limit(value int) *WebhookListService {
	s.limit = &value
	return s
}

func (s *WebhookListService) Cursor(value string) *WebhookListService {
	s.cursor = &value
	return s
}

func (s *WebhookListService) Do(ctx context.Context) (WebhookListResponse, error) {
	var response WebhookListResponse
	var queries map[string]string = nil
	if s.cursor != nil || s.limit != nil {
		queries = make(map[string]string)
		if s.cursor != nil {
			queries["cursor"] = *s.cursor
		}
		if s.limit != nil {
			queries["limit"] = fmt.Sprintf("%v", *s.limit)
		}
	}
	err := s.HttpService.Do(ctx, "GET", "/webhooks", nil, queries, 200, &response)
	return response, err
}