package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
)

type certificateDestinationFingerprintApproveService struct {
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

func (c *Client) NewCertificateDestinationFingerprintApprove() *certificateDestinationFingerprintApproveService {
	return &certificateDestinationFingerprintApproveService{c: c}
}

func (s *certificateDestinationFingerprintApproveService) request() certificateDestinationFingerprintApproveRequest {
	return certificateDestinationFingerprintApproveRequest{
		DestinationID: s.destinationID,
		Hash:          s.hash,
		PublicKey:     s.publicKey,
	}
}

func (s *certificateDestinationFingerprintApproveService) DestinationID(value string) *certificateDestinationFingerprintApproveService {
	s.destinationID = &value
	return s
}

func (s *certificateDestinationFingerprintApproveService) Hash(value string) *certificateDestinationFingerprintApproveService {
	s.hash = &value
	return s
}

func (s *certificateDestinationFingerprintApproveService) PublicKey(value string) *certificateDestinationFingerprintApproveService {
	s.publicKey = &value
	return s
}

func (s *certificateDestinationFingerprintApproveService) Do(ctx context.Context) (CertificateDestinationFingerprintApproveResponse, error) {
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

	r := Request{
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
