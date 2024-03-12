package destinations

import "github.com/fivetran/go-fivetran/common"

type DestinationDetailsBase struct {
	ID             				string `json:"id"`
	GroupID        				string `json:"group_id"`
	Service        				string `json:"service"`
	Region         				string `json:"region"`
	TimeZoneOffset 				string `json:"time_zone_offset"`
	SetupStatus    				string `json:"setup_status"`
	DaylightSavingTimeEnabled   bool   `json:"daylight_saving_time_enabled"`
}

type DestinationDetailsWithSetupTestsResponse struct {
	common.CommonResponse
	Data struct {
		DestinationDetailsBase
		Config     DestinationConfigResponse  `json:"config"`
		SetupTests []common.SetupTestResponse `json:"setup_tests"`
	} `json:"data"`
}

type DestinationDetailsWithSetupTestsCustomResponse struct {
	common.CommonResponse
	Data struct {
		DestinationDetailsBase
		Config     map[string]interface{}     `json:"config"`
		SetupTests []common.SetupTestResponse `json:"setup_tests"`
	} `json:"data"`
}

type DestinationDetailsResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		DestinationDetailsBase
		Config DestinationConfigResponse `json:"config"`
	} `json:"data"`
}

type DestinationDetailsCustomResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		DestinationDetailsBase
		Config map[string]interface{} `json:"config"`
	} `json:"data"`
}

type destinationCreateRequest struct {
	GroupID           			*string `json:"group_id,omitempty"`
	Service           			*string `json:"service,omitempty"`
	Region           			*string `json:"region,omitempty"`
	TimeZoneOffset    			*string `json:"time_zone_offset,omitempty"`
	Config            			any     `json:"config,omitempty"`
	TrustCertificates 			*bool   `json:"trust_certificates,omitempty"`
	TrustFingerprints 			*bool   `json:"trust_fingerprints,omitempty"`
	RunSetupTests    			*bool   `json:"run_setup_tests,omitempty"`
	DaylightSavingTimeEnabled   *bool   `json:"daylight_saving_time_enabled,omitempty"`
	
}

type destinationModifyRequest struct {
	Region            			*string `json:"region,omitempty"`
	TimeZoneOffset    			*string `json:"time_zone_offset,omitempty"`
	Config            			any     `json:"config,omitempty"`
	TrustCertificates 			*bool   `json:"trust_certificates,omitempty"`
	TrustFingerprints 			*bool   `json:"trust_fingerprints,omitempty"`
	RunSetupTests     			*bool   `json:"run_setup_tests,omitempty"`
	DaylightSavingTimeEnabled   *bool   `json:"daylight_saving_time_enabled,omitempty"`
}

type destinationSetupTestsRequest struct {
	TrustCertificates *bool `json:"trust_certificates,omitempty"`
	TrustFingerprints *bool `json:"trust_fingerprints,omitempty"`
}
