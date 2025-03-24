package connectors

import (
    "context"
    "fmt"

    httputils "github.com/fivetran/go-fivetran/http_utils"
)

// ConnectorTableConfigModifyService implements the Connector Management, Modify a Connector Table Config
// Ref. https://fivetran.com/docs/rest-api/connectors#modifyaconnectortableconfig
type ConnectorTableConfigModifyService struct {
    httputils.HttpService
    connectorId         *string
    schema              *string
    table               *string
    enabled             *bool
    sync_mode           *string
    columns             map[string]*ConnectorSchemaConfigColumn
}

func (csu *ConnectorTableConfigModifyService) request() *connectorTableConfigModifyRequest {
    var columns map[string]*ConnectorSchemaConfigColumnRequest
    if csu.columns != nil && len(csu.columns) != 0 {
        columns = make(map[string]*ConnectorSchemaConfigColumnRequest)
        for k, v := range csu.columns {
            columns[k] = v.Request()
        }
    }

    return &connectorTableConfigModifyRequest{
        Enabled:       csu.enabled,
        SyncMode:      csu.sync_mode,
        Columns:       columns,
    }
}

func (csu *ConnectorTableConfigModifyService) ConnectorId(value string) *ConnectorTableConfigModifyService {
    csu.connectorId = &value
    return csu
}

func (csu *ConnectorTableConfigModifyService) Schema(value string) *ConnectorTableConfigModifyService {
    csu.schema = &value
    return csu
}

func (csu *ConnectorTableConfigModifyService) Table(value string) *ConnectorTableConfigModifyService {
    csu.table = &value
    return csu
}

func (csu *ConnectorTableConfigModifyService) Enabled(value bool) *ConnectorTableConfigModifyService {
    csu.enabled = &value
    return csu
}

func (csu *ConnectorTableConfigModifyService) SyncMode(value string) *ConnectorTableConfigModifyService {
    csu.sync_mode = &value
    return csu
}

func (csu *ConnectorTableConfigModifyService) Columns(name string, table *ConnectorSchemaConfigColumn) *ConnectorTableConfigModifyService {
    if csu.columns == nil {
        csu.columns = make(map[string]*ConnectorSchemaConfigColumn)
    }
    csu.columns[name] = table
    return csu
}

func (csu *ConnectorTableConfigModifyService) Do(ctx context.Context) (ConnectorSchemaDetailsResponse, error) {
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
    url := fmt.Sprintf("/connectors/%v/schemas/%v/tables/%v", *csu.connectorId, *csu.schema, *csu.table)
    err := csu.HttpService.Do(ctx, "PATCH", url, csu.request(), nil, 200, &response)
    return response, err
}