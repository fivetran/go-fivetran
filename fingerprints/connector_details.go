package fingerprints

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type ConnectorFingerprintDetailsService struct {
	httputils.HttpService
	connectorID *string
	hash        *string
}

func (s *ConnectorFingerprintDetailsService) ConnectorID(value string) *ConnectorFingerprintDetailsService {
	s.connectorID = &value
	return s
}

func (s *ConnectorFingerprintDetailsService) Hash(value string) *ConnectorFingerprintDetailsService {
	s.hash = &value
	return s
}

func (s *ConnectorFingerprintDetailsService) Do(ctx context.Context) (FingerprintResponse, error) {
	var response FingerprintResponse
	url := fmt.Sprintf("/connectors/%v/fingerprints/%v", *s.connectorID, *s.hash)
	err := s.HttpService.Do(ctx, "GET", url, nil, nil, 200, &response)
	return response, err
}
