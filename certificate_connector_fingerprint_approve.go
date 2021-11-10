package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
)

// CertificateConnectorFingerprintApproveService implements the Certificate Management, Approve a connector fingerprint API.
// Ref. https://fivetran.com/docs/rest-api/certificates#approveaconnectorfingerprint
type CertificateConnectorFingerprintApproveService struct {
	c           *Client
	connectorID *string
	hash        *string
	publicKey   *string
}

type certificateConnectorFingerprintApproveRequest struct {
	ConnectorID *string `json:"connector_id,omitempty"`
	Hash        *string `json:"hash,omitempty"`
	PublicKey   *string `json:"public_key,omitempty"`
}

type CertificateConnectorFingerprintApproveResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (c *Client) NewCertificateConnectorFingerprintApprove() *CertificateConnectorFingerprintApproveService {
	return &CertificateConnectorFingerprintApproveService{c: c}
}

func (s *CertificateConnectorFingerprintApproveService) request() *certificateConnectorFingerprintApproveRequest {
	return &certificateConnectorFingerprintApproveRequest{
		ConnectorID: s.connectorID,
		Hash:        s.hash,
		PublicKey:   s.publicKey,
	}
}

func (s *CertificateConnectorFingerprintApproveService) ConnectorID(value string) *CertificateConnectorFingerprintApproveService {
	s.connectorID = &value
	return s
}

func (s *CertificateConnectorFingerprintApproveService) Hash(value string) *CertificateConnectorFingerprintApproveService {
	s.hash = &value
	return s
}

func (s *CertificateConnectorFingerprintApproveService) PublicKey(value string) *CertificateConnectorFingerprintApproveService {
	s.publicKey = &value
	return s
}

func (s *CertificateConnectorFingerprintApproveService) Do(ctx context.Context) (CertificateConnectorFingerprintApproveResponse, error) {
	var response CertificateConnectorFingerprintApproveResponse
	url := fmt.Sprintf("%v/fingerprints", s.c.baseURL)
	expectedStatus := 200

	headers := s.c.fillHeaders()
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
		err := fmt.Errorf("status code: %v; expected: %v", respStatus, expectedStatus)
		return response, err
	}

	return response, nil
}
