package connections

import (
	"context"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type ConnectionSyncService struct {
	httputils.HttpService
	connectionID *string
}

func (s *ConnectionSyncService) ConnectionID(connectionID string) *ConnectionSyncService {
	s.connectionID = &connectionID
	return s
}

func (s *ConnectionSyncService) Do(ctx context.Context) (common.CommonResponse, error) {
	var response common.CommonResponse

	if s.connectionID == nil {
		return response, fmt.Errorf("missing required ConnectionID")
	}

	err := s.HttpService.Do(ctx, "POST", fmt.Sprintf("/connections/%v/force", *s.connectionID), nil, nil, 200, &response)
	return response, err
}
