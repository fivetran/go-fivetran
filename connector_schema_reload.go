package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
)

// ConnectorSchemaReloadService implements the Connector Management, Reload a Connector Schema Config API.
// Ref. https://fivetran.com/docs/rest-api/connectors#reloadaconnectorschemaconfig
type ConnectorSchemaReloadService struct {
	c           *Client
	connectorID *string
	excludeMode *string
}

type connectorSchemaReloadRequest struct {
	ExcludeMode *string `json:"exclude_mode,omitempty"`
}

func (c *Client) NewConnectorSchemaReload() *ConnectorSchemaReloadService {
	return &ConnectorSchemaReloadService{c: c}
}

func (s *ConnectorSchemaReloadService) request() *connectorSchemaReloadRequest {
	mode := "PRESERVE"
	if s.excludeMode != nil {
		mode = *s.excludeMode
	}
	return &connectorSchemaReloadRequest{ExcludeMode: &mode}
}

func (s *ConnectorSchemaReloadService) ConnectorID(value string) *ConnectorSchemaReloadService {
	s.connectorID = &value
	return s
}

func (s *ConnectorSchemaReloadService) ExcludeMode(value string) *ConnectorSchemaReloadService {
	s.excludeMode = &value
	return s
}

func (s *ConnectorSchemaReloadService) Do(ctx context.Context) (ConnectorSchemaDetailsResponse, error) {
	var response ConnectorSchemaDetailsResponse

	if s.connectorID == nil {
		return response, fmt.Errorf("missing required ConnectorID")
	}

	reqBody, err := json.Marshal(s.request())
	if err != nil {
		return response, err
	}

	url := fmt.Sprintf("%v/connectors/%v/schemas/reload", s.c.baseURL, *s.connectorID)
	expectedStatus := 200

	headers := s.c.commonHeaders()
	headers["Content-Type"] = "application/json"
	headers["Accept"] = restAPIv2

	r := request{
		method:  "POST",
		url:     url,
		body:    reqBody,
		queries: nil,
		headers: headers,
		client:  s.c.httpClient,
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
