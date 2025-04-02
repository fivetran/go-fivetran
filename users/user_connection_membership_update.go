package users

import (
	"context"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type UserConnectionMembershipUpdateService struct {
	httputils.HttpService
	userId      *string
	connectionId *string
	role        *string
}

func (s *UserConnectionMembershipUpdateService) request() *userConnectionMembershipUpdateRequest {
	return &userConnectionMembershipUpdateRequest{
		Role: s.role,
	}
}

func (s *UserConnectionMembershipUpdateService) UserId(value string) *UserConnectionMembershipUpdateService {
	s.userId = &value
	return s
}

func (s *UserConnectionMembershipUpdateService) ConnectionId(value string) *UserConnectionMembershipUpdateService {
	s.connectionId = &value
	return s
}

func (s *UserConnectionMembershipUpdateService) Role(value string) *UserConnectionMembershipUpdateService {
	s.role = &value
	return s
}

func (s *UserConnectionMembershipUpdateService) Do(ctx context.Context) (common.CommonResponse, error) {
	var response common.CommonResponse

	if s.userId == nil {
		return response, fmt.Errorf("missing required userId")
	}

	if s.connectionId == nil {
		return response, fmt.Errorf("missing required connectionId")
	}

	url := fmt.Sprintf("/users/%v/connections/%v", *s.userId, *s.connectionId)
	err := s.HttpService.Do(ctx, "PATCH", url, s.request(), nil, 200, &response)
	return response, err
}
