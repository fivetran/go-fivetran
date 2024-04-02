package proxy

import (
	"context"
	"fmt"
	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// ProxyConnectionMembershipCreateService implements the Proxy Agent Management, Attach connection to the proxy agent
// Ref. https://fivetran.com/docs/rest-api/proxy-management#attachconnectiontotheproxyagent
type ProxyConnectionMembershipCreateService struct {
	httputils.HttpService
	proxyId      	*string
	connectionId    *string
}

func (s *ProxyConnectionMembershipCreateService) request() *proxyConnectionMembershipCreateRequest {
	return &proxyConnectionMembershipCreateRequest{
		ConnectionId: 	s.connectionId,
	}
}

func (s *ProxyConnectionMembershipCreateService) ProxyId(value string) *ProxyConnectionMembershipCreateService {
	s.proxyId = &value
	return s
}

func (s *ProxyConnectionMembershipCreateService) ConnectionId(value string) *ProxyConnectionMembershipCreateService {
	s.connectionId = &value
	return s
}

func (s *ProxyConnectionMembershipCreateService) Do(ctx context.Context) (common.CommonResponse, error) {
	var response common.CommonResponse
	if s.proxyId == nil {
		return response, fmt.Errorf("missing required proxyId")
	}
	if s.connectionId == nil {
		return response, fmt.Errorf("missing required connectionId")
	}
	url := fmt.Sprintf("/proxy/%v/connections", *s.proxyId)
	err := s.HttpService.Do(ctx, "POST", url, s.request(), nil, 200, &response)
	return response, err
}
