package certificates

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type ConnectionCertificateDetailsService struct {
	httputils.HttpService
	connectionID *string
	hash        *string
}

func (s *ConnectionCertificateDetailsService) ConnectionID(value string) *ConnectionCertificateDetailsService {
	s.connectionID = &value
	return s
}

func (s *ConnectionCertificateDetailsService) Hash(value string) *ConnectionCertificateDetailsService {
	s.hash = &value
	return s
}

func (s *ConnectionCertificateDetailsService) Do(ctx context.Context) (CertificateResponse, error) {
	var response CertificateResponse
	url := fmt.Sprintf("/connections/%v/certificates/%v", *s.connectionID, *s.hash)
	err := s.HttpService.Do(ctx, "GET", url, nil, nil, 200, &response)
	return response, err
}
