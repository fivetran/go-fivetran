package teams

import (
	"context"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// TeamUserMembershipDeleteService implements the Team Management, Delete a user from a team
// Ref. https://fivetran.com/docs/rest-api/teams#deleteauserfromateam
type TeamUserMembershipDeleteService struct {
	httputils.HttpService
	teamId *string
	userId *string
}

func (s *TeamUserMembershipDeleteService) TeamId(value string) *TeamUserMembershipDeleteService {
	s.teamId = &value
	return s
}

func (s *TeamUserMembershipDeleteService) UserId(value string) *TeamUserMembershipDeleteService {
	s.userId = &value
	return s
}

func (s *TeamUserMembershipDeleteService) Do(ctx context.Context) (common.CommonResponse, error) {
	var response common.CommonResponse
	
	if s.teamId == nil {
		return response, fmt.Errorf("missing required teamId")
	}

	if s.userId == nil {
		return response, fmt.Errorf("missing required userId")
	}

	url := fmt.Sprintf("/teams/%v/users/%v", *s.teamId, *s.userId)
	err := s.HttpService.Do(ctx, "DELETE", url, nil, nil, 200, &response)
	return response, err
}