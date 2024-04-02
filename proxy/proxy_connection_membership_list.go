package proxy

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// ProxyConnectionMembershipsListService implements the Proxy Agent Management, Return all connections attached to the proxy agent
// Ref. https://fivetran.com/docs/rest-api/proxy-management#returnallconnectionsattachedtotheproxyagent
type ProxyConnectionMembershipsListService struct {
	httputils.HttpService
	proxyId *string
	limit  *int
	cursor *string
}

func (s *ProxyConnectionMembershipsListService) ProxyId(value string) *ProxyConnectionMembershipsListService {
	s.proxyId = &value
	return s
}

func (s *ProxyConnectionMembershipsListService) Limit(value int) *ProxyConnectionMembershipsListService {
	s.limit = &value
	return s
}

func (s *ProxyConnectionMembershipsListService) Cursor(value string) *ProxyConnectionMembershipsListService {
	s.cursor = &value
	return s
}

func (s *ProxyConnectionMembershipsListService) Do(ctx context.Context) (ProxyConnectionMembershipsListResponse, error) {
	var response ProxyConnectionMembershipsListResponse

	if s.proxyId == nil {
		return response, fmt.Errorf("missing required proxyId")
	}

	url := fmt.Sprintf("/proxy/%v/connections", *s.proxyId)
	var queries map[string]string = nil
	if s.cursor != nil || s.limit != nil {
		queries = make(map[string]string)
		if s.cursor != nil {
			queries["cursor"] = *s.cursor
		}
		if s.limit != nil {
			queries["limit"] = fmt.Sprintf("%v", *s.limit)
		}
	}
	err := s.HttpService.Do(ctx, "GET", url, nil, queries, 200, &response)
	return response, err
}