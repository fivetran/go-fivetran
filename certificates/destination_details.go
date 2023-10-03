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

func NewDestinationCertificateDetailsRequestParams() httputils.HttpParams {
	return httputils.HttpParams{
		Method:         "GET",
		ExpectedStatus: 200,
	}
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

	err := s.HttpService.Do(ctx, url, nil, nil, &response)
	return response, err
}
