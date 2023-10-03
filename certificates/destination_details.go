package certificates

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type DestinationCertificateDetailsService struct {
	httputils.HttpService
	destinationID *string
	hash          *string
}

func (s *DestinationCertificateDetailsService) DestinationID(value string) *DestinationCertificateDetailsService {
	s.destinationID = &value
	return s
}

func (s *DestinationCertificateDetailsService) Hash(value string) *DestinationCertificateDetailsService {
	s.hash = &value
	return s
}

func (s *DestinationCertificateDetailsService) Do(ctx context.Context) (CertificateResponse, error) {
	var response CertificateResponse
	url := fmt.Sprintf("/destinations/%v/certificates/%v", *s.destinationID, *s.hash)

	err := s.HttpService.Do(ctx, "GET", url, nil, nil, 200, &response)
	return response, err
}
