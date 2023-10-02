package fivetran

import (
	"context"
	"encoding/json"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
	"github.com/fivetran/go-fivetran/teams"
)

// TeamUserMembershipDetailsService implements the Team Management, retrieve Team Details.
// Ref. https://fivetran.com/docs/rest-api/teams#retrieveusermembershipinateam
type TeamUserMembershipDetailsService struct {
	c      *Client
	teamId *string
	userId *string
}

type TeamUserMembershipDetailsResponse struct {
	Code string                   `json:"code"`
	Data teams.TeamUserMembership `json:"data"`
}

func (c *Client) NewTeamUserMembershipDetails() *TeamUserMembershipDetailsService {
	return &TeamUserMembershipDetailsService{c: c}
}

func (s *TeamUserMembershipDetailsService) TeamId(value string) *TeamUserMembershipDetailsService {
	s.teamId = &value
	return s
}

func (s *TeamUserMembershipDetailsService) UserId(value string) *TeamUserMembershipDetailsService {
	s.userId = &value
	return s
}

func (s *TeamUserMembershipDetailsService) Do(ctx context.Context) (TeamUserMembershipDetailsResponse, error) {
	var response TeamUserMembershipDetailsResponse

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
