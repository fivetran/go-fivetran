package destinations

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// DestinationModifyService implements the Destination Management, Modify a Destination API.
// Ref. https://fivetran.com/docs/rest-api/destinations#modifyadestination
type DestinationModifyService struct {
	httputils.HttpService
	destinationID     *string
	region            *string
	timeZoneOffset    *string
	config            *DestinationConfig
	trustCertificates *bool
	trustFingerprints *bool
	runSetupTests     *bool
}

func (s *DestinationModifyService) request() *destinationModifyRequest {
	var config interface{}

	if s.config != nil {
		config = s.config.Request()
	}

	return &destinationModifyRequest{
		Region:            s.region,
		TimeZoneOffset:    s.timeZoneOffset,
		Config:            config,
		TrustCertificates: s.trustCertificates,
		TrustFingerprints: s.trustFingerprints,
		RunSetupTests:     s.runSetupTests,
	}
}

func (s *DestinationModifyService) DestinationID(value string) *DestinationModifyService {
	s.destinationID = &value
	return s
}

func (s *DestinationModifyService) Region(value string) *DestinationModifyService {
	s.region = &value
	return s
}

func (s *DestinationModifyService) TimeZoneOffset(value string) *DestinationModifyService {
	s.timeZoneOffset = &value
	return s
}

func (s *DestinationModifyService) Config(value *DestinationConfig) *DestinationModifyService {
	s.config = value
	return s
}

func (s *DestinationModifyService) TrustCertificates(value bool) *DestinationModifyService {
	s.trustCertificates = &value
	return s
}

func (s *DestinationModifyService) TrustFingerprints(value bool) *DestinationModifyService {
	s.trustFingerprints = &value
	return s
}

func (s *DestinationModifyService) RunSetupTests(value bool) *DestinationModifyService {
	s.runSetupTests = &value
	return s
}

func (s *DestinationModifyService) Do(ctx context.Context) (DestinationDetailsWithSetupTestsResponse, error) {
	var response DestinationDetailsWithSetupTestsResponse

	if s.destinationID == nil {
		return response, fmt.Errorf("missing required DestinationID")
	}

	url := fmt.Sprintf("/destinations/%v", *s.destinationID)
	err := s.HttpService.Do(ctx, "PATCH", url, s.request(), nil, 200, &response)
	return response, err
}
