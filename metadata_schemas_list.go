package fivetran

import (
	"context"
	"encoding/json"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// MetadataSchemasListService implements the Metadata Management, Retrieve schema metadata
// Ref. https://fivetran.com/docs/rest-api/metadata#retrieveschemametadata
type MetadataSchemasListService struct {
	c           *Client
	limit       *int
	cursor      *string
	connectorId *string
}

type MetadataSchemasListResponse struct {
	Code string `json:"code"`
	Data struct {
		Items []struct {
			Id                string `json:"id"`
			NameInSource      string `json:"name_in_source"`
			NameInDestination string `json:"name_in_destination"`
		} `json:"items"`
		NextCursor string `json:"next_cursor"`
	} `json:"data"`
}

func (c *Client) NewMetadataSchemasList() *MetadataSchemasListService {
	return &MetadataSchemasListService{c: c}
}

func (s *MetadataSchemasListService) ConnectorId(value string) *MetadataSchemasListService {
	s.connectorId = &value
	return s
}

func (s *MetadataSchemasListService) Limit(value int) *MetadataSchemasListService {
	s.limit = &value
	return s
}

func (s *MetadataSchemasListService) Cursor(value string) *MetadataSchemasListService {
	s.cursor = &value
	return s
}

func (s *MetadataSchemasListService) Do(ctx context.Context) (MetadataSchemasListResponse, error) {
	var response MetadataSchemasListResponse
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
