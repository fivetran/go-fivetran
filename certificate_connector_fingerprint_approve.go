package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
)

type certificateConnectorFingerprintApproveService struct {
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

func (c *Client) NewCertificateConnectorFingerprintApprove() *certificateConnectorFingerprintApproveService {
	return &certificateConnectorFingerprintApproveService{c: c}
}

func (s *certificateConnectorFingerprintApproveService) request() certificateConnectorFingerprintApproveRequest {
	return certificateConnectorFingerprintApproveRequest{
		ConnectorID: s.connectorID,
		Hash:        s.hash,
		PublicKey:   s.publicKey,
	}
}

func (s *certificateConnectorFingerprintApproveService) ConnectorID(value string) *certificateConnectorFingerprintApproveService {
	s.connectorID = &value
	return s
}

func (s *certificateConnectorFingerprintApproveService) Hash(value string) *certificateConnectorFingerprintApproveService {
	s.hash = &value
	return s
}

func (s *certificateConnectorFingerprintApproveService) PublicKey(value string) *certificateConnectorFingerprintApproveService {
	s.publicKey = &value
	return s
}

func (s *certificateConnectorFingerprintApproveService) Do(ctx context.Context) (CertificateConnectorFingerprintApproveResponse, error) {
	var response CertificateConnectorFingerprintApproveResponse
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
