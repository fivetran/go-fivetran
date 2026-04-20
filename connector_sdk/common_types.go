package connectorsdk

import (
	"time"

	"github.com/fivetran/go-fivetran/common"
)

type ConnectorSdkPackageData struct {
	ID             string    `json:"id"`
	ConnectionID   string    `json:"connection_id"`
	CreatedBy      string    `json:"created_by"`
	LastUpdatedBy  string    `json:"last_updated_by"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	FileSha256Hash string    `json:"file_sha256_hash"`
}

type ConnectorSdkPackageResponse struct {
	common.CommonResponse
	Data ConnectorSdkPackageData `json:"data"`
}

type ConnectorSdkPackagesListResponse struct {
	common.CommonResponse
	Data struct {
		Items      []ConnectorSdkPackageData `json:"items"`
		NextCursor string                    `json:"next_cursor"`
	} `json:"data"`
}
