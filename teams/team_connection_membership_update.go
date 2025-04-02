package teams

import (
	"context"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type TeamConnectionMembershipUpdateService struct {
	httputils.HttpService
	teamId      *string
	connectionId *string
	role        *string
}

func (s *TeamConnectionMembershipUpdateService) request() *teamConnectionMembershipUpdateRequest {
	return &teamConnectionMembershipUpdateRequest{
		Role: s.role,
	}
}

func (s *TeamConnectionMembershipUpdateService) TeamId(value string) *TeamConnectionMembershipUpdateService {
	s.teamId = &value
	return s
}

func (s *TeamConnectionMembershipUpdateService) ConnectionId(value string) *TeamConnectionMembershipUpdateService {
	s.connectionId = &value
	return s
}

func (s *TeamConnectionMembershipUpdateService) Role(value string) *TeamConnectionMembershipUpdateService {
	s.role = &value
	return s
}

func (s *TeamConnectionMembershipUpdateService) Do(ctx context.Context) (common.CommonResponse, error) {
	var response common.CommonResponse

	if s.teamId == nil {
		return response, fmt.Errorf("missing required teamId")
	}

	if s.connectionId == nil {
		return response, fmt.Errorf("missing required connectionId")
	}

	url := fmt.Sprintf("/teams/%v/connections/%v", *s.teamId, *s.connectionId)
	err := s.HttpService.Do(ctx, "PATCH", url, s.request(), nil, 200, &response)
	return response, err
}