package fivetran

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// TeamGroupMembershipModifyService implements the Team Management, Update group membership
// Ref. https://fivetran.com/docs/rest-api/teams#updategroupmembership
type TeamGroupMembershipModifyService struct {
	c       *Client
	teamId  *string
	groupId *string
	role    *string
}

type teamGroupMembershipModifyRequest struct {
	Role *string `json:"role,omitempty"`
}

func (c *Client) NewTeamGroupMembershipModify() *TeamGroupMembershipModifyService {
	return &TeamGroupMembershipModifyService{c: c}
}

func (s *TeamGroupMembershipModifyService) request() *teamGroupMembershipModifyRequest {
	return &teamGroupMembershipModifyRequest{
		Role: s.role,
	}
}

func (s *TeamGroupMembershipModifyService) TeamId(value string) *TeamGroupMembershipModifyService {
	s.teamId = &value
	return s
}

func (s *TeamGroupMembershipModifyService) GroupId(value string) *TeamGroupMembershipModifyService {
	s.groupId = &value
	return s
}

func (s *TeamGroupMembershipModifyService) Role(value string) *TeamGroupMembershipModifyService {
	s.role = &value
	return s
}

func (s *TeamGroupMembershipModifyService) Do(ctx context.Context) (common.CommonResponse, error) {
	var response common.CommonResponse

	if s.teamId == nil {
		return response, fmt.Errorf("missing required teamId")
	}

	if s.groupId == nil {
		return response, fmt.Errorf("missing required groupId")
	}

	url := fmt.Sprintf("%v/teams/%v/groups/%v", s.c.baseURL, *s.teamId, *s.groupId)
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
