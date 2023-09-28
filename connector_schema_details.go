package fivetran

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fivetran/go-fivetran/connectors"
)

// ConnectorSchemaDetailsService implements the Connector Management, Retrieve a Connector Schema Config API.
// Ref. https://fivetran.com/docs/rest-api/connectors#retrieveaconnectorschemaconfig
type ConnectorSchemaDetailsService struct {
	c           *Client
	connectorID *string
}

func (c *Client) NewConnectorSchemaDetails() *ConnectorSchemaDetailsService {
	return &ConnectorSchemaDetailsService{c: c}
}

func (s *ConnectorSchemaDetailsService) ConnectorID(value string) *ConnectorSchemaDetailsService {
	s.connectorID = &value
	return s
}

func (s *ConnectorSchemaDetailsService) Do(ctx context.Context) (connectors.ConnectorSchemaDetailsResponse, error) {
	var response connectors.ConnectorSchemaDetailsResponse

	if s.connectorID == nil {
		return response, fmt.Errorf("missing required ConnectorID")
	}

	url := fmt.Sprintf("%v/connectors/%v/schemas", s.c.baseURL, *s.connectorID)
	expectedStatus := 200

	headers := s.c.commonHeaders()
	headers["Accept"] = restAPIv2

	r := request{
		method:           "GET",
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
