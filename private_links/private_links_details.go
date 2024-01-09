package privatelinks

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// PrivateLinksDetailsService implements the Private Links Management, Retrieve Private Link details API.
// Ref. https://fivetran.com/docs/rest-api/private-links-management#retrieveprivatelinkdetails
type PrivateLinksDetailsService struct {
	httputils.HttpService
	privateLinkId 	  *string
}

func (s *PrivateLinksDetailsService) PrivateLinkId(value string) *PrivateLinksDetailsService {
	s.privateLinkId = &value
	return s
}

func (s *PrivateLinksDetailsService) Do(ctx context.Context) (PrivateLinksResponse, error) {
	var response PrivateLinksResponse
	if s.privateLinkId == nil {
		return response, fmt.Errorf("missing required privateLinkId")
	}

	url := fmt.Sprintf("/private-links/%v", *s.privateLinkId)
	err := s.HttpService.Do(ctx, "GET", url, nil, nil, 200, &response)
	return response, err
}