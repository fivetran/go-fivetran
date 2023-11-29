package privatelinks

// PrivateLinksConfig builds Private Links Management, Private Link Config.
// Ref. https://fivetran.com/docs/rest-api/private-links-management#privatelinksetupconfigurations
type PrivateLinksConfig struct {
	connectionServiceName  	*string
	connectionServiceId   	*string
	workspaceUrl 			*string
	accountName      		*string
	accountUrl   			*string
	vpceId       			*string
	plsId       			*string
	awsAccountId    		*string
	clusterIdentifier       *string
	subResourceName     	*string
}

type privateLinksConfigRequest struct {
	ConnectionServiceName  	*string `json:"connection_service_name,omitempty"`
	ConnectionServiceId   	*string `json:"connection_service_id,omitempty"`
	WorkspaceUrl 			*string `json:"workspace_url,omitempty"`
	AccountName  			*string `json:"account_name,omitempty"`
	AccountUrl   			*string `json:"account_url,omitempty"`
	VpceId       			*string `json:"vpceId,omitempty"`
	PlsId       			*string `json:"pls_id,omitempty"`
	AwsAccountId    		*string `json:"aws_account_id,omitempty"`
	ClusterIdentifier       *string `json:"cluster_identifier,omitempty"`
	SubResourceName     	*string `json:"sub_resource_name,omitempty"`
}
 
type PrivateLinksConfigResponse struct {
	ConnectionServiceName 	string `json:"connection_service_name"`
	ConnectionServiceId   	string `json:"connection_service_id"`
	WorkspaceUrl 			string `json:"workspace_url"`
	AccountName      		string `json:"account_name"`
	AccountUrl   			string `json:"account_url"`
	VpceId       			string `json:"vpceId"`
	PlsId       			string `json:"pls_id"`
	AwsAccountId    		string `json:"aws_account_id"`
	ClusterIdentifier       string `json:"cluster_identifier"`
	SubResourceName     	string `json:"sub_resource_name"`
}

func (plc *PrivateLinksConfig) Request() *privateLinksConfigRequest {
	return &privateLinksConfigRequest{
		ConnectionServiceName:  plc.connectionServiceName,
		ConnectionServiceId:  	plc.connectionServiceId,
		WorkspaceUrl: 			plc.workspaceUrl,
		AccountName:      		plc.accountName,
		AccountUrl:   			plc.accountUrl,
		VpceId:      			plc.vpceId,
		PlsId:       			plc.plsId,
		AwsAccountId:    		plc.awsAccountId,
		ClusterIdentifier:      plc.clusterIdentifier,
		SubResourceName:     	plc.subResourceName,
	}
}

func (plc *PrivateLinksConfig) ConnectionServiceName(value string) *PrivateLinksConfig {
	plc.connectionServiceName = &value
	return plc
}

func (plc *PrivateLinksConfig) ConnectionServiceId(value string) *PrivateLinksConfig {
	plc.connectionServiceId = &value
	return plc
}

func (plc *PrivateLinksConfig) WorkspaceUrl(value string) *PrivateLinksConfig {
	plc.workspaceUrl = &value
	return plc
}

func (plc *PrivateLinksConfig) AccountName(value string) *PrivateLinksConfig {
	plc.accountName = &value
	return plc
}

func (plc *PrivateLinksConfig) AccountUrl(value string) *PrivateLinksConfig {
	plc.accountUrl = &value
	return plc
}

func (plc *PrivateLinksConfig) VpceId(value string) *PrivateLinksConfig {
	plc.vpceId = &value
	return plc
}

func (plc *PrivateLinksConfig) PlsId(value string) *PrivateLinksConfig {
	plc.plsId = &value
	return plc
}

func (plc *PrivateLinksConfig) AwsAccountId(value string) *PrivateLinksConfig {
	plc.awsAccountId = &value
	return plc
}

func (plc *PrivateLinksConfig) ClusterIdentifier(value string) *PrivateLinksConfig {
	plc.clusterIdentifier = &value
	return plc
}

func (plc *PrivateLinksConfig) SubResourceName(value string) *PrivateLinksConfig {
	plc.subResourceName = &value
	return plc
}
