package fivetran

import (
	"context"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

func (c *Client) NewHttpService() httputils.HttpService {
	return httputils.HttpService{
		Method:           "POST",
		CommonHeaders:    c.commonHeadersByMethod("POST"),
		BaseUrl:          c.baseURL,
		MaxRetryAttempts: c.maxRetryAttempts,
		HandleRateLimits: c.handleRateLimits,
		Client:           c.httpClient,
		ExpectedStatus:   200,
	}
}

func (c *Client) NewConnectorSync() *ConnectorSyncService {
	return &ConnectorSyncService{
		HttpService: c.NewHttpService(),
	}
}

// ConnectorSyncService implements the Connector Management, Sync Connector Data API.
// Ref. https://fivetran.com/docs/rest-api/connectors#syncconnectordata
type ConnectorSyncService struct {
	connectorID *string
	httputils.HttpService
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

	err := s.HttpService.Do(ctx, fmt.Sprintf("/connectors/%v/force", *s.connectorID), nil, nil, &response)

	return response, err
}
