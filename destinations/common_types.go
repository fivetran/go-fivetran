package destinations

import "github.com/fivetran/go-fivetran/common"

type DestinationDetailsBase struct {
	ID             string                    `json:"id"`
	GroupID        string                    `json:"group_id"`
	Service        string                    `json:"service"`
	Region         string                    `json:"region"`
	TimeZoneOffset string                    `json:"time_zone_offset"`
	SetupStatus    string                    `json:"setup_status"`
	Config         DestinationConfigResponse `json:"config"`
}

type DestinationDetailsWithSetupTestsResponse struct {
	common.CommonResponse
	Data struct {
		DestinationDetailsBase
		SetupTests []common.SetupTestResponse `json:"setup_tests"`
	} `json:"data"`
}

type DestinationDetailsResponse struct {
	Code    string                 `json:"code"`
	Message string                 `json:"message"`
	Data    DestinationDetailsBase `json:"data"`
}
