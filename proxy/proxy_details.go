package proxy

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// ProxyDetailsService implements the Proxy Agent Management, Retrieve Proxy Agent Details.
// Ref. https://fivetran.com/docs/rest-api/proxy-management#retrieveproxyagentdetails
type ProxyDetailsService struct {
	httputils.HttpService
	proxyId 	*string
}

func (s *ProxyDetailsService) ProxyId(value string) *ProxyDetailsService {
	s.proxyId = &value
	return s
}

func (s *ProxyDetailsService) Do(ctx context.Context) (ProxyDetailsResponse, error) {
	var response ProxyDetailsResponse

	if s.proxyId == nil {
		return response, fmt.Errorf("missing required proxyId")
	}

	url := fmt.Sprintf("/proxy/%v", *s.proxyId)
	err := s.HttpService.Do(ctx, "GET", url, nil, nil, 200, &response)
	return response, err
}