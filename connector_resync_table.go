package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
)

// F stands for Field
// needs to be exported because of json.Marshal()
type ConnectorReSyncTableService struct {
	c           *Client
	connectorID *string
	schema      *string
	table       *string
}

type ConnectorReSyncTable struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (c *Client) NewConnectorReSyncTableService() *ConnectorReSyncTableService {
	return &ConnectorReSyncTableService{c: c}
}

func (s *ConnectorReSyncTableService) ConnectorID(connectorID string) *ConnectorReSyncTableService {
	s.connectorID = &connectorID
	return s
}

func (s *ConnectorReSyncTableService) Schema(schema string) *ConnectorReSyncTableService {
	s.schema = &schema
	return s
}

func (s *ConnectorReSyncTableService) Table(table string) *ConnectorReSyncTableService {
	s.table = &table
	return s
}

func (s *ConnectorReSyncTableService) Do(ctx context.Context) (ConnectorReSyncTable, error) {
	if s.connectorID == nil { // we don't validate business rules (unless it is strictly necessary)
		err := fmt.Errorf("missing required ConnectorID")
		return ConnectorReSyncTable{}, err
	}

	if s.schema == nil { // we don't validate business rules (unless it is strictly necessary)
		err := fmt.Errorf("missing required Schema")
		return ConnectorReSyncTable{}, err
	}

	if s.table == nil { // we don't validate business rules (unless it is strictly necessary)
		err := fmt.Errorf("missing required Table")
		return ConnectorReSyncTable{}, err
	}

	url := fmt.Sprintf("%v/connectors/%v/schemas/%v/tables/%v/resync", s.c.baseURL, *s.connectorID, *s.schema, *s.table)
	expectedStatus := 200
	headers := make(map[string]string)

	headers["Authorization"] = s.c.authorization

	r := Request{
		method:  "POST",
		url:     url,
		body:    nil,
		queries: nil,
		headers: headers,
	}

	respBody, respStatus, err := httpRequest(r, ctx)
	if err != nil {
		return ConnectorReSyncTable{}, err
	}

	var connectorReSyncTable ConnectorReSyncTable
	if err := json.Unmarshal(respBody, &connectorReSyncTable); err != nil {
		return ConnectorReSyncTable{}, err
	}

	if respStatus != expectedStatus {
		err := fmt.Errorf("status code: %v; expected %v", respStatus, expectedStatus)
		return connectorReSyncTable, err
	}

	return connectorReSyncTable, nil
}
