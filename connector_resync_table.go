package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
)

type connectorReSyncTableService struct {
	c           *Client
	connectorID *string
	schema      *string
	table       *string
}

type ConnectorReSyncTableResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (c *Client) NewConnectorReSyncTable() *connectorReSyncTableService {
	return &connectorReSyncTableService{c: c}
}

func (s *connectorReSyncTableService) ConnectorID(value string) *connectorReSyncTableService {
	s.connectorID = &value
	return s
}

func (s *connectorReSyncTableService) Schema(value string) *connectorReSyncTableService {
	s.schema = &value
	return s
}

func (s *connectorReSyncTableService) Table(value string) *connectorReSyncTableService {
	s.table = &value
	return s
}

func (s *connectorReSyncTableService) Do(ctx context.Context) (ConnectorReSyncTableResponse, error) {
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
		return response, err
	}

	if err := json.Unmarshal(respBody, &response); err != nil {
		return response, err
	}

	if respStatus != expectedStatus {
		err := fmt.Errorf("status code: %v; expected %v", respStatus, expectedStatus)
		return response, err
	}

	return response, nil
}
