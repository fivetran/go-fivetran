package teams

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type TeamsUpdateService struct {
	httputils.HttpService
	teamId      *string
	name        *string
	description *string
	role        *string
}

func (s *TeamsUpdateService) request() *teamsUpdateRequest {
	return &teamsUpdateRequest{
		Name:        s.name,
		Description: s.description,
		Role:        s.role,
	}
}

func (s *TeamsUpdateService) TeamId(value string) *TeamsUpdateService {
	s.teamId = &value
	return s
}

func (s *TeamsUpdateService) Name(value string) *TeamsUpdateService {
	s.name = &value
	return s
}

func (s *TeamsUpdateService) Role(value string) *TeamsUpdateService {
	s.role = &value
	return s
}

func (s *TeamsUpdateService) Description(value string) *TeamsUpdateService {
	s.description = &value
	return s
}

func (s *TeamsUpdateService) Do(ctx context.Context) (TeamsUpdateResponse, error) {
	var response TeamsUpdateResponse

	if s.teamId == nil {
		return response, fmt.Errorf("missing required teamId")
	}

	url := fmt.Sprintf("/teams/%v", *s.teamId)
	err := s.HttpService.Do(ctx, "PATCH", url, s.request(), nil, 200, &response)
	return response, err
}