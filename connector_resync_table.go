package fivetran

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// ConnectorReSyncTableService implements the Connector Management, Re-sync Connector Table Data API.
// Ref. https://fivetran.com/docs/rest-api/connectors#resyncconnectortabledata
type ConnectorReSyncTableService struct {
	c           *Client
	connectorID *string
	schema      *string
	table       *string
}

func (c *Client) NewConnectorReSyncTable() *ConnectorReSyncTableService {
	return &ConnectorReSyncTableService{c: c}
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

	url := fmt.Sprintf("%v/connectors/%v/schemas/%v/tables/%v/resync", s.c.baseURL, *s.connectorID, *s.schema, *s.table)
	expectedStatus := 200

	headers := s.c.commonHeaders()

	r := httputils.Request{
		Method:           "POST",
		Url:              url,
		Body:             nil,
		Queries:          nil,
		Headers:          headers,
		Client:           s.c.httpClient,
		HandleRateLimits: s.c.handleRateLimits,
		MaxRetryAttempts: s.c.maxRetryAttempts,
	}

	respBody, respStatus, err := r.Do(ctx)
	if err != nil {
		return response, err
	}

	if err := json.Unmarshal(respBody, &response); err != nil {
		return response, err
	}

	if respStatus != expectedStatus {
		err := fmt.Errorf("status code: %v; expected: %v", respStatus, expectedStatus)
		return response, err
	}

	return response, nil
}
