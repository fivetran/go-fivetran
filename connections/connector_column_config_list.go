package connectors

import (
    "context"
    "fmt"

    httputils "github.com/fivetran/go-fivetran/http_utils"
)

// ConnectorColumnConfigListService implements the Connector Management, Retrieve Source Table Columns Config
// Ref. https://fivetran.com/docs/rest-api/connectors#retrievesourcetablecolumnsconfig
type ConnectorColumnConfigListService struct {
    httputils.HttpService
    connectorId         *string
    schema              *string
    table               *string
}

func (s *ConnectorColumnConfigListService) ConnectorId(value string) *ConnectorColumnConfigListService {
    s.connectorId = &value
    return s
}

func (s *ConnectorColumnConfigListService) Schema(value string) *ConnectorColumnConfigListService {
    s.schema = &value
    return s
}

func (s *ConnectorColumnConfigListService) Table(value string) *ConnectorColumnConfigListService {
    s.table = &value
    return s
}

func (s *ConnectorColumnConfigListService) Do(ctx context.Context) (ConnectorColumnConfigListResponse, error) {
    var response ConnectorColumnConfigListResponse
    url := fmt.Sprintf("/connectors/%v/schemas/%v/tables/%v/columns", *s.connectorId, *s.schema, *s.table)
    err := s.HttpService.Do(ctx, "GET", url, nil, nil, 200, &response)
    return response, err
}
