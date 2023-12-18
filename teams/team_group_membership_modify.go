package teams

import (
	"context"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// TeamGroupMembershipModifyService implements the Team Management, Update group membership
// Ref. https://fivetran.com/docs/rest-api/teams#updategroupmembership
type TeamGroupMembershipModifyService struct {
	httputils.HttpService
	teamId  *string
	groupId *string
	role    *string
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

	url := fmt.Sprintf("/teams/%v/groups/%v", *s.teamId, *s.groupId)
	err := s.HttpService.Do(ctx, "PATCH", url, s.request(), nil, 200, &response)
	return response, err
}