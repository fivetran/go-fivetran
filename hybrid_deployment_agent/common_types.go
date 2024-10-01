package hybriddeploymentagent

type hybridDeploymentAgentCreateRequest struct {
    GroupId       *string `json:"group_id,omitempty"`
    DisplayName   *string `json:"display_name,omitempty"`
    EnvType       *string `json:"env_type,omitempty"`
    AcceptTerms   *bool   `json:"accept_terms,omitempty"`
}

type HybridDeploymentAgentDetails struct {
    Id              string `json:"id"`
    DisplayName     string `json:"display_name"`
    GroupId         string `json:"group_id"`
    RegisteredAt    string `json:"registered_at"`
}

type HybridDeploymentAgentUsageDetails struct {
    ConnectionId   string `json:"connection_id"`
    Schema         string `json:"schema"`
    Service        string `json:"service"`
}

type HybridDeploymentAgentData struct {
    HybridDeploymentAgentDetails
    Usage []HybridDeploymentAgentUsageDetails `json:"usage"`
}

type HybridDeploymentAgentCreateResponse struct {
    Code    string `json:"code"`
    Message string `json:"message"`
    Data    struct {
        HybridDeploymentAgentDetails
        Files struct {
          ConfigJson        string `json:"config_json"`
          AuthJson          string `json:"auth_json"`
          DockerComposeYaml string `json:"docker_compose_yaml"`
        } `json:"files"`
    } `json:"data"`
}

type HybridDeploymentAgentDetailsResponse struct {
    Code string                     `json:"code"`
    Data HybridDeploymentAgentData   `json:"data"`
}

type HybridDeploymentAgentListResponse struct {
    Code string `json:"code"`
    Data struct {
        Items      []HybridDeploymentAgentData `json:"items"`
        NextCursor string                     `json:"next_cursor"`
    } `json:"data"`
}