package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
)

// ConnectorDeleteService implements the Connector Management, Delete a Connector API.
// Ref. https://fivetran.com/docs/rest-api/connectors#deleteaconnector
type ConnectorDeleteService struct {
	c           *Client
	connectorID *string
}

type ConnectorDeleteResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (c *Client) NewConnectorDelete() *ConnectorDeleteService {
	return &ConnectorDeleteService{c: c}
}

func (s *ConnectorDeleteService) ConnectorID(connectorID string) *ConnectorDeleteService {
	s.connectorID = &connectorID
	return s
}

func (s *ConnectorDeleteService) Do(ctx context.Context) (ConnectorDeleteResponse, error) {
	var response ConnectorDeleteResponse

	if s.connectorID == nil {
		return response, fmt.Errorf("missing required ConnectorID")
	}

	url := fmt.Sprintf("%v/connectors/%v", s.c.baseURL, *s.connectorID)
	expectedStatus := 200

	headers := s.c.commonHeaders()

	r := request{
		method:           "DELETE",
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
