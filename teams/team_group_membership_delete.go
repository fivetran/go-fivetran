package teams

import (
	"context"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// TeamGroupMembershipDeleteService implements the Team Management, Delete group membership
// Ref. https://fivetran.com/docs/rest-api/teams#deletegroupmembership
type TeamGroupMembershipDeleteService struct {
	httputils.HttpService
	teamId  *string
	groupId *string
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

	url := fmt.Sprintf("/teams/%v/groups/%v", *s.teamId, *s.groupId)
	err := s.HttpService.Do(ctx, "DELETE", url, nil, nil, 200, &response)
	return response, err
}