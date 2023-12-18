package teams

import (
	"context"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// TeamsDeleteService implements the Team Management, Delete a Team.
// Ref. https://fivetran.com/docs/rest-api/teams#deleteateam
type TeamsDeleteService struct {
	httputils.HttpService
	teamId *string
}

func (s *TeamsDeleteService) TeamId(value string) *TeamsDeleteService {
	s.teamId = &value
	return s
}

func (s *TeamsDeleteService) Do(ctx context.Context) (common.CommonResponse, error) {
	var response common.CommonResponse
	
	if s.teamId == nil {
		return response, fmt.Errorf("missing required teamId")
	}

	url := fmt.Sprintf("/teams/%v", *s.teamId)
	err := s.HttpService.Do(ctx, "DELETE", url, nil, nil, 200, &response)
	return response, err
}