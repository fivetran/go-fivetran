package fingerprints

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type ConnectorFingerprintsListService struct {
	httputils.HttpService
	connectorID *string
	cursor      *string
	limit       *int
}

func (s *ConnectorFingerprintsListService) ConnectorID(value string) *ConnectorFingerprintsListService {
	s.connectorID = &value
	return s
}

func (s *ConnectorFingerprintsListService) Cursor(value string) *ConnectorFingerprintsListService {
	s.cursor = &value
	return s
}

func (s *ConnectorFingerprintsListService) Limit(value int) *ConnectorFingerprintsListService {
	s.limit = &value
	return s
}

func (s *ConnectorFingerprintsListService) Do(ctx context.Context) (FingerprintsListResponse, error) {
	var response FingerprintsListResponse
	url := fmt.Sprintf("/connectors/%v/fingerprints", *s.connectorID)
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
	err := s.HttpService.Do(ctx, "GET", url, nil, queries, 200, &response)
	return response, err
}
