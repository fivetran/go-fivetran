package connectors

import (
	"context"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// ConnectorDeleteService implements the Connector Management, Delete a Connector API.
// Ref. https://fivetran.com/docs/rest-api/connectors#deleteaconnector
type ConnectorDeleteService struct {
    httputils.HttpService
	connectorID *string
}

func (s *ConnectorDeleteService) ConnectorID(connectorID string) *ConnectorDeleteService {
	s.connectorID = &connectorID
	return s
}

func (s *ConnectorDeleteService) Do(ctx context.Context) (common.CommonResponse, error) {
	var response common.CommonResponse
	
	if s.connectorID == nil {
		return response, fmt.Errorf("missing required connectorID")
	}

	url := fmt.Sprintf("/connectors/%v", *s.connectorID)
	err := s.HttpService.Do(ctx, "DELETE", url, nil, nil, 200, &response)
	return response, err
}