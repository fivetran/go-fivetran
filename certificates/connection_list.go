package certificates

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type ConnectionCertificatesListService struct {
	httputils.HttpService
	connectionID *string
	cursor      *string
	limit       *int
}

func (s *ConnectionCertificatesListService) ConnectionID(value string) *ConnectionCertificatesListService {
	s.connectionID = &value
	return s
}

func (s *ConnectionCertificatesListService) Cursor(value string) *ConnectionCertificatesListService {
	s.cursor = &value
	return s
}

func (s *ConnectionCertificatesListService) Limit(value int) *ConnectionCertificatesListService {
	s.limit = &value
	return s
}

func (s *ConnectionCertificatesListService) Do(ctx context.Context) (CertificatesListResponse, error) {
	var response CertificatesListResponse
	url := fmt.Sprintf("/connections/%v/certificates", *s.connectionID)
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
