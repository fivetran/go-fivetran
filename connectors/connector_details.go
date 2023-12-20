package connectors

import (
	"context"
	"fmt"

	"github.com/fivetran/go-fivetran/utils"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// ConnectorDetailsService implements the Connector Management, Retrieve Connector Details API.
// Ref. https://fivetran.com/docs/rest-api/connectors#retrieveconnectordetails
type ConnectorDetailsService struct {
    httputils.HttpService
	connectorID *string
}

func (s *ConnectorDetailsService) ConnectorID(value string) *ConnectorDetailsService {
	s.connectorID = &value
	return s
}

func (s *ConnectorDetailsService) do(ctx context.Context, response any) error {
	if s.connectorID == nil {
		return fmt.Errorf("missing required ConnectorID")
	}
    url := fmt.Sprintf("/connectors/%v", *s.connectorID)
    err := s.HttpService.Do(ctx, "GET", url, nil, nil, 200, &response)
    return err
}

func (s *ConnectorDetailsService) Do(ctx context.Context) (DetailsWithConfigNoTestsResponse, error) {
	var response DetailsWithConfigNoTestsResponse

	err := s.do(ctx, &response)

	return response, err
}

func (s *ConnectorDetailsService) DoCustom(ctx context.Context) (DetailsWithCustomConfigNoTestsResponse, error) {
	var response DetailsWithCustomConfigNoTestsResponse

	err := s.do(ctx, &response)

	return response, err
}

func (s *ConnectorDetailsService) DoCustomMerged(ctx context.Context) (DetailsWithCustomMergedConfigNoTestsResponse, error) {
	var response DetailsWithCustomMergedConfigNoTestsResponse

	err := s.do(ctx, &response)

	if err == nil {
		err = utils.FetchFromMap(&response.Data.CustomConfig, &response.Data.Config)
	}

	return response, err
}
