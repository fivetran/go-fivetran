package transformations

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type TransformationsListService struct {
	httputils.HttpService
	limit  *int
	cursor *string
}

func (s *TransformationsListService) Limit(value int) *TransformationsListService {
	s.limit = &value
	return s
}

func (s *TransformationsListService) Cursor(value string) *TransformationsListService {
	s.cursor = &value
	return s
}

func (s *TransformationsListService) Do(ctx context.Context) (TransformationsListResponse, error) {
	var response TransformationsListResponse
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
	err := s.HttpService.Do(ctx, "GET", "/transformations", nil, queries, 200, &response)
	return response, err
}
