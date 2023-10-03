package fingerprints

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type DestinationFingerprintDetailsService struct {
	httputils.HttpService
	destinationID *string
	hash          *string
}

func (s *DestinationFingerprintDetailsService) DestinationID(value string) *DestinationFingerprintDetailsService {
	s.destinationID = &value
	return s
}

func (s *DestinationFingerprintDetailsService) Hash(value string) *DestinationFingerprintDetailsService {
	s.hash = &value
	return s
}

func (s *DestinationFingerprintDetailsService) Do(ctx context.Context) (FingerprintResponse, error) {
	var response FingerprintResponse
	url := fmt.Sprintf("/destinations/%v/fingerprints/%v", *s.destinationID, *s.hash)
	err := s.HttpService.Do(ctx, "GET", url, nil, nil, 200, &response)
	return response, err
}
