package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
)

// ConnectorReSyncTableService implements the Connector Management, Re-sync Connector Table Data API.
// Ref. https://fivetran.com/docs/rest-api/connectors#resyncconnectortabledata
type ConnectorReSyncTableService struct {
	c           *Client
	connectorID *string
	schema      *string
	table       *string
}

type ConnectorReSyncTableResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
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

func (s *ConnectorReSyncTableService) Do(ctx context.Context) (ConnectorReSyncTableResponse, error) {
	var response ConnectorReSyncTableResponse

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

	r := request{
		method:           "POST",
		url:              url,
		body:             nil,
		queries:          nil,
		headers:          headers,
		client:           s.c.httpClient,
		handleRateLimits: s.c.handleRateLimits,
		maxRetryAttempts: s.c.maxRetryAttempts,
	}

	respBody, respStatus, err := r.httpRequest(ctx)
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
