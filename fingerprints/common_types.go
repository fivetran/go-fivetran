package fingerprints

import "github.com/fivetran/go-fivetran/common"

type FingerprintResponse struct {
	common.CommonResponse
	Data FingerprintDetails `json:"data"`
}

type FingerprintDetails struct {
	Hash          string `json:"hash"`
	PublicKey     string `json:"public_key"`
	ValidatedDate string `json:"validated_date"`
	ValidatedBy   string `json:"validated_by"`
}

type FingerprintsListResponse struct {
	common.CommonResponse
	Data struct {
		Items      []FingerprintDetails `json:"items"`
		NextCursor string               `json:"next_cursor"`
	} `json:"data"`
}
