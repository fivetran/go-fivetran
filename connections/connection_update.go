package connections

import (
    "context"
    "fmt"

    "github.com/fivetran/go-fivetran/utils"
    httputils "github.com/fivetran/go-fivetran/http_utils"
)

type ConnectionUpdateService struct {
    httputils.HttpService
    connectionID              *string
    paused                   *bool
    syncFrequency            *int
    dailySyncTime            *string
    trustCertificates        *bool
    trustFingerprints        *bool
    isHistoricalSync         *bool
    scheduleType             *string
    runSetupTests            *bool
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

func (s *ConnectionUpdateService) requestBase() connectionUpdateRequestBase {
    return connectionUpdateRequestBase{
        Paused:                     s.paused,
        SyncFrequency:              s.syncFrequency,
        DailySyncTime:              s.dailySyncTime,
        TrustCertificates:          s.trustCertificates,
        TrustFingerprints:          s.trustFingerprints,
        IsHistoricalSync:           s.isHistoricalSync,
        ScheduleType:               s.scheduleType,
        RunSetupTests:              s.runSetupTests,
        PauseAfterTrial:            s.pauseAfterTrial,
        PrivateLinkId:              s.privateLinkId,
        HybridDeploymentAgentId:    s.hybridDeploymentAgentId,
        NetworkingMethod:           s.networkingMethod,
        ProxyAgentId:               s.proxyAgentId,
        DataDelaySensitivity:       s.dataDelaySensitivity,
        DataDelayThreshold:         s.dataDelayThreshold,
    }
}

func (s *ConnectionUpdateService) request() *connectionUpdateRequest {
    var config interface{}
    if s.config != nil {
        config = s.config.Request()
    }

    var auth interface{}
    if s.auth != nil {
        auth = s.auth.Request()
    }

    return &connectionUpdateRequest{
        connectionUpdateRequestBase: s.requestBase(),
        Config:                     config,
        Auth:                       auth,
    }
}

func (s *ConnectionUpdateService) requestCustom() *connectionCustomUpdateRequest {
    return &connectionCustomUpdateRequest{
        connectionUpdateRequestBase: s.requestBase(),
        Config:                     s.configCustom,
        Auth:                       s.authCustom,
    }
}

func (s *ConnectionUpdateService) requestCustomMerged() (*connectionCustomUpdateRequest, error) {
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

    return &connectionCustomUpdateRequest{
        connectionUpdateRequestBase: s.requestBase(),
        Config:                     currentConfig,
        Auth:                       currentAuth,
    }, nil
}

func (s *ConnectionUpdateService) ConnectionID(value string) *ConnectionUpdateService {
    s.connectionID = &value
    return s
}

func (s *ConnectionUpdateService) HybridDeploymentAgentId(value string) *ConnectionUpdateService {
    s.hybridDeploymentAgentId = &value
    return s
}

func (s *ConnectionUpdateService) ProxyAgentId(value string) *ConnectionUpdateService {
    s.proxyAgentId = &value
    return s
}

func (s *ConnectionUpdateService) PrivateLinkId(value string) *ConnectionUpdateService {
    s.privateLinkId = &value
    return s
}

func (s *ConnectionUpdateService) NetworkingMethod(value string) *ConnectionUpdateService {
    s.networkingMethod = &value
    return s
}

func (s *ConnectionUpdateService) Paused(value bool) *ConnectionUpdateService {
    s.paused = &value
    return s
}

func (s *ConnectionUpdateService) SyncFrequency(value *int) *ConnectionUpdateService {
    s.syncFrequency = value
    return s
}

func (s *ConnectionUpdateService) DailySyncTime(value string) *ConnectionUpdateService {
    s.dailySyncTime = &value
    return s
}

func (s *ConnectionUpdateService) Config(value *ConnectionConfig) *ConnectionUpdateService {
    s.config = value
    return s
}

func (s *ConnectionUpdateService) Auth(value *ConnectionAuth) *ConnectionUpdateService {
    s.auth = value
    return s
}

func (s *ConnectionUpdateService) ConfigCustom(value *map[string]interface{}) *ConnectionUpdateService {
    s.configCustom = value
    return s
}

func (s *ConnectionUpdateService) AuthCustom(value *map[string]interface{}) *ConnectionUpdateService {
    s.authCustom = value
    return s
}

func (s *ConnectionUpdateService) TrustCertificates(value bool) *ConnectionUpdateService {
    s.trustCertificates = &value
    return s
}

func (s *ConnectionUpdateService) TrustFingerprints(value bool) *ConnectionUpdateService {
    s.trustFingerprints = &value
    return s
}

func (s *ConnectionUpdateService) IsHistoricalSync(value bool) *ConnectionUpdateService {
    s.isHistoricalSync = &value
    return s
}

func (s *ConnectionUpdateService) ScheduleType(value string) *ConnectionUpdateService {
    s.scheduleType = &value
    return s
}

func (s *ConnectionUpdateService) RunSetupTests(value bool) *ConnectionUpdateService {
    s.runSetupTests = &value
    return s
}

func (s *ConnectionUpdateService) PauseAfterTrial(value bool) *ConnectionUpdateService {
    s.pauseAfterTrial = &value
    return s
}
func (s *ConnectionUpdateService) DataDelayThreshold(value *int) *ConnectionUpdateService {
    s.dataDelayThreshold = value
    return s
}
func (s *ConnectionUpdateService) DataDelaySensitivity(value string) *ConnectionUpdateService {
    s.dataDelaySensitivity = &value
    return s
}
func (s *ConnectionUpdateService) do(ctx context.Context, req, response any) error {
    if s.connectionID == nil {
        return fmt.Errorf("missing required connectionID")
    }
    url := fmt.Sprintf("/connections/%v", *s.connectionID)
    err := s.HttpService.Do(ctx, "PATCH", url, req, nil, 200, &response)
    return err
}

func (s *ConnectionUpdateService) Do(ctx context.Context) (DetailsWithConfigResponse, error) {
    var response DetailsWithConfigResponse

    err := s.do(ctx, s.request(), &response)

    return response, err
}

func (s *ConnectionUpdateService) DoCustom(ctx context.Context) (DetailsWithCustomConfigResponse, error) {
    var response DetailsWithCustomConfigResponse

    err := s.do(ctx, s.requestCustom(), &response)

    return response, err
}

func (s *ConnectionUpdateService) DoCustomMerged(ctx context.Context) (DetailsWithCustomMergedConfigResponse, error) {
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
