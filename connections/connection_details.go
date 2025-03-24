package connections

import (
	"context"
	"fmt"

	"github.com/fivetran/go-fivetran/utils"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type ConnectionDetailsService struct {
    httputils.HttpService
	connectionID *string
}

func (s *ConnectionDetailsService) ConnectionID(value string) *ConnectionDetailsService {
	s.connectionID = &value
	return s
}

func (s *ConnectionDetailsService) do(ctx context.Context, response any) error {
	if s.connectionID == nil {
		return fmt.Errorf("missing required ConnectionID")
	}
    url := fmt.Sprintf("/connections/%v", *s.connectionID)
    err := s.HttpService.Do(ctx, "GET", url, nil, nil, 200, &response)
    return err
}

func (s *ConnectionDetailsService) Do(ctx context.Context) (DetailsWithConfigNoTestsResponse, error) {
	var response DetailsWithConfigNoTestsResponse

	err := s.do(ctx, &response)

	return response, err
}

func (s *ConnectionDetailsService) DoCustom(ctx context.Context) (DetailsWithCustomConfigNoTestsResponse, error) {
	var response DetailsWithCustomConfigNoTestsResponse

	err := s.do(ctx, &response)

	return response, err
}

func (s *ConnectionDetailsService) DoCustomMerged(ctx context.Context) (DetailsWithCustomMergedConfigNoTestsResponse, error) {
	var response DetailsWithCustomMergedConfigNoTestsResponse

	err := s.do(ctx, &response)

	if err == nil {
		err = utils.FetchFromMap(&response.Data.CustomConfig, &response.Data.Config)
	}

	return response, err
}
