package connectors

import (
    "context"

    "github.com/fivetran/go-fivetran/utils"
    httputils "github.com/fivetran/go-fivetran/http_utils"
)

// ConnectorCreateService implements the Connector Management, Create a Connector API.
// Ref. https://fivetran.com/docs/rest-api/connectors#createaconnector

type ConnectorCreateService struct {
    httputils.HttpService
    service                  *string
    groupID                  *string
    trustCertificates        *bool
    trustFingerprints        *bool
    runSetupTests            *bool
    paused                   *bool
    syncFrequency            *int
    dailySyncTime            *string
    pauseAfterTrial          *bool
    hybridDeploymentAgentId  *string
    networkingMethod         *string
    privateLinkId            *string
    proxyAgentId             *string
    dataDelaySensitivity     *string
    dataDelayThreshold       *int
    config                   *ConnectorConfig
    auth                     *ConnectorAuth
    configCustom             *map[string]interface{}
    authCustom               *map[string]interface{}
}

func (s *ConnectorCreateService) requestBase() connectorCreateRequestBase {
    return connectorCreateRequestBase{
        Service:                    s.service,
        GroupID:                    s.groupID,
        TrustCertificates:          s.trustCertificates,
        TrustFingerprints:          s.trustFingerprints,
        RunSetupTests:              s.runSetupTests,
        Paused:                     s.paused,
        SyncFrequency:              s.syncFrequency,
        DailySyncTime:              s.dailySyncTime,
        PauseAfterTrial:            s.pauseAfterTrial,
        PrivateLinkId:              s.privateLinkId,
        HybridDeploymentAgentId:    s.hybridDeploymentAgentId,
        NetworkingMethod:           s.networkingMethod,
        ProxyAgentId:               s.proxyAgentId,
        DataDelaySensitivity:       s.dataDelaySensitivity,
        DataDelayThreshold:         s.dataDelayThreshold,
    }
}

func (s *ConnectorCreateService) request() *connectorCreateRequest {
    var config interface{}
    if s.config != nil {
        config = s.config.Request()
    }

    var auth interface{}
    if s.auth != nil {
        auth = s.auth.Request()
    }

    r := &connectorCreateRequest{
        connectorCreateRequestBase: s.requestBase(),
        Config:                     config,
        Auth:                       auth,
    }

    return r
}

func (s *ConnectorCreateService) requestCustom() *connectorCustomCreateRequest {
    return &connectorCustomCreateRequest{
        connectorCreateRequestBase: s.requestBase(),
        Config:                     s.configCustom,
        Auth:                       s.authCustom,
    }
}

func (s *ConnectorCreateService) requestCustomMerged() (*connectorCustomCreateRequest, error) {
    currentConfig := s.configCustom

    if s.config != nil {
        var err error
        currentConfig, err = s.config.Merge(currentConfig)
        if err != nil {
            return nil, err
        }
    }

    currentAuth := s.authCustom
    if s.auth != nil {
        var err error
        currentAuth, err = s.auth.Merge(currentAuth)
        if err != nil {
            return nil, err
        }
    }

    return &connectorCustomCreateRequest{
        connectorCreateRequestBase: s.requestBase(),
        Config:                     currentConfig,
        Auth:                       currentAuth,
    }, nil
}

func (s *ConnectorCreateService) Service(value string) *ConnectorCreateService {
    s.service = &value
    return s
}

func (s *ConnectorCreateService) GroupID(value string) *ConnectorCreateService {
    s.groupID = &value
    return s
}

func (s *ConnectorCreateService) HybridDeploymentAgentId(value string) *ConnectorCreateService {
    s.hybridDeploymentAgentId = &value
    return s
}

func (s *ConnectorCreateService) ProxyAgentId(value string) *ConnectorCreateService {
    s.proxyAgentId = &value
    return s
}

func (s *ConnectorCreateService) PrivateLinkId(value string) *ConnectorCreateService {
    s.privateLinkId = &value
    return s
}

func (s *ConnectorCreateService) NetworkingMethod(value string) *ConnectorCreateService {
    s.networkingMethod = &value
    return s
}

func (s *ConnectorCreateService) TrustCertificates(value bool) *ConnectorCreateService {
    s.trustCertificates = &value
    return s
}

func (s *ConnectorCreateService) TrustFingerprints(value bool) *ConnectorCreateService {
    s.trustFingerprints = &value
    return s
}

func (s *ConnectorCreateService) RunSetupTests(value bool) *ConnectorCreateService {
    s.runSetupTests = &value
    return s
}

func (s *ConnectorCreateService) Paused(value bool) *ConnectorCreateService {
    s.paused = &value
    return s
}

func (s *ConnectorCreateService) SyncFrequency(value *int) *ConnectorCreateService {
    s.syncFrequency = value
    return s
}

func (s *ConnectorCreateService) DailySyncTime(value string) *ConnectorCreateService {
    s.dailySyncTime = &value
    return s
}

func (s *ConnectorCreateService) PauseAfterTrial(value bool) *ConnectorCreateService {
    s.pauseAfterTrial = &value
    return s
}

func (s *ConnectorCreateService) Config(value *ConnectorConfig) *ConnectorCreateService {
    s.config = value
    return s
}

func (s *ConnectorCreateService) ConfigCustom(value *map[string]interface{}) *ConnectorCreateService {
    s.configCustom = value
    return s
}

func (s *ConnectorCreateService) Auth(value *ConnectorAuth) *ConnectorCreateService {
    s.auth = value
    return s
}

func (s *ConnectorCreateService) AuthCustom(value *map[string]interface{}) *ConnectorCreateService {
    s.authCustom = value
    return s
}
func (s *ConnectorCreateService) DataDelayThreshold(value *int) *ConnectorCreateService {
    s.dataDelayThreshold = value
    return s
}
func (s *ConnectorCreateService) DataDelaySensitivity(value string) *ConnectorCreateService {
    s.dataDelaySensitivity = &value
    return s
}
func (s *ConnectorCreateService) do(ctx context.Context, req, response any) error {
    err := s.HttpService.Do(ctx, "POST", "/connectors", req, nil, 201, &response)
    return err
}

func (s *ConnectorCreateService) Do(ctx context.Context) (DetailsWithConfigResponse, error) {
    var response DetailsWithConfigResponse

    err := s.do(ctx, s.request(), &response)

    return response, err
}

func (s *ConnectorCreateService) DoCustom(ctx context.Context) (DetailsWithCustomConfigResponse, error) {
    var response DetailsWithCustomConfigResponse

    err := s.do(ctx, s.requestCustom(), &response)

    return response, err
}

func (s *ConnectorCreateService) DoCustomMerged(ctx context.Context) (DetailsWithCustomMergedConfigResponse, error) {
    var response DetailsWithCustomMergedConfigResponse

    req, err := s.requestCustomMerged()

    if err != nil {
        return response, err
    }

    err = s.do(ctx, req, &response)

    if err == nil {
        err = utils.FetchFromMap(&response.Data.CustomConfig, &response.Data.Config)
    }

    return response, err
}
