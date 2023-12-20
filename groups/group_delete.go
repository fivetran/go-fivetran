package groups

import (
	"context"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type GroupDeleteService struct {
	httputils.HttpService
	groupID *string
}

func (s *GroupDeleteService) GroupID(value string) *GroupDeleteService {
	s.groupID = &value
	return s
}

func (s *GroupDeleteService) Do(ctx context.Context) (common.CommonResponse, error) {
	var response common.CommonResponse
	
	if s.groupID == nil {
		return response, fmt.Errorf("missing required groupID")
	}

	url := fmt.Sprintf("/groups/%v", *s.groupID)
	err := s.HttpService.Do(ctx, "DELETE", url, nil, nil, 200, &response)
	return response, err
}
