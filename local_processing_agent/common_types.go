package localprocessingagent

type localProcessingAgentCreateRequest struct {
    GroupId       *string `json:"group_id,omitempty"`
    DisplayName   *string `json:"display_name,omitempty"`
    EnvType       *string `json:"env_type,omitempty"`
    AcceptTerms   *bool   `json:"accept_terms,omitempty"`
}

type LocalProcessingAgentDetails struct {
    Id              string `json:"id"`
    DisplayName     string `json:"display_name"`
    GroupId         string `json:"group_id"`
    RegisteredAt    string `json:"registered_at"`
}

type LocalProcessingAgentUsageDetails struct {
    ConnectionId   string `json:"connection_id"`
    Schema         string `json:"schema"`
    Service        string `json:"service"`
}

type LocalProcessingAgentData struct {
    LocalProcessingAgentDetails
    Usage []LocalProcessingAgentUsageDetails `json:"usage"`
}

type LocalProcessingAgentCreateResponse struct {
    Code    string `json:"code"`
    Message string `json:"message"`
    Data    struct {
        LocalProcessingAgentDetails
        Files struct {
          ConfigJson        string `json:"config_json"`
          AuthJson          string `json:"auth_json"`
          DockerComposeYaml string `json:"docker_compose_yaml"`
        } `json:"files"`
    } `json:"data"`
}

type LocalProcessingAgentDetailsResponse struct {
    Code string                     `json:"code"`
    Data LocalProcessingAgentData   `json:"data"`
}

type LocalProcessingAgentListResponse struct {
    Code string `json:"code"`
    Data struct {
        Items      []LocalProcessingAgentData `json:"items"`
        NextCursor string                     `json:"next_cursor"`
    } `json:"data"`
}


type externalLoggingCreateRequestBase struct {
    Id      *string `json:"id,omitempty"`
    GroupId *string `json:"group_id,omitempty"`
    Service *string `json:"service,omitempty"`
    Enabled *bool   `json:"enabled,omitempty"`
}