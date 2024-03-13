package destinations

import (
	"context"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// DestinationCreateService implements the Destination Management, Create a Destination API.
// Ref. https://fivetran.com/docs/rest-api/destinations#createadestination
type DestinationCreateService struct {
	httputils.HttpService
	groupID           			*string
	service           			*string
	region            			*string
	timeZoneOffset    			*string
	config            			*DestinationConfig
	configCustom      			*map[string]interface{}
	trustCertificates 			*bool
	trustFingerprints 			*bool
	runSetupTests     			*bool
	daylightSavingTimeEnabled   *bool
}

func (s *DestinationCreateService) GroupID(value string) *DestinationCreateService {
	s.groupID = &value
	return s
}

func (s *DestinationCreateService) Service(value string) *DestinationCreateService {
	s.service = &value
	return s
}

func (s *DestinationCreateService) Region(value string) *DestinationCreateService {
	s.region = &value
	return s
}

func (s *DestinationCreateService) TimeZoneOffset(value string) *DestinationCreateService {
	s.timeZoneOffset = &value
	return s
}

func (s *DestinationCreateService) Config(value *DestinationConfig) *DestinationCreateService {
	s.config = value
	return s
}

func (s *DestinationCreateService) ConfigCustom(value *map[string]interface{}) *DestinationCreateService {
	s.configCustom = value
	return s
}

func (s *DestinationCreateService) TrustCertificates(value bool) *DestinationCreateService {
	s.trustCertificates = &value
	return s
}

func (s *DestinationCreateService) TrustFingerprints(value bool) *DestinationCreateService {
	s.trustFingerprints = &value
	return s
}

func (s *DestinationCreateService) RunSetupTests(value bool) *DestinationCreateService {
	s.runSetupTests = &value
	return s
}

func (s *DestinationCreateService) DaylightSavingTimeEnabled(value bool) *DestinationCreateService {
	s.daylightSavingTimeEnabled = &value
	return s
}

func (s *DestinationCreateService) Do(ctx context.Context) (DestinationDetailsWithSetupTestsResponse, error) {
	var response DestinationDetailsWithSetupTestsResponse
	err := s.HttpService.Do(ctx, "POST", "/destinations", s.request(), nil, 201, &response)
	return response, err
}

func (s *DestinationCreateService) DoCustom(ctx context.Context) (DestinationDetailsWithSetupTestsCustomResponse, error) {
	var response DestinationDetailsWithSetupTestsCustomResponse
	err := s.HttpService.Do(ctx, "POST", "/destinations", s.requestCustom(), nil, 201, &response)
	return response, err
}

func (s *DestinationCreateService) request() *destinationCreateRequest {
	var config interface{}
	if s.config != nil {
		config = s.config.Request()
	}

	return &destinationCreateRequest{
		GroupID:           				s.groupID,
		Service:           				s.service,
		Region:            				s.region,
		TimeZoneOffset:    				s.timeZoneOffset,
		Config:           				config,
		TrustCertificates: 				s.trustCertificates,
		TrustFingerprints: 				s.trustFingerprints,
		RunSetupTests:     			    s.runSetupTests,
		DaylightSavingTimeEnabled:      s.daylightSavingTimeEnabled,
	}
}

func (s *DestinationCreateService) requestCustom() *destinationCreateRequest {
	return &destinationCreateRequest{
		GroupID:           				s.groupID,
		Service:           				s.service,
		Region:            				s.region,
		TimeZoneOffset:    				s.timeZoneOffset,
		Config:            				s.configCustom,
		TrustCertificates: 				s.trustCertificates,
		TrustFingerprints: 				s.trustFingerprints,
		RunSetupTests:    				s.runSetupTests,
		DaylightSavingTimeEnabled:      s.daylightSavingTimeEnabled,
	}
}
