package externallogging

type ExternalLoggingResponseBase struct {
	Id      string `json:"id"`
	Service string `json:"service"`
	Enabled bool   `json:"enabled"`
}

type ExternalLoggingResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ExternalLoggingResponseBase
		Config ExternalLoggingConfigResponse `json:"config"`
	} `json:"data"`
}

type ExternalLoggingCustomResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ExternalLoggingResponseBase
		Config map[string]interface{} `json:"config"`
	} `json:"data"`
}

type ExternalLoggingCustomMergedResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ExternalLoggingResponseBase
		CustomConfig map[string]interface{}        `json:"config"`
		Config       ExternalLoggingConfigResponse // no mapping here
	} `json:"data"`
}
