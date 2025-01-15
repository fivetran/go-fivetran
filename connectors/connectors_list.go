package connectors

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type ConnectorsListService struct {
	httputils.HttpService
	limit  *int
	cursor *string
}

func (s *ConnectorsListService) Limit(value int) *ConnectorsListService {
	s.limit = &value
	return s
}

func (s *ConnectorsListService) Cursor(value string) *ConnectorsListService {
	s.cursor = &value
	return s
}

func (s *ConnectorsListService) Do(ctx context.Context) (ConnectorsListResponse, error) {
	var response ConnectorsListResponse
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
	err := s.HttpService.Do(ctx, "GET", "/connectors", nil, queries, 200, &response)
	return response, err
}
