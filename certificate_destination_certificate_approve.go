package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
)

type certificateDestinationCertificateApproveService struct {
	c             *Client
	destinationID *string
	hash          *string
	encodedCert   *string
}

type certificateDestinationCertificateApproveRequest struct {
	DestinationID *string `json:"destination_id,omitempty"`
	Hash          *string `json:"hash,omitempty"`
	EncodedCert   *string `json:"encoded_cert,omitempty"`
}

type CertificateDestinationCertificateApproveResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (c *Client) NewCertificateDestinationCertificateApprove() *certificateDestinationCertificateApproveService {
	return &certificateDestinationCertificateApproveService{c: c}
}

func (s *certificateDestinationCertificateApproveService) request() certificateDestinationCertificateApproveRequest {
	return certificateDestinationCertificateApproveRequest{
		DestinationID: s.destinationID,
		Hash:          s.hash,
		EncodedCert:   s.encodedCert,
	}
}

func (s *certificateDestinationCertificateApproveService) DestinationID(value string) *certificateDestinationCertificateApproveService {
	s.destinationID = &value
	return s
}

func (s *certificateDestinationCertificateApproveService) Hash(value string) *certificateDestinationCertificateApproveService {
	s.hash = &value
	return s
}

func (s *certificateDestinationCertificateApproveService) EncodedCert(value string) *certificateDestinationCertificateApproveService {
	s.encodedCert = &value
	return s
}

func (s *certificateDestinationCertificateApproveService) Do(ctx context.Context) (CertificateDestinationCertificateApproveResponse, error) {
	var response CertificateDestinationCertificateApproveResponse
	url := fmt.Sprintf("%v/certificates", s.c.baseURL)
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
