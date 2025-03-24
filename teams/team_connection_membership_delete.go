package teams

import (
	"context"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type TeamConnectionMembershipDeleteService struct {
	httputils.HttpService
	teamId      *string
	connectionId *string
}

func (s *TeamConnectionMembershipDeleteService) TeamId(value string) *TeamConnectionMembershipDeleteService {
	s.teamId = &value
	return s
}

func (s *TeamConnectionMembershipDeleteService) ConnectionId(value string) *TeamConnectionMembershipDeleteService {
	s.connectionId = &value
	return s
}

func (s *TeamConnectionMembershipDeleteService) Do(ctx context.Context) (common.CommonResponse, error) {
	var response common.CommonResponse
	
	if s.teamId == nil {
		return response, fmt.Errorf("missing required teamId")
	}

	if s.connectionId == nil {
		return response, fmt.Errorf("missing required ConnectionId")
	}

	url := fmt.Sprintf("/teams/%v/Connections/%v", *s.teamId, *s.connectionId)
	err := s.HttpService.Do(ctx, "DELETE", url, nil, nil, 200, &response)
	return response, err
}