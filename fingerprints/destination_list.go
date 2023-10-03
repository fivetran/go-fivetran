package fingerprints

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type DestinationFingerprintsListService struct {
	httputils.HttpService
	destinationID *string
	cursor        *string
	limit         *int
}

func (s *DestinationFingerprintsListService) DestinationID(value string) *DestinationFingerprintsListService {
	s.destinationID = &value
	return s
}

func (s *DestinationFingerprintsListService) Cursor(value string) *DestinationFingerprintsListService {
	s.cursor = &value
	return s
}

func (s *DestinationFingerprintsListService) Limit(value int) *DestinationFingerprintsListService {
	s.limit = &value
	return s
}

func (s *DestinationFingerprintsListService) Do(ctx context.Context) (FingerprintsListResponse, error) {
	var response FingerprintsListResponse
	url := fmt.Sprintf("/destinations/%v/fingerprints", *s.destinationID)
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
