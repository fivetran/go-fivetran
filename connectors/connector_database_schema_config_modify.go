package connectors

import (
    "context"
    "fmt"

    httputils "github.com/fivetran/go-fivetran/http_utils"
)

// ConnectorDatabaseSchemaConfigModifyService implements the Connector Management, Modify a Connector Database Schema Config
// Ref. https://fivetran.com/docs/rest-api/connectors#modifyaconnectordatabaseschemaconfig
type ConnectorDatabaseSchemaConfigModifyService struct {
    httputils.HttpService
    connectorId         *string
    schema              *string
    enabled             *bool
    tables              map[string]*ConnectorSchemaConfigTable
}

func (csu *ConnectorDatabaseSchemaConfigModifyService) request() *connectorSchemaConfigModifyRequest {
    var tables map[string]*ConnectorSchemaConfigTableRequest
    if csu.tables != nil && len(csu.tables) != 0 {
        tables = make(map[string]*ConnectorSchemaConfigTableRequest)
        for k, v := range csu.tables {
            tables[k] = v.Request()
        }
    }

    return &connectorSchemaConfigModifyRequest{
        Enabled:       csu.enabled,
        Tables:       tables,
    }
}

func (csu *ConnectorDatabaseSchemaConfigModifyService) ConnectorId(value string) *ConnectorDatabaseSchemaConfigModifyService {
    csu.connectorId = &value
    return csu
}

func (csu *ConnectorDatabaseSchemaConfigModifyService) Schema(value string) *ConnectorDatabaseSchemaConfigModifyService {
    csu.schema = &value
    return csu
}

func (csu *ConnectorDatabaseSchemaConfigModifyService) Enabled(value bool) *ConnectorDatabaseSchemaConfigModifyService {
    csu.enabled = &value
    return csu
}

func (csu *ConnectorDatabaseSchemaConfigModifyService) Tables(name string, table *ConnectorSchemaConfigTable) *ConnectorDatabaseSchemaConfigModifyService {
    if csu.tables == nil {
        csu.tables = make(map[string]*ConnectorSchemaConfigTable)
    }
    csu.tables[name] = table
    return csu
}

func (csu *ConnectorDatabaseSchemaConfigModifyService) Do(ctx context.Context) (ConnectorSchemaDetailsResponse, error) {
    var response ConnectorSchemaDetailsResponse

    if csu.connectorId == nil {
        return response, fmt.Errorf("missing required connectorId")
    }

    if csu.schema == nil {
        return response, fmt.Errorf("missing required schema")
    }

    url := fmt.Sprintf("/connectors/%v/schemas/%v", *csu.connectorId, *csu.schema)
    err := csu.HttpService.Do(ctx, "PATCH", url, csu.request(), nil, 200, &response)
    return response, err
}