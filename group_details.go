package fivetran

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fivetran/go-fivetran/groups"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// GroupDetailsService implements the Group Management, Retrieve Group Details API.
// Ref. https://fivetran.com/docs/rest-api/groups#retrievegroupdetails
type GroupDetailsService struct {
	c       *Client
	groupID *string
}

func (c *Client) NewGroupDetails() *GroupDetailsService {
	return &GroupDetailsService{c: c}
}

func (s *GroupDetailsService) GroupID(value string) *GroupDetailsService {
	s.groupID = &value
	return s
}

func (s *GroupDetailsService) Do(ctx context.Context) (groups.GroupDetailsResponse, error) {
	var response groups.GroupDetailsResponse

	if s.groupID == nil {
		return response, fmt.Errorf("missing required GroupID")
	}

	url := fmt.Sprintf("%v/groups/%v", s.c.baseURL, *s.groupID)
	expectedStatus := 200

	headers := s.c.commonHeaders()

	r := httputils.Request{
		Method:           "GET",
		Url:              url,
		Body:             nil,
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
