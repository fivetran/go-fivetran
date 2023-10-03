package certificates

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// CertificateConnectorCertificateApproveService implements the Certificate Management, Approve a connector certificate API.
// Ref. https://fivetran.com/docs/rest-api/certificates#approveaconnectorcertificate
type ConnectorCertificateApproveService struct {
	httputils.HttpService
	connectorID *string
	hash        *string
	encodedCert *string
}

func (s *ConnectorCertificateApproveService) request() *certificateApproveRequest {
	return &certificateApproveRequest{
		Hash:        s.hash,
		EncodedCert: s.encodedCert,
	}
}

func (s *ConnectorCertificateApproveService) ConnectorID(value string) *ConnectorCertificateApproveService {
	s.connectorID = &value
	return s
}

func (s *ConnectorCertificateApproveService) Hash(value string) *ConnectorCertificateApproveService {
	s.hash = &value
	return s
}

func (s *ConnectorCertificateApproveService) EncodedCert(value string) *ConnectorCertificateApproveService {
	s.encodedCert = &value
	return s
}

func (s *ConnectorCertificateApproveService) Do(ctx context.Context) (CertificateResponse, error) {
	var response CertificateResponse
	url := fmt.Sprintf("/connectors/%v/certificates", *s.connectorID)
	err := s.HttpService.Do(ctx, "POST", url, s.request(), nil, 201, &response)
	return response, err
}
