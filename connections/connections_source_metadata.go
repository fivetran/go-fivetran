package connections

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type ConnectionsSourceMetadataService struct {
	httputils.HttpService
	limit  *int
	cursor *string
}

func (s *ConnectionsSourceMetadataService) Limit(value int) *ConnectionsSourceMetadataService {
	s.limit = &value
	return s
}

func (s *ConnectionsSourceMetadataService) Cursor(value string) *ConnectionsSourceMetadataService {
	s.cursor = &value
	return s
}

func (s *ConnectionsSourceMetadataService) Do(ctx context.Context) (ConnectionsSourceMetadataResponse, error) {
	var response ConnectionsSourceMetadataResponse
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

	err := s.HttpService.Do(ctx, "GET", "/metadata/connections", nil, queries, 200, &response)
	return response, err
}