package certificates

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type ConnectionCertificateApproveService struct {
	httputils.HttpService
	connectionID *string
	hash        *string
	encodedCert *string
}

func (s *ConnectionCertificateApproveService) request() *certificateApproveRequest {
	return &certificateApproveRequest{
		Hash:        s.hash,
		EncodedCert: s.encodedCert,
	}
}

func (s *ConnectionCertificateApproveService) ConnectionID(value string) *ConnectionCertificateApproveService {
	s.connectionID = &value
	return s
}

func (s *ConnectionCertificateApproveService) Hash(value string) *ConnectionCertificateApproveService {
	s.hash = &value
	return s
}

func (s *ConnectionCertificateApproveService) EncodedCert(value string) *ConnectionCertificateApproveService {
	s.encodedCert = &value
	return s
}

func (s *ConnectionCertificateApproveService) Do(ctx context.Context) (CertificateResponse, error) {
	var response CertificateResponse
	url := fmt.Sprintf("/connections/%v/certificates", *s.connectionID)
	err := s.HttpService.Do(ctx, "POST", url, s.request(), nil, 201, &response)
	return response, err
}
