package connections

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type ConnectionSchemaReloadService struct {
    httputils.HttpService
	connectionID *string
	excludeMode *string
}

func (s *ConnectionSchemaReloadService) request() *connectionSchemaReloadRequest {
	mode := "PRESERVE"
	if s.excludeMode != nil {
		mode = *s.excludeMode
	}
	return &connectionSchemaReloadRequest{ExcludeMode: &mode}
}

func (s *ConnectionSchemaReloadService) ConnectionID(value string) *ConnectionSchemaReloadService {
	s.connectionID = &value
	return s
}

func (s *ConnectionSchemaReloadService) ExcludeMode(value string) *ConnectionSchemaReloadService {
	s.excludeMode = &value
	return s
}

func (s *ConnectionSchemaReloadService) Do(ctx context.Context) (ConnectionSchemaDetailsResponse, error) {
	var response ConnectionSchemaDetailsResponse

	if s.connectionID == nil {
		return response, fmt.Errorf("missing required ConnectionID")
	}
	
	url := fmt.Sprintf("/connections/%v/schemas/reload", *s.connectionID)
	err := s.HttpService.Do(ctx, "POST", url, s.request(), nil, 200, &response)
	return response, err
}