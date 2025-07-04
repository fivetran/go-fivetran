package metadata

import (
    "github.com/fivetran/go-fivetran/common"
)

type SupportedFeatures struct {
    Id          string  `json:"id"`
    Notes       string  `json:"notes"`
}

type Property struct {
    Type           string               `json:"type"`
    Description    string               `json:"description"`
    Title          string               `json:"title"`
    Readonly       bool                 `json:"readonly"`
    Properties     map[string]*Property `json:"properties"`
    Enum           []string             `json:"enum"`
    Items          *Property            `json:"items"`     // for array properties
    Required       []string             `json:"required"`
}

type ConnectorMetadata struct {
    ID                      string                  `json:"id"`
    Name                    string                  `json:"name"`
    Type                    string                  `json:"type"`
    Description             string                  `json:"description"`
    IconURL                 string                  `json:"icon_url"`
    LinkToDocs              string                  `json:"link_to_docs"`
    LinkToErd               string                  `json:"link_to_erd"`
    Icons                   []string                `json:"icons"`
    ConnectorClass          string                  `json:"connector_class"`
    ServiceStatus           string                  `json:"service_status"`
    ServiceStatusUpdatedAt  string                  `json:"service_status_updated_at"`
    SupportedFeatures       []SupportedFeatures     `json:"supported_features"`
    Config                  Property                `json:"config"`
    Auth                    Property                `json:"auth"`
}

type ConnectorMetadataResponse struct {
    common.CommonResponse
    Data struct {
        ConnectorMetadata
    } `json:"data"`
}

type ConnectorMetadataListResponse struct {
    common.CommonResponse
    Data struct {
        Items      []ConnectorMetadata `json:"items"`
        NextCursor string              `json:"next_cursor"`
    } `json:"data"`
}