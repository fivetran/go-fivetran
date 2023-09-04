package fivetran

import (
    "context"
    "encoding/json"
    "fmt"
)

// ExternalLoggingCreateService implements the Log Management, Create a Log Service API.
// Ref. https://fivetran.com/docs/rest-api/log-service-management#createalogservice
type ExternalLoggingCreateService struct {
    c                 *Client
    groupId           *string
    service           *string
    enabled           *bool
    config            *ExternalLoggingConfig
    configCustom      *map[string]interface{}
}

/* requests */
type externalLoggingCreateRequestBase struct {
    Id                *string                          `json:"id,omitempty"`
    GroupId           *string                          `json:"group_id,omitempty"`
    Service           *string                          `json:"service,omitempty"`
    Enabled           *bool                            `json:"enabled,omitempty"`   
}

type externalLoggingCreateRequest struct {
    externalLoggingCreateRequestBase
    Config          *externalLoggingConfigRequest `json:"config,omitempty"`
}

type externalLoggingCustomCreateRequest struct {
    externalLoggingCreateRequestBase
    Config          *map[string]interface{} `json:"config,omitempty"`
}

/* responses */

type ExternalLoggingCreateResponseBase struct {
    Id             string `json:"id"`
    Service        string `json:"service"`
    Enabled        bool   `json:"enabled"`
}

type ExternalLoggingCreateResponse struct {
    Code    string `json:"code"`
    Message string `json:"message"`
    Data    struct {
        ExternalLoggingCreateResponseBase
        Config            ExternalLoggingConfigResponse     `json:"config"`
    } `json:"data"`
}

type ExternalLoggingCustomCreateResponse struct {
    Code    string `json:"code"`
    Message string `json:"message"`
    Data    struct {
        ExternalLoggingCreateResponseBase
        Config            ExternalLoggingConfigResponse     `json:"config"`
    } `json:"data"`
}

type ExternalLoggingCustomMergedCreateResponse struct {
    Code    string `json:"code"`
    Message string `json:"message"`
    Data    struct {
        ExternalLoggingCreateResponseBase
        CustomConfig      map[string]interface{}         `json:"config"`
        Config            ExternalLoggingConfigResponse // no mapping here
    } `json:"data"`
}

func (c *Client) NewExternalLoggingCreate() *ExternalLoggingCreateService {
    return &ExternalLoggingCreateService{c: c}
}

func (s *ExternalLoggingCreateService) requestBase() externalLoggingCreateRequestBase {
    return externalLoggingCreateRequestBase{
        GroupId:           s.groupId,
        Service:           s.service,
        Enabled:           s.enabled,
    }
}

func (s *ExternalLoggingCreateService) request() *externalLoggingCreateRequest {
    var config *externalLoggingConfigRequest
    if s.config != nil {
        config = s.config.request()
    }

    r := &externalLoggingCreateRequest{
        externalLoggingCreateRequestBase: s.requestBase(),
        Config:                           config,
    }

    return r
}

func (s *ExternalLoggingCreateService) requestCustom() *externalLoggingCustomCreateRequest {
    return &externalLoggingCustomCreateRequest{
        externalLoggingCreateRequestBase: s.requestBase(),
        Config:                           s.configCustom,
    }
}

func (s *ExternalLoggingCreateService) requestCustomMerged() (*externalLoggingCustomCreateRequest, error) {
    currentConfig := s.configCustom

    if s.config != nil {
        var err error
        currentConfig, err = s.config.merge(currentConfig)
        if err != nil {
            return nil, err
        }
    }

    return &externalLoggingCustomCreateRequest{
        externalLoggingCreateRequestBase: s.requestBase(),
        Config:                           currentConfig,
    }, nil
}

func (s *ExternalLoggingCreateService) GroupId(value string) *ExternalLoggingCreateService {
    s.groupId = &value
    return s
}

func (s *ExternalLoggingCreateService) Service(value string) *ExternalLoggingCreateService {
    s.service = &value
    return s
}

func (s *ExternalLoggingCreateService) Enabled(value bool) *ExternalLoggingCreateService {
    s.enabled = &value
    return s
}

func (s *ExternalLoggingCreateService) Config(value *ExternalLoggingConfig) *ExternalLoggingCreateService {
    s.config = value
    return s
}

func (s *ExternalLoggingCreateService) ConfigCustom(value *map[string]interface{}) *ExternalLoggingCreateService {
    s.configCustom = value
    return s
}

func (s *ExternalLoggingCreateService) do(ctx context.Context, req, response any) error {
    url := fmt.Sprintf("%v/external-logging", s.c.baseURL)
    expectedStatus := 201

    headers := s.c.commonHeaders()
    headers["Content-Type"] = "application/json"
    headers["Accept"] = restAPIv2

    reqBody, err := json.Marshal(req)
    if err != nil {
        return err
    }

    r := request{
        method:  "POST",
        url:     url,
        body:    reqBody,
        queries: nil,
        headers: headers,
        client:  s.c.httpClient,
    }

    respBody, respStatus, err := r.httpRequest(ctx)
    if err != nil {
        return err
    }

    if err := json.Unmarshal(respBody, &response); err != nil {
        return err
    }

    if respStatus != expectedStatus {
        err := fmt.Errorf("status code: %v; expected: %v", respStatus, expectedStatus)
        return err
    }

    return nil
}

func (s *ExternalLoggingCreateService) Do(ctx context.Context) (ExternalLoggingCreateResponse, error) {
    var response ExternalLoggingCreateResponse

    err := s.do(ctx, s.request(), &response)

    return response, err
}

func (s *ExternalLoggingCreateService) DoCustom(ctx context.Context) (ExternalLoggingCustomCreateResponse, error) {
    var response ExternalLoggingCustomCreateResponse

    err := s.do(ctx, s.requestCustom(), &response)
    
    return response, err
}

func (s *ExternalLoggingCreateService) DoCustomMerged(ctx context.Context) (ExternalLoggingCustomMergedCreateResponse, error) {
    var response ExternalLoggingCustomMergedCreateResponse

    req, err := s.requestCustomMerged()

    if err != nil {
        return response, err
    }

    err = s.do(ctx, req, &response)

    if err == nil {
        err = FetchFromMap(&response.Data.CustomConfig, &response.Data.Config)
    }

    return response, err
}
