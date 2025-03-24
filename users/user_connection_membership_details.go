package users

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type UserConnectionMembershipDetailsService struct {
	httputils.HttpService
	userId      *string
	connectionId *string
}

func (s *UserConnectionMembershipDetailsService) UserId(value string) *UserConnectionMembershipDetailsService {
	s.userId = &value
	return s
}

func (s *UserConnectionMembershipDetailsService) ConnectionId(value string) *UserConnectionMembershipDetailsService {
	s.connectionId = &value
	return s
}

func (s *UserConnectionMembershipDetailsService) Do(ctx context.Context) (UserConnectionMembershipDetailsResponse, error) {
	var response UserConnectionMembershipDetailsResponse

	if s.userId == nil {
		return response, fmt.Errorf("missing required userId")
	}

	if s.connectionId == nil {
		return response, fmt.Errorf("missing required connectionId")
	}

	url := fmt.Sprintf("/users/%v/connections/%v", *s.userId, *s.connectionId)
	err := s.HttpService.Do(ctx, "GET", url, nil, nil, 200, &response)
	return response, err
}
