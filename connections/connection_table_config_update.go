package connections

import (
    "context"
    "fmt"

    httputils "github.com/fivetran/go-fivetran/http_utils"
)

type ConnectionTableConfigUpdateService struct {
    httputils.HttpService
    connectionId         *string
    schema              *string
    table               *string
    enabled             *bool
    sync_mode           *string
    columns             map[string]*ConnectionSchemaConfigColumn
}

func (csu *ConnectionTableConfigUpdateService) request() *connectionTableConfigUpdateRequest {
    var columns map[string]*ConnectionSchemaConfigColumnRequest
    if csu.columns != nil && len(csu.columns) != 0 {
        columns = make(map[string]*ConnectionSchemaConfigColumnRequest)
        for k, v := range csu.columns {
            columns[k] = v.Request()
        }
    }

    return &connectionTableConfigUpdateRequest{
        Enabled:       csu.enabled,
        SyncMode:      csu.sync_mode,
        Columns:       columns,
    }
}

func (csu *ConnectionTableConfigUpdateService) ConnectionId(value string) *ConnectionTableConfigUpdateService {
    csu.connectionId = &value
    return csu
}

func (csu *ConnectionTableConfigUpdateService) Schema(value string) *ConnectionTableConfigUpdateService {
    csu.schema = &value
    return csu
}

func (csu *ConnectionTableConfigUpdateService) Table(value string) *ConnectionTableConfigUpdateService {
    csu.table = &value
    return csu
}

func (csu *ConnectionTableConfigUpdateService) Enabled(value bool) *ConnectionTableConfigUpdateService {
    csu.enabled = &value
    return csu
}

func (csu *ConnectionTableConfigUpdateService) SyncMode(value string) *ConnectionTableConfigUpdateService {
    csu.sync_mode = &value
    return csu
}

func (csu *ConnectionTableConfigUpdateService) Columns(name string, table *ConnectionSchemaConfigColumn) *ConnectionTableConfigUpdateService {
    if csu.columns == nil {
        csu.columns = make(map[string]*ConnectionSchemaConfigColumn)
    }
    csu.columns[name] = table
    return csu
}

func (csu *ConnectionTableConfigUpdateService) Do(ctx context.Context) (ConnectionSchemaDetailsResponse, error) {
    var response ConnectionSchemaDetailsResponse
    if csu.connectionId == nil {
        return response, fmt.Errorf("missing required connectionId")
    }

    if csu.schema == nil {
        return response, fmt.Errorf("missing required schema")
    }

    if csu.table == nil {
        return response, fmt.Errorf("missing required table")
    }
    url := fmt.Sprintf("/connections/%v/schemas/%v/tables/%v", *csu.connectionId, *csu.schema, *csu.table)
    err := csu.HttpService.Do(ctx, "PATCH", url, csu.request(), nil, 200, &response)
    return response, err
}