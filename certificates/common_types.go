package certificates

import "github.com/fivetran/go-fivetran/common"

type CertificateDetails struct {
	PublicKey     string `json:"public_key"`
	Name          string `json:"name"`
	Hash          string `json:"hash"`
	Type          string `json:"type"`
	ValidatedDate string `json:"validated_date"`
	ValidatedBy   string `json:"validated_by"`
	Sha1          string `json:"sha1"`
	Sha256        string `json:"sha256"`
}

type CertificateResponse struct {
	common.CommonResponse
	Data CertificateDetails `json:"data"`
}

type CertificatesListResponse struct {
	common.CommonResponse
	Data struct {
		Items      []CertificateDetails `json:"items"`
		NextCursor string               `json:"next_cursor"`
	} `json:"data"`
}

type certificateApproveRequest struct {
	Hash        *string `json:"hash,omitempty"`
	EncodedCert *string `json:"encoded_cert,omitempty"`
}
