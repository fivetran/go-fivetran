package teams

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type TeamConnectionMembershipCreateService struct {
	httputils.HttpService
	teamId      *string
	connectionId *string
	role        *string
}

func (s *TeamConnectionMembershipCreateService) request() *teamConnectionMembershipCreateRequest {
	return &teamConnectionMembershipCreateRequest{
		ConnectionId: s.connectionId,
		Role:        s.role,
	}
}

func (s *TeamConnectionMembershipCreateService) TeamId(value string) *TeamConnectionMembershipCreateService {
	s.teamId = &value
	return s
}

func (s *TeamConnectionMembershipCreateService) ConnectionId(value string) *TeamConnectionMembershipCreateService {
	s.connectionId = &value
	return s
}

func (s *TeamConnectionMembershipCreateService) Role(value string) *TeamConnectionMembershipCreateService {
	s.role = &value
	return s
}

func (s *TeamConnectionMembershipCreateService) Do(ctx context.Context) (TeamConnectionMembershipCreateResponse, error) {
	var response TeamConnectionMembershipCreateResponse
	if s.teamId == nil {
		return response, fmt.Errorf("missing required teamId")
	}
	url := fmt.Sprintf("/teams/%v/connections", *s.teamId)
	err := s.HttpService.Do(ctx, "POST", url, s.request(), nil, 201, &response)
	return response, err
}
