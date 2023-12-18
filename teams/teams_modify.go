package teams

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// TeamsModifyService implements the Team Management, Modify a Team.
// Ref. https://fivetran.com/docs/rest-api/teams#modifyateam
type TeamsModifyService struct {
	httputils.HttpService
	teamId      *string
	name        *string
	description *string
	role        *string
}

func (s *TeamsModifyService) request() *teamsModifyRequest {
	return &teamsModifyRequest{
		Name:        s.name,
		Description: s.description,
		Role:        s.role,
	}
}

func (s *TeamsModifyService) TeamId(value string) *TeamsModifyService {
	s.teamId = &value
	return s
}

func (s *TeamsModifyService) Name(value string) *TeamsModifyService {
	s.name = &value
	return s
}

func (s *TeamsModifyService) Role(value string) *TeamsModifyService {
	s.role = &value
	return s
}

func (s *TeamsModifyService) Description(value string) *TeamsModifyService {
	s.description = &value
	return s
}

func (s *TeamsModifyService) Do(ctx context.Context) (TeamsModifyResponse, error) {
	var response TeamsModifyResponse

	if s.teamId == nil {
		return response, fmt.Errorf("missing required teamId")
	}

	url := fmt.Sprintf("/teams/%v", *s.teamId)
	err := s.HttpService.Do(ctx, "PATCH", url, s.request(), nil, 200, &response)
	return response, err
}