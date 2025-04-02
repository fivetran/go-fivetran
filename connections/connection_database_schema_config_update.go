package connections

import (
    "context"
    "fmt"

    httputils "github.com/fivetran/go-fivetran/http_utils"
)

type ConnectionDatabaseSchemaConfigUpdateService struct {
    httputils.HttpService
    connectionId         *string
    schema              *string
    enabled             *bool
    tables              map[string]*ConnectionSchemaConfigTable
}

func (csu *ConnectionDatabaseSchemaConfigUpdateService) request() *connectionSchemaConfigTableUpdateRequest {
    var tables map[string]*ConnectionSchemaConfigTableRequest
    if csu.tables != nil && len(csu.tables) != 0 {
        tables = make(map[string]*ConnectionSchemaConfigTableRequest)
        for k, v := range csu.tables {
            tables[k] = v.Request()
        }
    }

    return &connectionSchemaConfigTableUpdateRequest{
        Enabled:       csu.enabled,
        Tables:       tables,
    }
}

func (csu *ConnectionDatabaseSchemaConfigUpdateService) ConnectionId(value string) *ConnectionDatabaseSchemaConfigUpdateService {
    csu.connectionId = &value
    return csu
}

func (csu *ConnectionDatabaseSchemaConfigUpdateService) Schema(value string) *ConnectionDatabaseSchemaConfigUpdateService {
    csu.schema = &value
    return csu
}

func (csu *ConnectionDatabaseSchemaConfigUpdateService) Enabled(value bool) *ConnectionDatabaseSchemaConfigUpdateService {
    csu.enabled = &value
    return csu
}

func (csu *ConnectionDatabaseSchemaConfigUpdateService) Tables(name string, table *ConnectionSchemaConfigTable) *ConnectionDatabaseSchemaConfigUpdateService {
    if csu.tables == nil {
        csu.tables = make(map[string]*ConnectionSchemaConfigTable)
    }
    csu.tables[name] = table
    return csu
}

func (csu *ConnectionDatabaseSchemaConfigUpdateService) Do(ctx context.Context) (ConnectionSchemaDetailsResponse, error) {
    var response ConnectionSchemaDetailsResponse

    if csu.connectionId == nil {
        return response, fmt.Errorf("missing required connectionId")
    }

    if csu.schema == nil {
        return response, fmt.Errorf("missing required schema")
    }

    url := fmt.Sprintf("/connections/%v/schemas/%v", *csu.connectionId, *csu.schema)
    err := csu.HttpService.Do(ctx, "PATCH", url, csu.request(), nil, 200, &response)
    return response, err
}