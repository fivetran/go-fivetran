package teams

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// TeamUserMembershipDetailsService implements the Team Management, retrieve Team Details.
// Ref. https://fivetran.com/docs/rest-api/teams#retrieveusermembershipinateam
type TeamUserMembershipDetailsService struct {
	httputils.HttpService
	teamId *string
	userId *string
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

	url := fmt.Sprintf("/teams/%v/users/%v", *s.teamId, *s.userId)
	err := s.HttpService.Do(ctx, "GET", url, nil, nil, 200, &response)
	return response, err
}