package fivetran

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/fivetran/go-fivetran/certificates"
	connectcard "github.com/fivetran/go-fivetran/connect_card"
	"github.com/fivetran/go-fivetran/connections"
	"github.com/fivetran/go-fivetran/destinations"
	externallogging "github.com/fivetran/go-fivetran/external_logging"
	"github.com/fivetran/go-fivetran/fingerprints"
	"github.com/fivetran/go-fivetran/groups"
	httputils "github.com/fivetran/go-fivetran/http_utils"
	hybriddeploymentagent "github.com/fivetran/go-fivetran/hybrid_deployment_agent"
	"github.com/fivetran/go-fivetran/metadata"
	privatelink "github.com/fivetran/go-fivetran/private_link"
	"github.com/fivetran/go-fivetran/proxy"
	"github.com/fivetran/go-fivetran/roles"
	"github.com/fivetran/go-fivetran/teams"
	"github.com/fivetran/go-fivetran/transformations"
	"github.com/fivetran/go-fivetran/users"
	"github.com/fivetran/go-fivetran/webhooks"
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
const defaultUserAgent = "Go-Fivetran/1.2.3"

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

func (c *Client) NewConnectionSync() *connections.ConnectionSyncService {
	return &connections.ConnectionSyncService{
		HttpService: c.NewHttpService(),
	}
}

func (c *Client) NewCertificateConnectionCertificateApprove() *certificates.ConnectionCertificateApproveService {
	return &certificates.ConnectionCertificateApproveService{
		HttpService: c.NewHttpService(),
	}
}

func (c *Client) NewCertificateDestinationCertificateApprove() *certificates.DestinationCertificateApproveService {
	return &certificates.DestinationCertificateApproveService{
		HttpService: c.NewHttpService(),
	}
}

func (c *Client) NewCertificateConnectionFingerprintApprove() *fingerprints.ConnectionFingerprintApproveService {
	return &fingerprints.ConnectionFingerprintApproveService{
		HttpService: c.NewHttpService(),
	}
}

func (c *Client) NewCertificateDestinationFingerprintApprove() *fingerprints.DestinationFingerprintApproveService {
	return &fingerprints.DestinationFingerprintApproveService{
		HttpService: c.NewHttpService(),
	}
}

func (c *Client) NewConnectionCertificateRevoke() *certificates.ConnectionCertificateRevokeService {
	return &certificates.ConnectionCertificateRevokeService{
		HttpService: c.NewHttpService(),
	}
}

func (c *Client) NewDestinationCertificateRevoke() *certificates.DestinationCertificateRevokeService {
	return &certificates.DestinationCertificateRevokeService{
		HttpService: c.NewHttpService(),
	}
}

func (c *Client) NewConnectionCertificatesList() *certificates.ConnectionCertificatesListService {
	return &certificates.ConnectionCertificatesListService{
		HttpService: c.NewHttpService(),
	}
}

func (c *Client) NewDestinationCertificatesList() *certificates.DestinationCertificatesListService {
	return &certificates.DestinationCertificatesListService{
		HttpService: c.NewHttpService(),
	}
}

func (c *Client) NewConnectionCertificateDetails() *certificates.ConnectionCertificateDetailsService {
	return &certificates.ConnectionCertificateDetailsService{
		HttpService: c.NewHttpService(),
	}
}

func (c *Client) NewDestinationCertificateDetails() *certificates.DestinationCertificateDetailsService {
	return &certificates.DestinationCertificateDetailsService{
		HttpService: c.NewHttpService(),
	}
}

func (c *Client) NewConnectionFingerprintRevoke() *fingerprints.ConnectionFingerprintRevokeService {
	return &fingerprints.ConnectionFingerprintRevokeService{
		HttpService: c.NewHttpService(),
	}
}

func (c *Client) NewDestinationFingerprintRevoke() *fingerprints.DestinationFingerprintRevokeService {
	return &fingerprints.DestinationFingerprintRevokeService{
		HttpService: c.NewHttpService(),
	}
}

func (c *Client) NewConnectionFingerprintsList() *fingerprints.ConnectionFingerprintsListService {
	return &fingerprints.ConnectionFingerprintsListService{
		HttpService: c.NewHttpService(),
	}
}

func (c *Client) NewDestinationFingerprintsList() *fingerprints.DestinationFingerprintsListService {
	return &fingerprints.DestinationFingerprintsListService{
		HttpService: c.NewHttpService(),
	}
}

func (c *Client) NewConnectionFingerprintDetails() *fingerprints.ConnectionFingerprintDetailsService {
	return &fingerprints.ConnectionFingerprintDetailsService{
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

func (c *Client) NewGroupUpdate() *groups.GroupUpdateService {
	return &groups.GroupUpdateService{HttpService: c.NewHttpService()}
}

func (c *Client) NewGroupListConnections() *groups.GroupListConnectionsService {
	return &groups.GroupListConnectionsService{HttpService: c.NewHttpService()}
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

/* External Logging */
func (c *Client) NewExternalLoggingCreate() *externallogging.ExternalLoggingCreateService {
	return &externallogging.ExternalLoggingCreateService{HttpService: c.NewHttpService()}
}

func (c *Client) NewExternalLoggingDelete() *externallogging.ExternalLoggingDeleteService {
	return &externallogging.ExternalLoggingDeleteService{HttpService: c.NewHttpService()}
}

func (c *Client) NewExternalLoggingDetails() *externallogging.ExternalLoggingDetailsService {
	return &externallogging.ExternalLoggingDetailsService{HttpService: c.NewHttpService()}
}

func (c *Client) NewExternalLoggingUpdate() *externallogging.ExternalLoggingUpdateService {
	return &externallogging.ExternalLoggingUpdateService{HttpService: c.NewHttpService()}
}

func (c *Client) NewExternalLoggingSetupTests() *externallogging.ExternalLoggingSetupTestsService {
	return &externallogging.ExternalLoggingSetupTestsService{HttpService: c.NewHttpService()}
}

func (c *Client) NewExternalLoggingList() *externallogging.ExternalLoggingListService {
	return &externallogging.ExternalLoggingListService{HttpService: c.NewHttpService()}
}

/* Destinations */
func (c *Client) NewDestinationCreate() *destinations.DestinationCreateService {
	http := c.NewHttpService()
	http.CommonHeaders["Accept"] = restAPIv2
	return &destinations.DestinationCreateService{HttpService: http}
}

func (c *Client) NewDestinationDelete() *destinations.DestinationDeleteService {
	return &destinations.DestinationDeleteService{HttpService: c.NewHttpService()}
}

func (c *Client) NewDestinationDetails() *destinations.DestinationDetailsService {
	http := c.NewHttpService()
	http.CommonHeaders["Accept"] = restAPIv2
	return &destinations.DestinationDetailsService{HttpService: http}
}

func (c *Client) NewDestinationUpdate() *destinations.DestinationUpdateService {
	http := c.NewHttpService()
	http.CommonHeaders["Accept"] = restAPIv2
	return &destinations.DestinationUpdateService{HttpService: http}
}

func (c *Client) NewDestinationSetupTests() *destinations.DestinationSetupTestsService {
	http := c.NewHttpService()
	http.CommonHeaders["Accept"] = restAPIv2
	return &destinations.DestinationSetupTestsService{HttpService: http}
}

func (c *Client) NewDestinationsList() *destinations.DestinationsListService {
	return &destinations.DestinationsListService{HttpService: c.NewHttpService()}
}

/* Users */
func (c *Client) NewUsersList() *users.UsersListService {
	return &users.UsersListService{HttpService: c.NewHttpService()}
}

func (c *Client) NewUserDetails() *users.UserDetailsService {
	return &users.UserDetailsService{HttpService: c.NewHttpService()}
}

func (c *Client) NewUserUpdate() *users.UserUpdateService {
	return &users.UserUpdateService{HttpService: c.NewHttpService()}
}

func (c *Client) NewUserInvite() *users.UserInviteService {
	return &users.UserInviteService{HttpService: c.NewHttpService()}
}

func (c *Client) NewUserDelete() *users.UserDeleteService {
	return &users.UserDeleteService{HttpService: c.NewHttpService()}
}

func (c *Client) NewUserGroupMembershipCreate() *users.UserGroupMembershipCreateService {
	return &users.UserGroupMembershipCreateService{HttpService: c.NewHttpService()}
}

func (c *Client) NewUserGroupMembershipDelete() *users.UserGroupMembershipDeleteService {
	return &users.UserGroupMembershipDeleteService{HttpService: c.NewHttpService()}
}

func (c *Client) NewUserGroupMembershipDetails() *users.UserGroupMembershipDetailsService {
	return &users.UserGroupMembershipDetailsService{HttpService: c.NewHttpService()}
}

func (c *Client) NewUserGroupMembershipsList() *users.UserGroupMembershipsListService {
	return &users.UserGroupMembershipsListService{HttpService: c.NewHttpService()}
}

func (c *Client) NewUserGroupMembershipUpdate() *users.UserGroupMembershipUpdateService {
	return &users.UserGroupMembershipUpdateService{HttpService: c.NewHttpService()}
}

func (c *Client) NewUserConnectionMembershipsList() *users.UserConnectionMembershipsListService {
	return &users.UserConnectionMembershipsListService{HttpService: c.NewHttpService()}
}

func (c *Client) NewUserConnectionMembershipUpdate() *users.UserConnectionMembershipUpdateService {
	return &users.UserConnectionMembershipUpdateService{HttpService: c.NewHttpService()}
}

func (c *Client) NewUserConnectionMembershipCreate() *users.UserConnectionMembershipCreateService {
	return &users.UserConnectionMembershipCreateService{HttpService: c.NewHttpService()}
}

func (c *Client) NewUserConnectionMembershipDelete() *users.UserConnectionMembershipDeleteService {
	return &users.UserConnectionMembershipDeleteService{HttpService: c.NewHttpService()}
}

func (c *Client) NewUserConnectionMembershipDetails() *users.UserConnectionMembershipDetailsService {
	return &users.UserConnectionMembershipDetailsService{HttpService: c.NewHttpService()}
}

/* Webhooks */
func (c *Client) NewWebhookDelete() *webhooks.WebhookDeleteService {
	return &webhooks.WebhookDeleteService{HttpService: c.NewHttpService()}
}

func (c *Client) NewWebhookUpdate() *webhooks.WebhookUpdateService {
	return &webhooks.WebhookUpdateService{HttpService: c.NewHttpService()}
}

func (c *Client) NewWebhookGroupCreate() *webhooks.WebhookGroupCreateService {
	return &webhooks.WebhookGroupCreateService{HttpService: c.NewHttpService()}
}

func (c *Client) NewWebhookAccountCreate() *webhooks.WebhookAccountCreateService {
	return &webhooks.WebhookAccountCreateService{HttpService: c.NewHttpService()}
}

func (c *Client) NewWebhookList() *webhooks.WebhookListService {
	return &webhooks.WebhookListService{HttpService: c.NewHttpService()}
}

func (c *Client) NewWebhookDetails() *webhooks.WebhookDetailsService {
	return &webhooks.WebhookDetailsService{HttpService: c.NewHttpService()}
}

func (c *Client) NewWebhookTest() *webhooks.WebhookTestService {
	return &webhooks.WebhookTestService{HttpService: c.NewHttpService()}
}

/* Teams */
func (c *Client) NewTeamsCreate() *teams.TeamsCreateService {
	return &teams.TeamsCreateService{HttpService: c.NewHttpService()}
}

func (c *Client) NewTeamsDelete() *teams.TeamsDeleteService {
	return &teams.TeamsDeleteService{HttpService: c.NewHttpService()}
}

func (c *Client) NewTeamsDetails() *teams.TeamsDetailsService {
	return &teams.TeamsDetailsService{HttpService: c.NewHttpService()}
}

func (c *Client) NewTeamsList() *teams.TeamsListService {
	return &teams.TeamsListService{HttpService: c.NewHttpService()}
}

func (c *Client) NewTeamsUpdate() *teams.TeamsUpdateService {
	return &teams.TeamsUpdateService{HttpService: c.NewHttpService()}
}

func (c *Client) NewTeamConnectionMembershipCreate() *teams.TeamConnectionMembershipCreateService {
	return &teams.TeamConnectionMembershipCreateService{HttpService: c.NewHttpService()}
}

func (c *Client) NewTeamConnectionMembershipDelete() *teams.TeamConnectionMembershipDeleteService {
	return &teams.TeamConnectionMembershipDeleteService{HttpService: c.NewHttpService()}
}

func (c *Client) NewTeamConnectionMembershipDetails() *teams.TeamConnectionMembershipDetailsService {
	return &teams.TeamConnectionMembershipDetailsService{HttpService: c.NewHttpService()}
}

func (c *Client) NewTeamConnectionMembershipsList() *teams.TeamConnectionMembershipsListService {
	return &teams.TeamConnectionMembershipsListService{HttpService: c.NewHttpService()}
}

func (c *Client) NewTeamConnectionMembershipUpdate() *teams.TeamConnectionMembershipUpdateService {
	return &teams.TeamConnectionMembershipUpdateService{HttpService: c.NewHttpService()}
}

func (c *Client) NewTeamGroupMembershipCreate() *teams.TeamGroupMembershipCreateService {
	return &teams.TeamGroupMembershipCreateService{HttpService: c.NewHttpService()}
}

func (c *Client) NewTeamGroupMembershipDelete() *teams.TeamGroupMembershipDeleteService {
	return &teams.TeamGroupMembershipDeleteService{HttpService: c.NewHttpService()}
}

func (c *Client) NewTeamGroupMembershipDetails() *teams.TeamGroupMembershipDetailsService {
	return &teams.TeamGroupMembershipDetailsService{HttpService: c.NewHttpService()}
}

func (c *Client) NewTeamGroupMembershipsList() *teams.TeamGroupMembershipsListService {
	return &teams.TeamGroupMembershipsListService{HttpService: c.NewHttpService()}
}

func (c *Client) NewTeamGroupMembershipUpdate() *teams.TeamGroupMembershipUpdateService {
	return &teams.TeamGroupMembershipUpdateService{HttpService: c.NewHttpService()}
}

func (c *Client) NewTeamUserMembershipCreate() *teams.TeamUserMembershipCreateService {
	return &teams.TeamUserMembershipCreateService{HttpService: c.NewHttpService()}
}

func (c *Client) NewTeamUserMembershipDelete() *teams.TeamUserMembershipDeleteService {
	return &teams.TeamUserMembershipDeleteService{HttpService: c.NewHttpService()}
}

func (c *Client) NewTeamUserMembershipDetails() *teams.TeamUserMembershipDetailsService {
	return &teams.TeamUserMembershipDetailsService{HttpService: c.NewHttpService()}
}

func (c *Client) NewTeamUserMembershipsList() *teams.TeamUserMembershipsListService {
	return &teams.TeamUserMembershipsListService{HttpService: c.NewHttpService()}
}

func (c *Client) NewTeamUserMembershipUpdate() *teams.TeamUserMembershipUpdateService {
	return &teams.TeamUserMembershipUpdateService{HttpService: c.NewHttpService()}
}

func (c *Client) NewTeamsDeleteRoleInAccount() *teams.TeamsDeleteRoleInAccountService {
	return &teams.TeamsDeleteRoleInAccountService{HttpService: c.NewHttpService()}
}

func (c *Client) NewRolesList() *roles.RolesListService {
	return &roles.RolesListService{HttpService: c.NewHttpService()}
}

func (c *Client) NewConnectCard() *connectcard.ConnectCardService {
	return &connectcard.ConnectCardService{HttpService: c.NewHttpService()}
}

/* Connections */
func (c *Client) NewConnectionCreate() *connections.ConnectionCreateService {
	return &connections.ConnectionCreateService{HttpService: c.NewHttpService()}
}

func (c *Client) NewConnectionDetails() *connections.ConnectionDetailsService {
	http := c.NewHttpService()
	http.CommonHeaders["Accept"] = restAPIv2
	return &connections.ConnectionDetailsService{HttpService: http}
}

func (c *Client) NewConnectionUpdate() *connections.ConnectionUpdateService {
	return &connections.ConnectionUpdateService{HttpService: c.NewHttpService()}
}

func (c *Client) NewConnectionDelete() *connections.ConnectionDeleteService {
	return &connections.ConnectionDeleteService{HttpService: c.NewHttpService()}
}

func (c *Client) NewConnectionSetupTests() *connections.ConnectionSetupTestsService {
	return &connections.ConnectionSetupTestsService{HttpService: c.NewHttpService()}
}

func (c *Client) NewConnectionColumnConfigListService() *connections.ConnectionColumnConfigListService {
	return &connections.ConnectionColumnConfigListService{HttpService: c.NewHttpService()}
}

func (c *Client) NewConnectionColumnConfigUpdateService() *connections.ConnectionColumnConfigUpdateService {
	return &connections.ConnectionColumnConfigUpdateService{HttpService: c.NewHttpService()}
}

func (c *Client) NewConnectionDatabaseSchemaConfigUpdateService() *connections.ConnectionDatabaseSchemaConfigUpdateService {
	return &connections.ConnectionDatabaseSchemaConfigUpdateService{HttpService: c.NewHttpService()}
}

func (c *Client) NewConnectionReSyncTable() *connections.ConnectionReSyncTableService {
	return &connections.ConnectionReSyncTableService{HttpService: c.NewHttpService()}
}

func (c *Client) NewConnectionSchemaDetails() *connections.ConnectionSchemaDetailsService {
	return &connections.ConnectionSchemaDetailsService{HttpService: c.NewHttpService()}
}

func (c *Client) NewConnectionSchemaReload() *connections.ConnectionSchemaReloadService {
	return &connections.ConnectionSchemaReloadService{HttpService: c.NewHttpService()}
}

func (c *Client) NewConnectionSchemaUpdateService() *connections.ConnectionSchemaConfigUpdateService {
	return &connections.ConnectionSchemaConfigUpdateService{HttpService: c.NewHttpService()}
}

func (c *Client) NewConnectionSchemaCreateService() *connections.ConnectionSchemaConfigCreateService {
	return &connections.ConnectionSchemaConfigCreateService{HttpService: c.NewHttpService()}
}

func (c *Client) NewConnectionTableConfigUpdateService() *connections.ConnectionTableConfigUpdateService {
	return &connections.ConnectionTableConfigUpdateService{HttpService: c.NewHttpService()}
}

func (c *Client) NewConnectionsList() *connections.ConnectionsListService {
	return &connections.ConnectionsListService{HttpService: c.NewHttpService()}
}

/* Private Links */
func (c *Client) NewPrivateLinkCreate() *privatelink.PrivateLinkCreateService {
	return &privatelink.PrivateLinkCreateService{HttpService: c.NewHttpService()}
}

func (c *Client) NewPrivateLinkDelete() *privatelink.PrivateLinkDeleteService {
	return &privatelink.PrivateLinkDeleteService{HttpService: c.NewHttpService()}
}

func (c *Client) NewPrivateLinkList() *privatelink.PrivateLinkListService {
	return &privatelink.PrivateLinkListService{HttpService: c.NewHttpService()}
}

func (c *Client) NewPrivateLinkDetails() *privatelink.PrivateLinkDetailsService {
	return &privatelink.PrivateLinkDetailsService{HttpService: c.NewHttpService()}
}

func (c *Client) NewPrivateLinkUpdate() *privatelink.PrivateLinkUpdateService {
	return &privatelink.PrivateLinkUpdateService{HttpService: c.NewHttpService()}
}

/* Proxy */
func (c *Client) NewProxyCreate() *proxy.ProxyCreateService {
	return &proxy.ProxyCreateService{HttpService: c.NewHttpService()}
}

func (c *Client) NewProxyList() *proxy.ProxyListService {
	return &proxy.ProxyListService{HttpService: c.NewHttpService()}
}

func (c *Client) NewProxyDetails() *proxy.ProxyDetailsService {
	return &proxy.ProxyDetailsService{HttpService: c.NewHttpService()}
}

func (c *Client) NewProxyDelete() *proxy.ProxyDeleteService {
	return &proxy.ProxyDeleteService{HttpService: c.NewHttpService()}
}

func (c *Client) NewProxyRegenerateSecrets() *proxy.ProxyRegenerateSecretsService {
	return &proxy.ProxyRegenerateSecretsService{HttpService: c.NewHttpService()}
}

func (c *Client) NewProxyConnectionMembershipsList() *proxy.ProxyConnectionMembershipsListService {
	return &proxy.ProxyConnectionMembershipsListService{HttpService: c.NewHttpService()}
}

/* Hybrid Deployment Agent */
func (c *Client) NewHybridDeploymentAgentCreate() *hybriddeploymentagent.HybridDeploymentAgentCreateService {
	return &hybriddeploymentagent.HybridDeploymentAgentCreateService{HttpService: c.NewHttpService()}
}

func (c *Client) NewHybridDeploymentAgentDelete() *hybriddeploymentagent.HybridDeploymentAgentDeleteService {
	return &hybriddeploymentagent.HybridDeploymentAgentDeleteService{HttpService: c.NewHttpService()}
}

func (c *Client) NewHybridDeploymentAgentDetails() *hybriddeploymentagent.HybridDeploymentAgentDetailsService {
	return &hybriddeploymentagent.HybridDeploymentAgentDetailsService{HttpService: c.NewHttpService()}
}

func (c *Client) NewHybridDeploymentAgentList() *hybriddeploymentagent.HybridDeploymentAgentListService {
	return &hybriddeploymentagent.HybridDeploymentAgentListService{HttpService: c.NewHttpService()}
}

func (c *Client) NewHybridDeploymentAgentReAuth() *hybriddeploymentagent.HybridDeploymentAgentReAuthService {
	return &hybriddeploymentagent.HybridDeploymentAgentReAuthService{HttpService: c.NewHttpService()}
}

func (c *Client) NewHybridDeploymentAgentResetCredentials() *hybriddeploymentagent.HybridDeploymentAgentResetCredentialsService {
	return &hybriddeploymentagent.HybridDeploymentAgentResetCredentialsService{HttpService: c.NewHttpService()}
}

/* Transformations */
func (c *Client) NewTransformationProjectDetails() *transformations.TransformationProjectDetailsService {
	return &transformations.TransformationProjectDetailsService{HttpService: c.NewHttpService()}
}

func (c *Client) NewTransformationProjectCreate() *transformations.TransformationProjectCreateService {
	return &transformations.TransformationProjectCreateService{HttpService: c.NewHttpService()}
}

func (c *Client) NewTransformationProjectsList() *transformations.TransformationProjectsListService {
	return &transformations.TransformationProjectsListService{HttpService: c.NewHttpService()}
}

func (c *Client) NewTransformationProjectDelete() *transformations.TransformationProjectDeleteService {
	return &transformations.TransformationProjectDeleteService{HttpService: c.NewHttpService()}
}

func (c *Client) NewTransformationProjectUpdate() *transformations.TransformationProjectUpdateService {
	return &transformations.TransformationProjectUpdateService{HttpService: c.NewHttpService()}
}

func (c *Client) NewTransformationCreate() *transformations.TransformationCreateService {
	return &transformations.TransformationCreateService{HttpService: c.NewHttpService()}
}

func (c *Client) NewTransformationUpdate() *transformations.TransformationUpdateService {
	return &transformations.TransformationUpdateService{HttpService: c.NewHttpService()}
}

func (c *Client) NewTransformationDelete() *transformations.TransformationDeleteService {
	return &transformations.TransformationDeleteService{HttpService: c.NewHttpService()}
}

func (c *Client) NewTransformationsList() *transformations.TransformationsListService {
	return &transformations.TransformationsListService{HttpService: c.NewHttpService()}
}

func (c *Client) NewTransformationDetails() *transformations.TransformationDetailsService {
	return &transformations.TransformationDetailsService{HttpService: c.NewHttpService()}
}

func (c *Client) NewTransformationRun() *transformations.TransformationRunService {
	return &transformations.TransformationRunService{HttpService: c.NewHttpService()}
}

func (c *Client) NewTransformationCancel() *transformations.TransformationCancelService {
	return &transformations.TransformationCancelService{HttpService: c.NewHttpService()}
}

func (c *Client) NewTransformationUpgradePackage() *transformations.TransformationUpgradePackageService {
	return &transformations.TransformationUpgradePackageService{HttpService: c.NewHttpService()}
}

func (c *Client) NewQuickstartPackagesList() *transformations.QuickstartPackagesListService {
	return &transformations.QuickstartPackagesListService{HttpService: c.NewHttpService()}
}

func (c *Client) NewQuickstartPackageDetails() *transformations.QuickstartPackageDetailsService {
	return &transformations.QuickstartPackageDetailsService{HttpService: c.NewHttpService()}
}

func (c *Client) NewMetadataDetails() *metadata.MetadataDetailsService {
	return &metadata.MetadataDetailsService{HttpService: c.NewHttpService()}
}

func (c *Client) NewMetadataList() *metadata.MetadataListService {
	return &metadata.MetadataListService{HttpService: c.NewHttpService()}
}
