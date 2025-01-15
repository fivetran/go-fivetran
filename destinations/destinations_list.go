package destinations

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type DestinationsListService struct {
	httputils.HttpService
	limit  *int
	cursor *string
}

func (s *DestinationsListService) Limit(value int) *DestinationsListService {
	s.limit = &value
	return s
}

func (s *DestinationsListService) Cursor(value string) *DestinationsListService {
	s.cursor = &value
	return s
}

func (s *DestinationsListService) Do(ctx context.Context) (DestinationsListResponse, error) {
	var response DestinationsListResponse
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
	err := s.HttpService.Do(ctx, "GET", "/destinations", nil, queries, 200, &response)
	return response, err
}
