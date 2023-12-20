package dbt

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type DbtProjectsListService struct {
	httputils.HttpService
	limit  *int
	cursor *string
}

func (s *DbtProjectsListService) Limit(value int) *DbtProjectsListService {
	s.limit = &value
	return s
}

func (s *DbtProjectsListService) Cursor(value string) *DbtProjectsListService {
	s.cursor = &value
	return s
}

func (s *DbtProjectsListService) Do(ctx context.Context) (DbtProjectsListResponse, error) {
	var response DbtProjectsListResponse
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

	err := s.HttpService.Do(ctx, "GET", "/dbt/projects", nil, queries, 200, &response)
	return response, err
}
