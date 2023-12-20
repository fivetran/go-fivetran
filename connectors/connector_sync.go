package connectors

import (
	"context"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// ConnectorSyncService implements the Connector Management, Sync Connector Data API.
// Ref. https://fivetran.com/docs/rest-api/connectors#syncconnectordata
type ConnectorSyncService struct {
	httputils.HttpService
	connectorID *string
}

func (s *ConnectorSyncService) ConnectorID(connectorID string) *ConnectorSyncService {
	s.connectorID = &connectorID
	return s
}

func (s *ConnectorSyncService) Do(ctx context.Context) (common.CommonResponse, error) {
	var response common.CommonResponse

	if s.connectorID == nil {
		return response, fmt.Errorf("missing required ConnectorID")
	}

	err := s.HttpService.Do(ctx, "POST", fmt.Sprintf("/connectors/%v/force", *s.connectorID), nil, nil, 200, &response)
	return response, err
}
