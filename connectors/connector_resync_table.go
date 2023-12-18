package connectors

import (
	"context"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// ConnectorReSyncTableService implements the Connector Management, Re-sync Connector Table Data API.
// Ref. https://fivetran.com/docs/rest-api/connectors#resyncconnectortabledata
type ConnectorReSyncTableService struct {
    httputils.HttpService
	connectorID *string
	schema      *string
	table       *string
}

func (s *ConnectorReSyncTableService) ConnectorID(value string) *ConnectorReSyncTableService {
	s.connectorID = &value
	return s
}

func (s *ConnectorReSyncTableService) Schema(value string) *ConnectorReSyncTableService {
	s.schema = &value
	return s
}

func (s *ConnectorReSyncTableService) Table(value string) *ConnectorReSyncTableService {
	s.table = &value
	return s
}

func (s *ConnectorReSyncTableService) Do(ctx context.Context) (common.CommonResponse, error) {
	var response common.CommonResponse

	if s.connectorID == nil {
		return response, fmt.Errorf("missing required ConnectorID")
	}
	if s.schema == nil {
		return response, fmt.Errorf("missing required Schema")
	}
	if s.table == nil {
		return response, fmt.Errorf("missing required Table")
	}
	
	url := fmt.Sprintf("/connectors/%v/schemas/%v/tables/%v/resync", *s.connectorID, *s.schema, *s.table)
	err := s.HttpService.Do(ctx, "POST", url, nil, nil, 200, &response)
	return response, err
}