package privatelink

import (
	"context"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// PrivateLinkDeleteService implements the Private Link Management, Delete a Private Link
// Ref. https://fivetran.com/docs/rest-api/private-link-management#deleteaprivatelink
type PrivateLinkDeleteService struct {
	httputils.HttpService
	privateLinkId 	  *string
}

func (s *PrivateLinkDeleteService) PrivateLinkId(value string) *PrivateLinkDeleteService {
	s.privateLinkId = &value
	return s
}

func (s *PrivateLinkDeleteService) Do(ctx context.Context) (common.CommonResponse, error) {
	var response common.CommonResponse
	if s.privateLinkId == nil {
		return response, fmt.Errorf("missing required privateLinkId")
	}

	url := fmt.Sprintf("/private-links/%v", *s.privateLinkId)
	err := s.HttpService.Do(ctx, "DELETE", url, nil, nil, 200, &response)
	return response, err
}