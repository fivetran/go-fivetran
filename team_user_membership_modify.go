package fivetran

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// TeamUsersModifyService implements the Team Management, Modify a user membership
// Ref. https://fivetran.com/docs/rest-api/teams#modifyausermembership
type TeamUserMembershipModifyService struct {
	c      *Client
	teamId *string
	userId *string
	role   *string
}

type teamUserMembershipModifyRequest struct {
	Role *string `json:"role,omitempty"`
}

func (c *Client) NewTeamUserMembershipModify() *TeamUserMembershipModifyService {
	return &TeamUserMembershipModifyService{c: c}
}

func (s *TeamUserMembershipModifyService) request() *teamUserMembershipModifyRequest {
	return &teamUserMembershipModifyRequest{
		Role: s.role,
	}
}

func (s *TeamUserMembershipModifyService) TeamId(value string) *TeamUserMembershipModifyService {
	s.teamId = &value
	return s
}

func (s *TeamUserMembershipModifyService) UserId(value string) *TeamUserMembershipModifyService {
	s.userId = &value
	return s
}

func (s *TeamUserMembershipModifyService) Role(value string) *TeamUserMembershipModifyService {
	s.role = &value
	return s
}

func (s *TeamUserMembershipModifyService) Do(ctx context.Context) (common.CommonResponse, error) {
	var response common.CommonResponse

	if s.teamId == nil {
		return response, fmt.Errorf("missing required teamId")
	}

	if s.userId == nil {
		return response, fmt.Errorf("missing required userId")
	}

	url := fmt.Sprintf("%v/teams/%v/users/%v", s.c.baseURL, *s.teamId, *s.userId)
	expectedStatus := 200

	headers := s.c.commonHeaders()
	headers["Content-Type"] = "application/json"
	headers["Accept"] = restAPIv2

	reqBody, err := json.Marshal(s.request())
	if err != nil {
		return response, err
	}

	r := httputils.Request{
		Method:           "PATCH",
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
