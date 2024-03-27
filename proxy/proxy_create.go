package proxy

import (
	"context"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// ProxyCreateService implements the Proxy Agent Management, Create a Proxy Agent.
// Ref. https://fivetran.com/docs/rest-api/proxy-management#createaproxyagent
type ProxyCreateService struct {
	httputils.HttpService
	display_name     *string
	group_id		 *string
}

func (s *ProxyCreateService) request() *proxyCreateRequest {
	return &proxyCreateRequest{
		DisplayName: s.display_name,
		GroupId: 	 s.group_id,
	}
}

func (s *ProxyCreateService) DisplayName(value string) *ProxyCreateService {
	s.display_name = &value
	return s
}

func (s *ProxyCreateService) GroupId(value string) *ProxyCreateService {
	s.group_id = &value
	return s
}

func (s *ProxyCreateService) Do(ctx context.Context) (ProxyCreateResponse, error) {
	var response ProxyCreateResponse
	err := s.HttpService.Do(ctx, "POST", "/proxy", s.request(), nil, 201, &response)
	return response, err
}