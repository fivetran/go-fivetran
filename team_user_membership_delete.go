package fivetran

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// TeamUserMembershipDeleteService implements the Team Management, Delete a user from a team
// Ref. https://fivetran.com/docs/rest-api/teams#deleteauserfromateam
type TeamUserMembershipDeleteService struct {
	c      *Client
	teamId *string
	userId *string
}

func (c *Client) NewTeamUserMembershipDelete() *TeamUserMembershipDeleteService {
	return &TeamUserMembershipDeleteService{c: c}
}

func (s *TeamUserMembershipDeleteService) TeamId(value string) *TeamUserMembershipDeleteService {
	s.teamId = &value
	return s
}

func (s *TeamUserMembershipDeleteService) UserId(value string) *TeamUserMembershipDeleteService {
	s.userId = &value
	return s
}

func (s *TeamUserMembershipDeleteService) Do(ctx context.Context) (common.CommonResponse, error) {
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
