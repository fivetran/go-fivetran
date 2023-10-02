package fivetran

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// GroupRemoveUserService implements the Group Management, Remove a User from a Group API.
// Ref. https://fivetran.com/docs/rest-api/groups#removeauserfromagroup
type GroupRemoveUserService struct {
	c       *Client
	groupID *string
	userID  *string
}

func (c *Client) NewGroupRemoveUser() *GroupRemoveUserService {
	return &GroupRemoveUserService{c: c}
}

func (s *GroupRemoveUserService) GroupID(value string) *GroupRemoveUserService {
	s.groupID = &value
	return s
}

func (s *GroupRemoveUserService) UserID(value string) *GroupRemoveUserService {
	s.userID = &value
	return s
}

func (s *GroupRemoveUserService) Do(ctx context.Context) (common.CommonResponse, error) {
	var response common.CommonResponse

	if s.groupID == nil {
		return response, fmt.Errorf("missing required GroupID")
	}
	if s.userID == nil {
		return response, fmt.Errorf("missing required UserID")
	}

	url := fmt.Sprintf("%v/groups/%v/users/%v", s.c.baseURL, *s.groupID, *s.userID)
	expectedStatus := 200

	headers := s.c.commonHeaders()

	r := httputils.Request{
		Method:           "DELETE",
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
