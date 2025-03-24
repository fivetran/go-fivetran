package fingerprints

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type ConnectionFingerprintApproveService struct {
	httputils.HttpService
	connectionID *string
	hash        *string
	publicKey   *string
}

type connectionFingerprintApproveRequest struct {
	Hash      *string `json:"hash,omitempty"`
	PublicKey *string `json:"public_key,omitempty"`
}

func (s *ConnectionFingerprintApproveService) request() *connectionFingerprintApproveRequest {
	return &connectionFingerprintApproveRequest{
		Hash:      s.hash,
		PublicKey: s.publicKey,
	}
}

func (s *ConnectionFingerprintApproveService) ConnectionID(value string) *ConnectionFingerprintApproveService {
	s.connectionID = &value
	return s
}

func (s *ConnectionFingerprintApproveService) Hash(value string) *ConnectionFingerprintApproveService {
	s.hash = &value
	return s
}

func (s *ConnectionFingerprintApproveService) PublicKey(value string) *ConnectionFingerprintApproveService {
	s.publicKey = &value
	return s
}

func (s *ConnectionFingerprintApproveService) Do(ctx context.Context) (FingerprintResponse, error) {
	var response FingerprintResponse
	url := fmt.Sprintf("/connections/%v/fingerprints", *s.connectionID)
	err := s.HttpService.Do(ctx, "POST", url, s.request(), nil, 201, &response)
	return response, err
}
