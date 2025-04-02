package connections

import (
    "context"

    "github.com/fivetran/go-fivetran/utils"
    httputils "github.com/fivetran/go-fivetran/http_utils"
)

type ConnectionCreateService struct {
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
    config                   *ConnectionConfig
    auth                     *ConnectionAuth
    configCustom             *map[string]interface{}
    authCustom               *map[string]interface{}
}

func (s *ConnectionCreateService) requestBase() connectionCreateRequestBase {
    return connectionCreateRequestBase{
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

func (s *ConnectionCreateService) request() *connectionCreateRequest {
    var config interface{}
    if s.config != nil {
        config = s.config.Request()
    }

    var auth interface{}
    if s.auth != nil {
        auth = s.auth.Request()
    }

    r := &connectionCreateRequest{
        connectionCreateRequestBase: s.requestBase(),
        Config:                     config,
        Auth:                       auth,
    }

    return r
}

func (s *ConnectionCreateService) requestCustom() *connectionCustomCreateRequest {
    return &connectionCustomCreateRequest{
        connectionCreateRequestBase: s.requestBase(),
        Config:                     s.configCustom,
        Auth:                       s.authCustom,
    }
}

func (s *ConnectionCreateService) requestCustomMerged() (*connectionCustomCreateRequest, error) {
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

    return &connectionCustomCreateRequest{
        connectionCreateRequestBase: s.requestBase(),
        Config:                     currentConfig,
        Auth:                       currentAuth,
    }, nil
}

func (s *ConnectionCreateService) Service(value string) *ConnectionCreateService {
    s.service = &value
    return s
}

func (s *ConnectionCreateService) GroupID(value string) *ConnectionCreateService {
    s.groupID = &value
    return s
}

func (s *ConnectionCreateService) HybridDeploymentAgentId(value string) *ConnectionCreateService {
    s.hybridDeploymentAgentId = &value
    return s
}

func (s *ConnectionCreateService) ProxyAgentId(value string) *ConnectionCreateService {
    s.proxyAgentId = &value
    return s
}

func (s *ConnectionCreateService) PrivateLinkId(value string) *ConnectionCreateService {
    s.privateLinkId = &value
    return s
}

func (s *ConnectionCreateService) NetworkingMethod(value string) *ConnectionCreateService {
    s.networkingMethod = &value
    return s
}

func (s *ConnectionCreateService) TrustCertificates(value bool) *ConnectionCreateService {
    s.trustCertificates = &value
    return s
}

func (s *ConnectionCreateService) TrustFingerprints(value bool) *ConnectionCreateService {
    s.trustFingerprints = &value
    return s
}

func (s *ConnectionCreateService) RunSetupTests(value bool) *ConnectionCreateService {
    s.runSetupTests = &value
    return s
}

func (s *ConnectionCreateService) Paused(value bool) *ConnectionCreateService {
    s.paused = &value
    return s
}

func (s *ConnectionCreateService) SyncFrequency(value *int) *ConnectionCreateService {
    s.syncFrequency = value
    return s
}

func (s *ConnectionCreateService) DailySyncTime(value string) *ConnectionCreateService {
    s.dailySyncTime = &value
    return s
}

func (s *ConnectionCreateService) PauseAfterTrial(value bool) *ConnectionCreateService {
    s.pauseAfterTrial = &value
    return s
}

func (s *ConnectionCreateService) Config(value *ConnectionConfig) *ConnectionCreateService {
    s.config = value
    return s
}

func (s *ConnectionCreateService) ConfigCustom(value *map[string]interface{}) *ConnectionCreateService {
    s.configCustom = value
    return s
}

func (s *ConnectionCreateService) Auth(value *ConnectionAuth) *ConnectionCreateService {
    s.auth = value
    return s
}

func (s *ConnectionCreateService) AuthCustom(value *map[string]interface{}) *ConnectionCreateService {
    s.authCustom = value
    return s
}
func (s *ConnectionCreateService) DataDelayThreshold(value *int) *ConnectionCreateService {
    s.dataDelayThreshold = value
    return s
}
func (s *ConnectionCreateService) DataDelaySensitivity(value string) *ConnectionCreateService {
    s.dataDelaySensitivity = &value
    return s
}
func (s *ConnectionCreateService) do(ctx context.Context, req, response any) error {
    err := s.HttpService.Do(ctx, "POST", "/connections", req, nil, 201, &response)
    return err
}

func (s *ConnectionCreateService) Do(ctx context.Context) (DetailsWithConfigResponse, error) {
    var response DetailsWithConfigResponse

    err := s.do(ctx, s.request(), &response)

    return response, err
}

func (s *ConnectionCreateService) DoCustom(ctx context.Context) (DetailsWithCustomConfigResponse, error) {
    var response DetailsWithCustomConfigResponse

    err := s.do(ctx, s.requestCustom(), &response)

    return response, err
}

func (s *ConnectionCreateService) DoCustomMerged(ctx context.Context) (DetailsWithCustomMergedConfigResponse, error) {
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
