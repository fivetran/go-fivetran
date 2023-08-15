package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
)

// ConnectorSyncService implements the Connector Management, Sync Connector Data API.
// Ref. https://fivetran.com/docs/rest-api/connectors#syncconnectordata
type ConnectorSyncService struct {
	c           *Client
	connectorID *string
}

type ConnectorSyncResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (c *Client) NewConnectorSync() *ConnectorSyncService {
	return &ConnectorSyncService{c: c}
}

func (s *ConnectorSyncService) ConnectorID(connectorID string) *ConnectorSyncService {
	s.connectorID = &connectorID
	return s
}

func (s *ConnectorSyncService) Do(ctx context.Context) (ConnectorSyncResponse, error) {
	var response ConnectorSyncResponse

	if s.connectorID == nil {
		return response, fmt.Errorf("missing required ConnectorID")
	}

	url := fmt.Sprintf("%v/connectors/%v/force", s.c.baseURL, *s.connectorID)
	expectedStatus := 200

	headers := s.c.commonHeaders()
	headers["Content-Type"] = "application/json"

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
