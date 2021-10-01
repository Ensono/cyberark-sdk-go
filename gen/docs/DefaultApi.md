# \DefaultApi

All URIs are relative to *http://example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdHocConnectthroughPSM**](DefaultApi.md#AdHocConnectthroughPSM) | **Post** /PasswordVault/api/Accounts/AdHocConnect | Ad-Hoc Connect through PSM
[**AddDiscoveredAccountsV108**](DefaultApi.md#AddDiscoveredAccountsV108) | **Post** /PasswordVault/api/DiscoveredAccounts | Add Discovered Accounts (v10.8+)
[**AddSafe**](DefaultApi.md#AddSafe) | **Post** /PasswordVault/api/Safes | Add Safe
[**AddSafeMember**](DefaultApi.md#AddSafeMember) | **Post** /PasswordVault/api/Safes/{Safe}/Members | Add Safe Member
[**ChangePassword**](DefaultApi.md#ChangePassword) | **Put** /PasswordVault/WebServices/PIMServices.svc/Accounts/{AccountID}/ChangeCredentials | Change Password
[**ChangePasswordImmediately**](DefaultApi.md#ChangePasswordImmediately) | **Post** /PasswordVault/API/Accounts/{AccountID}/Change | Change Password Immediately
[**ChangePasswordSetNextPassword**](DefaultApi.md#ChangePasswordSetNextPassword) | **Post** /passwordvault/api/Accounts/{AccountID}/SetNextPassword | Change Password, Set Next Password
[**ChangePasswordintheVaultOnly**](DefaultApi.md#ChangePasswordintheVaultOnly) | **Post** /PasswordVault/api/Accounts/{AccountID}/Password/Update | Change Password in the Vault Only
[**CheckInanExclusiveAccount**](DefaultApi.md#CheckInanExclusiveAccount) | **Post** /PasswordVault/API/Accounts/{AccountID}/CheckIn | Check In an Exclusive Account
[**ConnectUsingPSM**](DefaultApi.md#ConnectUsingPSM) | **Post** /PasswordVault/API/Accounts/{AccountID}/PSMConnect | Connect Using PSM
[**DeleteSafe**](DefaultApi.md#DeleteSafe) | **Delete** /PasswordVault/api/Safes/{Safe} | Delete Safe
[**GetAllSafes**](DefaultApi.md#GetAllSafes) | **Get** /PasswordVault/api/Safes | Get All Safes
[**GetDiscoveredAccountDetails**](DefaultApi.md#GetDiscoveredAccountDetails) | **Get** /passwordvault/api/DiscoveredAccounts/{AccountID} | Get Discovered Account Details
[**GetDiscoveredAccounts**](DefaultApi.md#GetDiscoveredAccounts) | **Get** /passwordvault/api/DiscoveredAccounts | Get Discovered Accounts
[**GetJustinTimeAccess**](DefaultApi.md#GetJustinTimeAccess) | **Post** /PasswordVault/api/Accounts/{AccountID}/grantAdministrativeAccess | Get Just in Time Access
[**GetPasswordValue**](DefaultApi.md#GetPasswordValue) | **Post** /PasswordVault/api/Accounts/{AccountID}/Password/Retrieve | Get Password Value
[**GetSafeAccountGroups**](DefaultApi.md#GetSafeAccountGroups) | **Get** /PasswordVault/API/Safes/{Safe}/AccountGroups | Get Safe Account Groups
[**GetSafeDetails**](DefaultApi.md#GetSafeDetails) | **Get** /PasswordVault/api/Safes/{Safe} | Get Safe Details
[**GetSafeMembers**](DefaultApi.md#GetSafeMembers) | **Get** /PasswordVault/api/Safes/{Safe}/Members | Get Safe Members
[**GetSafebyPlatformID**](DefaultApi.md#GetSafebyPlatformID) | **Get** /PasswordVault/api/Platforms/{PlatformName}/Safes | Get Safe by Platform ID
[**ReconcilePassword**](DefaultApi.md#ReconcilePassword) | **Post** /PasswordVault/API/Accounts/{AccountID}/Reconcile | Reconcile Password
[**VerifyPassword**](DefaultApi.md#VerifyPassword) | **Post** /PasswordVault/API/Accounts/{AccountID}/Verify | Verify Password



## AdHocConnectthroughPSM

> AdHocConnectthroughPSM(ctx).Authorization(authorization).ContentType(contentType).Body(body).Execute()

Ad-Hoc Connect through PSM



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    authorization := "{{CyberArkLogonResult}}" // string | 
    contentType := "application/json" // string | 
    body := "{
	"UserName":"<User Name>",
	"Secret":"<password>",
	"Address":"<Address>",
	"PlatformId":"<Secure Connect Platform>",
	"extraFields":{},
	"PSMConnectPrerequisites": {
		"ConnectionComponent":"<Connection Component ID>",
		"ConnectionType":"<RDPFile or PSMGW>"
	}
}" // string | This method allows you to connect through PSM without using an existing account, by returning settings that can be used with an RDP client application or for the HTML5 gateway.  You must enable Privileged Session Monitoring and ad-hoc connection via PVWA configuration. For more details, see [Ad Hoc Connections](https://docs.cyberark.com/Product-Doc/OnlineHelp/PAS/Latest/en/Content/PASIMP/Configuring-Secure-Connect.htm).

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.AdHocConnectthroughPSM(context.Background()).Authorization(authorization).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AdHocConnectthroughPSM``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdHocConnectthroughPSMRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **contentType** | **string** |  | 
 **body** | **string** | This method allows you to connect through PSM without using an existing account, by returning settings that can be used with an RDP client application or for the HTML5 gateway.  You must enable Privileged Session Monitoring and ad-hoc connection via PVWA configuration. For more details, see [Ad Hoc Connections](https://docs.cyberark.com/Product-Doc/OnlineHelp/PAS/Latest/en/Content/PASIMP/Configuring-Secure-Connect.htm). | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddDiscoveredAccountsV108

> AddDiscoveredAccountsV108(ctx).Authorization(authorization).ContentType(contentType).Body(body).Execute()

Add Discovered Accounts (v10.8+)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    authorization := "{{CyberArkLogonResult}}" // string | 
    contentType := "application/json" // string | 
    body := "[
    {
        "userName": "{{Username}}",
        "address": "{{Address}}",
        "platformTypeAccountProperties": {
            "SID": "S-1-5-21-304654729-3147011263-1431158397-3154"
        },
        "accountEnabled": true,
        "osGroups": "Backup Operators,IIS_IUSRS,Network Configuration Operators",
        "platformType": "{{PlatformName}}",
        "domain": "example.com",
        "lastLogonDateTime": "1530635686",
        "lastPasswordSetDateTime": "1530635786",
        "passwordNeverExpires": false,
        "osVersion": "Windows Server 2012 R2 Standard",
        "privileged": false,
        "userDisplayName": "User Display Name",
        "description": "User Description",
        "passwordExpirationDateTime": "1530645686",
        "osFamily": "Server",
        "OrganizationalUnit": "CN=Users,DC=example,DC=com",
        "additionalProperties": {
            "Port": 445,
            "UserDN": "CN=user1,CN=Users,DC=example,DC=com"
        },
        "Dependencies": [
            {
                "name": "ServiceDep",
                "address": "{{Address}}",
                "type": "Windows Service"
            }
        ]
    },
    {
        "userName": "amazon_accesskey_user",
        "address": "aws.com",
        "discoveryDate": "2018-05-03T13:00:00Z",
        "platformType": "AWS Access Keys",
        "privileged": true,
        "platformTypeAccountProperties": {
            "awsAccessKeyID": "ASASASASSA",
            "awsAccountID": "123123123123"
        }
    },
    {
        "userName": "amazon_dashboard_user",
        "address": "aws.com",
        "discoveryDate": "2018-05-03T13:00:00Z",
        "platformType": "AWS",
        "privileged": true,
        "platformTypeAccountProperties": {
            "awsAccountID": "123123123123"
        }
    },
    {
        "userName": string,
        "address": string,
        "discoveryDate": "1530635689",
        "platformType": "Azure Password Management",
        "accountEnabled": true,
        "privileged": true,
        "privilegedCriteria": string,
        "platformTypeAccountProperties": {
            "activeDirectoryID": string
        }
    }
]" // string | This method adds newly discovered accounts and their dependencies.  | Action | Description | | --- | --- | | Discovered account | Discovered accounts are added to the Pending Accounts list. | | Discovered dependency | Discovered dependencies are added to the Pending Accounts list or are automatically added to the corresponding master account that has already been onboarded. | | Automatic onboarding | Discovered accounts are onboarded according to matching onboarding rules. If automatic onboarding fails for any reason, the account is added to the Pending Accounts list. |  Currently, Private SSH Keys are not supported.  The user who runs this web service requires the following users and permissions:  To add pending accounts:  * **User:** Owner of PasswordManager_Pending Safe * **Permissions:**   * Add account   * List files   * Update account properties  To onboard the account:  * **User:** Owner of the target Safe of the onboarding rule. * **Permissions:**   * Add account   * Update account properties   * Initiate CPM account management operations

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.AddDiscoveredAccountsV108(context.Background()).Authorization(authorization).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AddDiscoveredAccountsV108``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddDiscoveredAccountsV108Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **contentType** | **string** |  | 
 **body** | **string** | This method adds newly discovered accounts and their dependencies.  | Action | Description | | --- | --- | | Discovered account | Discovered accounts are added to the Pending Accounts list. | | Discovered dependency | Discovered dependencies are added to the Pending Accounts list or are automatically added to the corresponding master account that has already been onboarded. | | Automatic onboarding | Discovered accounts are onboarded according to matching onboarding rules. If automatic onboarding fails for any reason, the account is added to the Pending Accounts list. |  Currently, Private SSH Keys are not supported.  The user who runs this web service requires the following users and permissions:  To add pending accounts:  * **User:** Owner of PasswordManager_Pending Safe * **Permissions:**   * Add account   * List files   * Update account properties  To onboard the account:  * **User:** Owner of the target Safe of the onboarding rule. * **Permissions:**   * Add account   * Update account properties   * Initiate CPM account management operations | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddSafe

> AddSafe(ctx).Authorization(authorization).ContentType(contentType).Body(body).Execute()

Add Safe



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    authorization := "{{CyberArkLogonResult}}" // string | 
    contentType := "application/json" // string | 
    body := "{
    "SafeName": "testSafe1",
    "Description": "Desc1",
    "OLACEnabled": false,
    "ManagingCPM": "PasswordManager",
    "NumberOfVersionsRetention": null,
    "NumberOfDaysRetention": 0,
    "AutoPurgeEnabled": false,
    "Location": "\\"
}" // string | This method adds a new Safe to the Vault.  The user who runs this web service requires **Add Safes** permissions in the Vault.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.AddSafe(context.Background()).Authorization(authorization).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AddSafe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddSafeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **contentType** | **string** |  | 
 **body** | **string** | This method adds a new Safe to the Vault.  The user who runs this web service requires **Add Safes** permissions in the Vault. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddSafeMember

> AddSafeMember(ctx, safe).Authorization(authorization).ContentType(contentType).Body(body).Execute()

Add Safe Member



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    authorization := "{{CyberArkLogonResult}}" // string | Session Authorization Token
    contentType := "application/json" // string | 
    safe := "Safe" // string | 
    body := "{
    "memberName": "JonDoe",
    "searchIn": "Vault",
    "membershipExpirationDate": 123456,
    "permissions": {
        "useAccounts": false,
        "retrieveAccounts": false,
        "listAccounts": false,
        "addAccounts": false,
        "updateAccountContent": false,
        "updateAccountProperties": false,
        "initiateCPMAccountManagementOperations": false,
        "specifyNextAccountContent": false,
        "renameAccounts": false,
        "deleteAccounts": false,
        "unlockAccounts": false,
        "manageSafe": false,
        "manageSafeMembers": false,
        "backupSafe": false,
        "viewAuditLog": false,
        "viewSafeMembers": false,
        "accessWithoutConfirmation": false,
        "createFolders": false,
        "deleteFolders": false,
        "moveAccountsAndFolders": false,
        "requestsAuthorizationLevel1": false,
        "requestsAuthorizationLevel2": false
    }
}" // string | This method adds an existing user or group as a Safe member.  The user who runs this web service requires the following permissions in the Vault:  * Manage Safe Members

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.AddSafeMember(context.Background(), safe).Authorization(authorization).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AddSafeMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**safe** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddSafeMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Session Authorization Token | 
 **contentType** | **string** |  | 

 **body** | **string** | This method adds an existing user or group as a Safe member.  The user who runs this web service requires the following permissions in the Vault:  * Manage Safe Members | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChangePassword

> ChangePassword(ctx, accountID).Authorization(authorization).ContentType(contentType).Body(body).Execute()

Change Password



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    authorization := "{{CyberArkLogonResult}}" // string | 
    contentType := "application/json" // string | 
    accountID := "AccountID" // string | 
    body := "{
	"ImmediateChangeByCPM": <true/false>,
	"ChangeCredsForGroup": <true/false>
}" // string | This method marks the account for an immediate password change by the CPM to a new random password.  The user who runs this web service requires the following permission in the Safe where the privileged account is stored:  * Initiate CPM password management operations

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ChangePassword(context.Background(), accountID).Authorization(authorization).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ChangePassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangePasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **contentType** | **string** |  | 

 **body** | **string** | This method marks the account for an immediate password change by the CPM to a new random password.  The user who runs this web service requires the following permission in the Safe where the privileged account is stored:  * Initiate CPM password management operations | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChangePasswordImmediately

> ChangePasswordImmediately(ctx, accountID).Authorization(authorization).ContentType(contentType).Body(body).Execute()

Change Password Immediately



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    authorization := "{{CyberArkLogonResult}}" // string | 
    contentType := "application/json" // string | 
    accountID := "AccountID" // string | 
    body := "{
	"ChangeEntireGroup" : true
}" // string | This method marks an account for an immediate credentials change by the CPM to a new random value.  The user who runs this web service requires the following permission in the Safe where the privileged account is stored:  * Initiate CPM password management operations

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ChangePasswordImmediately(context.Background(), accountID).Authorization(authorization).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ChangePasswordImmediately``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangePasswordImmediatelyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **contentType** | **string** |  | 

 **body** | **string** | This method marks an account for an immediate credentials change by the CPM to a new random value.  The user who runs this web service requires the following permission in the Safe where the privileged account is stored:  * Initiate CPM password management operations | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChangePasswordSetNextPassword

> ChangePasswordSetNextPassword(ctx, accountID).Authorization(authorization).ContentType(contentType).Body(body).Execute()

Change Password, Set Next Password



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    authorization := "{{CyberArkLogonResult}}" // string | 
    contentType := "application/json" // string | 
    accountID := "AccountID" // string | 
    body := "{
	"ChangeImmediately" : true,
	"NewCredentials": "<credentials>"
}" // string | This method enables users to set the account's credentials to use for the next CPM change.  The user who runs this web service requires the following permissions in the Safe where the privileged account is stored:  * Initiate CPM password management operations * Specify next password value

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ChangePasswordSetNextPassword(context.Background(), accountID).Authorization(authorization).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ChangePasswordSetNextPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangePasswordSetNextPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **contentType** | **string** |  | 

 **body** | **string** | This method enables users to set the account&#39;s credentials to use for the next CPM change.  The user who runs this web service requires the following permissions in the Safe where the privileged account is stored:  * Initiate CPM password management operations * Specify next password value | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChangePasswordintheVaultOnly

> ChangePasswordintheVaultOnly(ctx, accountID).Authorization(authorization).ContentType(contentType).Body(body).Execute()

Change Password in the Vault Only



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    authorization := "{{CyberArkLogonResult}}" // string | 
    contentType := "application/json" // string | 
    accountID := "AccountID" // string | 
    body := "{
	"ChangeEntireGroup": false,
	"NewCredentials": "<string>"
}" // string | This method enables users to set account credentials and change them in the Vault. This will not affect credentials on the target device.  The user who runs this web service requires **Update password value** permission in the Safe where the privileged account is stored.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ChangePasswordintheVaultOnly(context.Background(), accountID).Authorization(authorization).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ChangePasswordintheVaultOnly``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangePasswordintheVaultOnlyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **contentType** | **string** |  | 

 **body** | **string** | This method enables users to set account credentials and change them in the Vault. This will not affect credentials on the target device.  The user who runs this web service requires **Update password value** permission in the Safe where the privileged account is stored. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckInanExclusiveAccount

> CheckInanExclusiveAccount(ctx, accountID).Authorization(authorization).ContentType(contentType).Execute()

Check In an Exclusive Account



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    authorization := "{{CyberArkLogonResult}}" // string | 
    contentType := "application/json" // string | 
    accountID := "AccountID" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CheckInanExclusiveAccount(context.Background(), accountID).Authorization(authorization).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CheckInanExclusiveAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckInanExclusiveAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **contentType** | **string** |  | 


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectUsingPSM

> ConnectUsingPSM(ctx, accountID).Authorization(authorization).ContentType(contentType).Accept(accept).Body(body).Execute()

Connect Using PSM



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    authorization := "{{CyberArkLogonResult}}" // string | 
    contentType := "application/json" // string | 
    accept := "RDP" // string | 
    accountID := "AccountID" // string | 
    body := "{
	"reason":"<Reason>",
	"TicketingSystemName":"<Ticketing system>",
	"TicketId":"<Ticketid>",
	"ConnectionComponent":"<Connection component id>",
	"ConnectionParams": {
		"<Connection parameter name>": {
			"value":"<Connection parameter value>",
			"ShouldSave":<true\false>
		},
		"<Connection parameter name>": {
			"value":"<Connection parameter value>",
			"ShouldSave":<true\false>
		}
	}
}" // string | This method enables you to connect to an account through PSM (PSMConnect) using a connection method defined in the PVWA.  A response header defines which connection method is returned.  For more information, refer to [Privileged Session Management Interface](https://docs.cyberark.com/Product-Doc/OnlineHelp/PAS/Latest/en/Content/PASIMP/Configuring-the-Privileged-Session-Management-Interface.htm).  ## Header Parameter  Parameter: Accept  Type: String  Description: The table below describes the expected response format depending on the value of the Accept header in the request, per connection method configuration (RDP File or PSM Gateway).  | PVWA configuration | Optional values | Connection method | | --- | --- | --- | | RDP | application/json, application/octet-stream (default), `*_/ *` | RDPFile (JSON), RDPFile (octet-stream raw) | | PSMGW | `* / *` | PSMGW (JSON) - Returns the HTML5 connection data. |  **Note:** To use PSMGW, PSMGW must be configured before using this REST API in order to receive a PSMGW response.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ConnectUsingPSM(context.Background(), accountID).Authorization(authorization).ContentType(contentType).Accept(accept).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ConnectUsingPSM``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectUsingPSMRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **contentType** | **string** |  | 
 **accept** | **string** |  | 

 **body** | **string** | This method enables you to connect to an account through PSM (PSMConnect) using a connection method defined in the PVWA.  A response header defines which connection method is returned.  For more information, refer to [Privileged Session Management Interface](https://docs.cyberark.com/Product-Doc/OnlineHelp/PAS/Latest/en/Content/PASIMP/Configuring-the-Privileged-Session-Management-Interface.htm).  ## Header Parameter  Parameter: Accept  Type: String  Description: The table below describes the expected response format depending on the value of the Accept header in the request, per connection method configuration (RDP File or PSM Gateway).  | PVWA configuration | Optional values | Connection method | | --- | --- | --- | | RDP | application/json, application/octet-stream (default), &#x60;*_/ *&#x60; | RDPFile (JSON), RDPFile (octet-stream raw) | | PSMGW | &#x60;* / *&#x60; | PSMGW (JSON) - Returns the HTML5 connection data. |  **Note:** To use PSMGW, PSMGW must be configured before using this REST API in order to receive a PSMGW response. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSafe

> DeleteSafe(ctx, safe).Authorization(authorization).ContentType(contentType).Execute()

Delete Safe



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    authorization := "{{CyberArkLogonResult}}" // string | 
    contentType := "application/json" // string | 
    safe := "Safe" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteSafe(context.Background(), safe).Authorization(authorization).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSafe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**safe** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSafeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **contentType** | **string** |  | 


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllSafes

> Model200 GetAllSafes(ctx).Search(search).Offset(offset).Limit(limit).Sort(sort).IncludeAccounts(includeAccounts).ExtendedDetails(extendedDetails).Authorization(authorization).ContentType(contentType).Execute()

Get All Safes



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    search := "searchWord" // string | Searches according to the Safe name. Search is performed according to the REST standard (search=\"search word\").
    offset := float32(8.14) // float32 | Offset of the first Safe that is returned in the collection of results.
    limit := float32(8.14) // float32 | The maximum number of Safes that are returned. When used together with the offset parameter, this value determines the number of Safes to return, starting from the first Safe that is returned.
    sort := "safeName asc" // string | Sorts according to the safeName property in ascending order (default) or descending order to control the sort direction.
    includeAccounts := "false" // string | Whether or not to return accounts for each Safe as part of the response. If not sent, the value is False.
    extendedDetails := "true" // string | Whether or not to return all Safe details or only safeName as part of the response. If not sent, the value is True.
    authorization := "{{CyberArkLogonResult}}" // string | 
    contentType := "application/json" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetAllSafes(context.Background()).Search(search).Offset(offset).Limit(limit).Sort(sort).IncludeAccounts(includeAccounts).ExtendedDetails(extendedDetails).Authorization(authorization).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetAllSafes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllSafes`: Model200
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetAllSafes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllSafesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | Searches according to the Safe name. Search is performed according to the REST standard (search&#x3D;\&quot;search word\&quot;). | 
 **offset** | **float32** | Offset of the first Safe that is returned in the collection of results. | 
 **limit** | **float32** | The maximum number of Safes that are returned. When used together with the offset parameter, this value determines the number of Safes to return, starting from the first Safe that is returned. | 
 **sort** | **string** | Sorts according to the safeName property in ascending order (default) or descending order to control the sort direction. | 
 **includeAccounts** | **string** | Whether or not to return accounts for each Safe as part of the response. If not sent, the value is False. | 
 **extendedDetails** | **string** | Whether or not to return all Safe details or only safeName as part of the response. If not sent, the value is True. | 
 **authorization** | **string** |  | 
 **contentType** | **string** |  | 

### Return type

[**Model200**](Model200.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDiscoveredAccountDetails

> GetDiscoveredAccountDetails(ctx, accountID).Authorization(authorization).ContentType(contentType).Execute()

Get Discovered Account Details



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    authorization := "{{CyberArkLogonResult}}" // string | Session Authorization Token
    contentType := "application/json" // string | 
    accountID := "AccountID" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetDiscoveredAccountDetails(context.Background(), accountID).Authorization(authorization).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetDiscoveredAccountDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDiscoveredAccountDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Session Authorization Token | 
 **contentType** | **string** |  | 


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDiscoveredAccounts

> GetDiscoveredAccounts(ctx).Filter(filter).Search(search).SearchType(searchType).Offset(offset).Limit(limit).Authorization(authorization).ContentType(contentType).Execute()

Get Discovered Accounts



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    filter := "platformType eq Windows Server Local AND privileged eq true AND accountEnabled eq true" // string | search accounts using platformType, privileged, and/or accountEnabled values
    search := "admin" // string | search is supported for username and address
    searchType := "contains" // string | keyword is contained (contains, DEFAULT) or beginning (startswith)
    offset := float32(8.14) // float32 | the offset of the first returned account in the list of results
    limit := float32(8.14) // float32 | the maximum number of accounts to return (maximum value allowed is 1000)
    authorization := "{{CyberArkLogonResult}}" // string | Session Authorization Token
    contentType := "application/json" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetDiscoveredAccounts(context.Background()).Filter(filter).Search(search).SearchType(searchType).Offset(offset).Limit(limit).Authorization(authorization).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetDiscoveredAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDiscoveredAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | search accounts using platformType, privileged, and/or accountEnabled values | 
 **search** | **string** | search is supported for username and address | 
 **searchType** | **string** | keyword is contained (contains, DEFAULT) or beginning (startswith) | 
 **offset** | **float32** | the offset of the first returned account in the list of results | 
 **limit** | **float32** | the maximum number of accounts to return (maximum value allowed is 1000) | 
 **authorization** | **string** | Session Authorization Token | 
 **contentType** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJustinTimeAccess

> GetJustinTimeAccess(ctx, accountID).Authorization(authorization).ContentType(contentType).Execute()

Get Just in Time Access



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    authorization := "{{CyberArkLogonResult}}" // string | 
    contentType := "application/json" // string | 
    accountID := "AccountID" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetJustinTimeAccess(context.Background(), accountID).Authorization(authorization).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetJustinTimeAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJustinTimeAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **contentType** | **string** |  | 


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPasswordValue

> GetPasswordValue(ctx, accountID).Authorization(authorization).ContentType(contentType).Body(body).Execute()

Get Password Value



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    authorization := "{{CyberArkLogonResult}}" // string | 
    contentType := "application/json" // string | 
    accountID := "AccountID" // string | 
    body := "{
	"reason":"<Reason>",
	"TicketingSystemName": "<Ticketing system>",
	"TicketId": "<Ticketid>",
	"Version": <version number>,
	"ActionType": "<action type - show\copy\connect>
	"isUse": <true\false>,
	"Machine": "<my remote machine address>"
}" // string | This method enables users to retrieve the password or SSH key of an existing account that is identified by its Account ID. It enables users to specify a reason and ticket ID, if required.  **Note:** The ability to retrieve credentials using this REST API is intended for human use only and is not recommended for applications or automated processes, where application-based authentication is required.  For application or automated processes use cases, refer to the [AAM Credential Providers Online Help](https://docs.cyberark.com/Product-Doc/OnlineHelp/AAM-CP/Latest/en/Default.htm).

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetPasswordValue(context.Background(), accountID).Authorization(authorization).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetPasswordValue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPasswordValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **contentType** | **string** |  | 

 **body** | **string** | This method enables users to retrieve the password or SSH key of an existing account that is identified by its Account ID. It enables users to specify a reason and ticket ID, if required.  **Note:** The ability to retrieve credentials using this REST API is intended for human use only and is not recommended for applications or automated processes, where application-based authentication is required.  For application or automated processes use cases, refer to the [AAM Credential Providers Online Help](https://docs.cyberark.com/Product-Doc/OnlineHelp/AAM-CP/Latest/en/Default.htm). | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSafeAccountGroups

> GetSafeAccountGroups(ctx, safe).Execute()

Get Safe Account Groups



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    safe := "Safe" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetSafeAccountGroups(context.Background(), safe).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSafeAccountGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**safe** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSafeAccountGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSafeDetails

> Model200 GetSafeDetails(ctx, safe).Authorization(authorization).ContentType(contentType).Execute()

Get Safe Details



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    authorization := "{{CyberArkLogonResult}}" // string | 
    contentType := "application/json" // string | 
    safe := "Safe" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetSafeDetails(context.Background(), safe).Authorization(authorization).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSafeDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSafeDetails`: Model200
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSafeDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**safe** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSafeDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **contentType** | **string** |  | 


### Return type

[**Model200**](Model200.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSafeMembers

> GetSafeMembers(ctx, safe).Filter(filter).Search(search).Offset(offset).Limit(limit).Sort(sort).Authorization(authorization).ContentType(contentType).Execute()

Get Safe Members



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    filter := "memberType eq user AND membershipExpired eq false AND includePreDefinedUsers eq false" // string | Filters are according to the REST standard. Search for Safe members using the following filters. Multiple filters can be applied using the AND operator. memberType, membershipExpired, includePreDefinedUsers
    search := "searchWord" // string | Searches according to the Safe name. Search is performed according to the REST standard (search=\"search word\").
    offset := float32(8.14) // float32 | Offset of the first member that is returned in the collection of results.
    limit := float32(8.14) // float32 | The maximum number of members that are returned. When used together with the offset parameter, this value determines the number of Safes to return, starting from the first Safe that is returned.
    sort := "asc" // string | Sorts according to the memberName property in ascending order (default) or descending order to control the sort direction.
    authorization := "{{CyberArkLogonResult}}" // string | Session Authorization Token
    contentType := "application/json" // string | 
    safe := "Safe" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetSafeMembers(context.Background(), safe).Filter(filter).Search(search).Offset(offset).Limit(limit).Sort(sort).Authorization(authorization).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSafeMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**safe** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSafeMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Filters are according to the REST standard. Search for Safe members using the following filters. Multiple filters can be applied using the AND operator. memberType, membershipExpired, includePreDefinedUsers | 
 **search** | **string** | Searches according to the Safe name. Search is performed according to the REST standard (search&#x3D;\&quot;search word\&quot;). | 
 **offset** | **float32** | Offset of the first member that is returned in the collection of results. | 
 **limit** | **float32** | The maximum number of members that are returned. When used together with the offset parameter, this value determines the number of Safes to return, starting from the first Safe that is returned. | 
 **sort** | **string** | Sorts according to the memberName property in ascending order (default) or descending order to control the sort direction. | 
 **authorization** | **string** | Session Authorization Token | 
 **contentType** | **string** |  | 


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSafebyPlatformID

> Model200 GetSafebyPlatformID(ctx, platformName).Authorization(authorization).ContentType(contentType).Execute()

Get Safe by Platform ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    authorization := "{{CyberArkLogonResult}}" // string | 
    contentType := "application/json" // string | 
    platformName := "PlatformName" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetSafebyPlatformID(context.Background(), platformName).Authorization(authorization).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSafebyPlatformID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSafebyPlatformID`: Model200
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSafebyPlatformID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platformName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSafebyPlatformIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **contentType** | **string** |  | 


### Return type

[**Model200**](Model200.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReconcilePassword

> ReconcilePassword(ctx, accountID).Authorization(authorization).ContentType(contentType).Execute()

Reconcile Password



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    authorization := "{{CyberArkLogonResult}}" // string | 
    contentType := "application/json" // string | 
    accountID := "AccountID" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ReconcilePassword(context.Background(), accountID).Authorization(authorization).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ReconcilePassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReconcilePasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **contentType** | **string** |  | 


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyPassword

> VerifyPassword(ctx, accountID).Authorization(authorization).ContentType(contentType).Execute()

Verify Password



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    authorization := "{{CyberArkLogonResult}}" // string | 
    contentType := "application/json" // string | 
    accountID := "AccountID" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.VerifyPassword(context.Background(), accountID).Authorization(authorization).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.VerifyPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **contentType** | **string** |  | 


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

