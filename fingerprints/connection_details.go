package fingerprints

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type ConnectionFingerprintDetailsService struct {
	httputils.HttpService
	connectionID *string
	hash        *string
}

func (s *ConnectionFingerprintDetailsService) ConnectionID(value string) *ConnectionFingerprintDetailsService {
	s.connectionID = &value
	return s
}

func (s *ConnectionFingerprintDetailsService) Hash(value string) *ConnectionFingerprintDetailsService {
	s.hash = &value
	return s
}

func (s *ConnectionFingerprintDetailsService) Do(ctx context.Context) (FingerprintResponse, error) {
	var response FingerprintResponse
	url := fmt.Sprintf("/connections/%v/fingerprints/%v", *s.connectionID, *s.hash)
	err := s.HttpService.Do(ctx, "GET", url, nil, nil, 200, &response)
	return response, err
}
