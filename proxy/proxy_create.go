package proxy

import (
    "context"

    httputils "github.com/fivetran/go-fivetran/http_utils"
)

type ProxyCreateService struct {
    httputils.HttpService
    display_name     *string
    group_region     *string
}

func (s *ProxyCreateService) request() *proxyCreateRequest {
    return &proxyCreateRequest{
        DisplayName: s.display_name,
        GroupRegion: s.group_region,
    }
}

func (s *ProxyCreateService) DisplayName(value string) *ProxyCreateService {
    s.display_name = &value
    return s
}

func (s *ProxyCreateService) GroupRegion(value string) *ProxyCreateService {
    s.group_region = &value
    return s
}

func (s *ProxyCreateService) Do(ctx context.Context) (ProxyCreateResponse, error) {
    var response ProxyCreateResponse
    err := s.HttpService.Do(ctx, "POST", "/proxy", s.request(), nil, 201, &response)
    return response, err
}