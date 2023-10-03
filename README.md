# Fivetran SDK for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/fivetran/go-fivetran.svg)](https://pkg.go.dev/github.com/fivetran/go-fivetran)

`go-fivetran` is the official Fivetran SDK for the Go programming language.

Checkout our [CHANGELOG](CHANGELOG.md) for information about the latest bug fixes, updates, and features added to the SDK. 

Make sure you read the Fivetran REST API [documentation](https://fivetran.com/docs/rest-api) before using the SDK.

**NOTE**: `go-fivetran` is still in [ALPHA](https://en.wikipedia.org/wiki/Software_release_life_cycle#Alpha) development stage. Future versions may introduce breaking changes. 

## Installation

```
go get github.com/fivetran/go-fivetran
```

## Importing

```go
import (
    "github.com/fivetran/go-fivetran"
)
```

## Getting started

Initialize a new Fivetran client: 

```go
	// get apiKey and apiSecret from environment variables
	apiKey := os.Getenv("FIVETRAN_APIKEY")
	apiSecret := os.Getenv("FIVETRAN_APISECRET")

	// initialize a new client
	client := fivetran.New(apiKey, apiSecret)
```

Each REST API endpoint has a service. Initialize a new Service: 

```go
	// initialize a new UsersList service
	svc := client.NewUsersList()
```

Call the API:

```go
	// call the REST API
	resp, err := svc.Do(context.Background())
	if err != nil {
		...
	}
```

Or you can simply call API in chain style. Call `Do()` at the end to send an HTTP request to the REST API:

```go
	resp, err := client.NewUsersList().
		Limit(3).
		Do(context.Background())

```

## Examples

You can find examples for all services in the [examples](examples/) directory.

## API List

The following [Fivetran REST API](https://fivetran.com/docs/rest-api) endpoints are implemented by the Fivetran SDK for Go: 

### [User Management API](https://fivetran.com/docs/rest-api/users)

REST API Endpoint | REST API Version | SDK Service
--- | --- | ---
[List all Users](https://fivetran.com/docs/rest-api/users#listallusers) | v1 | [UsersListService](https://pkg.go.dev/github.com/fivetran/go-fivetran#UsersListService)
[Retrieve user details](https://fivetran.com/docs/rest-api/users#retrieveuserdetails) | v1 | [UserDetailsService](https://pkg.go.dev/github.com/fivetran/go-fivetran#UserDetailsService)
[Invite a user](https://fivetran.com/docs/rest-api/users#inviteauser) | v1 | [UserInviteService](https://pkg.go.dev/github.com/fivetran/go-fivetran#UserInviteService) 
[Modify a user](https://fivetran.com/docs/rest-api/users#modifyauser) | v1 | [UserModifyService](https://pkg.go.dev/github.com/fivetran/go-fivetran#UserModifyService)
[Delete a user](https://fivetran.com/docs/rest-api/users#deleteauser) | v1 | [UserDeleteService](https://pkg.go.dev/github.com/fivetran/go-fivetran#UserDeleteService)

### [Group Management API](https://fivetran.com/docs/rest-api/groups)

REST API Endpoint | REST API Version | SDK Service
--- | --- | ---
[Create a group](https://fivetran.com/docs/rest-api/groups#createagroup) | v1 | [GroupCreateService](https://pkg.go.dev/github.com/fivetran/go-fivetran#GroupCreateService)
[List all groups](https://fivetran.com/docs/rest-api/groups#listallgroups) | v1 | [GroupsListService](https://pkg.go.dev/github.com/fivetran/go-fivetran#GroupsListService)
[Retrieve group details](https://fivetran.com/docs/rest-api/groups#retrievegroupdetails) | v1 | [GroupDetailsService](https://pkg.go.dev/github.com/fivetran/go-fivetran#GroupDetailsService)
[Modify a group](https://fivetran.com/docs/rest-api/groups#modifyagroup) | v1 | [GroupModifyService](https://pkg.go.dev/github.com/fivetran/go-fivetran#GroupModifyService)
[List all connectors within a group](https://fivetran.com/docs/rest-api/groups#listallconnectorswithinagroup) | v1 | [GroupListConnectorsService](https://pkg.go.dev/github.com/fivetran/go-fivetran#GroupListConnectorsService)
[List all users within a group](https://fivetran.com/docs/rest-api/groups#listalluserswithinagroup) | v1 | [GroupListUsersService](https://pkg.go.dev/github.com/fivetran/go-fivetran#GroupListUsersService)
[Add a user to a group](https://fivetran.com/docs/rest-api/groups#addausertoagroup) | v1 | [GroupAddUserService](https://pkg.go.dev/github.com/fivetran/go-fivetran#GroupAddUserService)
[Remove a user from a group](https://fivetran.com/docs/rest-api/groups#removeauserfromagroup) | v1 | [GroupRemoveUserService](https://pkg.go.dev/github.com/fivetran/go-fivetran#GroupRemoveUserService)
[Delete a group](https://fivetran.com/docs/rest-api/groups#deleteagroup) | v1 | [GroupDeleteService](https://pkg.go.dev/github.com/fivetran/go-fivetran#GroupDeleteService)

### [Destination Management API](https://fivetran.com/docs/rest-api/destinations)

REST API Endpoint | REST API Version | SDK Service/Config
--- | --- | ---
[Create a destination](https://fivetran.com/docs/rest-api/destinations#createadestination) | v1 | [DestinationCreateService](https://pkg.go.dev/github.com/fivetran/go-fivetran#DestinationCreateService)
[Retrieve destination details](https://fivetran.com/docs/rest-api/destinations#retrievedestinationdetails) | v1 | [DestinationDetailsService](https://pkg.go.dev/github.com/fivetran/go-fivetran#DestinationDetailsService)
[Modify a destination](https://fivetran.com/docs/rest-api/destinations#modifyadestination) | v1 | [DestinationModifyService](https://pkg.go.dev/github.com/fivetran/go-fivetran#DestinationModifyService)
[Run destination setup tests](https://fivetran.com/docs/rest-api/destinations#rundestinationsetuptests) | v1 | [DestinationSetupTestsService](https://pkg.go.dev/github.com/fivetran/go-fivetran#DestinationSetupTestsService)
[Delete a destination](https://fivetran.com/docs/rest-api/destinations#deleteadestination) | v1 | [DestinationDeleteService](https://pkg.go.dev/github.com/fivetran/go-fivetran#DestinationDeleteService)
[Destination Config](https://fivetran.com/docs/rest-api/destinations/config) | v1 | [DestinationConfig](https://pkg.go.dev/github.com/fivetran/go-fivetran#DestinationConfig)

### [Connector Management API](https://fivetran.com/docs/rest-api/connectors)

REST API Endpoint | REST API Version | SDK Service/Config/Auth
--- | --- | ---
[Retrieve source metadata](https://fivetran.com/docs/rest-api/connectors#retrievesourcemetadata) | v1 | [ConnectorsSourceMetadataService](https://pkg.go.dev/github.com/fivetran/go-fivetran#ConnectorsSourceMetadataService)
[Create a connector](https://fivetran.com/docs/rest-api/connectors#createaconnector) | v2 | [ConnectorCreateService](https://pkg.go.dev/github.com/fivetran/go-fivetran#ConnectorCreateService)
[Retrieve connector details](https://fivetran.com/docs/rest-api/connectors#retrieveconnectordetails) | v2 | [ConnectorDetailsService](https://pkg.go.dev/github.com/fivetran/go-fivetran#ConnectorDetailsService)
[Modify a connector](https://fivetran.com/docs/rest-api/connectors#modifyaconnector) | v2 | [ConnectorModifyService](https://pkg.go.dev/github.com/fivetran/go-fivetran#ConnectorModifyService)
[Sync connector data](https://fivetran.com/docs/rest-api/connectors#syncconnectordata) | v1 | [ConnectorSyncService](https://pkg.go.dev/github.com/fivetran/go-fivetran#ConnectorSyncService)
[Re-sync connector table data](https://fivetran.com/docs/rest-api/connectors#resyncconnectortabledata) | v1 | [ConnectorReSyncTableService](https://pkg.go.dev/github.com/fivetran/go-fivetran#ConnectorReSyncTableService)
[Run connector setup tests](https://fivetran.com/docs/rest-api/connectors#runconnectorsetuptests) | v2 | [ConnectorSetupTestsService](https://pkg.go.dev/github.com/fivetran/go-fivetran#ConnectorSetupTestsService)
[Delete a connector](https://fivetran.com/docs/rest-api/connectors#deleteaconnector) | v1 | [ConnectorDeleteService](https://pkg.go.dev/github.com/fivetran/go-fivetran#ConnectorDeleteService)
[Retrieve a connector schema config](https://fivetran.com/docs/rest-api/connectors#retrieveaconnectorschemaconfig) | v1 | [ConnectorSchemaDetailsService](https://pkg.go.dev/github.com/fivetran/go-fivetran#ConnectorSchemaDetailsService)
[Retrieve source table columns config](https://fivetran.com/docs/rest-api/connectors#retrievesourcetablecolumnsconfig) | v1 | [ConnectorColumnConfigListService](https://pkg.go.dev/github.com/fivetran/go-fivetran#ConnectorColumnConfigListService)
[Reload a connector schema config](https://fivetran.com/docs/rest-api/connectors#reloadaconnectorschemaconfig) | v1 | [ConnectorSchemaReloadService](https://pkg.go.dev/github.com/fivetran/go-fivetran#ConnectorSchemaReloadService)
[Modify a connector schema config](https://fivetran.com/docs/rest-api/connectors#modifyaconnectorschemaconfig) | v1 | [ConnectorSchemaConfigUpdateService](https://pkg.go.dev/github.com/fivetran/go-fivetran#ConnectorSchemaConfigUpdateService)
[Modify a connector database schema config](https://fivetran.com/docs/rest-api/connectors#modifyaconnectordatabaseschemaconfig) | v1 | [ConnectorDatabaseSchemaConfigModifyService](https://pkg.go.dev/github.com/fivetran/go-fivetran#ConnectorDatabaseSchemaConfigModifyService)
[Modify a connector table config](https://fivetran.com/docs/rest-api/connectors#modifyaconnectortableconfig) | v1 | [ConnectorTableConfigModifyService](https://pkg.go.dev/github.com/fivetran/go-fivetran#ConnectorTableConfigModifyService)
[Modify a connector column config](https://fivetran.com/docs/rest-api/connectors#modifyaconnectorcolumnconfig) | v1 | [ConnectorColumnConfigModifyService](https://pkg.go.dev/github.com/fivetran/go-fivetran#ConnectorColumnConfigModifyService)
[Connector Config](https://fivetran.com/docs/rest-api/connectors/config) | v1 | [ConnectorConfig](https://pkg.go.dev/github.com/fivetran/go-fivetran#ConnectorConfig)<br> [ConnectorConfigReports](https://pkg.go.dev/github.com/fivetran/go-fivetran#ConnectorConfigReports)<br> [ConnectorConfigProjectCredentials](https://pkg.go.dev/github.com/fivetran/go-fivetran#ConnectorConfigProjectCredentials)<br> [ConnectorConfigCustomTables](https://pkg.go.dev/github.com/fivetran/go-fivetran#ConnectorConfigCustomTables)
[Connector Auth](https://fivetran.com/docs/rest-api/connectors) | v1 | [ConnectorAuth](https://pkg.go.dev/github.com/fivetran/go-fivetran#ConnectorAuth)<br> [ConnectorAuthClientAccess](https://pkg.go.dev/github.com/fivetran/go-fivetran#ConnectorAuthClientAccess)
[Connect Card](https://fivetran.com/docs/rest-api/connectors/connect-card) | v1 | [ConnectCardService](https://pkg.go.dev/github.com/fivetran/go-fivetran#ConnectCardService)

### [Transformations Management API](https://fivetran.com/docs/rest-api/dbt-transformation-management)
REST API Endpoint | REST API Version | SDK Service
--- | --- | ---
[Create dbt Project](https://fivetran.com/docs/rest-api/dbt-transformation-management#createdbtproject) | v1 | [DbtProjectCreateService](https://pkg.go.dev/github.com/fivetran/go-fivetran#DbtProjectCreateService)
[Retrieve dbt Project Details](https://fivetran.com/docs/rest-api/dbt-transformation-management#retrievedbtprojectdetails) | v1 | [DbtProjectDetailsService](https://pkg.go.dev/github.com/fivetran/go-fivetran#DbtProjectDetailsService)
[Retrieve dbt Projects List](https://fivetran.com/docs/rest-api/dbt-transformation-management#retrievedbtprojects) | v1 | [DbtProjectsListService](https://pkg.go.dev/github.com/fivetran/go-fivetran#DbtProjectsListService)
[Retrieve dbt Model Details](https://fivetran.com/docs/rest-api/dbt-transformation-management#retrievedbtprojectmodeldetails) | v1 | [DbtModelDetailsService](https://pkg.go.dev/github.com/fivetran/go-fivetran#DbtModelDetailsService)
[Retrieve dbt Models List](https://fivetran.com/docs/rest-api/dbt-transformation-management#retrievedbtprojectmodels) | v1 | [DbtModelsListService](https://pkg.go.dev/github.com/fivetran/go-fivetran#DbtModelsListService)
[Update dbt Project](https://fivetran.com/docs/rest-api/dbt-transformation-management#modifydbtproject) | v1 | [DbtProjectModifyService](https://pkg.go.dev/github.com/fivetran/go-fivetran#DbtProjectModifyService)
[Delete dbt Project](https://fivetran.com/docs/rest-api/dbt-transformation-management#deletedbtproject) | v1 | [DbtProjectDeleteService](https://pkg.go.dev/github.com/fivetran/go-fivetran#DbtProjectDeleteService)
[Create dbt Transformation](https://fivetran.com/docs/rest-api/dbt-transformation-management#createdbttransformation) | v1 | [DbtTransformationCreateService](https://pkg.go.dev/github.com/fivetran/go-fivetran#DbtTransformationCreateService)
[Retrieve dbt Transformation Details](https://fivetran.com/docs/rest-api/dbt-transformation-management#retrievedbttransformationdetails) | v1 | [DbtTransformationDetailsService](https://pkg.go.dev/github.com/fivetran/go-fivetran#DbtTransformationDetailsService)
[Update dbt Transformation](https://fivetran.com/docs/rest-api/dbt-transformation-management#modifydbttransformation) | v1 | [DbtTransformationModifyService](https://pkg.go.dev/github.com/fivetran/go-fivetran#DbtTransformationModifyService)
[Delete dbt Transformation](https://fivetran.com/docs/rest-api/dbt-transformation-management#deletedbttransformation) | v1 | [DbtTransformationDeleteService](https://pkg.go.dev/github.com/fivetran/go-fivetran#DbtTransformationDeleteService)

### [Certificate Management API](https://fivetran.com/docs/rest-api/certificates)
REST API Endpoint | REST API Version | SDK Service
--- | --- | ---
[Approve a connector certificate](https://fivetran.com/docs/rest-api/certificates#approveaconnectorcertificate) | v1 | [CertificateConnectorCertificateApproveService](https://pkg.go.dev/github.com/fivetran/go-fivetran#CertificateConnectorCertificateApproveService)
[Approve a connector fingerprint](https://fivetran.com/docs/rest-api/certificates#approveaconnectorfingerprint) | v1 | [CertificateConnectorFingerprintApproveService](https://pkg.go.dev/github.com/fivetran/go-fivetran#CertificateConnectorFingerprintApproveService)
[Approve a destination certificate](https://fivetran.com/docs/rest-api/certificates#approveadestinationcertificate) | v1 | [CertificateDestinationCertificateApproveService](https://pkg.go.dev/github.com/fivetran/go-fivetran#CertificateDestinationCertificateApproveService)
[Approve a destination fingerprint](https://fivetran.com/docs/rest-api/certificates#approveadestinationfingerprint) | v1 | [CertificateDestinationFingerprintApproveService](https://pkg.go.dev/github.com/fivetran/go-fivetran#CertificateDestinationFingerprintApproveService)

### [Log Service Management](https://fivetran.com/docs/rest-api/log-service-management)

REST API Endpoint | REST API Version | SDK Service/Config
--- | --- | ---
[Create a Log Service](https://fivetran.com/docs/rest-api/log-service-management#createalogservice) | v1 | [ExternalLoggingCreateService](https://pkg.go.dev/github.com/fivetran/go-fivetran#ExternalLoggingCreateService)
[Retrieve Log Service Details](https://fivetran.com/docs/rest-api/log-service-management#retrievelogservicedetails) | v1 | [ExternalLoggingDetailsService](https://pkg.go.dev/github.com/fivetran/go-fivetran#ExternalLoggingDetailsService)
[Update a Log Service](https://fivetran.com/docs/rest-api/log-service-management#updatealogservice) | v1 | [ExternalLoggingModifyService](https://pkg.go.dev/github.com/fivetran/go-fivetran#ExternalLoggingModifyService)
[Delete a Log Service](https://fivetran.com/docs/rest-api/log-service-management#deletealogservice) | v1 | [ExternalLoggingDeleteService](https://pkg.go.dev/github.com/fivetran/go-fivetran#ExternalLoggingDeleteService)
[Run Log Service Setup Tests](https://fivetran.com/docs/rest-api/log-service-management#runlogservicesetuptests) | v1 | [ExternalLoggingSetupTestsService](https://pkg.go.dev/github.com/fivetran/go-fivetran#ExternalLoggingSetupTestsService)

### [Metadata API](https://fivetran.com/docs/rest-api/metadata#metadataapi)

REST API Endpoint | REST API Version | SDK Service/Config
--- | --- | ---
[Retrieve schema metadata](https://fivetran.com/docs/rest-api/metadata#retrieveschemametadata) | v1 | [MetadataSchemaListService](https://pkg.go.dev/github.com/fivetran/go-fivetran#ExternalLoggingCreateService)
[Retrieve table metadata](https://fivetran.com/docs/rest-api/metadata#retrievetablemetadata) | v1 | [MetadataTablesListService](https://pkg.go.dev/github.com/fivetran/go-fivetran#ExternalLoggingCreateService)
[Retrieve column metadata](https://fivetran.com/docs/rest-api/metadata#retrievecolumnmetadata) | v1 | [MetadataColumnListService](https://pkg.go.dev/github.com/fivetran/go-fivetran#ExternalLoggingCreateService)

### [Webhook Management](https://fivetran.com/docs/rest-api/webhooks#webhookmanagement)

REST API Endpoint | REST API Version | SDK Service/Config
--- | --- | ---
[Create account webhook](https://fivetran.com/docs/rest-api/webhooks#createaccountwebhook) | v1 | [WebhookAccountCreateService](https://pkg.go.dev/github.com/fivetran/go-fivetran#WebhookAccountCreateService)
[Create group webhook](https://fivetran.com/docs/rest-api/webhooks#creategroupwebhook) | v1 | [WebhookGroupCreateService](https://pkg.go.dev/github.com/fivetran/go-fivetran#WebhookGroupCreateService)
[Retrieve webhook details](https://fivetran.com/docs/rest-api/webhooks#retrievewebhookdetails) | v1 | [WebhookDetailsService](https://pkg.go.dev/github.com/fivetran/go-fivetran#WebhookDetailsService)
[Update webhook](https://fivetran.com/docs/rest-api/webhooks#updatewebhook) | v1 | [WebhookModifyService](https://pkg.go.dev/github.com/fivetran/go-fivetran#WebhookModifyService)
[Delete webhook](https://fivetran.com/docs/rest-api/webhooks#deletewebhook) | v1 | [WebhookDeleteService](https://pkg.go.dev/github.com/fivetran/go-fivetran#WebhookDeleteService)
[Retrieve the list of webhooks](https://fivetran.com/docs/rest-api/webhooks#retrievethelistofwebhooks) | v1 | [WebhookListService](https://pkg.go.dev/github.com/fivetran/go-fivetran#WebhookListService)
[Test webhook](https://fivetran.com/docs/rest-api/webhooks#testwebhook) | v1 | [WebhookTestService](https://pkg.go.dev/github.com/fivetran/go-fivetran#WebhookTestService)

### [Role Management](https://fivetran.com/docs/rest-api/roles)

REST API Endpoint | REST API Version | SDK Service/Config
--- | --- | ---
[List all roles](https://fivetran.com/docs/rest-api/roles#listallroles) | v1 | [RolesListService](https://pkg.go.dev/github.com/fivetran/go-fivetran#RolesListService)

### [Team Management](https://fivetran.com/docs/rest-api/teams)

REST API Endpoint | REST API Version | SDK Service/Config
--- | --- | ---
[List all teams](https://fivetran.com/docs/rest-api/teams#listallteams) | v1 | [TeamsListService](https://pkg.go.dev/github.com/fivetran/go-fivetran#TeamsListService)
[Retrieve team details](https://fivetran.com/docs/rest-api/teams#retrieveteamdetails) | v1 | [TeamsDetailsService](https://pkg.go.dev/github.com/fivetran/go-fivetran#TeamsDetailsService)
[Create a team](https://fivetran.com/docs/rest-api/teams#createateam) | v1 | [TeamsCreateService](https://pkg.go.dev/github.com/fivetran/go-fivetran#TeamsCreateService)
[Modify a team](https://fivetran.com/docs/rest-api/teams#modifyateam) | v1 | [TeamsModifyService](https://pkg.go.dev/github.com/fivetran/go-fivetran#TeamsModifyService)
[Delete a team role in the account](https://fivetran.com/docs/rest-api/teams#deleteteamroleinaccount) | v1 | [TeamsDeleteRoleInAccountService](https://pkg.go.dev/github.com/fivetran/go-fivetran#TeamsDeleteRoleInAccountService)
[Delete a team](https://fivetran.com/docs/rest-api/teams#deleteateam) | v1 | [TeamsDeleteService](https://pkg.go.dev/github.com/fivetran/go-fivetran#TeamsDeleteService)

### [Team Management User memberships](https://fivetran.com/docs/rest-api/teams#usermemberships)
REST API Endpoint | REST API Version | SDK Service/Config
--- | --- | ---
[List all user memberships](https://fivetran.com/docs/rest-api/teams#listallusermemberships) | v1 | [TeamUserMembershipsListService](https://pkg.go.dev/github.com/fivetran/go-fivetran#TeamUserMembershipsListService)
[Retrieve user membership](https://fivetran.com/docs/rest-api/teams#retrieveusermembershipinateam) | v1 | [TeamUserMembershipDetailsService](https://pkg.go.dev/github.com/fivetran/go-fivetran#TeamUserMembershipDetailsService)
[Add a user to a team](https://fivetran.com/docs/rest-api/teams#addausertoateam) | v1 | [TeamUserMembershipCreateService](https://pkg.go.dev/github.com/fivetran/go-fivetran#TeamUserMembershipCreateService)
[Modify a user membership](https://fivetran.com/docs/rest-api/teams#modifyausermembership) | v1 | [TeamUserMembershipModifyService](https://pkg.go.dev/github.com/fivetran/go-fivetran#TeamUserMembershipModifyService)
[Delete a user from a team](https://fivetran.com/docs/rest-api/teams#deleteauserfromateam) | v1 | [TeamUserMembershipDeleteService](https://pkg.go.dev/github.com/fivetran/go-fivetran#TeamUserMembershipDeleteService)

### [Team Management Connector memberships](https://fivetran.com/docs/rest-api/teams#connectormemberships)
REST API Endpoint | REST API Version | SDK Service/Config
--- | --- | ---
[List all connector memberships](https://fivetran.com/docs/rest-api/teams#listallconnectormemberships) | v1 | [TeamConnectorMembershipsListService](https://pkg.go.dev/github.com/fivetran/go-fivetran#TeamConnectorMembershipsListService)
[Retrieve connector membership](https://fivetran.com/docs/rest-api/teams#retrieveconnectormembership) | v1 | [TeamConnectorMembershipDetailsService](https://pkg.go.dev/github.com/fivetran/go-fivetran#TeamConnectorMembershipDetailsService)
[Add connector membership](https://fivetran.com/docs/rest-api/teams#addconnectormembership) | v1 | [TeamConnectorMembershipCreateService](https://pkg.go.dev/github.com/fivetran/go-fivetran#TeamConnectorMembershipCreateService)
[Update connector membership](https://fivetran.com/docs/rest-api/teams#updateconnectormembership) | v1 | [TeamConnectorMembershipModifyService](https://pkg.go.dev/github.com/fivetran/go-fivetran#TeamConnectorMembershipModifyService)
[Delete connector membership](https://fivetran.com/docs/rest-api/teams#deleteconnectormembership) | v1 | [TeamConnectorMembershipDeleteService](https://pkg.go.dev/github.com/fivetran/go-fivetran#TeamConnectorMembershipDeleteService)

### [Team Management Group memberships](https://fivetran.com/docs/rest-api/teams#groupmemberships)
REST API Endpoint | REST API Version | SDK Service/Config
--- | --- | ---
[List all group memberships](https://fivetran.com/docs/rest-api/teams#listallgroupmemberships) | v1 | [TeamGroupMembershipsService](https://pkg.go.dev/github.com/fivetran/go-fivetran#TeamGroupMembershipsService)
[Retrieve group membership](https://fivetran.com/docs/rest-api/teams#retrievegroupmembership) | v1 | [TeamGroupMembershipDetailsService](https://pkg.go.dev/github.com/fivetran/go-fivetran#TeamGroupMembershipDetailsService)
[Add group membership](https://fivetran.com/docs/rest-api/teams#addgroupmembership) | v1 | [TeamGroupMembershipCreateService](https://pkg.go.dev/github.com/fivetran/go-fivetran#TeamGroupMembershipCreateService)
[Update group membership](https://fivetran.com/docs/rest-api/teams#updategroupmembership) | v1 | [TeamGroupMembershipModifyService](https://pkg.go.dev/github.com/fivetran/go-fivetran#TeamGroupMembershipModifyService)
[Delete group membership](https://fivetran.com/docs/rest-api/teams#deletegroupmembership) | v1 | [TeamGroupMembershipDeleteService](https://pkg.go.dev/github.com/fivetran/go-fivetran#TeamGroupMembershipDeleteService)

## Support

Please get in touch with us through our [Support Portal](https://support.fivetran.com/) if you 
have any comments, suggestions, support requests, or bug reports.  
