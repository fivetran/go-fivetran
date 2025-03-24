package connections

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type ConnectionSchemaDetailsService struct {
    httputils.HttpService
	connectionID *string
}

func (s *ConnectionSchemaDetailsService) ConnectionID(value string) *ConnectionSchemaDetailsService {
	s.connectionID = &value
	return s
}

func (s *ConnectionSchemaDetailsService) Do(ctx context.Context) (ConnectionSchemaDetailsResponse, error) {
	var response ConnectionSchemaDetailsResponse

	if s.connectionID == nil {
		return response, fmt.Errorf("missing required ConnectionID")
	}

	url := fmt.Sprintf("/connections/%v/schemas", *s.connectionID)
	err := s.HttpService.Do(ctx, "GET", url, nil, nil, 200, &response)
	return response, err
}