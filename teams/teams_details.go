package teams

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type TeamsDetailsService struct {
	httputils.HttpService
	teamId *string
}

func (s *TeamsDetailsService) TeamId(value string) *TeamsDetailsService {
	s.teamId = &value
	return s
}

func (s *TeamsDetailsService) Do(ctx context.Context) (TeamsDetailsResponse, error) {
	var response TeamsDetailsResponse

	if s.teamId == nil {
		return response, fmt.Errorf("missing required teamId")
	}

	url := fmt.Sprintf("/teams/%v", *s.teamId)
	err := s.HttpService.Do(ctx, "GET", url, nil, nil, 200, &response)
	return response, err
}