package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type groupDetailsService struct {
	c       *Client
	groupID *string
}

type GroupDetailsResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ID        string    `json:"id"`
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"created_at"`
	} `json:"data"`
}

func (c *Client) NewGroupDetails() *groupDetailsService {
	return &groupDetailsService{c: c}
}

func (s *groupDetailsService) GroupID(value string) *groupDetailsService {
	s.groupID = &value
	return s
}

func (s *groupDetailsService) Do(ctx context.Context) (GroupDetailsResponse, error) {
	var response GroupDetailsResponse

	if s.groupID == nil {
		return response, fmt.Errorf("missing required GroupID")
	}

	url := fmt.Sprintf("%v/groups/%v", s.c.baseURL, *s.groupID)
	expectedStatus := 200

	headers := make(map[string]string)
	headers["Authorization"] = s.c.authorization

	r := Request{
		method:  "GET",
		url:     url,
		body:    nil,
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
