package fivetran

import (
	"context"
	"encoding/json"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
	"github.com/fivetran/go-fivetran/teams"
)

// TeamListGroupsService implements the Team Management, List all group memberships
// Ref. https://fivetran.com/docs/rest-api/teams#listallgroupmemberships
type TeamGroupMembershipsListService struct {
	c      *Client
	teamId *string
	limit  *int
	cursor *string
}

type TeamGroupMembershipsListResponse struct {
	Code string `json:"code"`
	Data struct {
		Items      []teams.TeamGroupMembership `json:"items"`
		NextCursor string                      `json:"next_cursor"`
	} `json:"data"`
}

func (c *Client) NewTeamGroupMembershipsList() *TeamGroupMembershipsListService {
	return &TeamGroupMembershipsListService{c: c}
}

func (s *TeamGroupMembershipsListService) TeamId(value string) *TeamGroupMembershipsListService {
	s.teamId = &value
	return s
}

func (s *TeamGroupMembershipsListService) Limit(value int) *TeamGroupMembershipsListService {
	s.limit = &value
	return s
}

func (s *TeamGroupMembershipsListService) Cursor(value string) *TeamGroupMembershipsListService {
	s.cursor = &value
	return s
}

func (s *TeamGroupMembershipsListService) Do(ctx context.Context) (TeamGroupMembershipsListResponse, error) {
	var response TeamGroupMembershipsListResponse

	if s.teamId == nil {
		return response, fmt.Errorf("missing required teamId")
	}

	url := fmt.Sprintf("%v/teams/%v/groups", s.c.baseURL, *s.teamId)
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
