package teams

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type TeamGroupMembershipCreateService struct {
	httputils.HttpService
	teamId  *string
	groupId *string
	role    *string
}

func (s *TeamGroupMembershipCreateService) request() *teamGroupMembershipCreateRequest {
	return &teamGroupMembershipCreateRequest{
		GroupId: s.groupId,
		Role:    s.role,
	}
}

func (s *TeamGroupMembershipCreateService) TeamId(value string) *TeamGroupMembershipCreateService {
	s.teamId = &value
	return s
}

func (s *TeamGroupMembershipCreateService) GroupId(value string) *TeamGroupMembershipCreateService {
	s.groupId = &value
	return s
}

func (s *TeamGroupMembershipCreateService) Role(value string) *TeamGroupMembershipCreateService {
	s.role = &value
	return s
}

func (s *TeamGroupMembershipCreateService) Do(ctx context.Context) (TeamGroupMembershipCreateResponse, error) {
	var response TeamGroupMembershipCreateResponse
	if s.teamId == nil {
		return response, fmt.Errorf("missing required teamId")
	}
	url := fmt.Sprintf("/teams/%v/groups", *s.teamId)
	err := s.HttpService.Do(ctx, "POST", url, s.request(), nil, 201, &response)
	return response, err
}