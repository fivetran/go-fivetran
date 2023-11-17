package groups

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type GroupSshKeyService struct {
	httputils.HttpService
	groupID *string
}

func (s *GroupSshKeyService) GroupID(value string) *GroupSshKeyService {
	s.groupID = &value
	return s
}

func (s *GroupSshKeyService) Do(ctx context.Context) (GroupSshKeyResponse, error) {
	var response GroupSshKeyResponse
	url := fmt.Sprintf("/groups/%v/public-key", *s.groupID)
	err := s.HttpService.Do(ctx, "GET", url, nil, nil, 200, &response)
	return response, err
}
