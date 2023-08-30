package fivetran

import (
    "context"
    "encoding/json"
    "fmt"
)

// MetadataColumnListService implements the Metadata Management, Retrieve Column metadata
// Ref. https://fivetran.com/docs/rest-api/metadata#retrievecolumnmetadata
type MetadataColumnListService struct {
    c               *Client
    limit           *int
    cursor          *string
    connectorId     *string
}

type MetadataColumnListResponse struct {
    Code    string `json:"code"`
    Data    struct {
        Items []struct {
            Id                      string    `json:"id"`
            ParentId                string    `json:"parent_id"`
            NameInSource            string    `json:"name_in_source"`
            NameInDestination       string    `json:"name_in_destination"`
            TypeInSource            string    `json:"type_in_source"`
            TypeInDestination       string    `json:"type_in_destination"`
            IsPrimaryKey            *bool     `json:"is_primary_key"`
            IsForeignKey            *bool     `json:"is_foreign_key"`
        } `json:"items"`
        NextCursor string `json:"next_cursor"`
    } `json:"data"`
}

func (c *Client) NewMetadataColumnList() *MetadataColumnListService {
    return &MetadataColumnListService{c: c}
}

func (s *MetadataColumnListService) ConnectorId(value string) *MetadataColumnListService {
    s.connectorId = &value
    return s
}

func (s *MetadataColumnListService) Limit(value int) *MetadataColumnListService {
    s.limit = &value
    return s
}

func (s *MetadataColumnListService) Cursor(value string) *MetadataColumnListService {
    s.cursor = &value
    return s
}

func (s *MetadataColumnListService) Do(ctx context.Context) (MetadataColumnListResponse, error) {
    var response MetadataColumnListResponse
    url := fmt.Sprintf("%v/metadata/connectors/%v/columns", s.c.baseURL, *s.connectorId)
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
