package certificates

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// CertificateConnectorCertificateApproveService implements the Certificate Management, Approve a connector certificate API.
// Ref. https://fivetran.com/docs/rest-api/certificates#approveaconnectorcertificate
type DestinationCertificateApproveService struct {
	httputils.HttpService
	destinationID *string
	hash          *string
	encodedCert   *string
}

func (s *DestinationCertificateApproveService) request() *certificateApproveRequest {
	return &certificateApproveRequest{
		Hash:        s.hash,
		EncodedCert: s.encodedCert,
	}
}

func (s *DestinationCertificateApproveService) DestinationID(value string) *DestinationCertificateApproveService {
	s.destinationID = &value
	return s
}

func (s *DestinationCertificateApproveService) Hash(value string) *DestinationCertificateApproveService {
	s.hash = &value
	return s
}

func (s *DestinationCertificateApproveService) EncodedCert(value string) *DestinationCertificateApproveService {
	s.encodedCert = &value
	return s
}

func (s *DestinationCertificateApproveService) Do(ctx context.Context) (CertificateResponse, error) {
	var response CertificateResponse
	url := fmt.Sprintf("/destinations/%v/certificates", *s.destinationID)
	err := s.HttpService.Do(ctx, "POST", url, s.request(), nil, 201, &response)
	return response, err
}
