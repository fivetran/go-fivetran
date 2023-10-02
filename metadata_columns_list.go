package fivetran

import (
	"context"
	"encoding/json"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// MetadataColumnsListService implements the Metadata Management, Retrieve Column metadata
// Ref. https://fivetran.com/docs/rest-api/metadata#retrievecolumnmetadata
type MetadataColumnsListService struct {
	c           *Client
	limit       *int
	cursor      *string
	connectorId *string
}

type MetadataColumnsListResponse struct {
	Code string `json:"code"`
	Data struct {
		Items []struct {
			Id                string `json:"id"`
			ParentId          string `json:"parent_id"`
			NameInSource      string `json:"name_in_source"`
			NameInDestination string `json:"name_in_destination"`
			TypeInSource      string `json:"type_in_source"`
			TypeInDestination string `json:"type_in_destination"`
			IsPrimaryKey      *bool  `json:"is_primary_key"`
			IsForeignKey      *bool  `json:"is_foreign_key"`
		} `json:"items"`
		NextCursor string `json:"next_cursor"`
	} `json:"data"`
}

func (c *Client) NewMetadataColumnsList() *MetadataColumnsListService {
	return &MetadataColumnsListService{c: c}
}

func (s *MetadataColumnsListService) ConnectorId(value string) *MetadataColumnsListService {
	s.connectorId = &value
	return s
}

func (s *MetadataColumnsListService) Limit(value int) *MetadataColumnsListService {
	s.limit = &value
	return s
}

func (s *MetadataColumnsListService) Cursor(value string) *MetadataColumnsListService {
	s.cursor = &value
	return s
}

func (s *MetadataColumnsListService) Do(ctx context.Context) (MetadataColumnsListResponse, error) {
	var response MetadataColumnsListResponse
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
