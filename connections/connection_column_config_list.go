package connections

import (
    "context"
    "fmt"

    httputils "github.com/fivetran/go-fivetran/http_utils"
)

type ConnectionColumnConfigListService struct {
    httputils.HttpService
    connectionId         *string
    schema              *string
    table               *string
}

func (s *ConnectionColumnConfigListService) ConnectionId(value string) *ConnectionColumnConfigListService {
    s.connectionId = &value
    return s
}

func (s *ConnectionColumnConfigListService) Schema(value string) *ConnectionColumnConfigListService {
    s.schema = &value
    return s
}

func (s *ConnectionColumnConfigListService) Table(value string) *ConnectionColumnConfigListService {
    s.table = &value
    return s
}

func (s *ConnectionColumnConfigListService) Do(ctx context.Context) (ConnectionColumnConfigListResponse, error) {
    var response ConnectionColumnConfigListResponse
    url := fmt.Sprintf("/connections/%v/schemas/%v/tables/%v/columns", *s.connectionId, *s.schema, *s.table)
    err := s.HttpService.Do(ctx, "GET", url, nil, nil, 200, &response)
    return response, err
}
