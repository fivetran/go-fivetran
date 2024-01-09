package privatelinks

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// PrivateLinksCreateService implements the Log Management, Create a Log Service API.
// Ref. https://fivetran.com/docs/rest-api/log-service-management#createalogservice
type PrivateLinksCreateService struct {
	httputils.HttpService
	name      	 *string
	service      *string
	groupId      *string
	config       *PrivateLinksConfig
}


func (s *PrivateLinksCreateService) request() privateLinksCreateRequest {
	var config interface{}
	if s.config != nil {
		config = s.config.Request()
	}

	return privateLinksCreateRequest{
		Name: 		s.name,
		GroupId: 	s.groupId,
		Service: 	s.service,
		Config:     config,
	}
}

func (s *PrivateLinksCreateService) GroupId(value string) *PrivateLinksCreateService {
	s.groupId = &value
	return s
}

func (s *PrivateLinksCreateService) Service(value string) *PrivateLinksCreateService {
	s.service = &value
	return s
}

func (s *PrivateLinksCreateService) Name(value string) *PrivateLinksCreateService {
	s.name = &value
	return s
}

func (s *PrivateLinksCreateService) Config(value *PrivateLinksConfig) *PrivateLinksCreateService {
	s.config = value
	return s
}

func (s *PrivateLinksCreateService) Do(ctx context.Context) (PrivateLinksResponse, error) {
	var response PrivateLinksResponse
	fmt.Printf("request %v",  s.request())
	err := s.HttpService.Do(ctx, "POST", "/private-links", s.request(), nil, 201, &response)
	return response, err
}