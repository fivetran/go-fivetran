package fivetran

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// ConnectorsSourceMetadataService implements the Connector Management, Retrieve source metadata API.
// Ref. https://fivetran.com/docs/rest-api/connectors#retrievesourcemetadata
type ConnectorsSourceMetadataService struct {
	c      *Client
	limit  *int
	cursor *string
}

type ConnectorsSourceMetadataResponse struct {
	common.CommonResponse
	Data struct {
		Items []struct {
			ID          string `json:"id"`
			Name        string `json:"name"`
			Type        string `json:"type"`
			Description string `json:"description"`
			IconURL     string `json:"icon_url"`
			LinkToDocs  string `json:"link_to_docs"`
			LinkToErd   string `json:"link_to_erd"`
		} `json:"items"`
		NextCursor string `json:"next_cursor"`
	} `json:"data"`
}

func (c *Client) NewConnectorsSourceMetadata() *ConnectorsSourceMetadataService {
	return &ConnectorsSourceMetadataService{c: c}
}

func (s *ConnectorsSourceMetadataService) Limit(value int) *ConnectorsSourceMetadataService {
	s.limit = &value
	return s
}

func (s *ConnectorsSourceMetadataService) Cursor(value string) *ConnectorsSourceMetadataService {
	s.cursor = &value
	return s
}

func (s *ConnectorsSourceMetadataService) Do(ctx context.Context) (ConnectorsSourceMetadataResponse, error) {
	var response ConnectorsSourceMetadataResponse
	url := fmt.Sprintf("%v/metadata/connectors", s.c.baseURL)
	expectedStatus := 200

	headers := s.c.commonHeaders()

	queries := make(map[string]string)
	if s.cursor != nil {
		queries["cursor"] = *s.cursor
	}
	if s.limit != nil {
		queries["limit"] = fmt.Sprint(*s.limit)
	}

	r := httputils.Request{
		Method:           "POST",
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
