package connectors

import (
    "context"
    "fmt"

    "github.com/fivetran/go-fivetran/utils"
    httputils "github.com/fivetran/go-fivetran/http_utils"
)

// ConnectorModifyService implements the Connector Management, Modify a Connector API.
// Ref. https://fivetran.com/docs/rest-api/connectors#modifyaconnector
type ConnectorModifyService struct {
    httputils.HttpService
    connectorID              *string
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
    config                   *ConnectorConfig
    auth                     *ConnectorAuth
    configCustom             *map[string]interface{}
    authCustom               *map[string]interface{}
}

func (s *ConnectorModifyService) requestBase() connectorModifyRequestBase {
    return connectorModifyRequestBase{
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

func (s *ConnectorModifyService) request() *connectorModifyRequest {
    var config interface{}
    if s.config != nil {
        config = s.config.Request()
    }

    var auth interface{}
    if s.auth != nil {
        auth = s.auth.Request()
    }

    return &connectorModifyRequest{
        connectorModifyRequestBase: s.requestBase(),
        Config:                     config,
        Auth:                       auth,
    }
}

func (s *ConnectorModifyService) requestCustom() *connectorCustomModifyRequest {
    return &connectorCustomModifyRequest{
        connectorModifyRequestBase: s.requestBase(),
        Config:                     s.configCustom,
        Auth:                       s.authCustom,
    }
}

func (s *ConnectorModifyService) requestCustomMerged() (*connectorCustomModifyRequest, error) {
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

    return &connectorCustomModifyRequest{
        connectorModifyRequestBase: s.requestBase(),
        Config:                     currentConfig,
        Auth:                       currentAuth,
    }, nil
}

func (s *ConnectorModifyService) ConnectorID(value string) *ConnectorModifyService {
    s.connectorID = &value
    return s
}

func (s *ConnectorModifyService) HybridDeploymentAgentId(value string) *ConnectorModifyService {
    s.hybridDeploymentAgentId = &value
    return s
}

func (s *ConnectorModifyService) ProxyAgentId(value string) *ConnectorModifyService {
    s.proxyAgentId = &value
    return s
}

func (s *ConnectorModifyService) PrivateLinkId(value string) *ConnectorModifyService {
    s.privateLinkId = &value
    return s
}

func (s *ConnectorModifyService) NetworkingMethod(value string) *ConnectorModifyService {
    s.networkingMethod = &value
    return s
}

func (s *ConnectorModifyService) Paused(value bool) *ConnectorModifyService {
    s.paused = &value
    return s
}

func (s *ConnectorModifyService) SyncFrequency(value *int) *ConnectorModifyService {
    s.syncFrequency = value
    return s
}

func (s *ConnectorModifyService) DailySyncTime(value string) *ConnectorModifyService {
    s.dailySyncTime = &value
    return s
}

func (s *ConnectorModifyService) Config(value *ConnectorConfig) *ConnectorModifyService {
    s.config = value
    return s
}

func (s *ConnectorModifyService) Auth(value *ConnectorAuth) *ConnectorModifyService {
    s.auth = value
    return s
}

func (s *ConnectorModifyService) ConfigCustom(value *map[string]interface{}) *ConnectorModifyService {
    s.configCustom = value
    return s
}

func (s *ConnectorModifyService) AuthCustom(value *map[string]interface{}) *ConnectorModifyService {
    s.authCustom = value
    return s
}

func (s *ConnectorModifyService) TrustCertificates(value bool) *ConnectorModifyService {
    s.trustCertificates = &value
    return s
}

func (s *ConnectorModifyService) TrustFingerprints(value bool) *ConnectorModifyService {
    s.trustFingerprints = &value
    return s
}

func (s *ConnectorModifyService) IsHistoricalSync(value bool) *ConnectorModifyService {
    s.isHistoricalSync = &value
    return s
}

func (s *ConnectorModifyService) ScheduleType(value string) *ConnectorModifyService {
    s.scheduleType = &value
    return s
}

func (s *ConnectorModifyService) RunSetupTests(value bool) *ConnectorModifyService {
    s.runSetupTests = &value
    return s
}

func (s *ConnectorModifyService) PauseAfterTrial(value bool) *ConnectorModifyService {
    s.pauseAfterTrial = &value
    return s
}
func (s *ConnectorModifyService) DataDelayThreshold(value *int) *ConnectorModifyService {
    s.dataDelayThreshold = value
    return s
}
func (s *ConnectorModifyService) DataDelaySensitivity(value string) *ConnectorModifyService {
    s.dataDelaySensitivity = &value
    return s
}
func (s *ConnectorModifyService) do(ctx context.Context, req, response any) error {
    if s.connectorID == nil {
        return fmt.Errorf("missing required connectorID")
    }
    url := fmt.Sprintf("/connectors/%v", *s.connectorID)
    err := s.HttpService.Do(ctx, "PATCH", url, req, nil, 200, &response)
    return err
}

func (s *ConnectorModifyService) Do(ctx context.Context) (DetailsWithConfigResponse, error) {
    var response DetailsWithConfigResponse

    err := s.do(ctx, s.request(), &response)

    return response, err
}

func (s *ConnectorModifyService) DoCustom(ctx context.Context) (DetailsWithCustomConfigResponse, error) {
    var response DetailsWithCustomConfigResponse

    err := s.do(ctx, s.requestCustom(), &response)

    return response, err
}

func (s *ConnectorModifyService) DoCustomMerged(ctx context.Context) (DetailsWithCustomMergedConfigResponse, error) {
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
