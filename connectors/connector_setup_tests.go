package connectors

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// ConnectorSetupTestsService implements the Connector Management, Run connector setup tests API.
// Ref. https://fivetran.com/docs/rest-api/connectors#runconnectorsetuptests
type ConnectorSetupTestsService struct {
    httputils.HttpService
	connectorID       *string
	trustCertificates *bool
	trustFingerprints *bool
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

func (s *ConnectorSetupTestsService) Do(ctx context.Context) (DetailsWithConfigResponse, error) {
	var response DetailsWithConfigResponse

	err := s.do(ctx, &response)

	return response, err
}

func (s *ConnectorSetupTestsService) DoCustom(ctx context.Context) (DetailsWithCustomConfigResponse, error) {
	var response DetailsWithCustomConfigResponse

	err := s.do(ctx, &response)

	return response, err
}

func (s *ConnectorSetupTestsService) DoCustomMerged(ctx context.Context) (DetailsWithCustomMergedConfigResponse, error) {
	var response DetailsWithCustomMergedConfigResponse

	err := s.do(ctx, &response)

	return response, err
}

func (s *ConnectorSetupTestsService) do(ctx context.Context, response any) error {
	if s.connectorID == nil {
		return fmt.Errorf("missing required ConnectorID")
	}
	url := fmt.Sprintf("/connectors/%v/test", *s.connectorID)
	err := s.HttpService.Do(ctx, "POST", url, s.request(), nil, 200, &response)
	return err
}