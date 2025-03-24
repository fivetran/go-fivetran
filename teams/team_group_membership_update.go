package teams

import (
	"context"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type TeamGroupMembershipUpdateService struct {
	httputils.HttpService
	teamId  *string
	groupId *string
	role    *string
}

func (s *TeamGroupMembershipUpdateService) request() *teamGroupMembershipUpdateRequest {
	return &teamGroupMembershipUpdateRequest{
		Role: s.role,
	}
}

func (s *TeamGroupMembershipUpdateService) TeamId(value string) *TeamGroupMembershipUpdateService {
	s.teamId = &value
	return s
}

func (s *TeamGroupMembershipUpdateService) GroupId(value string) *TeamGroupMembershipUpdateService {
	s.groupId = &value
	return s
}

func (s *TeamGroupMembershipUpdateService) Role(value string) *TeamGroupMembershipUpdateService {
	s.role = &value
	return s
}

func (s *TeamGroupMembershipUpdateService) Do(ctx context.Context) (common.CommonResponse, error) {
	var response common.CommonResponse

	if s.teamId == nil {
		return response, fmt.Errorf("missing required teamId")
	}

	if s.groupId == nil {
		return response, fmt.Errorf("missing required groupId")
	}

	url := fmt.Sprintf("/teams/%v/groups/%v", *s.teamId, *s.groupId)
	err := s.HttpService.Do(ctx, "PATCH", url, s.request(), nil, 200, &response)
	return response, err
}