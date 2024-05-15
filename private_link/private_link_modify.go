package privatelink

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// PrivateLinkModifyService implements the Private Link Management, Modify a Private Link Service API.
// Ref. https://fivetran.com/docs/rest-api/private-link-management#updateaprivatelink
type PrivateLinkModifyService struct {
	httputils.HttpService
	privateLinkId 	  *string
	config            *PrivateLinkConfig
}

func (s *PrivateLinkModifyService) request() *privateLinkModifyRequest {
	var config interface{}

	if s.config != nil {
		config = s.config.Request()
	}

	return &privateLinkModifyRequest{
		Config: 	config,
	}
}

func (s *PrivateLinkModifyService) PrivateLinkId(value string) *PrivateLinkModifyService {
	s.privateLinkId = &value
	return s
}

func (s *PrivateLinkModifyService) Config(value *PrivateLinkConfig) *PrivateLinkModifyService {
	s.config = value
	return s
}

func (s *PrivateLinkModifyService) Do(ctx context.Context) (PrivateLinkResponse, error) {
	var response PrivateLinkResponse
	if s.privateLinkId == nil {
		return response, fmt.Errorf("missing required privateLinkId")
	}

	url := fmt.Sprintf("/private-links/%v", *s.privateLinkId)
	err := s.HttpService.Do(ctx, "PATCH", url, s.request(), nil, 200, &response)
	return response, err
}