package certificates

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type ConnectorCertificateDetailsService struct {
	httputils.HttpService
	connectorID *string
	hash        *string
}

func NewConnectorCertificateDetailsRequestParams() httputils.HttpParams {
	return httputils.HttpParams{
		Method:         "GET",
		ExpectedStatus: 200,
	}
}

func (s *ConnectorCertificateDetailsService) ConnectorID(value string) *ConnectorCertificateDetailsService {
	s.connectorID = &value
	return s
}

func (s *ConnectorCertificateDetailsService) Hash(value string) *ConnectorCertificateDetailsService {
	s.hash = &value
	return s
}

func (s *ConnectorCertificateDetailsService) Do(ctx context.Context) (CertificateResponse, error) {
	var response CertificateResponse
	url := fmt.Sprintf("/connectors/%v/certificates/%v", *s.connectorID, *s.hash)

	err := s.HttpService.Do(ctx, url, nil, nil, &response)
	return response, err
}
