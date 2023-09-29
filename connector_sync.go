package fivetran

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// ConnectorSyncService implements the Connector Management, Sync Connector Data API.
// Ref. https://fivetran.com/docs/rest-api/connectors#syncconnectordata
type ConnectorSyncService struct {
	c           *Client
	connectorID *string
}

func (c *Client) NewConnectorSync() *ConnectorSyncService {
	return &ConnectorSyncService{c: c}
}

func (s *ConnectorSyncService) ConnectorID(connectorID string) *ConnectorSyncService {
	s.connectorID = &connectorID
	return s
}

func (s *ConnectorSyncService) Do(ctx context.Context) (common.CommonResponse, error) {
	var response common.CommonResponse

	if s.connectorID == nil {
		return response, fmt.Errorf("missing required ConnectorID")
	}

	url := fmt.Sprintf("%v/connectors/%v/force", s.c.baseURL, *s.connectorID)
	expectedStatus := 200

	headers := s.c.commonHeaders()
	headers["Content-Type"] = "application/json"

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
