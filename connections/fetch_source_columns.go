package connections

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type FetchSourceColumnsService struct {
	httputils.HttpService
	connectionId *string
	schema       *string
	tables       []string
}

func (s *FetchSourceColumnsService) ConnectionId(value string) *FetchSourceColumnsService {
	s.connectionId = &value
	return s
}

func (s *FetchSourceColumnsService) Schema(value string) *FetchSourceColumnsService {
	s.schema = &value
	return s
}

func (s *FetchSourceColumnsService) Tables(value []string) *FetchSourceColumnsService {
	s.tables = value
	return s
}

func (s *FetchSourceColumnsService) Do(ctx context.Context) (MultipleTableColumnsConfigResponse, error) {
	var response MultipleTableColumnsConfigResponse
	url := fmt.Sprintf("/connections/%v/schemas/%v/fetch-source-columns", *s.connectionId, *s.schema)

	request := struct {
		Tables []string `json:"tables"`
	}{
		Tables: s.tables,
	}

	err := s.HttpService.Do(ctx, "POST", url, request, nil, 200, &response)
	return response, err
}
