package fingerprints

import (
	"context"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type ConnectorFingerprintRevokeService struct {
	httputils.HttpService
	connectorID *string
	hash        *string
}

func (s *ConnectorFingerprintRevokeService) ConnectorID(value string) *ConnectorFingerprintRevokeService {
	s.connectorID = &value
	return s
}

func (s *ConnectorFingerprintRevokeService) Hash(value string) *ConnectorFingerprintRevokeService {
	s.hash = &value
	return s
}

func (s *ConnectorFingerprintRevokeService) Do(ctx context.Context) (common.CommonResponse, error) {
	var response common.CommonResponse
	url := fmt.Sprintf("/connectors/%v/fingerprints/%v", *s.connectorID, *s.hash)
	err := s.HttpService.Do(ctx, "DELETE", url, nil, nil, 200, &response)
	return response, err
}
