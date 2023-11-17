package groups

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type GroupServiceAccountService struct {
	httputils.HttpService
	groupID *string
}

func (s *GroupServiceAccountService) GroupID(value string) *GroupServiceAccountService {
	s.groupID = &value
	return s
}

func (s *GroupServiceAccountService) Do(ctx context.Context) (GroupServiceAccountResponse, error) {
	var response GroupServiceAccountResponse
	url := fmt.Sprintf("/groups/%v/service-account", *s.groupID)
	err := s.HttpService.Do(ctx, "GET", url, nil, nil, 200, &response)
	return response, err
}
