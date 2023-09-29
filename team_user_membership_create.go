package fivetran

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	"github.com/fivetran/go-fivetran/teams"
)

// TeamUserMembershipCreateService implements the Team Management, Add a user to a team
// Ref. https://fivetran.com/docs/rest-api/teams#addausertoateam
type TeamUserMembershipCreateService struct {
	c      *Client
	teamId *string
	userId *string
	role   *string
}

type teamUserMembershipCreateRequest struct {
	UserId *string `json:"user_id,omitempty"`
	Role   *string `json:"role,omitempty"`
}

type TeamUserMembershipCreateResponse struct {
	common.CommonResponse
	Data teams.TeamUserMembership `json:"data"`
}

func (c *Client) NewTeamUserMembershipCreate() *TeamUserMembershipCreateService {
	return &TeamUserMembershipCreateService{c: c}
}

func (s *TeamUserMembershipCreateService) request() *teamUserMembershipCreateRequest {
	return &teamUserMembershipCreateRequest{
		UserId: s.userId,
		Role:   s.role,
	}
}

func (s *TeamUserMembershipCreateService) TeamId(value string) *TeamUserMembershipCreateService {
	s.teamId = &value
	return s
}

func (s *TeamUserMembershipCreateService) UserId(value string) *TeamUserMembershipCreateService {
	s.userId = &value
	return s
}

func (s *TeamUserMembershipCreateService) Role(value string) *TeamUserMembershipCreateService {
	s.role = &value
	return s
}

func (s *TeamUserMembershipCreateService) Do(ctx context.Context) (TeamUserMembershipCreateResponse, error) {
	var response TeamUserMembershipCreateResponse

	if s.teamId == nil {
		return response, fmt.Errorf("missing required teamId")
	}

	url := fmt.Sprintf("%v/teams/%v/users", s.c.baseURL, *s.teamId)
	expectedStatus := 201

	headers := s.c.commonHeaders()
	headers["Content-Type"] = "application/json"
	headers["Accept"] = restAPIv2

	reqBody, err := json.Marshal(s.request())
	if err != nil {
		return response, err
	}

	r := request{
		method:  "POST",
		url:     url,
		body:    reqBody,
		queries: nil,
		headers: headers,
		client:  s.c.httpClient,
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
