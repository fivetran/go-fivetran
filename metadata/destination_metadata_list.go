package metadata

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type DestinationMetadataListService struct {
	httputils.HttpService
	limit  *int
	cursor *string
}

func (s *DestinationMetadataListService) Limit(value int) *DestinationMetadataListService {
	s.limit = &value
	return s
}

func (s *DestinationMetadataListService) Cursor(value string) *DestinationMetadataListService {
	s.cursor = &value
	return s
}

func (s *DestinationMetadataListService) Do(ctx context.Context) (DestinationMetadataListResponse, error) {
	var response DestinationMetadataListResponse

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
	err := s.HttpService.Do(ctx, "GET", "/metadata/destination-types", nil, queries, 200, &response)
	return response, err
}
