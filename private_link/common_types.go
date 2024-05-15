package privatelink

type PrivateLinkResponseBase struct {
    Id              string `json:"id"`
    Name            string `json:"name"`
    Region          string `json:"region"`
    Service         string `json:"service"`
    AccountId       string `json:"account_id"`
    CloudProvider   string `json:"cloud_provider"`
    State           string `json:"state"`
    StateSummary    string `json:"state_summary"`
    CreatedAt       string `json:"created_at"`
    CreatedBy       string `json:"created_by"`
}

type PrivateLinkResponse struct {
    Code    string `json:"code"`
    Message string `json:"message"`
    Data    struct {
        PrivateLinkResponseBase
        Config PrivateLinkConfigResponse `json:"config"`
    } `json:"data"`
}

type PrivateLinkListResponse struct {
    Code string `json:"code"`
    Data struct {
        Items      []PrivateLinkResponseBase `json:"items"`
        NextCursor string      `json:"next_cursor"`
    } `json:"data"`
}

type privateLinkCreateRequest struct {
    Name        *string `json:"name,omitempty"`
    Region      *string `json:"region,omitempty"`
    Service     *string `json:"service,omitempty"`
    Config      any     `json:"config,omitempty"`
}

type privateLinkModifyRequest struct {
    Config          any     `json:"config,omitempty"`
}