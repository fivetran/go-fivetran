package connectors

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// ConnectorSchemaReloadService implements the Connector Management, Reload a Connector Schema Config API.
// Ref. https://fivetran.com/docs/rest-api/connectors#reloadaconnectorschemaconfig
type ConnectorSchemaReloadService struct {
    httputils.HttpService
	connectorID *string
	excludeMode *string
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
	
	url := fmt.Sprintf("/connectors/%v/schemas/reload", *s.connectorID)
	err := s.HttpService.Do(ctx, "POST", url, s.request(), nil, 200, &response)
	return response, err
}