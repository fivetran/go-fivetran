package transformations

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type QuickstartPackagesListService struct {
	httputils.HttpService
	limit  *int
	cursor *string
}

func (s *QuickstartPackagesListService) Limit(value int) *QuickstartPackagesListService {
	s.limit = &value
	return s
}

func (s *QuickstartPackagesListService) Cursor(value string) *QuickstartPackagesListService {
	s.cursor = &value
	return s
}

func (s *QuickstartPackagesListService) Do(ctx context.Context) (QuickstartPackagesListResponse, error) {
	var response QuickstartPackagesListResponse
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
	err := s.HttpService.Do(ctx, "GET", "/transformations/package-metadata", nil, queries, 200, &response)
	return response, err
}
