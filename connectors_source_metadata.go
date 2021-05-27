package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
)

type ConnectorsSourceMetadataService struct {
	c      *Client
	limit  *int
	cursor *string
}

type ConnectorsSourceMetadata struct {
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

func (c *Client) NewConnectorsSourceMetadataService() *ConnectorsSourceMetadataService {
	return &ConnectorsSourceMetadataService{c: c}
}

func (s *ConnectorsSourceMetadataService) Limit(limit int) *ConnectorsSourceMetadataService {
	s.limit = &limit
	return s
}

func (s *ConnectorsSourceMetadataService) Cursor(cursor string) *ConnectorsSourceMetadataService {
	s.cursor = &cursor
	return s
}

func (s *ConnectorsSourceMetadataService) Do(ctx context.Context) (ConnectorsSourceMetadata, error) {
	url := fmt.Sprintf("%v/metadata/connectors", s.c.baseURL)
	expectedStatus := 200
	headers := make(map[string]string)
	queries := make(map[string]string)

	headers["Authorization"] = s.c.authorization

	if s.cursor != nil {
		queries["cursor"] = *s.cursor
	}

	if s.limit != nil {
		queries["limit"] = fmt.Sprint(*s.limit)
	}

	r := Request{
		method:  "GET",
		url:     url,
		body:    nil,
		queries: queries,
		headers: headers,
	}

	respBody, respStatus, err := httpRequest(r, ctx)
	if err != nil {
		return ConnectorsSourceMetadata{}, err
	}

	var connectorsSourceMetadata ConnectorsSourceMetadata
	if err := json.Unmarshal(respBody, &connectorsSourceMetadata); err != nil {
		return ConnectorsSourceMetadata{}, err
	}

	if respStatus != expectedStatus {
		err := fmt.Errorf("status code: %v; expected %v", respStatus, expectedStatus)
		return connectorsSourceMetadata, err
	}

	return connectorsSourceMetadata, nil
}
