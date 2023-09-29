package fivetran

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// TeamGroupMembershipDeleteService implements the Team Management, Delete group membership
// Ref. https://fivetran.com/docs/rest-api/teams#deletegroupmembership
type TeamGroupMembershipDeleteService struct {
	c       *Client
	teamId  *string
	groupId *string
}

func (c *Client) NewTeamGroupMembershipDelete() *TeamGroupMembershipDeleteService {
	return &TeamGroupMembershipDeleteService{c: c}
}

func (s *TeamGroupMembershipDeleteService) TeamId(value string) *TeamGroupMembershipDeleteService {
	s.teamId = &value
	return s
}

func (s *TeamGroupMembershipDeleteService) GroupId(value string) *TeamGroupMembershipDeleteService {
	s.groupId = &value
	return s
}

func (s *TeamGroupMembershipDeleteService) Do(ctx context.Context) (common.CommonResponse, error) {
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
