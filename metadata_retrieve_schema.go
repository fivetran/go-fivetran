package fivetran

import (
    "context"
    "encoding/json"
    "fmt"
)

// MetadataSchemaListService implements the Metadata Management, Retrieve schema metadata
// Ref. https://fivetran.com/docs/rest-api/metadata#retrieveschemametadata
type MetadataSchemaListService struct {
    c               *Client
    limit           *int
    cursor          *string
    connectorId     *string
}

type MetadataSchemaListResponse struct {
    Code    string `json:"code"`
    Data    struct {
        Items []struct {
            Id                     string    `json:"id"`
            NameInSource           string    `json:"name_in_source"`
            NameInDestination      string    `json:"name_in_destination"`
        } `json:"items"`
        NextCursor string `json:"next_cursor"`
    } `json:"data"`
}

func (c *Client) NewMetadataSchemaList() *MetadataSchemaListService {
    return &MetadataSchemaListService{c: c}
}

func (s *MetadataSchemaListService) ConnectorId(value string) *MetadataSchemaListService {
    s.connectorId = &value
    return s
}

func (s *MetadataSchemaListService) Limit(value int) *MetadataSchemaListService {
    s.limit = &value
    return s
}

func (s *MetadataSchemaListService) Cursor(value string) *MetadataSchemaListService {
    s.cursor = &value
    return s
}

func (s *MetadataSchemaListService) Do(ctx context.Context) (MetadataSchemaListResponse, error) {
    var response MetadataSchemaListResponse
    url := fmt.Sprintf("%v/metadata/connectors/%v/schemas", s.c.baseURL, *s.connectorId)
    expectedStatus := 200

    if s.connectorId == nil {
        return response, fmt.Errorf("missing required connectorId")
    }

    headers := s.c.commonHeaders()

    queries := make(map[string]string)
    if s.cursor != nil {
        queries["cursor"] = *s.cursor
    }
    if s.limit != nil {
        queries["limit"] = fmt.Sprint(*s.limit)
    }

    r := request{
        method:           "GET",
        url:              url,
        body:             nil,
        queries:          queries,
        headers:          headers,
        client:           s.c.httpClient,
        handleRateLimits: s.c.handleRateLimits,
        maxRetryAttempts: s.c.maxRetryAttempts,
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
