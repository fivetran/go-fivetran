package metadata

import (
	"github.com/fivetran/go-fivetran/common"
)

type SupportedFeatures struct {
	Id    string `json:"id"`
	Notes string `json:"notes"`
}

type Property struct {
	Type        string               `json:"type"`
	Description string               `json:"description"`
	Title       string               `json:"title"`
	Readonly    bool                 `json:"readonly"`
	Nullable    bool                 `json:"nullable"`
	Format      string               `json:"format"`
	Immutable   bool                 `json:"immutable"`
	FieldStatus string               `json:"fieldStatus"`
	Properties  map[string]*Property `json:"properties"`
	Enum        []string             `json:"enum"`
	Items       *Property            `json:"items"` // for array properties
	Required    []string             `json:"required"`
}

type ConnectorMetadata struct {
	ID                     string              `json:"id"`
	Name                   string              `json:"name"`
	Type                   string              `json:"type"`
	Description            string              `json:"description"`
	IconURL                string              `json:"icon_url"`
	LinkToDocs             string              `json:"link_to_docs"`
	LinkToErd              string              `json:"link_to_erd"`
	Icons                  []string            `json:"icons"`
	ConnectorClass         string              `json:"connector_class"`
	ServiceStatus          string              `json:"service_status"`
	ServiceStatusUpdatedAt string              `json:"service_status_updated_at"`
	SupportedFeatures      []SupportedFeatures `json:"supported_features"`
	SupportsSchemaReload   *bool               `json:"supports_schema_reload,omitempty"`
	Config                 Property            `json:"config"`
	Auth                   Property            `json:"auth"`
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

// DestinationMetadata describes a destination type: the fields accepted in its config and the
// capabilities (regions, networking, hybrid deployment, PrivateLink) the destination service supports.
// Config and ExternalSecretsKeysConfig are returned only by the details endpoint; the list endpoint
// omits them.
type DestinationMetadata struct {
	ID                         string   `json:"id"`
	Name                       string   `json:"name"`
	Type                       string   `json:"type"`
	Description                string   `json:"description"`
	IconURL                    string   `json:"icon_url"`
	LinkToDocs                 string   `json:"link_to_docs"`
	Icons                      []string `json:"icons"`
	ServiceStatus              string   `json:"service_status"`
	ServiceStatusUpdatedAt     string   `json:"service_status_updated_at"`
	SupportedRegions           []string `json:"supported_regions"`
	SupportsHybridDeployment   bool     `json:"supports_hybrid_deployment"`
	SupportedNetworkingMethods []string `json:"supported_networking_methods"`
	SupportsPrivateLink        bool     `json:"supports_private_link"`
	PrivateLinkTypes           []string `json:"private_link_types"`
	Config                     Property `json:"config"`
	// ExternalSecretsKeysConfig is returned for internal accounts only.
	ExternalSecretsKeysConfig *Property `json:"external_secrets_keys_config,omitempty"`
}

type DestinationMetadataResponse struct {
	common.CommonResponse
	Data struct {
		DestinationMetadata
	} `json:"data"`
}

type DestinationMetadataListResponse struct {
	common.CommonResponse
	Data struct {
		Items      []DestinationMetadata `json:"items"`
		NextCursor string                `json:"next_cursor"`
	} `json:"data"`
}
