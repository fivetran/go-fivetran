package fivetran

import (
    "context"
    "encoding/json"
    "fmt"

    "github.com/fivetran/go-fivetran/connectors"
    httputils "github.com/fivetran/go-fivetran/http_utils"
)

// ConnectorColumnConfigModifyService implements the Connector Management, Modify a Connector Table Config
// Ref. https://fivetran.com/docs/rest-api/connectors#modifyaconnectorcolumnconfig
type ConnectorColumnConfigModifyService struct {
    c                   *Client
    connectorId         *string
    schema              *string
    table               *string
    column              *string
    enabled             *bool
    hashed              *bool
}

type connectorColumnConfigModifyRequest struct {
    Enabled             *bool `json:"enabled,omitempty"`
    Hashed              *bool `json:"hashed,omitempty"`
}

func (c *Client) NewConnectorColumnConfigModifyService() *ConnectorColumnConfigModifyService {
    return &ConnectorColumnConfigModifyService{c: c}
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

func (csu *ConnectorColumnConfigModifyService) Do(ctx context.Context) (connectors.ConnectorSchemaDetailsResponse, error) {
    var response connectors.ConnectorSchemaDetailsResponse

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


    reqBody, err := json.Marshal(csu.request())
    if err != nil {
        return response, err
    }

    url := fmt.Sprintf("%v/connectors/%v/schemas/%v/tables/%v/columns/%v", csu.c.baseURL, *csu.connectorId, *csu.schema, *csu.table, *csu.column)
    expectedStatus := 200

    headers := csu.c.commonHeaders()
    headers["Content-Type"] = "application/json"
    headers["Accept"] = restAPIv2

    r := httputils.Request{
        Method:           "PATCH",
        Url:              url,
        Body:             reqBody,
        Queries:          nil,
        Headers:          headers,
        Client:           csu.c.httpClient,
        HandleRateLimits: csu.c.handleRateLimits,
        MaxRetryAttempts: csu.c.maxRetryAttempts,
    }

    respBody, respStatus, err := r.Do(ctx)
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
