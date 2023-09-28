package fivetran

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fivetran/go-fivetran/connectors"
)

// ConnectorSetupTestsService implements the Connector Management, Run connector setup tests API.
// Ref. https://fivetran.com/docs/rest-api/connectors#runconnectorsetuptests
type ConnectorSetupTestsService struct {
	c                 *Client
	connectorID       *string
	trustCertificates *bool
	trustFingerprints *bool
}

type connectorSetupTestsRequest struct {
	TrustCertificates *bool `json:"trust_certificates,omitempty"`
	TrustFingerprints *bool `json:"trust_fingerprints,omitempty"`
}

func (c *Client) NewConnectorSetupTests() *ConnectorSetupTestsService {
	return &ConnectorSetupTestsService{c: c}
}

func (s *ConnectorSetupTestsService) request() *connectorSetupTestsRequest {
	return &connectorSetupTestsRequest{
		TrustCertificates: s.trustCertificates,
		TrustFingerprints: s.trustFingerprints,
	}
}

func (s *ConnectorSetupTestsService) ConnectorID(value string) *ConnectorSetupTestsService {
	s.connectorID = &value
	return s
}

func (s *ConnectorSetupTestsService) TrustCertificates(value bool) *ConnectorSetupTestsService {
	s.trustCertificates = &value
	return s
}

func (s *ConnectorSetupTestsService) TrustFingerprints(value bool) *ConnectorSetupTestsService {
	s.trustFingerprints = &value
	return s
}

func (s *ConnectorSetupTestsService) Do(ctx context.Context) (connectors.DetailsWithConfigResponse, error) {
	var response connectors.DetailsWithConfigResponse

	err := s.do(ctx, &response)

	return response, err
}

func (s *ConnectorSetupTestsService) DoCustom(ctx context.Context) (connectors.DetailsWithCustomConfigResponse, error) {
	var response connectors.DetailsWithCustomConfigResponse

	err := s.do(ctx, &response)

	return response, err
}

func (s *ConnectorSetupTestsService) DoCustomMerged(ctx context.Context) (connectors.DetailsWithCustomMergedConfigResponse, error) {
	var response connectors.DetailsWithCustomMergedConfigResponse

	err := s.do(ctx, &response)

	return response, err
}

func (s *ConnectorSetupTestsService) do(ctx context.Context, response any) error {
	if s.connectorID == nil {
		return fmt.Errorf("missing required ConnectorID")
	}

	url := fmt.Sprintf("%v/connectors/%v/test", s.c.baseURL, *s.connectorID)
	expectedStatus := 200

	headers := s.c.commonHeaders()
	headers["Content-Type"] = "application/json"
	headers["Accept"] = restAPIv2

	reqBody, err := json.Marshal(s.request())
	if err != nil {
		return err
	}

	r := request{
		method:           "POST",
		url:              url,
		body:             reqBody,
		queries:          nil,
		headers:          headers,
		client:           s.c.httpClient,
		handleRateLimits: s.c.handleRateLimits,
		maxRetryAttempts: s.c.maxRetryAttempts,
	}

	respBody, respStatus, err := r.httpRequest(ctx)
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
