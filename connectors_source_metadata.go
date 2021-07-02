package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
)

// ConnectorsSourceMetadataService implements the Connector Management, Retrieve source metadata API.
// Ref. https://fivetran.com/docs/rest-api/connectors#retrievesourcemetadata
type ConnectorsSourceMetadataService struct {
	c      *Client
	limit  *int
	cursor *string
}

type ConnectorsSourceMetadataResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Items []struct {
			ID          string `json:"id"`
			Name        string `json:"name"`
			Type        string `json:"type"`
			Description string `json:"description"`
			IconURL     string `json:"icon_url"`
			LinkToDocs  string `json:"link_to_docs"`
			LinkToErd   string `json:"link_to_erd,omitempty"`
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

	headers := make(map[string]string)
	headers["Authorization"] = s.c.authorization

	queries := make(map[string]string)
	if s.cursor != nil {
		queries["cursor"] = *s.cursor
	}
	if s.limit != nil {
		queries["limit"] = fmt.Sprint(*s.limit)
	}

	r := request{
		method:  "GET",
		url:     url,
		body:    nil,
		queries: queries,
		headers: headers,
	}

	respBody, respStatus, err := httpRequest(r, ctx)
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
