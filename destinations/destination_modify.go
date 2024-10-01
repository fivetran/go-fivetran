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
    destinationID               *string
    region                      *string
    timeZoneOffset              *string
    config                      *DestinationConfig
    configCustom                *map[string]interface{}
    trustCertificates           *bool
    trustFingerprints           *bool
    runSetupTests               *bool
    daylightSavingTimeEnabled   *bool
    hybridDeploymentAgentId     *string
    networkingMethod            *string
    privateLinkId               *string
}

func (s *DestinationModifyService) request() *destinationModifyRequest {
    var config interface{}

    if s.config != nil {
        config = s.config.Request()
    }

    return &destinationModifyRequest{
        Region:                         s.region,
        TimeZoneOffset:                 s.timeZoneOffset,
        Config:                         config,
        TrustCertificates:              s.trustCertificates,
        TrustFingerprints:              s.trustFingerprints,
        RunSetupTests:                  s.runSetupTests,
        DaylightSavingTimeEnabled:      s.daylightSavingTimeEnabled,
        PrivateLinkId:                  s.privateLinkId,
        HybridDeploymentAgentId:        s.hybridDeploymentAgentId,
        NetworkingMethod:               s.networkingMethod,
    }
}

func (s *DestinationModifyService) requestCustom() *destinationModifyRequest {
    return &destinationModifyRequest{
        Region:                         s.region,
        TimeZoneOffset:                 s.timeZoneOffset,
        Config:                         s.configCustom,
        TrustCertificates:              s.trustCertificates,
        TrustFingerprints:              s.trustFingerprints,
        RunSetupTests:                  s.runSetupTests,
        DaylightSavingTimeEnabled:      s.daylightSavingTimeEnabled,
        PrivateLinkId:                  s.privateLinkId,
        HybridDeploymentAgentId:        s.hybridDeploymentAgentId,
        NetworkingMethod:               s.networkingMethod,
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

func (s *DestinationModifyService) ConfigCustom(value *map[string]interface{}) *DestinationModifyService {
    s.configCustom = value
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

func (s *DestinationModifyService) DaylightSavingTimeEnabled(value bool) *DestinationModifyService {
    s.daylightSavingTimeEnabled = &value
    return s
}

func (s *DestinationModifyService) HybridDeploymentAgentId(value string) *DestinationModifyService {
    s.hybridDeploymentAgentId = &value
    return s
}

func (s *DestinationModifyService) PrivateLinkId(value string) *DestinationModifyService {
    s.privateLinkId = &value
    return s
}

func (s *DestinationModifyService) NetworkingMethod(value string) *DestinationModifyService {
    s.networkingMethod = &value
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

func (s *DestinationModifyService) DoCustom(ctx context.Context) (DestinationDetailsWithSetupTestsCustomResponse, error) {
    var response DestinationDetailsWithSetupTestsCustomResponse

    if s.destinationID == nil {
        return response, fmt.Errorf("missing required DestinationID")
    }

    url := fmt.Sprintf("/destinations/%v", *s.destinationID)
    err := s.HttpService.Do(ctx, "PATCH", url, s.requestCustom(), nil, 200, &response)
    return response, err
}
