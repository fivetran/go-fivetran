package proxy

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type ProxyListService struct {
	httputils.HttpService
	limit  *int
	cursor *string
}

func (s *ProxyListService) Limit(value int) *ProxyListService {
	s.limit = &value
	return s
}

func (s *ProxyListService) Cursor(value string) *ProxyListService {
	s.cursor = &value
	return s
}

func (s *ProxyListService) Do(ctx context.Context) (ProxyListResponse, error) {
	var response ProxyListResponse

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
	err := s.HttpService.Do(ctx, "GET", "/proxy", nil, queries, 200, &response)
	return response, err
}