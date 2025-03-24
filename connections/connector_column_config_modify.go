package connectors

import (
    "context"
    "fmt"

    httputils "github.com/fivetran/go-fivetran/http_utils"
)

// ConnectorColumnConfigModifyService implements the Connector Management, Modify a Connector Table Config
// Ref. https://fivetran.com/docs/rest-api/connectors#modifyaconnectorcolumnconfig
type ConnectorColumnConfigModifyService struct {
    httputils.HttpService
    connectorId         *string
    schema              *string
    table               *string
    column              *string
    enabled             *bool
    hashed              *bool
}

func (csu *ConnectorColumnConfigModifyService) request() *connectorColumnConfigModifyRequest {
    return &connectorColumnConfigModifyRequest{
        Enabled:       csu.enabled,
        Hashed:        csu.hashed,
    }
}

func (csu *ConnectorColumnConfigModifyService) ConnectorId(value string) *ConnectorColumnConfigModifyService {
    csu.connectorId = &value
    return csu
}

func (csu *ConnectorColumnConfigModifyService) Schema(value string) *ConnectorColumnConfigModifyService {
    csu.schema = &value
    return csu
}

func (csu *ConnectorColumnConfigModifyService) Table(value string) *ConnectorColumnConfigModifyService {
    csu.table = &value
    return csu
}

func (csu *ConnectorColumnConfigModifyService) Column(value string) *ConnectorColumnConfigModifyService {
    csu.column = &value
    return csu
}

func (csu *ConnectorColumnConfigModifyService) Enabled(value bool) *ConnectorColumnConfigModifyService {
    csu.enabled = &value
    return csu
}

func (csu *ConnectorColumnConfigModifyService) Hashed(value bool) *ConnectorColumnConfigModifyService {
    csu.hashed = &value
    return csu
}

func (csu *ConnectorColumnConfigModifyService) Do(ctx context.Context) (ConnectorSchemaDetailsResponse, error) {
    var response ConnectorSchemaDetailsResponse
    if csu.connectorId == nil {
        return response, fmt.Errorf("missing required connectorId")
    }

    if csu.schema == nil {
        return response, fmt.Errorf("missing required schema")
    }

    if csu.table == nil {
        return response, fmt.Errorf("missing required table")
    }

    if csu.column == nil {
        return response, fmt.Errorf("missing required column")
    }
    url := fmt.Sprintf("/connectors/%v/schemas/%v/tables/%v/columns/%v", *csu.connectorId, *csu.schema, *csu.table, *csu.column)
    err := csu.HttpService.Do(ctx, "PATCH", url, csu.request(), nil, 200, &response)
    return response, err
}
