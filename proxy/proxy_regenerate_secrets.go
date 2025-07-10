package proxy

import (
    "context"
    "fmt"

    httputils "github.com/fivetran/go-fivetran/http_utils"
)

type ProxyRegenerateSecretsService struct {
    httputils.HttpService
    proxyId       *string
}

func (s *ProxyRegenerateSecretsService) ProxyId(value string) *ProxyRegenerateSecretsService {
    s.proxyId = &value
    return s
}

func (s *ProxyRegenerateSecretsService) Do(ctx context.Context) (ProxyCreateResponse, error) {
    var response ProxyCreateResponse
    
    if s.proxyId == nil {
        return response, fmt.Errorf("missing required proxyId")
    }

    url := fmt.Sprintf("/proxy/%v/regenerate-secrets", *s.proxyId)
    err := s.HttpService.Do(ctx, "POST", url, nil, nil, 200, &response)
    return response, err
}