package teams

import (
	"context"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type TeamUserMembershipUpdateService struct {
	httputils.HttpService
	teamId *string
	userId *string
	role   *string
}

func (s *TeamUserMembershipUpdateService) request() *teamUserMembershipUpdateRequest {
	return &teamUserMembershipUpdateRequest{
		Role: s.role,
	}
}

func (s *TeamUserMembershipUpdateService) TeamId(value string) *TeamUserMembershipUpdateService {
	s.teamId = &value
	return s
}

func (s *TeamUserMembershipUpdateService) UserId(value string) *TeamUserMembershipUpdateService {
	s.userId = &value
	return s
}

func (s *TeamUserMembershipUpdateService) Role(value string) *TeamUserMembershipUpdateService {
	s.role = &value
	return s
}

func (s *TeamUserMembershipUpdateService) Do(ctx context.Context) (common.CommonResponse, error) {
	var response common.CommonResponse

	if s.teamId == nil {
		return response, fmt.Errorf("missing required teamId")
	}

	if s.userId == nil {
		return response, fmt.Errorf("missing required userId")
	}

	url := fmt.Sprintf("/teams/%v/users/%v", *s.teamId, *s.userId)
	err := s.HttpService.Do(ctx, "PATCH", url, s.request(), nil, 200, &response)
	return response, err
}