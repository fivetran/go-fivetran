package fivetran

import (
    "context"
    "encoding/json"
    "fmt"
)

// ExternalLoggingDeleteService implements the Log Management, Delete a Log Service API.
// Ref. https://fivetran.com/docs/rest-api/log-service-management#deletealogservice
type ExternalLoggingDeleteService struct {
    c                   *Client
    externalLoggingId   *string
}

type ExternalLoggingDeleteResponse struct {
    Code    string `json:"code"`
    Message string `json:"message"`
}

func (c *Client) NewExternalLoggingDelete() *ExternalLoggingDeleteService {
    return &ExternalLoggingDeleteService{c: c}
}

func (s *ExternalLoggingDeleteService) ExternalLoggingId(value string) *ExternalLoggingDeleteService {
    s.externalLoggingId = &value
    return s
}

func (s *ExternalLoggingDeleteService) Do(ctx context.Context) (ExternalLoggingDeleteResponse, error) {
    var response ExternalLoggingDeleteResponse

    if s.externalLoggingId == nil {
        return response, fmt.Errorf("missing required ExternalLoggingId")
    }

    url := fmt.Sprintf("%v/external-logging/%v", s.c.baseURL, *s.externalLoggingId)
    expectedStatus := 200

    headers := s.c.commonHeaders()

    r := request{
        method:  "DELETE",
        url:     url,
        body:    nil,
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
