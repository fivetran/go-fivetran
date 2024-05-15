package privatelink

import (
	"context"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// PrivateLinkCreateService implements the Log Management, Create a Log Service API.
// Ref. https://fivetran.com/docs/rest-api/log-service-management#createalogservice
type PrivateLinkCreateService struct {
	httputils.HttpService
	name      	 *string
	region       *string
	service      *string
	config       *PrivateLinkConfig
}


func (s *PrivateLinkCreateService) request() privateLinkCreateRequest {
	var config interface{}
	if s.config != nil {
		config = s.config.Request()
	}

	return privateLinkCreateRequest{
		Name: 		s.name,
		Region: 	s.region,
		Service: 	s.service,
		Config:     config,
	}
}

func (s *PrivateLinkCreateService) Region(value string) *PrivateLinkCreateService {
	s.region = &value
	return s
}

func (s *PrivateLinkCreateService) Service(value string) *PrivateLinkCreateService {
	s.service = &value
	return s
}

func (s *PrivateLinkCreateService) Name(value string) *PrivateLinkCreateService {
	s.name = &value
	return s
}

func (s *PrivateLinkCreateService) Config(value *PrivateLinkConfig) *PrivateLinkCreateService {
	s.config = value
	return s
}

func (s *PrivateLinkCreateService) Do(ctx context.Context) (PrivateLinkResponse, error) {
	var response PrivateLinkResponse
	err := s.HttpService.Do(ctx, "POST", "/private-links", s.request(), nil, 201, &response)
	return response, err
}