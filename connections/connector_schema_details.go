package connectors

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// ConnectorSchemaDetailsService implements the Connector Management, Retrieve a Connector Schema Config API.
// Ref. https://fivetran.com/docs/rest-api/connectors#retrieveaconnectorschemaconfig
type ConnectorSchemaDetailsService struct {
    httputils.HttpService
	connectorID *string
}

func (s *ConnectorSchemaDetailsService) ConnectorID(value string) *ConnectorSchemaDetailsService {
	s.connectorID = &value
	return s
}

func (s *ConnectorSchemaDetailsService) Do(ctx context.Context) (ConnectorSchemaDetailsResponse, error) {
	var response ConnectorSchemaDetailsResponse

	if s.connectorID == nil {
		return response, fmt.Errorf("missing required ConnectorID")
	}

	url := fmt.Sprintf("/connectors/%v/schemas", *s.connectorID)
	err := s.HttpService.Do(ctx, "GET", url, nil, nil, 200, &response)
	return response, err
}