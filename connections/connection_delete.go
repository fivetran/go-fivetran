package connections

import (
	"context"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type ConnectionDeleteService struct {
    httputils.HttpService
	connectionID *string
}

func (s *ConnectionDeleteService) ConnectionID(connectionID string) *ConnectionDeleteService {
	s.connectionID = &connectionID
	return s
}

func (s *ConnectionDeleteService) Do(ctx context.Context) (common.CommonResponse, error) {
	var response common.CommonResponse
	
	if s.connectionID == nil {
		return response, fmt.Errorf("missing required connectionID")
	}

	url := fmt.Sprintf("/connections/%v", *s.connectionID)
	err := s.HttpService.Do(ctx, "DELETE", url, nil, nil, 200, &response)
	return response, err
}