package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
)

type connectorSyncService struct {
	c           *Client
	connectorID *string
}

type ConnectorSyncResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (c *Client) NewConnectorSync() *connectorSyncService {
	return &connectorSyncService{c: c}
}

func (s *connectorSyncService) ConnectorID(connectorID string) *connectorSyncService {
	s.connectorID = &connectorID
	return s
}

func (s *connectorSyncService) Do(ctx context.Context) (ConnectorSyncResponse, error) {
	var response ConnectorSyncResponse

	if s.connectorID == nil {
		return response, fmt.Errorf("missing required ConnectorID")
	}

	url := fmt.Sprintf("%v/connectors/%v/force", s.c.baseURL, *s.connectorID)
	expectedStatus := 200

	headers := make(map[string]string)
	headers["Authorization"] = s.c.authorization
	headers["Content-Type"] = "application/json"

	reqBody, err := json.Marshal(s)
	if err != nil {
		return response, err
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
