# Fivetran SDK for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/fivetran/go-fivetran.svg)](https://pkg.go.dev/github.com/fivetran/go-fivetran)

`go-fivetran` is the official Fivetran SDK for the Go programming language.

Checkout our [CHANGELOG](CHANGELOG.md) for information about the latest bug fixes, updates, and features added to the SDK. 

Make sure you read the Fivetran REST API [documentation](https://fivetran.com/docs/rest-api) before using the SDK.

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

### [User Management API](https://fivetran.com/docs/rest-api/api-reference/users)

REST API Endpoint | REST API Version | SDK Service
--- | --- | ---
[List all Users](https://fivetran.com/docs/rest-api/api-reference/users/list-all-users) | v1 | [UsersListService](https://pkg.go.dev/github.com/fivetran/go-fivetran#UsersListService)
[Retrieve user details](https://fivetran.com/docs/rest-api/api-reference/users/user-details) | v1 | [UserDetailsService](https://pkg.go.dev/github.com/fivetran/go-fivetran#UserDetailsService)
[Invite a user](https://fivetran.com/docs/rest-api/api-reference/users/create-user) | v1 | [UserInviteService](https://pkg.go.dev/github.com/fivetran/go-fivetran#UserInviteService) 
[Update a user](https://fivetran.com/docs/rest-api/api-reference/users/modify-user) | v1 | [UserModifyService](https://pkg.go.dev/github.com/fivetran/go-fivetran#UserModifyService)
[Delete a user](https://fivetran.com/docs/rest-api/api-reference/users/delete-user) | v1 | [UserDeleteService](https://pkg.go.dev/github.com/fivetran/go-fivetran#UserDeleteService)
[List all connection memberships](https://fivetran.com/docs/rest-api/api-reference/users/get-user-memberships-in-connections) | v1 | [UserConnectionMembershipsListService](https://pkg.go.dev/github.com/fivetran/go-fivetran#UserConnectionMembershipsListService)
[Retrieve connection membership](https://fivetran.com/docs/rest-api/api-reference/users/get-user-membership-in-connections) | v1 | [UserConnectionMembershipDetailsService](https://pkg.go.dev/github.com/fivetran/go-fivetran#UserConnectionMembershipDetailsService)
[Add connection membership](https://fivetran.com/docs/rest-api/api-reference/users/add-user-membership-in-connection) | v1 | [UserConnectionMembershipCreateService](https://pkg.go.dev/github.com/fivetran/go-UserConnectionMembershipCreateService#UserDeleteService)
[Update connection membership](https://fivetran.com/docs/rest-api/api-reference/users/update-user-membership-in-connection) | v1 | [UserConnectionMembershipModifyService](https://pkg.go.dev/github.com/fivetran/go-fivetran#UserConnectionMembershipModifyService)
[Delete connection membership](https://fivetran.com/docs/rest-api/api-reference/users/delete-user-membership-in-connection) | v1 | [UserConnectionMembershipDeleteService](https://pkg.go.dev/github.com/fivetran/go-fivetran#UserConnectionMembershipDeleteService)
[List all group memberships](https://fivetran.com/docs/rest-api/api-reference/users/get-user-memberships-in-groups) | v1 | [UserGroupMembershipsListService](https://pkg.go.dev/github.com/fivetran/go-fivetran#UserGroupMembershipsListService)
[Retrieve group membership](https://fivetran.com/docs/rest-api/api-reference/users/get-user-membership-in-group) | v1 | [UserGroupMembershipDetailsService](https://pkg.go.dev/github.com/fivetran/go-fivetran#UserGroupMembershipDetailsService)
[Add group membership](https://fivetran.com/docs/rest-api/api-reference/users/add-user-membership-in-group) | v1 | [UserGroupMembershipCreateService](https://pkg.go.dev/github.com/fivetran/go-fivetran#UserGroupMembershipCreateService)
[Update group membership](https://fivetran.com/docs/rest-api/api-reference/users/update-user-membership-in-group) | v1 | [UserGroupMembershipModifyService](https://pkg.go.dev/github.com/fivetran/go-fivetran#UserGroupMembershipModifyService)
[Delete group membership](https://fivetran.com/docs/rest-api/api-reference/users/delete-user-membership-in-group) | v1 | [UserGroupMembershipDeleteService](https://pkg.go.dev/github.com/fivetran/go-fivetran#UserGroupMembershipDeleteService)

### [Group Management API](https://fivetran.com/docs/rest-api/api-reference/groups)

REST API Endpoint | REST API Version | SDK Service
--- | --- | ---
[Create a group](https://fivetran.com/docs/rest-api/api-reference/groups/create-group) | v1 | [GroupCreateService](https://pkg.go.dev/github.com/fivetran/go-fivetran#GroupCreateService)
[List all groups](https://fivetran.com/docs/rest-api/api-reference/groups/list-all-groups) | v1 | [GroupsListService](https://pkg.go.dev/github.com/fivetran/go-fivetran#GroupsListService)
[Retrieve group details](https://fivetran.com/docs/rest-api/api-reference/groups/group-details) | v1 | [GroupDetailsService](https://pkg.go.dev/github.com/fivetran/go-fivetran#GroupDetailsService)
[Update a group](https://fivetran.com/docs/rest-api/api-reference/groups/modify-group) | v1 | [GroupModifyService](https://pkg.go.dev/github.com/fivetran/go-fivetran#GroupModifyService)
[List All Connections within a Group](https://fivetran.com/docs/rest-api/api-reference/groups/list-all-connections-in-group) | v1 | [GroupListConnectionsService](https://pkg.go.dev/github.com/fivetran/go-fivetran#GroupListConnectionsService)
[List all users within a group](https://fivetran.com/docs/rest-api/api-reference/groups/list-all-users-in-group) | v1 | [GroupListUsersService](https://pkg.go.dev/github.com/fivetran/go-fivetran#GroupListUsersService)
[Add a user to a group](https://fivetran.com/docs/rest-api/api-reference/groups/add-user-to-group) | v1 | [GroupAddUserService](https://pkg.go.dev/github.com/fivetran/go-fivetran#GroupAddUserService)
[Remove a user from a group](https://fivetran.com/docs/rest-api/api-reference/groups/delete-user-from-group) | v1 | [GroupRemoveUserService](https://pkg.go.dev/github.com/fivetran/go-fivetran#GroupRemoveUserService)
[Delete a group](https://fivetran.com/docs/rest-api/api-reference/groups/delete-group) | v1 | [GroupDeleteService](https://pkg.go.dev/github.com/fivetran/go-fivetran#GroupDeleteService)
[Retrieve Group Public SSH Key](https://fivetran.com/docs/rest-api/api-reference/groups/group-ssh-public-key) | v1 | [GroupSshKeyService](https://pkg.go.dev/github.com/fivetran/go-fivetran/groups#GroupSshKeyService)
[Retrieve Group Service Account](https://fivetran.com/docs/rest-api/api-reference/groups/group-service-account) | v1 | [GroupServiceAccountService](https://pkg.go.dev/github.com/fivetran/go-fivetran/groups#GroupServiceAccountService)

### [Destination Management API](https://fivetran.com/docs/rest-api/destinations)

REST API Endpoint | REST API Version | SDK Service/Config
--- | --- | ---
[List all destinations within account](https://fivetran.com/docs/rest-api/api-reference/destinations/list-destinations) | v1 | [DestinationsListService](https://pkg.go.dev/github.com/fivetran/go-fivetran#DestinationsListService)
[Create a destination](https://fivetran.com/docs/rest-api/destinations#createadestination) | v1 | [DestinationCreateService](https://pkg.go.dev/github.com/fivetran/go-fivetran#DestinationCreateService)
[Retrieve destination details](https://fivetran.com/docs/rest-api/destinations#retrievedestinationdetails) | v1 | [DestinationDetailsService](https://pkg.go.dev/github.com/fivetran/go-fivetran#DestinationDetailsService)
[Modify a destination](https://fivetran.com/docs/rest-api/destinations#modifyadestination) | v1 | [DestinationModifyService](https://pkg.go.dev/github.com/fivetran/go-fivetran#DestinationModifyService)
[Run destination setup tests](https://fivetran.com/docs/rest-api/destinations#rundestinationsetuptests) | v1 | [DestinationSetupTestsService](https://pkg.go.dev/github.com/fivetran/go-fivetran#DestinationSetupTestsService)
[Delete a destination](https://fivetran.com/docs/rest-api/destinations#deleteadestination) | v1 | [DestinationDeleteService](https://pkg.go.dev/github.com/fivetran/go-fivetran#DestinationDeleteService)
[Destination Config](https://fivetran.com/docs/rest-api/destinations/config) | v1 | [DestinationConfig](https://pkg.go.dev/github.com/fivetran/go-fivetran#DestinationConfig)

### [Transformations Management API](https://fivetran.com/docs/rest-api/api-reference/transformation-management)
REST API Endpoint | REST API Version | SDK Service
--- | --- | ---
[Create Transformation Project](https://fivetran.com/docs/rest-api/api-reference/transformation-projects-management/create-transformation-project) | v1 | [TransformationProjectCreateService](https://pkg.go.dev/github.com/fivetran/go-fivetran#TransformationProjectCreateService)
[Delete Transformation Project](https://fivetran.com/docs/rest-api/api-reference/transformation-projects-management/delete-transformation-project) | v1 | [TransformationProjectDeleteService](https://pkg.go.dev/github.com/fivetran/go-fivetran#TransformationProjectDeleteService)
[List all Transformation Projects within Account](https://fivetran.com/docs/rest-api/api-reference/transformation-projects-management/list-all-transformation-projects) | v1 | [TransformationProjectsListService](https://pkg.go.dev/github.com/fivetran/go-fivetran#TransformationProjectsListService)
[Retrieve Transformation Project Details](https://fivetran.com/docs/rest-api/api-reference/transformation-projects-management/transformation-project-details) | v1 | [TransformationProjectDetailsService](https://pkg.go.dev/github.com/fivetran/go-fivetran#TransformationProjectDetailsService)
[Test Transformation Project](https://fivetran.com/docs/rest-api/api-reference/transformation-projects-management/test-transformation-project) | v1 | [TransformationProjectTestsService](https://pkg.go.dev/github.com/fivetran/go-fivetran#TransformationProjectTestsService)
[Update Transformation Project](https://fivetran.com/docs/rest-api/api-reference/transformation-projects-management/modify-transformation-project) | v1 | [TransformationProjectUpdateService](https://pkg.go.dev/github.com/fivetran/go-fivetran#TransformationProjectUpdateService)
[Create Transformation](https://fivetran.com/docs/rest-api/api-reference/transformation-management/create-transformation) | v1 | [TransformationCreateService](https://pkg.go.dev/github.com/fivetran/go-fivetran#TransformationCreateService)
[Update Transformation](https://fivetran.com/docs/rest-api/api-reference/transformation-management/update-transformation) | v1 | [TransformationUpdateService](https://pkg.go.dev/github.com/fivetran/go-fivetran#TransformationUpdateService)
[Delete Transformation](https://fivetran.com/docs/rest-api/api-reference/transformation-management/delete-transformation) | v1 | [TransformationDeleteService](https://pkg.go.dev/github.com/fivetran/go-fivetran#TransformationDeleteService)
[Retrieve Transformation Details](https://fivetran.com/docs/rest-api/api-reference/transformation-management/transformation-details) | v1 | [TransformationDetailsService](https://pkg.go.dev/github.com/fivetran/go-fivetran#TransformationDetailsService)
[List all Transformations within Account](https://fivetran.com/docs/rest-api/api-reference/transformation-management/transformations-list) | v1 | [TransformationsListService](https://pkg.go.dev/github.com/fivetran/go-fivetran#TransformationsListService)
[Upgrade Transformation Package](https://fivetran.com/docs/rest-api/api-reference/transformation-management/upgrade-transformation-package) | v1 | [TransformationUpgradePackageService](https://pkg.go.dev/github.com/fivetran/go-fivetran#TransformationUpgradePackageService)
[Run Transformation](https://fivetran.com/docs/rest-api/api-reference/transformation-management/run-transformation) | v1 | [TransformationRunService](https://pkg.go.dev/github.com/fivetran/go-fivetran#TransformationRunService)
[Cancel Transformation](https://fivetran.com/docs/rest-api/api-reference/transformation-management/cancel-transformation) | v1 | [TransformationCancelService](https://pkg.go.dev/github.com/fivetran/go-fivetran#TransformationCancelService)
[List All Quickstart Package Metadata](https://fivetran.com/docs/rest-api/api-reference/transformation-management/transformation-package-metadata-list) | v1 | [QuickstartPackagesListService](https://pkg.go.dev/github.com/fivetran/go-fivetran#QuickstartPackagesListService)
[Retrieve Quickstart Package Metadata Details](https://fivetran.com/docs/rest-api/api-reference/transformation-management/transformation-package-metadata-details) | v1 | [QuickstartPackagesDetailsService](https://pkg.go.dev/github.com/fivetran/go-fivetran#QuickstartPackagesDetailsService)

### [Connection Management API](https://fivetran.com/docs/rest-api/connections)

REST API Endpoint | REST API Version | SDK Service/Config/Auth
--- | --- | ---
[List all connections within account](https://fivetran.com/docs/rest-api/api-reference/connections/list-connections) | v1 | [ConnectionsListService](https://pkg.go.dev/github.com/fivetran/go-fivetran#ConnectionsListService)
[Retrieve source metadata](https://fivetran.com/docs/rest-api/connections#retrievesourcemetadata) | v1 | [ConnectionsSourceMetadataService](https://pkg.go.dev/github.com/fivetran/go-fivetran#ConnectionsSourceMetadataService)
[Create a connection](https://fivetran.com/docs/rest-api/connections#createaconnection) | v2 | [ConnectionCreateService](https://pkg.go.dev/github.com/fivetran/go-fivetran#ConnectionCreateService)
[Retrieve connection details](https://fivetran.com/docs/rest-api/connections#retrieveconnectiondetails) | v2 | [ConnectionDetailsService](https://pkg.go.dev/github.com/fivetran/go-fivetran#ConnectionDetailsService)
[Modify a connection](https://fivetran.com/docs/rest-api/connections#modifyaconnection) | v2 | [ConnectionModifyService](https://pkg.go.dev/github.com/fivetran/go-fivetran#ConnectionModifyService)
[Sync connection data](https://fivetran.com/docs/rest-api/connections#syncconnectiondata) | v1 | [ConnectionSyncService](https://pkg.go.dev/github.com/fivetran/go-fivetran#ConnectionSyncService)
[Re-sync connection table data](https://fivetran.com/docs/rest-api/connections#resyncconnectiontabledata) | v1 | [ConnectionReSyncTableService](https://pkg.go.dev/github.com/fivetran/go-fivetran#ConnectionReSyncTableService)
[Run connection setup tests](https://fivetran.com/docs/rest-api/connections#runconnectionsetuptests) | v2 | [ConnectionSetupTestsService](https://pkg.go.dev/github.com/fivetran/go-fivetran#ConnectionSetupTestsService)
[Delete a connection](https://fivetran.com/docs/rest-api/connections#deleteaconnection) | v1 | [ConnectionDeleteService](https://pkg.go.dev/github.com/fivetran/go-fivetran#ConnectionDeleteService)
[Retrieve a connection schema config](https://fivetran.com/docs/rest-api/connections#retrieveaconnectionschemaconfig) | v1 | [ConnectionSchemaDetailsService](https://pkg.go.dev/github.com/fivetran/go-fivetran#ConnectionSchemaDetailsService)
[Retrieve source table columns config](https://fivetran.com/docs/rest-api/connections#retrievesourcetablecolumnsconfig) | v1 | [ConnectionColumnConfigListService](https://pkg.go.dev/github.com/fivetran/go-fivetran#ConnectionColumnConfigListService)
[Reload a connection schema config](https://fivetran.com/docs/rest-api/connections#reloadaconnectionschemaconfig) | v1 | [ConnectionSchemaReloadService](https://pkg.go.dev/github.com/fivetran/go-fivetran#ConnectionSchemaReloadService)
[Create a Connection Schema Config](https://fivetran.com/docs/rest-api/connections#createaconnectionschemaconfig) | v1 | [ConnectionSchemaConfigCreateService](https://pkg.go.dev/github.com/fivetran/go-fivetran#ConnectionSchemaConfigCreateService)
[Modify a connection schema config](https://fivetran.com/docs/rest-api/connections#modifyaconnectionschemaconfig) | v1 | [ConnectionSchemaConfigUpdateService](https://pkg.go.dev/github.com/fivetran/go-fivetran#ConnectionSchemaConfigUpdateService)
[Modify a connection database schema config](https://fivetran.com/docs/rest-api/connections#modifyaconnectiondatabaseschemaconfig) | v1 | [ConnectionDatabaseSchemaConfigModifyService](https://pkg.go.dev/github.com/fivetran/go-fivetran#ConnectionDatabaseSchemaConfigModifyService)
[Modify a connection table config](https://fivetran.com/docs/rest-api/connections#modifyaconnectiontableconfig) | v1 | [ConnectionTableConfigModifyService](https://pkg.go.dev/github.com/fivetran/go-fivetran#ConnectionTableConfigModifyService)
[Modify a connection column config](https://fivetran.com/docs/rest-api/connections#modifyaconnectioncolumnconfig) | v1 | [ConnectionColumnConfigModifyService](https://pkg.go.dev/github.com/fivetran/go-fivetran#ConnectionColumnConfigModifyService)
[Connection Config](https://fivetran.com/docs/rest-api/connections/config) | v1 | [ConnectionConfig](https://pkg.go.dev/github.com/fivetran/go-fivetran#ConnectionConfig)<br> [ConnectionConfigReports](https://pkg.go.dev/github.com/fivetran/go-fivetran#ConnectionConfigReports)<br> [ConnectionConfigProjectCredentials](https://pkg.go.dev/github.com/fivetran/go-fivetran#ConnectionConfigProjectCredentials)<br> [ConnectionConfigCustomTables](https://pkg.go.dev/github.com/fivetran/go-fivetran#ConnectionConfigCustomTables)
[Connection Auth](https://fivetran.com/docs/rest-api/connections) | v1 | [ConnectionAuth](https://pkg.go.dev/github.com/fivetran/go-fivetran#ConnectionAuth)<br> [ConnectionAuthClientAccess](https://pkg.go.dev/github.com/fivetran/go-fivetran#ConnectionAuthClientAccess)
[Connect Card](https://fivetran.com/docs/rest-api/connections/connect-card) | v1 | [ConnectCardService](https://pkg.go.dev/github.com/fivetran/go-fivetran#ConnectCardService)

### [Certificate Management API](https://fivetran.com/docs/rest-api/certificates)
REST API Endpoint | REST API Version | SDK Service
--- | --- | ---
[Approve a connection certificate](https://fivetran.com/docs/rest-api/certificates#approveaconnectioncertificate) | v1 | [ConnectionCertificateApproveService](https://pkg.go.dev/github.com/fivetran/go-fivetran/certificates#ConnectionCertificateApproveService)
[Approve a connection fingerprint](https://fivetran.com/docs/rest-api/certificates#approveaconnectionfingerprint) | v1 | [ConnectionFingerprintApproveService](https://pkg.go.dev/github.com/fivetran/go-fivetran/fingerprints#ConnectionFingerprintApproveService)
[List all approved certificates for connection](https://fivetran.com/docs/rest-api/certificates#listallapprovedcertificatesforconnection) | v1 | [ConnectionCertificatesListService](https://pkg.go.dev/github.com/fivetran/go-fivetran/certificates#ConnectionCertificatesListService)
[List all approved fingerprints for connection](https://fivetran.com/docs/rest-api/certificates#listallapprovedfingerprintsforconnection) | v1 | [ConnectionFingerprintsListService](https://pkg.go.dev/github.com/fivetran/go-fivetran/fingerprints#ConnectionFingerprintsListService)
[Retrieve a connection certificate details](https://fivetran.com/docs/rest-api/certificates#retrieveaconnectioncertificatedetails) | v1 | [ConnectionCertificateDetailsService](https://pkg.go.dev/github.com/fivetran/go-fivetran/certificates#ConnectionCertificateDetailsService)
[Retrieve a connection fingerprint details](https://fivetran.com/docs/rest-api/certificates#retrieveaconnectionfingerprintdetails) | v1 | [ConnectionFingerprintDetailsService](https://pkg.go.dev/github.com/fivetran/go-fivetran/fingerprints#ConnectionFingerprintDetailsService)
[Revoke a connection certificate](https://fivetran.com/docs/rest-api/certificates#revokeaconnectioncertificate) | v1 | [ConnectionCertificateRevokeService](https://pkg.go.dev/github.com/fivetran/go-fivetran/certificates#ConnectionCertificateRevokeService)
[Revoke a connection fingerprint](https://fivetran.com/docs/rest-api/certificates#revokeaconnectionfingerprint) | v1 | [ConnectionFingerprintRevokeService](https://pkg.go.dev/github.com/fivetran/go-fivetran/fingerprints#ConnectionFingerprintRevokeService)
[Approve a destination certificate](https://fivetran.com/docs/rest-api/certificates#approveadestinationcertificate) | v1 | [DestinationCertificateApproveService](https://pkg.go.dev/github.com/fivetran/go-fivetran/certificates#DestinationCertificateApproveService)
[Approve a destination fingerprint](https://fivetran.com/docs/rest-api/certificates#approveadestinationfingerprint) | v1 | [DestinationFingerprintApproveService](https://pkg.go.dev/github.com/fivetran/go-fivetran/fingerprints#DestinationFingerprintApproveService)
[List all approved certificates for destination](https://fivetran.com/docs/rest-api/certificates#listallapprovedcertificatesfordestination) | v1 | [DestinationCertificatesListService](https://pkg.go.dev/github.com/fivetran/go-fivetran/certificates#DestinationCertificatesListService)
[List all approved fingerprints for destination](https://fivetran.com/docs/rest-api/certificates#listallapprovedfingerprintsfordestination) | v1 | [DestinationFingerprintsListService](https://pkg.go.dev/github.com/fivetran/go-fivetran/fingerprints#DestinationFingerprintsListService)
[Retrieve a destination certificate details](https://fivetran.com/docs/rest-api/certificates#retrieveadestinationcertificatedetails) | v1 | [DestinationCertificateDetailsService](https://pkg.go.dev/github.com/fivetran/go-fivetran/certificates#DestinationCertificateDetailsService)
[Retrieve a destination fingerprint details](https://fivetran.com/docs/rest-api/certificates#retrieveadestinationfingerprintdetails) | v1 | [DestinationFingerprintDetailsService](https://pkg.go.dev/github.com/fivetran/go-fivetran/fingerprints#DestinationFingerprintDetailsService)
[Revoke a destination certificate](https://fivetran.com/docs/rest-api/certificates#revokeadestinationcertificate) | v1 | [DestinationCertificateRevokeService](https://pkg.go.dev/github.com/fivetran/go-fivetran/certificates#DestinationCertificateRevokeService)
[Revoke a destination fingerprint](https://fivetran.com/docs/rest-api/certificates#revokeadestinationfingerprint) | v1 | [DestinationFingerprintRevokeService](https://pkg.go.dev/github.com/fivetran/go-fivetran/fingerprints#DestinationFingerprintRevokeService)

### [Log Service Management](https://fivetran.com/docs/rest-api/log-service-management)

REST API Endpoint | REST API Version | SDK Service/Config
--- | --- | ---
[List all Log Services](https://fivetran.com/docs/rest-api/api-reference/log-service-management/list-log-services) | v1 | [ExternalLoggingListService](https://pkg.go.dev/github.com/fivetran/go-fivetran#ExternalLoggingListService)
[Create a Log Service](https://fivetran.com/docs/rest-api/log-service-management#createalogservice) | v1 | [ExternalLoggingCreateService](https://pkg.go.dev/github.com/fivetran/go-fivetran#ExternalLoggingCreateService)
[Retrieve Log Service Details](https://fivetran.com/docs/rest-api/log-service-management#retrievelogservicedetails) | v1 | [ExternalLoggingDetailsService](https://pkg.go.dev/github.com/fivetran/go-fivetran#ExternalLoggingDetailsService)
[Update a Log Service](https://fivetran.com/docs/rest-api/log-service-management#updatealogservice) | v1 | [ExternalLoggingModifyService](https://pkg.go.dev/github.com/fivetran/go-fivetran#ExternalLoggingModifyService)
[Delete a Log Service](https://fivetran.com/docs/rest-api/log-service-management#deletealogservice) | v1 | [ExternalLoggingDeleteService](https://pkg.go.dev/github.com/fivetran/go-fivetran#ExternalLoggingDeleteService)
[Run Log Service Setup Tests](https://fivetran.com/docs/rest-api/log-service-management#runlogservicesetuptests) | v1 | [ExternalLoggingSetupTestsService](https://pkg.go.dev/github.com/fivetran/go-fivetran#ExternalLoggingSetupTestsService)

### [Metadata API](https://fivetran.com/docs/rest-api/metadata#metadataapi)

REST API Endpoint | REST API Version | SDK Service/Config
--- | --- | ---
[Retrieve schema metadata](https://fivetran.com/docs/rest-api/metadata#retrieveschemametadata) | v1 | [MetadataSchemaListService](https://pkg.go.dev/github.com/fivetran/go-fivetran#MetadataSchemaListService)
[Retrieve table metadata](https://fivetran.com/docs/rest-api/metadata#retrievetablemetadata) | v1 | [MetadataTablesListService](https://pkg.go.dev/github.com/fivetran/go-fivetran#MetadataTablesListService)
[Retrieve column metadata](https://fivetran.com/docs/rest-api/metadata#retrievecolumnmetadata) | v1 | [MetadataColumnListService](https://pkg.go.dev/github.com/fivetran/go-fivetran#MetadataColumnListService)

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

### [Team Management Connection memberships](https://fivetran.com/docs/rest-api/teams#connectionmemberships)

REST API Endpoint | REST API Version | SDK Service/Config
--- | --- | ---
[List all connection memberships](https://fivetran.com/docs/rest-api/teams#listallconnectionmemberships) | v1 | [TeamConnectionMembershipsListService](https://pkg.go.dev/github.com/fivetran/go-fivetran#TeamConnectionMembershipsListService)
[Retrieve connection membership](https://fivetran.com/docs/rest-api/teams#retrieveconnectionmembership) | v1 | [TeamConnectionMembershipDetailsService](https://pkg.go.dev/github.com/fivetran/go-fivetran#TeamConnectionMembershipDetailsService)
[Add connection membership](https://fivetran.com/docs/rest-api/teams#addconnectionmembership) | v1 | [TeamConnectionMembershipCreateService](https://pkg.go.dev/github.com/fivetran/go-fivetran#TeamConnectionMembershipCreateService)
[Update connection membership](https://fivetran.com/docs/rest-api/teams#updateconnectionmembership) | v1 | [TeamConnectionMembershipModifyService](https://pkg.go.dev/github.com/fivetran/go-fivetran#TeamConnectionMembershipModifyService)
[Delete connection membership](https://fivetran.com/docs/rest-api/teams#deleteconnectionmembership) | v1 | [TeamConnectionMembershipDeleteService](https://pkg.go.dev/github.com/fivetran/go-fivetran#TeamConnectionMembershipDeleteService)

### [Team Management Group memberships](https://fivetran.com/docs/rest-api/teams#groupmemberships)

REST API Endpoint | REST API Version | SDK Service/Config
--- | --- | ---
[List all group memberships](https://fivetran.com/docs/rest-api/teams#listallgroupmemberships) | v1 | [TeamGroupMembershipsService](https://pkg.go.dev/github.com/fivetran/go-fivetran#TeamGroupMembershipsService)
[Retrieve group membership](https://fivetran.com/docs/rest-api/teams#retrievegroupmembership) | v1 | [TeamGroupMembershipDetailsService](https://pkg.go.dev/github.com/fivetran/go-fivetran#TeamGroupMembershipDetailsService)
[Add group membership](https://fivetran.com/docs/rest-api/teams#addgroupmembership) | v1 | [TeamGroupMembershipCreateService](https://pkg.go.dev/github.com/fivetran/go-fivetran#TeamGroupMembershipCreateService)
[Update group membership](https://fivetran.com/docs/rest-api/teams#updategroupmembership) | v1 | [TeamGroupMembershipModifyService](https://pkg.go.dev/github.com/fivetran/go-fivetran#TeamGroupMembershipModifyService)
[Delete group membership](https://fivetran.com/docs/rest-api/teams#deletegroupmembership) | v1 | [TeamGroupMembershipDeleteService](https://pkg.go.dev/github.com/fivetran/go-fivetran#TeamGroupMembershipDeleteService)

### [Private Links Management](https://fivetran.com/docs/rest-api/private-links-management)

REST API Endpoint | REST API Version | SDK Service/Config
--- | --- | ---
[Create a Private Link](https://fivetran.com/docs/rest-api/private-links-management#createaprivatelink) | v1 |  [PrivateLinksCreateService](https://pkg.go.dev/github.com/fivetran/go-fivetran#PrivateLinksCreateService)
[List all Private Links within Account](https://fivetran.com/docs/rest-api/private-links-management#listallprivatelinkswithinaccount) | v1 |  [PrivateLinkListService](https://pkg.go.dev/github.com/fivetran/go-fivetran#PrivateLinkListService)
[Retrieve Private Link Details](https://fivetran.com/docs/rest-api/private-links-management#retrieveprivatelinkdetails) | v1 |  [PrivateLinksDetailsService](https://pkg.go.dev/github.com/fivetran/go-fivetran#PrivateLinksDetailsService)
[Update a Private Link](https://fivetran.com/docs/rest-api/private-links-management#updateaprivatelink) | v1 |  [PrivateLinksModifyService](https://pkg.go.dev/github.com/fivetran/go-fivetran#PrivateLinksModifyService)
[Delete a Private Link](https://fivetran.com/docs/rest-api/private-links-management#deleteaprivatelink) | v1 |  [PrivateLinksDeleteService](https://pkg.go.dev/github.com/fivetran/go-fivetran#PrivateLinksDeleteService)

### [Proxy Agents Management](https://fivetran.com/docs/rest-api/proxy-management)

REST API Endpoint | REST API Version | SDK Service/Config
--- | --- | ---
[Create a Proxy Agent](/docs/rest-api/proxy-management#createaproxyagent) | v1 |  [ProxyCreateService](https://pkg.go.dev/github.com/fivetran/go-fivetran#ProxyCreateService)
[List all Proxy Agents](/docs/rest-api/proxy-management#listallproxyagents) | v1 |  [ProxyListService](https://pkg.go.dev/github.com/fivetran/go-fivetran#ProxyListService)
[Retrieve Proxy Agent Details](/docs/rest-api/proxy-management#retrieveproxyagentdetails) | v1 |  [ProxyDetailsService](https://pkg.go.dev/github.com/fivetran/go-fivetran#ProxyDetailsService)
[Delete a Proxy Agent](/docs/rest-api/proxy-management#deleteaproxyagent) | v1 |  [ProxyDeleteService](https://pkg.go.dev/github.com/fivetran/go-fivetran#ProxyDeleteService)
[Return all connections attached to the proxy agent](/docs/rest-api/proxy-management#returnsallconnectionsattachedtotheproxyagent) | v1 |  [ProxyConnectionMembershipsListService](https://pkg.go.dev/github.com/fivetran/go-fivetran#ProxyConnectionMembershipsListService)

### [Hybrid Deployment Agent Management](https://fivetran.com/docs/rest-api/hybrid-deployment-agent-management#hybriddeploymentagentmanagement)

REST API Endpoint | REST API Version | SDK Service/Config
--- | --- | ---
[Create a Hybrid Deployment Agent](https://fivetran.com/docs/rest-api/hybrid-deployment-agent-management#createahybriddeploymentagent) | v1 |  [HybridDeploymentAgentCreateService](https://pkg.go.dev/github.com/fivetran/go-fivetran#HybridDeploymentAgentCreateService)
[List Hybrid Deployment Agents](/docs/rest-api/hybrid-deployment-agent-management#listhybriddeploymentagents) | v1 |  [HybridDeploymentAgentListService](https://pkg.go.dev/github.com/fivetran/go-fivetran#HybridDeploymentAgentListService)
[Retrieve Hybrid Deployment Agent Details](/docs/rest-api/hybrid-deployment-agent-management#retrievehybriddeploymentagentdetails) | v1 |  [HybridDeploymentAgentDetailsService](https://pkg.go.dev/github.com/fivetran/go-fivetran#HybridDeploymentAgentDetailsService)
[Delete a Hybrid Deployment Agent](/docs/rest-api/lhybrid-deployment-agent-management#deleteahybriddeploymentagent) | v1 |  [HybridDeploymentAgentDeleteService](https://pkg.go.dev/github.com/fivetran/go-fivetran#HybridDeploymentAgentDeleteService)
[Regenerate authentication keys for a Hybrid Deployment Agent](/docs/rest-api/hybrid-deployment-agent-management#regeneratekeys) | v1 |  [HybridDeploymentAgentReAuthService](https://pkg.go.dev/github.com/fivetran/go-fivetran#HybridDeploymentAgentReAuthService)

## Support

Please get in touch with us through our [Support Portal](https://support.fivetran.com/) if you 
have any comments, suggestions, support requests, or bug reports.  
