package teams

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type TeamConnectionMembershipDetailsService struct {
	httputils.HttpService
	teamId      *string
	connectionId *string
}

func (s *TeamConnectionMembershipDetailsService) TeamId(value string) *TeamConnectionMembershipDetailsService {
	s.teamId = &value
	return s
}

func (s *TeamConnectionMembershipDetailsService) ConnectionId(value string) *TeamConnectionMembershipDetailsService {
	s.connectionId = &value
	return s
}

func (s *TeamConnectionMembershipDetailsService) Do(ctx context.Context) (TeamConnectionMembershipDetailsResponse, error) {
	var response TeamConnectionMembershipDetailsResponse

	if s.teamId == nil {
		return response, fmt.Errorf("missing required teamId")
	}

	if s.connectionId == nil {
		return response, fmt.Errorf("missing required connectionId")
	}

	url := fmt.Sprintf("/teams/%v/connections/%v", *s.teamId, *s.connectionId)
	err := s.HttpService.Do(ctx, "GET", url, nil, nil, 200, &response)
	return response, err
}