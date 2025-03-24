package certificates

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type DestinationCertificatesListService struct {
	httputils.HttpService
	destinationID *string
	cursor        *string
	limit         *int
}

func (s *DestinationCertificatesListService) DestinationID(value string) *DestinationCertificatesListService {
	s.destinationID = &value
	return s
}

func (s *DestinationCertificatesListService) Cursor(value string) *DestinationCertificatesListService {
	s.cursor = &value
	return s
}

func (s *DestinationCertificatesListService) Limit(value int) *DestinationCertificatesListService {
	s.limit = &value
	return s
}

func (s *DestinationCertificatesListService) Do(ctx context.Context) (CertificatesListResponse, error) {
	var response CertificatesListResponse
	url := fmt.Sprintf("/destinations/%v/certificates", *s.destinationID)
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
