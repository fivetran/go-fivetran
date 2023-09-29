package fivetran

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
	"github.com/fivetran/go-fivetran/users"
)

// GroupListUsersService implements the Group Management, List All Users within a Group API.
// Ref. https://fivetran.com/docs/rest-api/groups#listalluserswithinagroup
type GroupListUsersService struct {
	c       *Client
	groupID *string
	limit   *int
	cursor  *string
}

type GroupListUsersResponse struct {
	common.CommonResponse
	Data struct {
		Items      []users.UserDetailsData `json:"items"`
		NextCursor string                  `json:"next_cursor"`
	} `json:"data"`
}

func (c *Client) NewGroupListUsers() *GroupListUsersService {
	return &GroupListUsersService{c: c}
}

func (s *GroupListUsersService) GroupID(value string) *GroupListUsersService {
	s.groupID = &value
	return s
}

func (s *GroupListUsersService) Limit(value int) *GroupListUsersService {
	s.limit = &value
	return s
}

func (s *GroupListUsersService) Cursor(value string) *GroupListUsersService {
	s.cursor = &value
	return s
}

func (s *GroupListUsersService) Do(ctx context.Context) (GroupListUsersResponse, error) {
	var response GroupListUsersResponse

	if s.groupID == nil {
		return response, fmt.Errorf("missing required GroupID")
	}

	url := fmt.Sprintf("%v/groups/%v/users", s.c.baseURL, *s.groupID)
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
