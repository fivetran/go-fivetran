package privatelinks

type PrivateLinksResponseBase struct {
	Id      		string `json:"id"`
	Name 			string `json:"name"`
	GroupId 		string `json:"group_id"`
	CloudProvider 	string `json:"cloud_provider"`
	Service 		string `json:"service"`
	Region 			string `json:"region"`
	State 			string `json:"state"`
	StateSummary 	string `json:"state_summary"`
	CreatedAt 		string `json:"created_at"`
	CreatedBy 		string `json:"created_by"`
}

type PrivateLinksResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		PrivateLinksResponseBase
		Config PrivateLinksConfigResponse `json:"config"`
	} `json:"data"`
}

type privateLinksCreateRequest struct {
	Name 		*string `json:"name,omitempty"`
	Service 	*string `json:"service,omitempty"`
	GroupId 	*string `json:"group_id,omitempty"`
	Config 		any     `json:"config,omitempty"`
}

type privateLinksModifyRequest struct {
	Config 			any     `json:"config,omitempty"`
}