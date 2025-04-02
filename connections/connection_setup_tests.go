package connections

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type ConnectionSetupTestsService struct {
    httputils.HttpService
	connectionID       *string
	trustCertificates *bool
	trustFingerprints *bool
}
func (s *ConnectionSetupTestsService) request() *connectionSetupTestsRequest {
	return &connectionSetupTestsRequest{
		TrustCertificates: s.trustCertificates,
		TrustFingerprints: s.trustFingerprints,
	}
}

func (s *ConnectionSetupTestsService) ConnectionID(value string) *ConnectionSetupTestsService {
	s.connectionID = &value
	return s
}

func (s *ConnectionSetupTestsService) TrustCertificates(value bool) *ConnectionSetupTestsService {
	s.trustCertificates = &value
	return s
}

func (s *ConnectionSetupTestsService) TrustFingerprints(value bool) *ConnectionSetupTestsService {
	s.trustFingerprints = &value
	return s
}

func (s *ConnectionSetupTestsService) Do(ctx context.Context) (DetailsWithConfigResponse, error) {
	var response DetailsWithConfigResponse

	err := s.do(ctx, &response)

	return response, err
}

func (s *ConnectionSetupTestsService) DoCustom(ctx context.Context) (DetailsWithCustomConfigResponse, error) {
	var response DetailsWithCustomConfigResponse

	err := s.do(ctx, &response)

	return response, err
}

func (s *ConnectionSetupTestsService) DoCustomMerged(ctx context.Context) (DetailsWithCustomMergedConfigResponse, error) {
	var response DetailsWithCustomMergedConfigResponse

	err := s.do(ctx, &response)

	return response, err
}

func (s *ConnectionSetupTestsService) do(ctx context.Context, response any) error {
	if s.connectionID == nil {
		return fmt.Errorf("missing required ConnectionID")
	}
	url := fmt.Sprintf("/connections/%v/test", *s.connectionID)
	err := s.HttpService.Do(ctx, "POST", url, s.request(), nil, 200, &response)
	return err
}