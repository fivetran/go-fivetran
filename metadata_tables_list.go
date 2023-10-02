package fivetran

import (
	"context"
	"encoding/json"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// MetadataTableListService implements the Metadata Management, Retrieve table metadata
// Ref. https://fivetran.com/docs/rest-api/metadata#retrievetablemetadata
type MetadataTablesListService struct {
	c           *Client
	limit       *int
	cursor      *string
	connectorId *string
}

type MetadataTablesListResponse struct {
	Code string `json:"code"`
	Data struct {
		Items []struct {
			Id                string `json:"id"`
			ParentId          string `json:"parent_id"`
			NameInSource      string `json:"name_in_source"`
			NameInDestination string `json:"name_in_destination"`
		} `json:"items"`
		NextCursor string `json:"next_cursor"`
	} `json:"data"`
}

func (c *Client) NewMetadataTablesList() *MetadataTablesListService {
	return &MetadataTablesListService{c: c}
}

func (s *MetadataTablesListService) ConnectorId(value string) *MetadataTablesListService {
	s.connectorId = &value
	return s
}

func (s *MetadataTablesListService) Limit(value int) *MetadataTablesListService {
	s.limit = &value
	return s
}

func (s *MetadataTablesListService) Cursor(value string) *MetadataTablesListService {
	s.cursor = &value
	return s
}

func (s *MetadataTablesListService) Do(ctx context.Context) (MetadataTablesListResponse, error) {
	var response MetadataTablesListResponse
	url := fmt.Sprintf("%v/metadata/connectors/%v/tables", s.c.baseURL, *s.connectorId)
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

	r := httputils.Request{
		Method:           "GET",
		Url:              url,
		Body:             nil,
		Queries:          queries,
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
