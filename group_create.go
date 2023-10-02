package fivetran

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fivetran/go-fivetran/groups"
	httputils "github.com/fivetran/go-fivetran/http_utils"
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

func (s *GroupCreateService) Do(ctx context.Context) (groups.GroupDetailsResponse, error) {
	var response groups.GroupDetailsResponse
	url := fmt.Sprintf("%v/groups", s.c.baseURL)
	expectedStatus := 201

	headers := s.c.commonHeaders()
	headers["Content-Type"] = "application/json"

	reqBody, err := json.Marshal(s.request())
	if err != nil {
		return response, err
	}

	r := httputils.Request{
		Method:           "POST",
		Url:              url,
		Body:             reqBody,
		Queries:          nil,
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
