package fivetran

import (
    "context"
    "encoding/json"
    "fmt"
)

// ConnectorTableConfigModifyService implements the Connector Management, Modify a Connector Table Config
// Ref. https://fivetran.com/docs/rest-api/connectors#modifyaconnectortableconfig
type ConnectorTableConfigModifyService struct {
    c                   *Client
    connectorId         *string
    schema              *string
    table               *string
    enabled             *bool
    sync_mode           *string
    columns             map[string]*ConnectorSchemaConfigColumn
}

type connectorTableConfigModifyRequest struct {
    Enabled             *bool                                         `json:"enabled,omitempty"`
    SyncMode            *string                                       `json:"sync_mode,omitempty"`
    Columns             map[string]*connectorSchemaConfigColumnRequest `json:"columns,omitempty"`
}

func (c *Client) NewConnectorTableConfigModifyService() *ConnectorTableConfigModifyService {
    return &ConnectorTableConfigModifyService{c: c}
}

func (csu *ConnectorTableConfigModifyService) request() *connectorTableConfigModifyRequest {
    var columns map[string]*connectorSchemaConfigColumnRequest
    if csu.columns != nil && len(csu.columns) != 0 {
        columns = make(map[string]*connectorSchemaConfigColumnRequest)
        for k, v := range csu.columns {
            columns[k] = v.request()
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

    reqBody, err := json.Marshal(csu.request())
    if err != nil {
        return response, err
    }

    url := fmt.Sprintf("%v/connectors/%v/schemas/%v/tables/%v", csu.c.baseURL, *csu.connectorId, *csu.schema, *csu.table)
    expectedStatus := 200

    headers := csu.c.commonHeaders()
    headers["Content-Type"] = "application/json"
    headers["Accept"] = restAPIv2

    r := request{
        method:           "PATCH",
        url:              url,
        body:             reqBody,
        queries:          nil,
        headers:          headers,
        client:           csu.c.httpClient,
        handleRateLimits: csu.c.handleRateLimits,
        maxRetryAttempts: csu.c.maxRetryAttempts,
    }

    respBody, respStatus, err := r.httpRequest(ctx)
    if err != nil {
        return response, err
    }

    if err := json.Unmarshal(respBody, &response); err != nil {
        return response, err
    }

    if respStatus != expectedStatus {
        err := fmt.Errorf("status code: %v; expected: %v", respStatus, expectedStatus)
        return response, err
    }

    return response, nil
}
