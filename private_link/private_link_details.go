package privatelink

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type PrivateLinkDetailsService struct {
	httputils.HttpService
	privateLinkId 	  *string
}

func (s *PrivateLinkDetailsService) PrivateLinkId(value string) *PrivateLinkDetailsService {
	s.privateLinkId = &value
	return s
}

func (s *PrivateLinkDetailsService) Do(ctx context.Context) (PrivateLinkResponse, error) {
	var response PrivateLinkResponse
	if s.privateLinkId == nil {
		return response, fmt.Errorf("missing required privateLinkId")
	}

	url := fmt.Sprintf("/private-links/%v", *s.privateLinkId)
	err := s.HttpService.Do(ctx, "GET", url, nil, nil, 200, &response)
	return response, err
}

func (s *PrivateLinkDetailsService) DoCustom(ctx context.Context) (PrivateLinkCustomResponse, error) {
	var response PrivateLinkCustomResponse
	if s.privateLinkId == nil {
		return response, fmt.Errorf("missing required privateLinkId")
	}

	url := fmt.Sprintf("/private-links/%v", *s.privateLinkId)
	err := s.HttpService.Do(ctx, "GET", url, nil, nil, 200, &response)
	return response, err
}