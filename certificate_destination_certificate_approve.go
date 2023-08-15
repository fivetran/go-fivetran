package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
)

// CertificateDestinationCertificateApproveService implements the Certificate Management, Approve a destination certificate API.
// Ref. https://fivetran.com/docs/rest-api/certificates#approveadestinationcertificate
type CertificateDestinationCertificateApproveService struct {
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

func (c *Client) NewCertificateDestinationCertificateApprove() *CertificateDestinationCertificateApproveService {
	return &CertificateDestinationCertificateApproveService{c: c}
}

func (s *CertificateDestinationCertificateApproveService) request() *certificateDestinationCertificateApproveRequest {
	return &certificateDestinationCertificateApproveRequest{
		DestinationID: s.destinationID,
		Hash:          s.hash,
		EncodedCert:   s.encodedCert,
	}
}

func (s *CertificateDestinationCertificateApproveService) DestinationID(value string) *CertificateDestinationCertificateApproveService {
	s.destinationID = &value
	return s
}

func (s *CertificateDestinationCertificateApproveService) Hash(value string) *CertificateDestinationCertificateApproveService {
	s.hash = &value
	return s
}

func (s *CertificateDestinationCertificateApproveService) EncodedCert(value string) *CertificateDestinationCertificateApproveService {
	s.encodedCert = &value
	return s
}

func (s *CertificateDestinationCertificateApproveService) Do(ctx context.Context) (CertificateDestinationCertificateApproveResponse, error) {
	var response CertificateDestinationCertificateApproveResponse
	url := fmt.Sprintf("%v/certificates", s.c.baseURL)
	expectedStatus := 200

	headers := s.c.commonHeaders()
	headers["Content-Type"] = "application/json"

	reqBody, err := json.Marshal(s.request())
	if err != nil {
		return response, err
	}

	r := request{
		method:           "POST",
		url:              url,
		body:             reqBody,
		queries:          nil,
		headers:          headers,
		client:           s.c.httpClient,
		handleRateLimits: s.c.handleRateLimits,
		maxRetryAttempts: s.c.maxRetryAttempts,
	}

	respBody, respStatus, err := r.httpRequest(ctx)
	if err != nil {
		return response, err
	}

	if err := json.Unmarshal(respBody, &response); err != nil {
		return response, err
	}

	if respStatus != expectedStatus {
		err := fmt.Errorf("status code: %v; expected: %v", respStatus, expectedStatus)
		return response, err
	}

	return response, nil
}
