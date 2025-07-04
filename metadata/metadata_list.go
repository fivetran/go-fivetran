package metadata

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type MetadataListService struct {
	httputils.HttpService
	limit   *int
	cursor  *string
}

func (s *MetadataListService) Limit(value int) *MetadataListService {
	s.limit = &value
	return s
}

func (s *MetadataListService) Cursor(value string) *MetadataListService {
	s.cursor = &value
	return s
}

func (s *MetadataListService) Do(ctx context.Context) (ConnectorMetadataListResponse, error) {
	var response ConnectorMetadataListResponse

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
	err := s.HttpService.Do(ctx, "GET", "/metadata/connector-types", nil, queries, 200, &response)
	return response, err
}