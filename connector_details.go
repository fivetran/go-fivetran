package fivetran

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fivetran/go-fivetran/connectors"
	httputils "github.com/fivetran/go-fivetran/http_utils"
	"github.com/fivetran/go-fivetran/utils"
)

// ConnectorDetailsService implements the Connector Management, Retrieve Connector Details API.
// Ref. https://fivetran.com/docs/rest-api/connectors#retrieveconnectordetails
type ConnectorDetailsService struct {
	c           *Client
	connectorID *string
}

func (c *Client) NewConnectorDetails() *ConnectorDetailsService {
	return &ConnectorDetailsService{c: c}
}

func (s *ConnectorDetailsService) ConnectorID(value string) *ConnectorDetailsService {
	s.connectorID = &value
	return s
}

func (s *ConnectorDetailsService) do(ctx context.Context, response any) error {
	if s.connectorID == nil {
		return fmt.Errorf("missing required ConnectorID")
	}

	url := fmt.Sprintf("%v/connectors/%v", s.c.baseURL, *s.connectorID)
	expectedStatus := 200

	headers := s.c.commonHeaders()
	headers["Accept"] = restAPIv2

	r := httputils.Request{
		Method:           "GET",
		Url:              url,
		Body:             nil,
		Queries:          nil,
		Headers:          headers,
		Client:           s.c.httpClient,
		HandleRateLimits: s.c.handleRateLimits,
		MaxRetryAttempts: s.c.maxRetryAttempts,
	}

	respBody, respStatus, err := r.Do(ctx)

	if err != nil {
		return err
	}

	if err := json.Unmarshal(respBody, &response); err != nil {
		return err
	}

	if respStatus != expectedStatus {
		err := fmt.Errorf("status code: %v; expected: %v", respStatus, expectedStatus)
		return err
	}

	return nil
}

func (s *ConnectorDetailsService) Do(ctx context.Context) (connectors.DetailsWithConfigNoTestsResponse, error) {
	var response connectors.DetailsWithConfigNoTestsResponse

	err := s.do(ctx, &response)

	return response, err
}

func (s *ConnectorDetailsService) DoCustom(ctx context.Context) (connectors.DetailsWithCustomConfigNoTestsResponse, error) {
	var response connectors.DetailsWithCustomConfigNoTestsResponse

	err := s.do(ctx, &response)

	return response, err
}

func (s *ConnectorDetailsService) DoCustomMerged(ctx context.Context) (connectors.DetailsWithCustomMergedConfigNoTestsResponse, error) {
	var response connectors.DetailsWithCustomMergedConfigNoTestsResponse

	err := s.do(ctx, &response)

	if err == nil {
		err = utils.FetchFromMap(&response.Data.CustomConfig, &response.Data.Config)
	}

	return response, err
}
