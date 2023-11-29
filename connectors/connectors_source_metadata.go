package connectors

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// ConnectorsSourceMetadataService implements the Connector Management, Retrieve source metadata API.
// Ref. https://fivetran.com/docs/rest-api/connectors#retrievesourcemetadata
type ConnectorsSourceMetadataService struct {
	httputils.HttpService
	limit  *int
	cursor *string
}

func (s *ConnectorsSourceMetadataService) Limit(value int) *ConnectorsSourceMetadataService {
	s.limit = &value
	return s
}

func (s *ConnectorsSourceMetadataService) Cursor(value string) *ConnectorsSourceMetadataService {
	s.cursor = &value
	return s
}

func (s *ConnectorsSourceMetadataService) Do(ctx context.Context) (ConnectorsSourceMetadataResponse, error) {
	var response ConnectorsSourceMetadataResponse
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

	err := s.HttpService.Do(ctx, "GET", "/metadata/connectors", nil, queries, 200, &response)
	return response, err
}