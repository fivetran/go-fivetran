package teams

import (
	"context"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type TeamsCreateService struct {
	httputils.HttpService
	name        *string
	description *string
	role        *string
}

func (s *TeamsCreateService) request() *teamsCreateRequest {
	return &teamsCreateRequest{
		Name:        s.name,
		Description: s.description,
		Role:        s.role,
	}
}

func (s *TeamsCreateService) Name(value string) *TeamsCreateService {
	s.name = &value
	return s
}

func (s *TeamsCreateService) Role(value string) *TeamsCreateService {
	s.role = &value
	return s
}

func (s *TeamsCreateService) Description(value string) *TeamsCreateService {
	s.description = &value
	return s
}

func (s *TeamsCreateService) Do(ctx context.Context) (TeamsCreateResponse, error) {
	var response TeamsCreateResponse
	err := s.HttpService.Do(ctx, "POST", "/teams", s.request(), nil, 201, &response)
	return response, err
}