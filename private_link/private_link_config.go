package privatelink

import "github.com/fivetran/go-fivetran/utils"

type PrivateLinkConfig struct {
    connectionServiceName       *string
    connectionServiceId         *string
    workspaceUrl                *string
    accountName                 *string
    accountUrl                  *string
    vpceId                      *string
    plsId                       *string
    awsAccountId                *string
    clusterIdentifier           *string
    subResourceName             *string
    privateDnsRegions           *string
    privateConnectionServiceId  *string
}

type privateLinkConfigRequest struct {
    ConnectionServiceName       *string `json:"connection_service_name,omitempty"`
    ConnectionServiceId         *string `json:"connection_service_id,omitempty"`
    WorkspaceUrl                *string `json:"workspace_url,omitempty"`
    AccountName                 *string `json:"account_name,omitempty"`
    AccountUrl                  *string `json:"account_url,omitempty"`
    VpceId                      *string `json:"vpce_id,omitempty"`
    PlsId                       *string `json:"pls_id,omitempty"`
    AwsAccountId                *string `json:"aws_account_id,omitempty"`
    ClusterIdentifier           *string `json:"cluster_identifier,omitempty"`
    SubResourceName             *string `json:"sub_resource_name,omitempty"`
    PrivateDnsRegions           *string `json:"private_dns_regions,omitempty"`
    PrivateConnectionServiceId  *string `json:"private_connection_service_id,omitempty"`
}
 
type PrivateLinkConfigResponse struct {
    ConnectionServiceName       string `json:"connection_service_name"`
    ConnectionServiceId         string `json:"connection_service_id"`
    WorkspaceUrl                string `json:"workspace_url"`
    AccountName                 string `json:"account_name"`
    AccountUrl                  string `json:"account_url"`
    VpceId                      string `json:"vpce_id"`
    PlsId                       string `json:"pls_id"`
    AwsAccountId                string `json:"aws_account_id"`
    ClusterIdentifier           string `json:"cluster_identifier"`
    SubResourceName             string `json:"sub_resource_name"`
    PrivateDnsRegions           string `json:"private_dns_regions"`
    PrivateConnectionServiceId  string `json:"private_connection_service_id,omitempty"`

}

func (plc *PrivateLinkConfig) Request() *privateLinkConfigRequest {
    return &privateLinkConfigRequest{
        ConnectionServiceName:              plc.connectionServiceName,
        ConnectionServiceId:                plc.connectionServiceId,
        WorkspaceUrl:                       plc.workspaceUrl,
        AccountName:                        plc.accountName,
        AccountUrl:                         plc.accountUrl,
        VpceId:                             plc.vpceId,
        PlsId:                              plc.plsId,
        AwsAccountId:                       plc.awsAccountId,
        ClusterIdentifier:                  plc.clusterIdentifier,
        SubResourceName:                    plc.subResourceName,
        PrivateDnsRegions:                  plc.privateDnsRegions,
        PrivateConnectionServiceId:         plc.privateConnectionServiceId,
    }
}

func (plc *PrivateLinkConfig) Merge(customConfig *map[string]interface{}) (*map[string]interface{}, error) {
    err := utils.MergeIntoMap(plc.Request(), customConfig)
    if err != nil {
        return nil, err
    }
    return customConfig, nil
}

func (plc *PrivateLinkConfig) ConnectionServiceName(value string) *PrivateLinkConfig {
    plc.connectionServiceName = &value
    return plc
}

func (plc *PrivateLinkConfig) ConnectionServiceId(value string) *PrivateLinkConfig {
    plc.connectionServiceId = &value
    return plc
}

func (plc *PrivateLinkConfig) WorkspaceUrl(value string) *PrivateLinkConfig {
    plc.workspaceUrl = &value
    return plc
}

func (plc *PrivateLinkConfig) AccountName(value string) *PrivateLinkConfig {
    plc.accountName = &value
    return plc
}

func (plc *PrivateLinkConfig) AccountUrl(value string) *PrivateLinkConfig {
    plc.accountUrl = &value
    return plc
}

func (plc *PrivateLinkConfig) VpceId(value string) *PrivateLinkConfig {
    plc.vpceId = &value
    return plc
}

func (plc *PrivateLinkConfig) PlsId(value string) *PrivateLinkConfig {
    plc.plsId = &value
    return plc
}

func (plc *PrivateLinkConfig) AwsAccountId(value string) *PrivateLinkConfig {
    plc.awsAccountId = &value
    return plc
}

func (plc *PrivateLinkConfig) ClusterIdentifier(value string) *PrivateLinkConfig {
    plc.clusterIdentifier = &value
    return plc
}

func (plc *PrivateLinkConfig) SubResourceName(value string) *PrivateLinkConfig {
    plc.subResourceName = &value
    return plc
}

func (plc *PrivateLinkConfig) PrivateDnsRegions(value string) *PrivateLinkConfig {
    plc.privateDnsRegions = &value
    return plc
}

func (plc *PrivateLinkConfig) PrivateConnectionServiceId(value string) *PrivateLinkConfig {
    plc.privateConnectionServiceId = &value
    return plc
}

