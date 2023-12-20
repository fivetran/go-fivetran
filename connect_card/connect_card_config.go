package connectcard

type ConnectCardConfig struct {
    redirectUri        *string
    hideSetupGuide     *bool
}

type connectCardConfigRequest struct {
    RedirectUri      *string `json:"redirect_uri,omitempty"`
    HideSetupGuide   *bool   `json:"hide_setup_guide,omitempty"`
}

type ConnectCardConfigResponse struct {
    RedirectUri       string `json:"redirect_uri"`
    HideSetupGuide    bool   `json:"hide_setup_guide"`
}

func (s *ConnectCardConfig) request() *connectCardConfigRequest {
	return &connectCardConfigRequest{
        RedirectUri:         s.redirectUri,
        HideSetupGuide:      s.hideSetupGuide,
	}
}

func (s *ConnectCardConfig) RedirectUri(value string) *ConnectCardConfig {
    s.redirectUri = &value
    return s
}

func (s *ConnectCardConfig) HideSetupGuide(value bool) *ConnectCardConfig {
    s.hideSetupGuide = &value
    return s
}
