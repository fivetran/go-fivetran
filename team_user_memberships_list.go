package fivetran

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fivetran/go-fivetran/teams"
)

// TeamUserMembershipsListService implements the Team Management, List all user memberships
// Ref. https://fivetran.com/docs/rest-api/teams#listallusermemberships
type TeamUserMembershipsListService struct {
	c      *Client
	teamId *string
	limit  *int
	cursor *string
}

type TeamUserMembershipsListResponse struct {
	Code string `json:"code"`
	Data struct {
		Items      []teams.TeamUserMembership `json:"items"`
		NextCursor string                     `json:"next_cursor"`
	} `json:"data"`
}

func (c *Client) NewTeamUserMembershipsList() *TeamUserMembershipsListService {
	return &TeamUserMembershipsListService{c: c}
}

func (s *TeamUserMembershipsListService) TeamId(value string) *TeamUserMembershipsListService {
	s.teamId = &value
	return s
}

func (s *TeamUserMembershipsListService) Limit(value int) *TeamUserMembershipsListService {
	s.limit = &value
	return s
}

func (s *TeamUserMembershipsListService) Cursor(value string) *TeamUserMembershipsListService {
	s.cursor = &value
	return s
}

func (s *TeamUserMembershipsListService) Do(ctx context.Context) (TeamUserMembershipsListResponse, error) {
	var response TeamUserMembershipsListResponse

	if s.teamId == nil {
		return response, fmt.Errorf("missing required teamId")
	}

	url := fmt.Sprintf("%v/teams/%v/users", s.c.baseURL, *s.teamId)
	expectedStatus := 200

	headers := s.c.commonHeaders()

	queries := make(map[string]string)
	if s.cursor != nil {
		queries["cursor"] = *s.cursor
	}
	if s.limit != nil {
		queries["limit"] = fmt.Sprint(*s.limit)
	}

	r := request{
		method:           "GET",
		url:              url,
		body:             nil,
		queries:          queries,
		headers:          headers,
		client:           s.c.httpClient,
		handleRateLimits: s.c.handleRateLimits,
		maxRetryAttempts: s.c.maxRetryAttempts,
	}

	respBody, respStatus, err := r.httpRequest(ctx)
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
