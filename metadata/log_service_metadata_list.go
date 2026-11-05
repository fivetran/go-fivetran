package metadata

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type LogServiceMetadataListService struct {
	httputils.HttpService
	limit  *int
	cursor *string
}

func (s *LogServiceMetadataListService) Limit(value int) *LogServiceMetadataListService {
	s.limit = &value
	return s
}

func (s *LogServiceMetadataListService) Cursor(value string) *LogServiceMetadataListService {
	s.cursor = &value
	return s
}

func (s *LogServiceMetadataListService) Do(ctx context.Context) (LogServiceMetadataListResponse, error) {
	var response LogServiceMetadataListResponse

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
	err := s.HttpService.Do(ctx, "GET", "/metadata/log-service-types", nil, queries, 200, &response)
	return response, err
}
