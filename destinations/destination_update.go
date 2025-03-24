package destinations

import (
    "context"
    "fmt"

    httputils "github.com/fivetran/go-fivetran/http_utils"
)

type DestinationUpdateService struct {
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

func (s *DestinationUpdateService) request() *destinationUpdateRequest {
    var config interface{}

    if s.config != nil {
        config = s.config.Request()
    }

    return &destinationUpdateRequest{
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

func (s *DestinationUpdateService) requestCustom() *destinationUpdateRequest {
    return &destinationUpdateRequest{
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

func (s *DestinationUpdateService) DestinationID(value string) *DestinationUpdateService {
    s.destinationID = &value
    return s
}

func (s *DestinationUpdateService) Region(value string) *DestinationUpdateService {
    s.region = &value
    return s
}

func (s *DestinationUpdateService) TimeZoneOffset(value string) *DestinationUpdateService {
    s.timeZoneOffset = &value
    return s
}

func (s *DestinationUpdateService) Config(value *DestinationConfig) *DestinationUpdateService {
    s.config = value
    return s
}

func (s *DestinationUpdateService) ConfigCustom(value *map[string]interface{}) *DestinationUpdateService {
    s.configCustom = value
    return s
}

func (s *DestinationUpdateService) TrustCertificates(value bool) *DestinationUpdateService {
    s.trustCertificates = &value
    return s
}

func (s *DestinationUpdateService) TrustFingerprints(value bool) *DestinationUpdateService {
    s.trustFingerprints = &value
    return s
}

func (s *DestinationUpdateService) RunSetupTests(value bool) *DestinationUpdateService {
    s.runSetupTests = &value
    return s
}

func (s *DestinationUpdateService) DaylightSavingTimeEnabled(value bool) *DestinationUpdateService {
    s.daylightSavingTimeEnabled = &value
    return s
}

func (s *DestinationUpdateService) HybridDeploymentAgentId(value string) *DestinationUpdateService {
    s.hybridDeploymentAgentId = &value
    return s
}

func (s *DestinationUpdateService) PrivateLinkId(value string) *DestinationUpdateService {
    s.privateLinkId = &value
    return s
}

func (s *DestinationUpdateService) NetworkingMethod(value string) *DestinationUpdateService {
    s.networkingMethod = &value
    return s
}

func (s *DestinationUpdateService) Do(ctx context.Context) (DestinationDetailsWithSetupTestsResponse, error) {
    var response DestinationDetailsWithSetupTestsResponse

    if s.destinationID == nil {
        return response, fmt.Errorf("missing required DestinationID")
    }

    url := fmt.Sprintf("/destinations/%v", *s.destinationID)
    err := s.HttpService.Do(ctx, "PATCH", url, s.request(), nil, 200, &response)
    return response, err
}

func (s *DestinationUpdateService) DoCustom(ctx context.Context) (DestinationDetailsWithSetupTestsCustomResponse, error) {
    var response DestinationDetailsWithSetupTestsCustomResponse

    if s.destinationID == nil {
        return response, fmt.Errorf("missing required DestinationID")
    }

    url := fmt.Sprintf("/destinations/%v", *s.destinationID)
    err := s.HttpService.Do(ctx, "PATCH", url, s.requestCustom(), nil, 200, &response)
    return response, err
}
