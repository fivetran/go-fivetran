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
    region               *string
    timeZoneOffset       *string
    config               *ExternalLoggingConfig
    trustCertificates    *bool
    trustFingerprints    *bool
    runSetupTests        *bool
}

type externalLoggingModifyRequest struct {
    Region            *string                       `json:"region,omitempty"`
    TimeZoneOffset    *string                       `json:"time_zone_offset,omitempty"`
    Config            *externalLoggingConfigRequest `json:"config,omitempty"`
    TrustCertificates *bool                         `json:"trust_certificates,omitempty"`
    TrustFingerprints *bool                         `json:"trust_fingerprints,omitempty"`
    RunSetupTests     *bool                         `json:"run_setup_tests,omitempty"`
}

type ExternalLoggingModifyResponse struct {
    Code    string `json:"code"`
    Message string `json:"message"`
    Data    struct {
        Id             string `json:"id"`
        GroupId        string `json:"group_id"`
        Service        string `json:"service"`
        Region         string `json:"region"`
        TimeZoneOffset string `json:"time_zone_offset"`
        SetupStatus    string `json:"setup_status"`
        SetupTests     []struct {
            Title   string `json:"title"`
            Status  string `json:"status"`
            Message string `json:"message"`
        } `json:"setup_tests"`
        Config ExternalLoggingConfigResponse `json:"config"`
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
        Region:            s.region,
        TimeZoneOffset:    s.timeZoneOffset,
        Config:            config,
        TrustCertificates: s.trustCertificates,
        TrustFingerprints: s.trustFingerprints,
        RunSetupTests:     s.runSetupTests,
    }
}

func (s *ExternalLoggingModifyService) ExternalLoggingID(value string) *ExternalLoggingModifyService {
    s.externalLoggingID = &value
    return s
}

func (s *ExternalLoggingModifyService) Region(value string) *ExternalLoggingModifyService {
    s.region = &value
    return s
}

func (s *ExternalLoggingModifyService) TimeZoneOffset(value string) *ExternalLoggingModifyService {
    s.timeZoneOffset = &value
    return s
}

func (s *ExternalLoggingModifyService) Config(value *ExternalLoggingConfig) *ExternalLoggingModifyService {
    s.config = value
    return s
}

func (s *ExternalLoggingModifyService) TrustCertificates(value bool) *ExternalLoggingModifyService {
    s.trustCertificates = &value
    return s
}

func (s *ExternalLoggingModifyService) TrustFingerprints(value bool) *ExternalLoggingModifyService {
    s.trustFingerprints = &value
    return s
}

func (s *ExternalLoggingModifyService) RunSetupTests(value bool) *ExternalLoggingModifyService {
    s.runSetupTests = &value
    return s
}

func (s *ExternalLoggingModifyService) Do(ctx context.Context) (ExternalLoggingModifyResponse, error) {
    var response ExternalLoggingModifyResponse

    if s.externalLoggingID == nil {
        return response, fmt.Errorf("missing required ExternalLoggingID")
    }

    url := fmt.Sprintf("%v/external-logging/%v", s.c.baseURL, *s.externalLoggingID)
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
