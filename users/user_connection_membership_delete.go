package users

import (
	"context"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type UserConnectionMembershipDeleteService struct {
	httputils.HttpService
	userId      *string
	connectionId *string
}

func (s *UserConnectionMembershipDeleteService) UserId(value string) *UserConnectionMembershipDeleteService {
	s.userId = &value
	return s
}

func (s *UserConnectionMembershipDeleteService) ConnectionId(value string) *UserConnectionMembershipDeleteService {
	s.connectionId = &value
	return s
}

func (s *UserConnectionMembershipDeleteService) Do(ctx context.Context) (common.CommonResponse, error) {
	var response common.CommonResponse
	
	if s.userId == nil {
		return response, fmt.Errorf("missing required userId")
	}

	if s.connectionId == nil {
		return response, fmt.Errorf("missing required connectionId")
	}

	url := fmt.Sprintf("/users/%v/connections/%v", *s.userId, *s.connectionId)
	err := s.HttpService.Do(ctx, "DELETE", url, nil, nil, 200, &response)
	return response, err
}
