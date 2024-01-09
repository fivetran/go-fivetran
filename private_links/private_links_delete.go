package privatelinks

import (
	"context"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// PrivateLinksDeleteService implements the Private Links Management, Delete a Private Link
// Ref. https://fivetran.com/docs/rest-api/private-links-management#deleteaprivatelink
type PrivateLinksDeleteService struct {
	httputils.HttpService
	privateLinkId 	  *string
}

func (s *PrivateLinksDeleteService) PrivateLinkId(value string) *PrivateLinksDeleteService {
	s.privateLinkId = &value
	return s
}

func (s *PrivateLinksDeleteService) Do(ctx context.Context) (common.CommonResponse, error) {
	var response common.CommonResponse
	if s.privateLinkId == nil {
		return response, fmt.Errorf("missing required privateLinkId")
	}

	url := fmt.Sprintf("/private-links/%v", *s.privateLinkId)
	err := s.HttpService.Do(ctx, "DELETE", url, nil, nil, 200, &response)
	return response, err
}