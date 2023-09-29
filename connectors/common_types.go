package connectors

import (
	"time"

	"github.com/fivetran/go-fivetran/common"
)

type DetailsResponseDataCommon struct {
	ID              string         `json:"id"`
	GroupID         string         `json:"group_id"`
	Service         string         `json:"service"`
	ServiceVersion  *int           `json:"service_version"`
	Schema          string         `json:"schema"`
	ConnectedBy     string         `json:"connected_by"`
	CreatedAt       time.Time      `json:"created_at"`
	SucceededAt     time.Time      `json:"succeeded_at"`
	FailedAt        time.Time      `json:"failed_at"`
	SyncFrequency   *int           `json:"sync_frequency"`
	ScheduleType    string         `json:"schedule_type"`
	Paused          *bool          `json:"paused"`
	PauseAfterTrial *bool          `json:"pause_after_trial"`
	DailySyncTime   string         `json:"daily_sync_time"`
	Status          StatusResponse `json:"status"`
}

type DetailsAndSetupTestsResponseDataCommon struct {
	DetailsResponseDataCommon
	SetupTests []common.SetupTestResponse `json:"setup_tests"`
}

type StatusResponse struct {
	SetupState       string                  `json:"setup_state"`
	SyncState        string                  `json:"sync_state"`
	UpdateState      string                  `json:"update_state"`
	IsHistoricalSync *bool                   `json:"is_historical_sync"`
	Tasks            []common.CommonResponse `json:"tasks"`
	Warnings         []common.CommonResponse `json:"warnings"`
}

type DetailsWithConfigResponse struct {
	common.CommonResponse
	Data struct {
		DetailsAndSetupTestsResponseDataCommon
		Config ConnectorConfigResponse `json:"config"`
	} `json:"data"`
}

type DetailsWithCustomConfigResponse struct {
	common.CommonResponse
	Data struct {
		DetailsAndSetupTestsResponseDataCommon
		Config map[string]interface{} `json:"config"`
	} `json:"data"`
}

type DetailsWithCustomMergedConfigResponse struct {
	common.CommonResponse
	Data struct {
		DetailsAndSetupTestsResponseDataCommon
		CustomConfig map[string]interface{}  `json:"config"`
		Config       ConnectorConfigResponse // no mapping here
	} `json:"data"`
}

type DetailsWithConfigNoTestsResponse struct {
	common.CommonResponse
	Data struct {
		DetailsResponseDataCommon
		Config ConnectorConfigResponse `json:"config"`
	} `json:"data"`
}

type DetailsWithCustomConfigNoTestsResponse struct {
	common.CommonResponse
	Data struct {
		DetailsResponseDataCommon
		Config map[string]interface{} `json:"config"`
	} `json:"data"`
}

type DetailsWithCustomMergedConfigNoTestsResponse struct {
	common.CommonResponse
	Data struct {
		DetailsResponseDataCommon
		CustomConfig map[string]interface{}  `json:"config"`
		Config       ConnectorConfigResponse // no mapping here
	} `json:"data"`
}

type ConnectorSchemaDetailsResponse struct {
	common.CommonResponse
	Data struct {
		SchemaChangeHandling string                                          `json:"schema_change_handling"`
		Schemas              map[string]*ConnectorSchemaConfigSchemaResponse `json:"schemas"`
	} `json:"data"`
}
