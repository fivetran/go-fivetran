package certificates

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// https://fivetran.com/docs/rest-api/certificates#listallapprovedcertificatesforconnector
type ConnectorCertificatesListService struct {
	httputils.HttpService
	connectorID *string
	cursor      *string
	limit       *int
}

func (s *ConnectorCertificatesListService) ConnectorID(value string) *ConnectorCertificatesListService {
	s.connectorID = &value
	return s
}

func (s *ConnectorCertificatesListService) Cursor(value string) *ConnectorCertificatesListService {
	s.cursor = &value
	return s
}

func (s *ConnectorCertificatesListService) Limit(value int) *ConnectorCertificatesListService {
	s.limit = &value
	return s
}

func (s *ConnectorCertificatesListService) Do(ctx context.Context) (CertificatesListResponse, error) {
	var response CertificatesListResponse
	url := fmt.Sprintf("/connectors/%v/certificates", *s.connectorID)
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
