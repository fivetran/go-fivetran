package destinations

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// DestinationSetupTestsService implements the Destination Management, Run destination setup tests API.
// Ref. https://fivetran.com/docs/rest-api/destinations#rundestinationsetuptests
type DestinationSetupTestsService struct {
	httputils.HttpService
	destinationID     *string
	trustCertificates *bool
	trustFingerprints *bool
}

func (s *DestinationSetupTestsService) request() *destinationSetupTestsRequest {
	return &destinationSetupTestsRequest{
		TrustCertificates: s.trustCertificates,
		TrustFingerprints: s.trustFingerprints,
	}
}

func (s *DestinationSetupTestsService) DestinationID(value string) *DestinationSetupTestsService {
	s.destinationID = &value
	return s
}

func (s *DestinationSetupTestsService) TrustCertificates(value bool) *DestinationSetupTestsService {
	s.trustCertificates = &value
	return s
}

func (s *DestinationSetupTestsService) TrustFingerprints(value bool) *DestinationSetupTestsService {
	s.trustFingerprints = &value
	return s
}

func (s *DestinationSetupTestsService) Do(ctx context.Context) (DestinationDetailsWithSetupTestsResponse, error) {
	var response DestinationDetailsWithSetupTestsResponse

	if s.destinationID == nil {
		return response, fmt.Errorf("missing required destinationID")
	}

	url := fmt.Sprintf("/destinations/%v/test", *s.destinationID)
	err := s.HttpService.Do(ctx, "POST", url, s.request(), nil, 200, &response)
	return response, err
}