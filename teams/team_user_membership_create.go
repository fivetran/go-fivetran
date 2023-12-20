package teams

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// TeamUserMembershipCreateService implements the Team Management, Add a user to a team
// Ref. https://fivetran.com/docs/rest-api/teams#addausertoateam
type TeamUserMembershipCreateService struct {
	httputils.HttpService
	teamId *string
	userId *string
	role   *string
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
	url := fmt.Sprintf("/teams/%v/users", *s.teamId)
	err := s.HttpService.Do(ctx, "POST", url, s.request(), nil, 201, &response)
	return response, err
}
