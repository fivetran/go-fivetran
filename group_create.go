package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// GroupCreateService implements the Group Management, Create a Group API.
// Ref. https://fivetran.com/docs/rest-api/groups#createagroup
type GroupCreateService struct {
	c    *Client
	name *string
}

type groupCreateRequest struct {
	Name *string `json:"name,omitempty"`
}

type GroupCreateResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ID        string    `json:"id"`
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"created_at"`
	} `json:"data"`
}

func (c *Client) NewGroupCreate() *GroupCreateService {
	return &GroupCreateService{c: c}
}

func (s *GroupCreateService) request() *groupCreateRequest {
	return &groupCreateRequest{
		Name: s.name,
	}
}

func (s *GroupCreateService) Name(value string) *GroupCreateService {
	s.name = &value
	return s
}

func (s *GroupCreateService) Do(ctx context.Context) (GroupCreateResponse, error) {
	var response GroupCreateResponse
	url := fmt.Sprintf("%v/groups", s.c.baseURL)
	expectedStatus := 201

	headers := s.c.commonHeaders()
	headers["Content-Type"] = "application/json"

	reqBody, err := json.Marshal(s.request())
	if err != nil {
		return response, err
	}

	r := request{
		method:  "POST",
		url:     url,
		body:    reqBody,
		queries: nil,
		headers: headers,
		client:  s.c.httpClient,
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
