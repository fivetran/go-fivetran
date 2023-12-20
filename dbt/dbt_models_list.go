package dbt

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type DbtModelsListService struct {
	httputils.HttpService
	projectId *string
	limit     *int
	cursor    *string
}

func (s *DbtModelsListService) ProjectId(value string) *DbtModelsListService {
	s.projectId = &value
	return s
}

func (s *DbtModelsListService) Limit(value int) *DbtModelsListService {
	s.limit = &value
	return s
}

func (s *DbtModelsListService) Cursor(value string) *DbtModelsListService {
	s.cursor = &value
	return s
}

func (s *DbtModelsListService) Do(ctx context.Context) (DbtModelsListResponse, error) {
	var response DbtModelsListResponse
	var queries map[string]string = nil
	if s.cursor != nil || s.limit != nil || s.projectId != nil {
		queries = make(map[string]string)
		if s.cursor != nil {
			queries["cursor"] = *s.cursor
		}
		if s.limit != nil {
			queries["limit"] = fmt.Sprintf("%v", *s.limit)
		}
		if s.projectId != nil {
			queries["project_id"] = *s.projectId
		}
	}

	err := s.HttpService.Do(ctx, "GET", "/dbt/models", nil, queries, 200, &response)
	return response, err
}