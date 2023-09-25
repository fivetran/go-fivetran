package fivetran

import (
    "context"
    "encoding/json"
    "fmt"
)

// ConnectorDatabaseSchemaConfigModifyService implements the Connector Management, Modify a Connector Database Schema Config
// Ref. https://fivetran.com/docs/rest-api/connectors#modifyaconnectordatabaseschemaconfig
type ConnectorDatabaseSchemaConfigModifyService struct {
    c                   *Client
    connectorId         *string
    schema              *string
    enabled             *bool
    tables              map[string]*ConnectorSchemaConfigTable
}

type connectorSchemaConfigModifyRequest struct {
    Enabled             *bool                                         `json:"enabled,omitempty"`
    Tables              map[string]*connectorSchemaConfigTableRequest `json:"tables,omitempty"`
}

func (c *Client) NewConnectorDatabaseSchemaConfigModifyService() *ConnectorDatabaseSchemaConfigModifyService {
    return &ConnectorDatabaseSchemaConfigModifyService{c: c}
}

func (csu *ConnectorDatabaseSchemaConfigModifyService) request() *connectorSchemaConfigModifyRequest {
    var tables map[string]*connectorSchemaConfigTableRequest
    if csu.tables != nil && len(csu.tables) != 0 {
        tables = make(map[string]*connectorSchemaConfigTableRequest)
        for k, v := range csu.tables {
            tables[k] = v.request()
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

    reqBody, err := json.Marshal(csu.request())
    if err != nil {
        return response, err
    }

    url := fmt.Sprintf("%v/connectors/%v/schemas/%v", csu.c.baseURL, *csu.connectorId, *csu.schema)
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
