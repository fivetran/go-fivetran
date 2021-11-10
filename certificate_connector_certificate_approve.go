package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
)

// CertificateConnectorCertificateApproveService implements the Certificate Management, Approve a connector certificate API.
// Ref. https://fivetran.com/docs/rest-api/certificates#approveaconnectorcertificate
type CertificateConnectorCertificateApproveService struct {
	c           *Client
	connectorID *string
	hash        *string
	encodedCert *string
}

type certificateConnectorCertificateApproveRequest struct {
	ConnectorID *string `json:"connector_id,omitempty"`
	Hash        *string `json:"hash,omitempty"`
	EncodedCert *string `json:"encoded_cert,omitempty"`
}

type CertificateConnectorCertificateApproveResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (c *Client) NewCertificateConnectorCertificateApprove() *CertificateConnectorCertificateApproveService {
	return &CertificateConnectorCertificateApproveService{c: c}
}

func (s *CertificateConnectorCertificateApproveService) request() *certificateConnectorCertificateApproveRequest {
	return &certificateConnectorCertificateApproveRequest{
		ConnectorID: s.connectorID,
		Hash:        s.hash,
		EncodedCert: s.encodedCert,
	}
}

func (s *CertificateConnectorCertificateApproveService) ConnectorID(value string) *CertificateConnectorCertificateApproveService {
	s.connectorID = &value
	return s
}

func (s *CertificateConnectorCertificateApproveService) Hash(value string) *CertificateConnectorCertificateApproveService {
	s.hash = &value
	return s
}

func (s *CertificateConnectorCertificateApproveService) EncodedCert(value string) *CertificateConnectorCertificateApproveService {
	s.encodedCert = &value
	return s
}

func (s *CertificateConnectorCertificateApproveService) Do(ctx context.Context) (CertificateConnectorCertificateApproveResponse, error) {
	var response CertificateConnectorCertificateApproveResponse
	url := fmt.Sprintf("%v/certificates", s.c.baseURL)
	expectedStatus := 200

	headers := s.c.commonHeaders()
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
