package fingerprints

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// CertificateConnectorFingerprintApproveService implements the Certificate Management, Approve a connector fingerprint API.
// Ref. https://fivetran.com/docs/rest-api/certificates#approveaconnectorfingerprint
type ConnectorFingerprintApproveService struct {
	httputils.HttpService
	connectorID *string
	hash        *string
	publicKey   *string
}

type connectorFingerprintApproveRequest struct {
	Hash      *string `json:"hash,omitempty"`
	PublicKey *string `json:"public_key,omitempty"`
}

func (s *ConnectorFingerprintApproveService) request() *connectorFingerprintApproveRequest {
	return &connectorFingerprintApproveRequest{
		Hash:      s.hash,
		PublicKey: s.publicKey,
	}
}

func (s *ConnectorFingerprintApproveService) ConnectorID(value string) *ConnectorFingerprintApproveService {
	s.connectorID = &value
	return s
}

func (s *ConnectorFingerprintApproveService) Hash(value string) *ConnectorFingerprintApproveService {
	s.hash = &value
	return s
}

func (s *ConnectorFingerprintApproveService) PublicKey(value string) *ConnectorFingerprintApproveService {
	s.publicKey = &value
	return s
}

func (s *ConnectorFingerprintApproveService) Do(ctx context.Context) (FingerprintResponse, error) {
	var response FingerprintResponse
	url := fmt.Sprintf("/connectors/%v/fingerprints", *s.connectorID)
	err := s.HttpService.Do(ctx, "POST", url, s.request(), nil, 201, &response)
	return response, err
}
