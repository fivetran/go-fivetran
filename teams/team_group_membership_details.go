package teams

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// TeamGroupMembershipDetailsService implements the Team Management, Retrieve group membership
// Ref. https://fivetran.com/docs/rest-api/teams#retrievegroupmembership
type TeamGroupMembershipDetailsService struct {
	httputils.HttpService
	teamId  *string
	groupId *string
}

func (s *TeamGroupMembershipDetailsService) TeamId(value string) *TeamGroupMembershipDetailsService {
	s.teamId = &value
	return s
}

func (s *TeamGroupMembershipDetailsService) GroupId(value string) *TeamGroupMembershipDetailsService {
	s.groupId = &value
	return s
}

func (s *TeamGroupMembershipDetailsService) Do(ctx context.Context) (TeamGroupMembershipDetailsResponse, error) {
	var response TeamGroupMembershipDetailsResponse

	if s.teamId == nil {
		return response, fmt.Errorf("missing required teamId")
	}

	if s.groupId == nil {
		return response, fmt.Errorf("missing required groupId")
	}

	url := fmt.Sprintf("/teams/%v/groups/%v", *s.teamId, *s.groupId)
	err := s.HttpService.Do(ctx, "GET", url, nil, nil, 200, &response)
	return response, err
}