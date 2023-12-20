package fivetran

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/fivetran/go-fivetran/connectors"
	"github.com/fivetran/go-fivetran/certificates"
	"github.com/fivetran/go-fivetran/fingerprints"
	"github.com/fivetran/go-fivetran/groups"
	"github.com/fivetran/go-fivetran/users"
	"github.com/fivetran/go-fivetran/external_logging"
	"github.com/fivetran/go-fivetran/destinations"
	"github.com/fivetran/go-fivetran/dbt"
	"github.com/fivetran/go-fivetran/webhooks"
	"github.com/fivetran/go-fivetran/teams"
	"github.com/fivetran/go-fivetran/roles"
	"github.com/fivetran/go-fivetran/connect_card"
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

func (c *Client) NewConnectorSync() *connectors.ConnectorSyncService {
	return &connectors.ConnectorSyncService{
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

func (c *Client) NewExternalLoggingModify() *externallogging.ExternalLoggingModifyService {
	return &externallogging.ExternalLoggingModifyService{HttpService: c.NewHttpService()}
}

func (c *Client) NewExternalLoggingSetupTests() *externallogging.ExternalLoggingSetupTestsService {
	return &externallogging.ExternalLoggingSetupTestsService{HttpService: c.NewHttpService()}
}

/* Destinations */
func (c *Client) NewDestinationCreate() *destinations.DestinationCreateService {
	return &destinations.DestinationCreateService{HttpService: c.NewHttpService()}
}

func (c *Client) NewDestinationDelete() *destinations.DestinationDeleteService {
	return &destinations.DestinationDeleteService{HttpService: c.NewHttpService()}
}

func (c *Client) NewDestinationDetails() *destinations.DestinationDetailsService {
	return &destinations.DestinationDetailsService{HttpService: c.NewHttpService()}
}

func (c *Client) NewDestinationModify() *destinations.DestinationModifyService {
	return &destinations.DestinationModifyService{HttpService: c.NewHttpService()}
}

func (c *Client) NewDestinationSetupTests() *destinations.DestinationSetupTestsService {
	return &destinations.DestinationSetupTestsService{HttpService: c.NewHttpService()}
}

/* Users */
func (c *Client) NewUsersList() *users.UsersListService {
	return &users.UsersListService{HttpService: c.NewHttpService()}
}

func (c *Client) NewUserDetails() *users.UserDetailsService {
	return &users.UserDetailsService{HttpService: c.NewHttpService()}
}

func (c *Client) NewUserModify() *users.UserModifyService {
	return &users.UserModifyService{HttpService: c.NewHttpService()}
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

func (c *Client) NewUserGroupMembershipModify() *users.UserGroupMembershipModifyService {
	return &users.UserGroupMembershipModifyService{HttpService: c.NewHttpService()}
}

func (c *Client) NewUserConnectorMembershipsList() *users.UserConnectorMembershipsListService {
	return &users.UserConnectorMembershipsListService{HttpService: c.NewHttpService()}
}

func (c *Client) NewUserConnectorMembershipModify() *users.UserConnectorMembershipModifyService {
	return &users.UserConnectorMembershipModifyService{HttpService: c.NewHttpService()}
}

func (c *Client) NewUserConnectorMembershipCreate() *users.UserConnectorMembershipCreateService {
	return &users.UserConnectorMembershipCreateService{HttpService: c.NewHttpService()}
}

func (c *Client) NewUserConnectorMembershipDelete() *users.UserConnectorMembershipDeleteService {
	return &users.UserConnectorMembershipDeleteService{HttpService: c.NewHttpService()}
}

func (c *Client) NewUserConnectorMembershipDetails() *users.UserConnectorMembershipDetailsService {
	return &users.UserConnectorMembershipDetailsService{HttpService: c.NewHttpService()}
}

/* DBT */
func (c *Client) NewDbtModelDetails() *dbt.DbtModelDetailsService {
	return &dbt.DbtModelDetailsService{HttpService: c.NewHttpService()}
}

func (c *Client) NewDbtModelsList() *dbt.DbtModelsListService {
	return &dbt.DbtModelsListService{HttpService: c.NewHttpService()}
}

func (c *Client) NewDbtProjectDetails() *dbt.DbtProjectDetailsService {
	return &dbt.DbtProjectDetailsService{HttpService: c.NewHttpService()}
}

func (c *Client) NewDbtProjectDelete() *dbt.DbtProjectDeleteService {
	return &dbt.DbtProjectDeleteService{HttpService: c.NewHttpService()}
}

func (c *Client) NewDbtProjectCreate() *dbt.DbtProjectCreateService {
	return &dbt.DbtProjectCreateService{HttpService: c.NewHttpService()}
}

func (c *Client) NewDbtProjectModify() *dbt.DbtProjectModifyService {
	return &dbt.DbtProjectModifyService{HttpService: c.NewHttpService()}
}

func (c *Client) NewDbtProjectsList() *dbt.DbtProjectsListService {
	return &dbt.DbtProjectsListService{HttpService: c.NewHttpService()}
}

func (c *Client) NewDbtTransformationCreateService() *dbt.DbtTransformationCreateService {
	return &dbt.DbtTransformationCreateService{HttpService: c.NewHttpService()}
}

func (c *Client) NewDbtTransformationDeleteService() *dbt.DbtTransformationDeleteService {
	return &dbt.DbtTransformationDeleteService{HttpService: c.NewHttpService()}
}

func (c *Client) NewDbtTransformationDetailsService() *dbt.DbtTransformationDetailsService {
	return &dbt.DbtTransformationDetailsService{HttpService: c.NewHttpService()}
}

func (c *Client) NewDbtTransformationModifyService() *dbt.DbtTransformationModifyService {
	return &dbt.DbtTransformationModifyService{HttpService: c.NewHttpService()}
}

/* Webhooks */
func (c *Client) NewWebhookDelete() *webhooks.WebhookDeleteService {
	return &webhooks.WebhookDeleteService{HttpService: c.NewHttpService()}
}

func (c *Client) NewWebhookModify() *webhooks.WebhookModifyService {
	return &webhooks.WebhookModifyService{HttpService: c.NewHttpService()}
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

func (c *Client) NewTeamsModify() *teams.TeamsModifyService {
	return &teams.TeamsModifyService{HttpService: c.NewHttpService()}
}

func (c *Client) NewTeamConnectorMembershipCreate() *teams.TeamConnectorMembershipCreateService {
	return &teams.TeamConnectorMembershipCreateService{HttpService: c.NewHttpService()}
}

func (c *Client) NewTeamConnectorMembershipDelete() *teams.TeamConnectorMembershipDeleteService {
	return &teams.TeamConnectorMembershipDeleteService{HttpService: c.NewHttpService()}
}

func (c *Client) NewTeamConnectorMembershipDetails() *teams.TeamConnectorMembershipDetailsService {
	return &teams.TeamConnectorMembershipDetailsService{HttpService: c.NewHttpService()}
}

func (c *Client) NewTeamConnectorMembershipsList() *teams.TeamConnectorMembershipsListService {
	return &teams.TeamConnectorMembershipsListService{HttpService: c.NewHttpService()}
}

func (c *Client) NewTeamConnectorMembershipModify() *teams.TeamConnectorMembershipModifyService {
	return &teams.TeamConnectorMembershipModifyService{HttpService: c.NewHttpService()}
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

func (c *Client) NewTeamGroupMembershipModify() *teams.TeamGroupMembershipModifyService {
	return &teams.TeamGroupMembershipModifyService{HttpService: c.NewHttpService()}
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

func (c *Client) NewTeamUserMembershipModify() *teams.TeamUserMembershipModifyService {
	return &teams.TeamUserMembershipModifyService{HttpService: c.NewHttpService()}
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

/* Connectors */
func (c *Client) NewConnectorCreate() *connectors.ConnectorCreateService {
	return &connectors.ConnectorCreateService{HttpService: c.NewHttpService()}
}

func (c *Client) NewConnectorDetails() *connectors.ConnectorDetailsService {
	return &connectors.ConnectorDetailsService{HttpService: c.NewHttpService()}
}

func (c *Client) NewConnectorModify() *connectors.ConnectorModifyService {
	return &connectors.ConnectorModifyService{HttpService: c.NewHttpService()}
}

func (c *Client) NewConnectorDelete() *connectors.ConnectorDeleteService {
	return &connectors.ConnectorDeleteService{HttpService: c.NewHttpService()}
}

func (c *Client) NewConnectorSetupTests() *connectors.ConnectorSetupTestsService {
	return &connectors.ConnectorSetupTestsService{HttpService: c.NewHttpService()}
}

func (c *Client) NewConnectorColumnConfigListService() *connectors.ConnectorColumnConfigListService {
    return &connectors.ConnectorColumnConfigListService{HttpService: c.NewHttpService()}
}

func (c *Client) NewConnectorColumnConfigModifyService() *connectors.ConnectorColumnConfigModifyService {
    return &connectors.ConnectorColumnConfigModifyService{HttpService: c.NewHttpService()}
}

func (c *Client) NewConnectorDatabaseSchemaConfigModifyService() *connectors.ConnectorDatabaseSchemaConfigModifyService {
    return &connectors.ConnectorDatabaseSchemaConfigModifyService{HttpService: c.NewHttpService()}
}

func (c *Client) NewConnectorReSyncTable() *connectors.ConnectorReSyncTableService {
	return &connectors.ConnectorReSyncTableService{HttpService: c.NewHttpService()}
}

func (c *Client) NewConnectorSchemaDetails() *connectors.ConnectorSchemaDetailsService {
	return &connectors.ConnectorSchemaDetailsService{HttpService: c.NewHttpService()}
}

func (c *Client) NewConnectorSchemaReload() *connectors.ConnectorSchemaReloadService {
	return &connectors.ConnectorSchemaReloadService{HttpService: c.NewHttpService()}
}

func (c *Client) NewConnectorSchemaUpdateService() *connectors.ConnectorSchemaConfigUpdateService {
	return &connectors.ConnectorSchemaConfigUpdateService{HttpService: c.NewHttpService()}
}

func (c *Client) NewConnectorTableConfigModifyService() *connectors.ConnectorTableConfigModifyService {
    return &connectors.ConnectorTableConfigModifyService{HttpService: c.NewHttpService()}
}