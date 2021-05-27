package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
)

// F stands for Field
// needs to be exported because of json.Marshal()
type ConnectorSyncService struct {
	c           *Client
	connectorID *string
}

type ConnectorSync struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (c *Client) NewConnectorSyncService() *ConnectorSyncService {
	return &ConnectorSyncService{c: c}
}

func (s *ConnectorSyncService) ConnectorID(connectorID string) *ConnectorSyncService {
	s.connectorID = &connectorID
	return s
}

func (s *ConnectorSyncService) Do(ctx context.Context) (ConnectorSync, error) {
	if s.connectorID == nil { // we don't validate business rules (unless it is strictly necessary)
		err := fmt.Errorf("missing required ConnectorID")
		return ConnectorSync{}, err
	}

	url := fmt.Sprintf("%v/connectors/%v/force", s.c.baseURL, *s.connectorID)
	expectedStatus := 200
	headers := make(map[string]string)

	headers["Authorization"] = s.c.authorization
	headers["Content-Type"] = "application/json"

	reqBody, err := json.Marshal(s)
	if err != nil {
		return ConnectorSync{}, err
	}

	r := Request{
		method:  "POST",
		url:     url,
		body:    reqBody,
		queries: nil,
		headers: headers,
	}

	respBody, respStatus, err := httpRequest(r, ctx)
	if err != nil {
		return ConnectorSync{}, err
	}

	var connectorSync ConnectorSync
	if err := json.Unmarshal(respBody, &connectorSync); err != nil {
		return ConnectorSync{}, err
	}

	if respStatus != expectedStatus {
		err := fmt.Errorf("status code: %v; expected %v", respStatus, expectedStatus)
		return connectorSync, err
	}

	return connectorSync, nil
}
