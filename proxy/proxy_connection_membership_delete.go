package proxy

import (
	"context"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// ProxyConnectionMembershipDeleteService implements the Proxy Agent Management, Detach connection from the proxy agent
// Ref. https://fivetran.com/docs/rest-api/proxy-management#detachconnectionfromtheproxyagent
type ProxyConnectionMembershipDeleteService struct {
	httputils.HttpService
	proxyId      *string
	connectionId *string
}

func (s *ProxyConnectionMembershipDeleteService) ProxyId(value string) *ProxyConnectionMembershipDeleteService {
	s.proxyId = &value
	return s
}

func (s *ProxyConnectionMembershipDeleteService) ConnectionId(value string) *ProxyConnectionMembershipDeleteService {
	s.connectionId = &value
	return s
}

func (s *ProxyConnectionMembershipDeleteService) Do(ctx context.Context) (common.CommonResponse, error) {
	var response common.CommonResponse
	
	if s.proxyId == nil {
		return response, fmt.Errorf("missing required proxyId")
	}

	if s.connectionId == nil {
		return response, fmt.Errorf("missing required connectionId")
	}

	url := fmt.Sprintf("/proxy/%v/connections/%v", *s.proxyId, *s.connectionId)
	err := s.HttpService.Do(ctx, "DELETE", url, nil, nil, 200, &response)
	return response, err
}