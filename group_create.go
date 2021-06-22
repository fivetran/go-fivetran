package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

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

func (s *GroupCreateService) request() groupCreateRequest {
	return groupCreateRequest{
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

	headers := make(map[string]string)
	headers["Authorization"] = s.c.authorization
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
	}

	respBody, respStatus, err := httpRequest(r, ctx)
	if err != nil {
		return response, err
	}

	if err := json.Unmarshal(respBody, &response); err != nil {
		return response, err
	}

	if respStatus != expectedStatus {
		err := fmt.Errorf("status code: %v; expected %v", respStatus, expectedStatus)
		return response, err
	}

	return response, nil
}
