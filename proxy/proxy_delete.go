package proxy

import (
	"context"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// ProxyDeleteService implements the Proxy Agent Management, Delete a Proxy Agent.
// Ref. https://fivetran.com/docs/rest-api/proxy-management#deleteaproxyagent
type ProxyDeleteService struct {
	httputils.HttpService
	proxyId       *string
}

func (s *ProxyDeleteService) ProxyId(value string) *ProxyDeleteService {
	s.proxyId = &value
	return s
}

func (s *ProxyDeleteService) Do(ctx context.Context) (common.CommonResponse, error) {
	var response common.CommonResponse
	
	if s.proxyId == nil {
		return response, fmt.Errorf("missing required proxyId")
	}

	url := fmt.Sprintf("/proxy/%v", *s.proxyId)
	err := s.HttpService.Do(ctx, "DELETE", url, nil, nil, 200, &response)
	return response, err
}