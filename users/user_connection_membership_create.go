package users

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type UserConnectionMembershipCreateService struct {
	httputils.HttpService
	userId      *string
	connectionId *string
	role        *string
}

func (s *UserConnectionMembershipCreateService) request() *userConnectionMembershipCreateRequest {
	return &userConnectionMembershipCreateRequest{
		ConnectionId: s.connectionId,
		Role:        s.role,
	}
}

func (s *UserConnectionMembershipCreateService) UserId(value string) *UserConnectionMembershipCreateService {
	s.userId = &value
	return s
}

func (s *UserConnectionMembershipCreateService) ConnectionId(value string) *UserConnectionMembershipCreateService {
	s.connectionId = &value
	return s
}

func (s *UserConnectionMembershipCreateService) Role(value string) *UserConnectionMembershipCreateService {
	s.role = &value
	return s
}

func (s *UserConnectionMembershipCreateService) Do(ctx context.Context) (UserConnectionMembershipCreateResponse, error) {
	var response UserConnectionMembershipCreateResponse
	url := fmt.Sprintf("/users/%v/connections", *s.userId)
	err := s.HttpService.Do(ctx, "POST", url, s.request(), nil, 201, &response)
	return response, err
}