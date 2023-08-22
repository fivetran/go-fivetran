package fivetran

import (
    "context"
    "encoding/json"
    "fmt"
)

// ExternalLoggingModifyService implements the Log Management, Modify a Log Service API.
// Ref. https://fivetran.com/docs/rest-api/log-service-management#updatealogservice
type ExternalLoggingModifyService struct {
    c                    *Client
    externalLoggingId    *string
    enabled              *bool
    config               *ExternalLoggingConfig
}

type externalLoggingModifyRequest struct {
    Enabled           *bool                         `json:"enabled,omitempty"`
    Config            *externalLoggingConfigRequest `json:"config,omitempty"`
}

type ExternalLoggingModifyResponse struct {
    Code    string `json:"code"`
    Message string `json:"message"`
    Data    struct {
        Id             string `json:"id"`
        Service        string `json:"service"`
        Enabled        bool   `json:"enabled"`
    } `json:"data"`
}

func (c *Client) NewExternalLoggingModify() *ExternalLoggingModifyService {
    return &ExternalLoggingModifyService{c: c}
}

func (s *ExternalLoggingModifyService) request() *externalLoggingModifyRequest {
    var config *externalLoggingConfigRequest

    if s.config != nil {
        config = s.config.request()
    }

    return &externalLoggingModifyRequest{
        Enabled:           s.enabled,
        Config:            config,
    }
}

func (s *ExternalLoggingModifyService) ExternalLoggingId(value string) *ExternalLoggingModifyService {
    s.externalLoggingId = &value
    return s
}

func (s *ExternalLoggingModifyService) Enabled(value bool) *ExternalLoggingModifyService {
    s.enabled = &value
    return s
}

func (s *ExternalLoggingModifyService) Config(value *ExternalLoggingConfig) *ExternalLoggingModifyService {
    s.config = value
    return s
}

func (s *ExternalLoggingModifyService) Do(ctx context.Context) (ExternalLoggingModifyResponse, error) {
    var response ExternalLoggingModifyResponse

    if s.externalLoggingId == nil {
        return response, fmt.Errorf("missing required ExternalLoggingID")
    }

    url := fmt.Sprintf("%v/external-logging/%v", s.c.baseURL, *s.externalLoggingId)
    expectedStatus := 200

    headers := s.c.commonHeaders()
    headers["Content-Type"] = "application/json"
    headers["Accept"] = restAPIv2

    reqBody, err := json.Marshal(s.request())
    if err != nil {
        return response, err
    }

    r := request{
        method:  "PATCH",
        url:     url,
        body:    reqBody,
        queries: nil,
        headers: headers,
        client:  s.c.httpClient,
    }

    respBody, respStatus, err := r.httpRequest(ctx)
    if err != nil {
        return response, err
    }

    if err := json.Unmarshal(respBody, &response); err != nil {
        return response, err
    }

    if respStatus != expectedStatus {
        err := fmt.Errorf("status code: %v; expected: %v", respStatus, expectedStatus)
        return response, err
    }

    return response, nil
}
