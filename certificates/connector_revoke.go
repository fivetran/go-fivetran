package certificates

import (
	"context"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// CertificateConnectorCertificateApproveService implements the Certificate Management, Revoke certificate for connector API.
// Ref. https://fivetran.com/docs/rest-api/certificates#revokeaconnectorcertificate
type ConnectorCertificateRevokeService struct {
	httputils.HttpService
	connectorID *string
	hash        *string
}

func (s *ConnectorCertificateRevokeService) ConnectorID(value string) *ConnectorCertificateRevokeService {
	s.connectorID = &value
	return s
}

func (s *ConnectorCertificateRevokeService) Hash(value string) *ConnectorCertificateRevokeService {
	s.hash = &value
	return s
}

func (s *ConnectorCertificateRevokeService) Do(ctx context.Context) (common.CommonResponse, error) {
	var response common.CommonResponse
	url := fmt.Sprintf("/connectors/%v/certificates/%v", *s.connectorID, *s.hash)
	err := s.HttpService.Do(ctx, "DELETE", url, nil, nil, 200, &response)
	return response, err
}
