package connectors

import (
    "time"

    "github.com/fivetran/go-fivetran/common"
)

type DetailsResponseDataCommon struct {
    ID                      string         `json:"id"`
    GroupID                 string         `json:"group_id"`
    Service                 string         `json:"service"`
    ServiceVersion          *int           `json:"service_version"`
    Schema                  string         `json:"schema"`
    ConnectedBy             string         `json:"connected_by"`
    CreatedAt               time.Time      `json:"created_at"`
    SucceededAt             time.Time      `json:"succeeded_at"`
    FailedAt                time.Time      `json:"failed_at"`
    SyncFrequency           *int           `json:"sync_frequency"`
    ScheduleType            string         `json:"schedule_type"`
    Paused                  *bool          `json:"paused"`
    PauseAfterTrial         *bool          `json:"pause_after_trial"`
    DailySyncTime           string         `json:"daily_sync_time"`
    PrivateLinkId           string         `json:"private_link_id"`
    HybridDeploymentAgentId string         `json:"hybrid_deployment_agent_id"`
    ProxyAgentId            string         `json:"proxy_agent_id"`
    NetworkingMethod        string         `json:"networking_method"`
    DataDelaySensitivity    string         `json:"data_delay_sensitivity"`
    DataDelayThreshold      *int           `json:"data_delay_threshold"`
    Status                  StatusResponse `json:"status"`
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

type connectorCreateRequestBase struct {
    Service                 *string `json:"service,omitempty"`
    GroupID                 *string `json:"group_id,omitempty"`
    TrustCertificates       *bool   `json:"trust_certificates,omitempty"`
    TrustFingerprints       *bool   `json:"trust_fingerprints,omitempty"`
    RunSetupTests           *bool   `json:"run_setup_tests,omitempty"`
    Paused                  *bool   `json:"paused,omitempty"`
    SyncFrequency           *int    `json:"sync_frequency,omitempty"`
    DailySyncTime           *string `json:"daily_sync_time,omitempty"`
    PauseAfterTrial         *bool   `json:"pause_after_trial,omitempty"`
    ProxyAgentId            *string `json:"proxy_agent_id,omitempty"`
    PrivateLinkId           *string `json:"private_link_id,omitempty"`
    HybridDeploymentAgentId *string `json:"hybrid_deployment_agent_id,omitempty"`
    NetworkingMethod        *string `json:"networking_method,omitempty"`
    DataDelaySensitivity    *string `json:"data_delay_sensitivity"`
    DataDelayThreshold      *int    `json:"data_delay_threshold"`
}

type connectorCreateRequest struct {
    connectorCreateRequestBase
    Config any `json:"config,omitempty"`
    Auth   any `json:"auth,omitempty"`
}

type connectorCustomCreateRequest struct {
    connectorCreateRequestBase
    Config *map[string]interface{} `json:"config,omitempty"`
    Auth   *map[string]interface{} `json:"auth,omitempty"`
}

type connectorSchemaConfigModifyRequest struct {
    Enabled *bool                                         `json:"enabled,omitempty"`
    Tables  map[string]*ConnectorSchemaConfigTableRequest `json:"tables,omitempty"`
}

type ConnectorColumnConfigListResponse struct {
    common.CommonResponse
    Data struct {
        Columns map[string]*ConnectorSchemaConfigColumnResponse `json:"columns"`
    } `json:"data"`
}

type connectorColumnConfigModifyRequest struct {
    Enabled *bool `json:"enabled,omitempty"`
    Hashed  *bool `json:"hashed,omitempty"`
}

type connectorModifyRequestBase struct {
    Paused                  *bool   `json:"paused,omitempty"`
    SyncFrequency           *int    `json:"sync_frequency,omitempty"`
    DailySyncTime           *string `json:"daily_sync_time,omitempty"`
    TrustCertificates       *bool   `json:"trust_certificates,omitempty"`
    TrustFingerprints       *bool   `json:"trust_fingerprints,omitempty"`
    IsHistoricalSync        *bool   `json:"is_historical_sync,omitempty"`
    ScheduleType            *string `json:"schedule_type,omitempty"`
    RunSetupTests           *bool   `json:"run_setup_tests,omitempty"`
    PauseAfterTrial         *bool   `json:"pause_after_trial,omitempty"`
    ProxyAgentId            *string `json:"proxy_agent_id,omitempty"`
    PrivateLinkId           *string `json:"private_link_id,omitempty"`
    HybridDeploymentAgentId *string `json:"hybrid_deployment_agent_id,omitempty"`
    NetworkingMethod        *string `json:"networking_method,omitempty"`
    DataDelaySensitivity    *string `json:"data_delay_sensitivity,omitempty"`
    DataDelayThreshold      *int    `json:"data_delay_threshold,omitempty"`
}

type connectorModifyRequest struct {
    connectorModifyRequestBase
    Config any `json:"config,omitempty"`
    Auth   any `json:"auth,omitempty"`
}

type connectorCustomModifyRequest struct {
    connectorModifyRequestBase
    Config *map[string]interface{} `json:"config,omitempty"`
    Auth   *map[string]interface{} `json:"auth,omitempty"`
}

type connectorSchemaReloadRequest struct {
    ExcludeMode *string `json:"exclude_mode,omitempty"`
}

type connectorSchemaConfigUpdateRequest struct {
    SchemaChangeHandling *string                                        `json:"schema_change_handling,omitempty"`
    Schemas              map[string]*ConnectorSchemaConfigSchemaRequest `json:"schemas,omitempty"`
}

type connectorSetupTestsRequest struct {
    TrustCertificates *bool `json:"trust_certificates,omitempty"`
    TrustFingerprints *bool `json:"trust_fingerprints,omitempty"`
}

type connectorTableConfigModifyRequest struct {
    Enabled  *bool                                          `json:"enabled,omitempty"`
    SyncMode *string                                        `json:"sync_mode,omitempty"`
    Columns  map[string]*ConnectorSchemaConfigColumnRequest `json:"columns,omitempty"`
}

type ConnectorsSourceMetadataResponse struct {
    common.CommonResponse
    Data struct {
        Items []struct {
            ID          string `json:"id"`
            Name        string `json:"name"`
            Type        string `json:"type"`
            Description string `json:"description"`
            IconURL     string `json:"icon_url"`
            LinkToDocs  string `json:"link_to_docs"`
            LinkToErd   string `json:"link_to_erd"`
        } `json:"items"`
        NextCursor string `json:"next_cursor"`
    } `json:"data"`
}

type ConnectorsListResponse struct {
    common.CommonResponse
    Data struct {
        Items      []DetailsResponseDataCommon `json:"items"`
        NextCursor string                      `json:"next_cursor"`
    } `json:"data"`
}