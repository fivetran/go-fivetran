package connections

import (
    "context"
    "fmt"

    httputils "github.com/fivetran/go-fivetran/http_utils"
)

type ConnectionColumnConfigUpdateService struct {
    httputils.HttpService
    connectionId         *string
    schema              *string
    table               *string
    column              *string
    enabled             *bool
    hashed              *bool
}

func (csu *ConnectionColumnConfigUpdateService) request() *connectionColumnConfigUpdateRequest {
    return &connectionColumnConfigUpdateRequest{
        Enabled:       csu.enabled,
        Hashed:        csu.hashed,
    }
}

func (csu *ConnectionColumnConfigUpdateService) ConnectionId(value string) *ConnectionColumnConfigUpdateService {
    csu.connectionId = &value
    return csu
}

func (csu *ConnectionColumnConfigUpdateService) Schema(value string) *ConnectionColumnConfigUpdateService {
    csu.schema = &value
    return csu
}

func (csu *ConnectionColumnConfigUpdateService) Table(value string) *ConnectionColumnConfigUpdateService {
    csu.table = &value
    return csu
}

func (csu *ConnectionColumnConfigUpdateService) Column(value string) *ConnectionColumnConfigUpdateService {
    csu.column = &value
    return csu
}

func (csu *ConnectionColumnConfigUpdateService) Enabled(value bool) *ConnectionColumnConfigUpdateService {
    csu.enabled = &value
    return csu
}

func (csu *ConnectionColumnConfigUpdateService) Hashed(value bool) *ConnectionColumnConfigUpdateService {
    csu.hashed = &value
    return csu
}

func (csu *ConnectionColumnConfigUpdateService) Do(ctx context.Context) (ConnectionSchemaDetailsResponse, error) {
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

    if csu.column == nil {
        return response, fmt.Errorf("missing required column")
    }
    url := fmt.Sprintf("/connections/%v/schemas/%v/tables/%v/columns/%v", *csu.connectionId, *csu.schema, *csu.table, *csu.column)
    err := csu.HttpService.Do(ctx, "PATCH", url, csu.request(), nil, 200, &response)
    return response, err
}
