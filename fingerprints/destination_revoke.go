package fingerprints

import (
	"context"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type DestinationFingerprintRevokeService struct {
	httputils.HttpService
	destinationID *string
	hash          *string
}

func (s *DestinationFingerprintRevokeService) DestinationID(value string) *DestinationFingerprintRevokeService {
	s.destinationID = &value
	return s
}

func (s *DestinationFingerprintRevokeService) Hash(value string) *DestinationFingerprintRevokeService {
	s.hash = &value
	return s
}

func (s *DestinationFingerprintRevokeService) Do(ctx context.Context) (common.CommonResponse, error) {
	var response common.CommonResponse
	url := fmt.Sprintf("/destinations/%v/fingerprints/%v", *s.destinationID, *s.hash)
	err := s.HttpService.Do(ctx, "DELETE", url, nil, nil, 200, &response)
	return response, err
}
