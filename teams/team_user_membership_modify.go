package teams

import (
	"context"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// TeamUsersModifyService implements the Team Management, Modify a user membership
// Ref. https://fivetran.com/docs/rest-api/teams#modifyausermembership
type TeamUserMembershipModifyService struct {
	httputils.HttpService
	teamId *string
	userId *string
	role   *string
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

	url := fmt.Sprintf("/teams/%v/users/%v", *s.teamId, *s.userId)
	err := s.HttpService.Do(ctx, "PATCH", url, s.request(), nil, 200, &response)
	return response, err
}