package teams

import (
	"context"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// TeamRoleDeleteService implements the Team Management, Delete team role in account
// Ref. https://fivetran.com/docs/rest-api/teams#deleteteamroleinaccount
type TeamsDeleteRoleInAccountService struct {
	httputils.HttpService
	teamId *string
}

func (s *TeamsDeleteRoleInAccountService) TeamId(value string) *TeamsDeleteRoleInAccountService {
	s.teamId = &value
	return s
}

func (s *TeamsDeleteRoleInAccountService) Do(ctx context.Context) (common.CommonResponse, error) {
	var response common.CommonResponse
	
	if s.teamId == nil {
		return response, fmt.Errorf("missing required teamId")
	}

	url := fmt.Sprintf("/teams/%v/role", *s.teamId)
	err := s.HttpService.Do(ctx, "DELETE", url, nil, nil, 200, &response)
	return response, err
}