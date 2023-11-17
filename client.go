package fivetran

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/fivetran/go-fivetran/certificates"
	"github.com/fivetran/go-fivetran/fingerprints"
	"github.com/fivetran/go-fivetran/groups"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// Client holds client configuration
type Client struct {
	baseURL          string
	authorization    string
	customUserAgent  string
	httpClient       httputils.HttpClient
	handleRateLimits bool
	maxRetryAttempts int
}

const defaultBaseURL = "https://api.fivetran.com/v1"
const restAPIv2 = "application/json;version=2"

// WARNING: Update Agent version on each release!
const defaultUserAgent = "Go-Fivetran/0.7.8"

// New receives API Key and API Secret, and returns a new Client with the
// default HTTP client
func New(apiKey, apiSecret string) *Client {
	credentials := fmt.Sprintf("Basic %v", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%v:%v", apiKey, apiSecret))))
	return &Client{
		baseURL:          defaultBaseURL,
		authorization:    credentials,
		httpClient:       &http.Client{},
		maxRetryAttempts: 2,
		handleRateLimits: true,
	}
}

// BaseURL changes Client base REST API endpoint URL
func (c *Client) BaseURL(baseURL string) {
	c.baseURL = baseURL
}

// CustomUserAgent sets custom User-Agent header in Client requests
func (c *Client) CustomUserAgent(customUserAgent string) {
	c.customUserAgent = customUserAgent
}

// SetHttpClient sets custom HTTP client to perform requests with
func (c *Client) SetHttpClient(httpClient httputils.HttpClient) {
	c.httpClient = httpClient
}

// SetHandleRateLimits sets custom HTTP client to handle rate limits automatically
func (c *Client) SetHandleRateLimits(handleRateLimits bool) {
	c.handleRateLimits = handleRateLimits
}

// SetMaxRetryAttempts sets custom HTTP client maximum retry attempts count
func (c *Client) SetMaxRetryAttempts(maxRetryAttempts int) {
	c.maxRetryAttempts = maxRetryAttempts
}

func (c *Client) NewHttpService() httputils.HttpService {
	return httputils.HttpService{
		CommonHeaders:    c.commonHeaders(),
		BaseUrl:          c.baseURL,
		MaxRetryAttempts: c.maxRetryAttempts,
		HandleRateLimits: c.handleRateLimits,
		Client:           c.httpClient,
	}
}

func (c *Client) commonHeaders() map[string]string {
	userAgent := defaultUserAgent

	if c.customUserAgent != "" {
		userAgent += " " + c.customUserAgent
	}

	return map[string]string{
		"Authorization": c.authorization,
		"User-Agent":    userAgent,
	}
}

func (c *Client) NewConnectorSync() *ConnectorSyncService {
	return &ConnectorSyncService{
		HttpService: c.NewHttpService(),
	}
}

func (c *Client) NewCertificateConnectorCertificateApprove() *certificates.ConnectorCertificateApproveService {
	return &certificates.ConnectorCertificateApproveService{
		HttpService: c.NewHttpService(),
	}
}

func (c *Client) NewCertificateDestinationCertificateApprove() *certificates.DestinationCertificateApproveService {
	return &certificates.DestinationCertificateApproveService{
		HttpService: c.NewHttpService(),
	}
}

func (c *Client) NewCertificateConnectorFingerprintApprove() *fingerprints.ConnectorFingerprintApproveService {
	return &fingerprints.ConnectorFingerprintApproveService{
		HttpService: c.NewHttpService(),
	}
}

func (c *Client) NewCertificateDestinationFingerprintApprove() *fingerprints.DestinationFingerprintApproveService {
	return &fingerprints.DestinationFingerprintApproveService{
		HttpService: c.NewHttpService(),
	}
}

func (c *Client) NewConnectorCertificateRevoke() *certificates.ConnectorCertificateRevokeService {
	return &certificates.ConnectorCertificateRevokeService{
		HttpService: c.NewHttpService(),
	}
}

func (c *Client) NewDestinationCertificateRevoke() *certificates.DestinationCertificateRevokeService {
	return &certificates.DestinationCertificateRevokeService{
		HttpService: c.NewHttpService(),
	}
}

func (c *Client) NewConnectorCertificatesList() *certificates.ConnectorCertificatesListService {
	return &certificates.ConnectorCertificatesListService{
		HttpService: c.NewHttpService(),
	}
}

func (c *Client) NewDestinationCertificatesList() *certificates.DestinationCertificatesListService {
	return &certificates.DestinationCertificatesListService{
		HttpService: c.NewHttpService(),
	}
}

func (c *Client) NewConnectorCertificateDetails() *certificates.ConnectorCertificateDetailsService {
	return &certificates.ConnectorCertificateDetailsService{
		HttpService: c.NewHttpService(),
	}
}

func (c *Client) NewDestinationCertificateDetails() *certificates.DestinationCertificateDetailsService {
	return &certificates.DestinationCertificateDetailsService{
		HttpService: c.NewHttpService(),
	}
}

func (c *Client) NewConnectorFingerprintRevoke() *fingerprints.ConnectorFingerprintRevokeService {
	return &fingerprints.ConnectorFingerprintRevokeService{
		HttpService: c.NewHttpService(),
	}
}

func (c *Client) NewDestinationFingerprintRevoke() *fingerprints.DestinationFingerprintRevokeService {
	return &fingerprints.DestinationFingerprintRevokeService{
		HttpService: c.NewHttpService(),
	}
}

func (c *Client) NewConnectorFingerprintsList() *fingerprints.ConnectorFingerprintsListService {
	return &fingerprints.ConnectorFingerprintsListService{
		HttpService: c.NewHttpService(),
	}
}

func (c *Client) NewDestinationFingerprintsList() *fingerprints.DestinationFingerprintsListService {
	return &fingerprints.DestinationFingerprintsListService{
		HttpService: c.NewHttpService(),
	}
}

func (c *Client) NewConnectorFingerprintDetails() *fingerprints.ConnectorFingerprintDetailsService {
	return &fingerprints.ConnectorFingerprintDetailsService{
		HttpService: c.NewHttpService(),
	}
}

func (c *Client) NewDestinationFingerprintDetails() *fingerprints.DestinationFingerprintDetailsService {
	return &fingerprints.DestinationFingerprintDetailsService{
		HttpService: c.NewHttpService(),
	}
}

func (c *Client) NewGroupCreate() *groups.GroupCreateService {
	return &groups.GroupCreateService{HttpService: c.NewHttpService()}
}

func (c *Client) NewGroupDetails() *groups.GroupDetailsService {
	return &groups.GroupDetailsService{HttpService: c.NewHttpService()}
}

func (c *Client) NewGroupAddUser() *groups.GroupAddUserService {
	return &groups.GroupAddUserService{HttpService: c.NewHttpService()}
}

func (c *Client) NewGroupRemoveUser() *groups.GroupRemoveUserService {
	return &groups.GroupRemoveUserService{HttpService: c.NewHttpService()}
}

func (c *Client) NewGroupDelete() *groups.GroupDeleteService {
	return &groups.GroupDeleteService{HttpService: c.NewHttpService()}
}

func (c *Client) NewGroupModify() *groups.GroupModifyService {
	return &groups.GroupModifyService{HttpService: c.NewHttpService()}
}

func (c *Client) NewGroupListConnectors() *groups.GroupListConnectorsService {
	return &groups.GroupListConnectorsService{HttpService: c.NewHttpService()}
}

func (c *Client) NewGroupListUsers() *groups.GroupListUsersService {
	return &groups.GroupListUsersService{HttpService: c.NewHttpService()}
}

func (c *Client) NewGroupsList() *groups.GroupsListService {
	return &groups.GroupsListService{HttpService: c.NewHttpService()}
}

func (c *Client) NewGroupSshPublicKey() *groups.GroupSshKeyService {
	return &groups.GroupSshKeyService{HttpService: c.NewHttpService()}
}

func (c *Client) NewGroupServiceAccount() *groups.GroupServiceAccountService {
	return &groups.GroupServiceAccountService{HttpService: c.NewHttpService()}
}
