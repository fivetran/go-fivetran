package privatelinks

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// PrivateLinksModifyService implements the Private Links Management, Modify a Private Link Service API.
// Ref. https://fivetran.com/docs/rest-api/private-links-management#updateaprivatelink
type PrivateLinksModifyService struct {
	httputils.HttpService
	privateLinkId 	  *string
	config            *PrivateLinksConfig
}

func (s *PrivateLinksModifyService) request() *privateLinksModifyRequest {
	var config interface{}

	if s.config != nil {
		config = s.config.Request()
	}

	return &privateLinksModifyRequest{
		Config: 	config,
	}
}

func (s *PrivateLinksModifyService) PrivateLinkId(value string) *PrivateLinksModifyService {
	s.privateLinkId = &value
	return s
}

func (s *PrivateLinksModifyService) Config(value *PrivateLinksConfig) *PrivateLinksModifyService {
	s.config = value
	return s
}

func (s *PrivateLinksModifyService) Do(ctx context.Context) (PrivateLinksResponse, error) {
	var response PrivateLinksResponse
	if s.privateLinkId == nil {
		return response, fmt.Errorf("missing required privateLinkId")
	}

	url := fmt.Sprintf("/private-links/%v", *s.privateLinkId)
	err := s.HttpService.Do(ctx, "PATCH", url, s.request(), nil, 200, &response)
	return response, err
}