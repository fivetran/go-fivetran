package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
)

type CertificateDestinationFingerprintApproveService struct {
	c             *Client
	destinationID *string
	hash          *string
	publicKey     *string
}

type certificateDestinationFingerprintApproveRequest struct {
	DestinationID *string `json:"destination_id,omitempty"`
	Hash          *string `json:"hash,omitempty"`
	PublicKey     *string `json:"public_key,omitempty"`
}

type CertificateDestinationFingerprintApproveResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (c *Client) NewCertificateDestinationFingerprintApprove() *CertificateDestinationFingerprintApproveService {
	return &CertificateDestinationFingerprintApproveService{c: c}
}

func (s *CertificateDestinationFingerprintApproveService) request() certificateDestinationFingerprintApproveRequest {
	return certificateDestinationFingerprintApproveRequest{
		DestinationID: s.destinationID,
		Hash:          s.hash,
		PublicKey:     s.publicKey,
	}
}

func (s *CertificateDestinationFingerprintApproveService) DestinationID(value string) *CertificateDestinationFingerprintApproveService {
	s.destinationID = &value
	return s
}

func (s *CertificateDestinationFingerprintApproveService) Hash(value string) *CertificateDestinationFingerprintApproveService {
	s.hash = &value
	return s
}

func (s *CertificateDestinationFingerprintApproveService) PublicKey(value string) *CertificateDestinationFingerprintApproveService {
	s.publicKey = &value
	return s
}

func (s *CertificateDestinationFingerprintApproveService) Do(ctx context.Context) (CertificateDestinationFingerprintApproveResponse, error) {
	var response CertificateDestinationFingerprintApproveResponse
	url := fmt.Sprintf("%v/fingerprints", s.c.baseURL)
	expectedStatus := 200

	headers := make(map[string]string)
	headers["Authorization"] = s.c.authorization
	headers["Content-Type"] = "application/json"

	reqBody, err := json.Marshal(s.request())
	if err != nil {
		return response, err
	}

	r := request{
		method:  "POST",
		url:     url,
		body:    reqBody,
		queries: nil,
		headers: headers,
	}

	respBody, respStatus, err := httpRequest(r, ctx)
	if err != nil {
		return response, err
	}

	if err := json.Unmarshal(respBody, &response); err != nil {
		return response, err
	}

	if respStatus != expectedStatus {
		err := fmt.Errorf("status code: %v; expected %v", respStatus, expectedStatus)
		return response, err
	}

	return response, nil
}
