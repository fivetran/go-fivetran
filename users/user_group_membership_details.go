package users

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type UserGroupMembershipDetailsService struct {
	httputils.HttpService
	userId  *string
	groupId *string
}

func (s *UserGroupMembershipDetailsService) UserId(value string) *UserGroupMembershipDetailsService {
	s.userId = &value
	return s
}

func (s *UserGroupMembershipDetailsService) GroupId(value string) *UserGroupMembershipDetailsService {
	s.groupId = &value
	return s
}

func (s *UserGroupMembershipDetailsService) Do(ctx context.Context) (UserGroupMembershipDetailsResponse, error) {
	var response UserGroupMembershipDetailsResponse

	if s.userId == nil {
		return response, fmt.Errorf("missing required userId")
	}

	if s.groupId == nil {
		return response, fmt.Errorf("missing required groupId")
	}

	url := fmt.Sprintf("/users/%v/groups/%v", *s.userId, *s.groupId)
	err := s.HttpService.Do(ctx, "GET", url, nil, nil, 200, &response)
	return response, err
}
