package certificates

import (
	"context"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type DestinationCertificateRevokeService struct {
	httputils.HttpService
	destinationID *string
	hash          *string
}

func (s *DestinationCertificateRevokeService) DestinationID(value string) *DestinationCertificateRevokeService {
	s.destinationID = &value
	return s
}

func (s *DestinationCertificateRevokeService) Hash(value string) *DestinationCertificateRevokeService {
	s.hash = &value
	return s
}

func (s *DestinationCertificateRevokeService) Do(ctx context.Context) (common.CommonResponse, error) {
	var response common.CommonResponse
	url := fmt.Sprintf("/destinations/%v/certificates/%v", *s.destinationID, *s.hash)
	err := s.HttpService.Do(ctx, "DELETE", url, nil, nil, 200, &response)
	return response, err
}
