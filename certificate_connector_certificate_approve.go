package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
)

type certificateConnectorCertificateApproveService struct {
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

func (c *Client) NewCertificateConnectorCertificateApprove() *certificateConnectorCertificateApproveService {
	return &certificateConnectorCertificateApproveService{c: c}
}

func (s *certificateConnectorCertificateApproveService) request() certificateConnectorCertificateApproveRequest {
	return certificateConnectorCertificateApproveRequest{
		ConnectorID: s.connectorID,
		Hash:        s.hash,
		EncodedCert: s.encodedCert,
	}
}

func (s *certificateConnectorCertificateApproveService) ConnectorID(value string) *certificateConnectorCertificateApproveService {
	s.connectorID = &value
	return s
}

func (s *certificateConnectorCertificateApproveService) Hash(value string) *certificateConnectorCertificateApproveService {
	s.hash = &value
	return s
}

func (s *certificateConnectorCertificateApproveService) EncodedCert(value string) *certificateConnectorCertificateApproveService {
	s.encodedCert = &value
	return s
}

func (s *certificateConnectorCertificateApproveService) Do(ctx context.Context) (CertificateConnectorCertificateApproveResponse, error) {
	var response CertificateConnectorCertificateApproveResponse
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
