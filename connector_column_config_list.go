package fivetran

import (
    "context"
    "encoding/json"
    "fmt"

    "github.com/fivetran/go-fivetran/connectors"
    httputils "github.com/fivetran/go-fivetran/http_utils"
)

// ConnectorColumnConfigListService implements the Connector Management, Retrieve Source Table Columns Config
// Ref. https://fivetran.com/docs/rest-api/connectors#retrievesourcetablecolumnsconfig
type ConnectorColumnConfigListService struct {
    c                   *Client
    connectorId         *string
    schema              *string
    table               *string
}

type ConnectorColumnConfigListResponse struct {
    Code                string `json:"code"`
    Columns             map[string]*connectors.ConnectorSchemaConfigColumnResponse `json:"columns"`
    
}

func (c *Client) NewConnectorColumnConfigListService() *ConnectorColumnConfigListService {
    return &ConnectorColumnConfigListService{c: c}
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

    if s.connectorId == nil {
        return response, fmt.Errorf("missing required connectorId")
    }

    if s.schema == nil {
        return response, fmt.Errorf("missing required schema")
    }

    if s.table == nil {
        return response, fmt.Errorf("missing required table")
    }

    url := fmt.Sprintf("%v/connectors/%v/schemas/%v/tables/%v/columns", s.c.baseURL, *s.connectorId, *s.schema, *s.table)
    expectedStatus := 200

    headers := s.c.commonHeaders()

    r := httputils.Request{
        Method:           "GET",
        Url:              url,
        Body:             nil,
        Queries:          nil,
        Headers:          headers,
        Client:           s.c.httpClient,
        HandleRateLimits: s.c.handleRateLimits,
        MaxRetryAttempts: s.c.maxRetryAttempts,
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
