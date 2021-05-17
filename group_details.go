package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type GroupDetailsService struct {
	c  *Client
	id string
}

type GroupDetails struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ID        string    `json:"id"`
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"created_at"`
	} `json:"data"`
}

func (c *Client) NewGroupDetailsService() *GroupDetailsService {
	return &GroupDetailsService{c: c}
}

func (s *GroupDetailsService) ID(id string) *GroupDetailsService {
	s.id = id
	return s
}

func (s *GroupDetailsService) Do(ctx context.Context) (GroupDetails, error) {
	if s.id == "" {
		err := fmt.Errorf("missing required ID")
		return GroupDetails{}, err
	}

	url := fmt.Sprintf("%v/groups/%v", s.c.baseURL, s.id)
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
		return GroupDetails{}, err
	}

	var groupDetails GroupDetails
	if err := json.Unmarshal(respBody, &groupDetails); err != nil {
		return GroupDetails{}, err
	}

	if respStatus != expectedStatus {
		err := fmt.Errorf("status code: %v; expected %v", respStatus, expectedStatus)
		return groupDetails, err
	}

	return groupDetails, nil
}
