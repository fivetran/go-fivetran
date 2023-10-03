package fingerprints

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type DestinationFingerprintApproveService struct {
	httputils.HttpService

	destinationID *string
	hash          *string
	publicKey     *string
}

type destinationFingerprintApproveRequest struct {
	Hash      *string `json:"hash,omitempty"`
	PublicKey *string `json:"public_key,omitempty"`
}

func (s *DestinationFingerprintApproveService) request() *destinationFingerprintApproveRequest {
	return &destinationFingerprintApproveRequest{
		Hash:      s.hash,
		PublicKey: s.publicKey,
	}
}

func (s *DestinationFingerprintApproveService) DestinationID(value string) *DestinationFingerprintApproveService {
	s.destinationID = &value
	return s
}

func (s *DestinationFingerprintApproveService) Hash(value string) *DestinationFingerprintApproveService {
	s.hash = &value
	return s
}

func (s *DestinationFingerprintApproveService) PublicKey(value string) *DestinationFingerprintApproveService {
	s.publicKey = &value
	return s
}

func (s *DestinationFingerprintApproveService) Do(ctx context.Context) (FingerprintResponse, error) {
	var response FingerprintResponse
	url := fmt.Sprintf("/destinations/%v/fingerprints", *s.destinationID)
	err := s.HttpService.Do(ctx, "POST", url, s.request(), nil, 201, &response)
	return response, err
}
