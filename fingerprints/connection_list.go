package fingerprints

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type ConnectionFingerprintsListService struct {
	httputils.HttpService
	connectionID *string
	cursor      *string
	limit       *int
}

func (s *ConnectionFingerprintsListService) ConnectionID(value string) *ConnectionFingerprintsListService {
	s.connectionID = &value
	return s
}

func (s *ConnectionFingerprintsListService) Cursor(value string) *ConnectionFingerprintsListService {
	s.cursor = &value
	return s
}

func (s *ConnectionFingerprintsListService) Limit(value int) *ConnectionFingerprintsListService {
	s.limit = &value
	return s
}

func (s *ConnectionFingerprintsListService) Do(ctx context.Context) (FingerprintsListResponse, error) {
	var response FingerprintsListResponse
	url := fmt.Sprintf("/connections/%v/fingerprints", *s.connectionID)
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
