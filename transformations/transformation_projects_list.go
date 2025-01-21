package transformations

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type TransformationProjectsListService struct {
	httputils.HttpService
	limit  *int
	cursor *string
}

func (s *TransformationProjectsListService) Limit(value int) *TransformationProjectsListService {
	s.limit = &value
	return s
}

func (s *TransformationProjectsListService) Cursor(value string) *TransformationProjectsListService {
	s.cursor = &value
	return s
}

func (s *TransformationProjectsListService) Do(ctx context.Context) (TransformationProjectsListResponse, error) {
	var response TransformationProjectsListResponse
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
	err := s.HttpService.Do(ctx, "GET", "/transformation-projects", nil, queries, 200, &response)
	return response, err
}
