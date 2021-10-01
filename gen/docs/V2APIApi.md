# \V2APIApi

All URIs are relative to *http://example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAccount**](V2APIApi.md#AddAccount) | **Post** /PasswordVault/api/Accounts | Add Account
[**AddPendingAccount**](V2APIApi.md#AddPendingAccount) | **Post** /PasswordVault/WebServices/PIMServices.svc/PendingAccounts | Add Pending Account
[**DeleteAccount**](V2APIApi.md#DeleteAccount) | **Delete** /PasswordVault/api/Accounts/{AccountID} | Delete Account
[**ExtendedAccountOverview**](V2APIApi.md#ExtendedAccountOverview) | **Get** /PasswordVault/api/ExtendedAccounts/{AccountID}/overview | Extended Account Overview
[**GetAccountActivity**](V2APIApi.md#GetAccountActivity) | **Get** /PasswordVault/WebServices/PIMServices.svc/Accounts/{AccountID}/Activities | Get Account Activity
[**GetAccountDetails**](V2APIApi.md#GetAccountDetails) | **Get** /PasswordVault/api/Accounts/{AccountID} | Get Account Details
[**GetAccounts**](V2APIApi.md#GetAccounts) | **Get** /PasswordVault/api/Accounts | Get Accounts
[**LinkanAccount**](V2APIApi.md#LinkanAccount) | **Post** /PasswordVault/api/Accounts/{AccountID}/LinkAccount | Link an Account
[**UpdateAccount**](V2APIApi.md#UpdateAccount) | **Patch** /PasswordVault/api/Accounts/{AccountID} | Update Account



## AddAccount

> AddAccount(ctx).Authorization(authorization).ContentType(contentType).Body(body).Execute()

Add Account



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
	"name": "string",
	"address": "string",
	"userName": "string",
	"platformId": "string",
	"safeName": "string",
	"secretType": "password",
	"secret": "string",
	"platformAccountProperties": {
		"LogonDomain": "string",
		"Port": "integer"
	},
	"secretManagement": {
		"automaticManagementEnabled": true,
		"manualManagementReason": "string"
	},
	"remoteMachinesAccess": {
		"remoteMachines": "string",
		"accessRestrictedToRemoteMachines": true
	}
}" // string | This method adds a new privileged account or SSH key to the Vault.  To run this web service, a user requires the following permission in the Vault:  * Add account  AND either  * Update password  OR  * Update password properties  **Note:** You require an additional license to add SSH keys to the Vault. For more information, contact your CyberArk representative.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2APIApi.AddAccount(context.Background()).Authorization(authorization).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2APIApi.AddAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **contentType** | **string** |  | 
 **body** | **string** | This method adds a new privileged account or SSH key to the Vault.  To run this web service, a user requires the following permission in the Vault:  * Add account  AND either  * Update password  OR  * Update password properties  **Note:** You require an additional license to add SSH keys to the Vault. For more information, contact your CyberArk representative. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddPendingAccount

> AddPendingAccount(ctx).Authorization(authorization).ContentType(contentType).Body(body).Execute()

Add Pending Account



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
	"pendingAccount": {
		"UserName":"<user name>",
		"Address":"<address>",
		"AccountDiscoveryDate":"<YYYY-MM-DDThh:mm:ssZ>",
		"AccountEnabled":"<enabled/disabled>", 
		"AccountOSGroups":"<group name>",
		"AccountType":"<domain/local>",
		"Domain":"<domain name>",
		"PasswordNeverExpires":"<true/false>",
		"OSVersion":"<OS version>",
		"OU":"<OU>",
		"AccountCategory":"<Privileged/Non-privileged>",
		"UserDisplayName":"<user display name>",
		"AccountDescription":"<description>",
		"GID":"<GID>",
		"UID":"<UID>",
		"OSType":"<Windows/Unix>",
		"DiscoveryPlatformType":"<platform name>",
		"MachineOSFamily":"<workstation/server>",
		"LastLogonDate":"<YYYY-MM-DDThh:mm:ssZ>",
		"LastPasswordSetDate":"<YYYY-MM-DDThh:mm:ssZ>",
		"AccountExpirationDate":"<YYYY-MM-DDThh:mm:ssZ>",
		"AccountCategoryCriteria":"<criteria>"
	}
}" // string | This method enables an account or SSH key that is discovered by an external scanner to be added as a pending account to the Accounts Feed. This facilitates the privileged account workflow, during which users can identify privileged accounts and determine which are onboarded to the Vault.  **Note:** In order to add SSH keys to the Vault, you require an additional license. For more information, contact your CyberArk representative.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2APIApi.AddPendingAccount(context.Background()).Authorization(authorization).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2APIApi.AddPendingAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddPendingAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **contentType** | **string** |  | 
 **body** | **string** | This method enables an account or SSH key that is discovered by an external scanner to be added as a pending account to the Accounts Feed. This facilitates the privileged account workflow, during which users can identify privileged accounts and determine which are onboarded to the Vault.  **Note:** In order to add SSH keys to the Vault, you require an additional license. For more information, contact your CyberArk representative. | 

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


## DeleteAccount

> DeleteAccount(ctx, accountID).Authorization(authorization).ContentType(contentType).Execute()

Delete Account



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
    resp, r, err := api_client.V2APIApi.DeleteAccount(context.Background(), accountID).Authorization(authorization).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2APIApi.DeleteAccount``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteAccountRequest struct via the builder pattern


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


## ExtendedAccountOverview

> ExtendedAccountOverview(ctx, accountID).Authorization(authorization).ContentType(contentType).Execute()

Extended Account Overview

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
    resp, r, err := api_client.V2APIApi.ExtendedAccountOverview(context.Background(), accountID).Authorization(authorization).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2APIApi.ExtendedAccountOverview``: %v\n", err)
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

Other parameters are passed through a pointer to a apiExtendedAccountOverviewRequest struct via the builder pattern


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


## GetAccountActivity

> Model200 GetAccountActivity(ctx, accountID).Authorization(authorization).ContentType(contentType).Execute()

Get Account Activity



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
    resp, r, err := api_client.V2APIApi.GetAccountActivity(context.Background(), accountID).Authorization(authorization).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2APIApi.GetAccountActivity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountActivity`: Model200
    fmt.Fprintf(os.Stdout, "Response from `V2APIApi.GetAccountActivity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountActivityRequest struct via the builder pattern


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


## GetAccountDetails

> Model200 GetAccountDetails(ctx, accountID).Authorization(authorization).ContentType(contentType).Execute()

Get Account Details



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
    resp, r, err := api_client.V2APIApi.GetAccountDetails(context.Background(), accountID).Authorization(authorization).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2APIApi.GetAccountDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountDetails`: Model200
    fmt.Fprintf(os.Stdout, "Response from `V2APIApi.GetAccountDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountDetailsRequest struct via the builder pattern


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


## GetAccounts

> GetAccounts(ctx).Search(search).SearchType(searchType).Sort(sort).Offset(offset).Limit(limit).Filter(filter).ContentType(contentType).Authorization(authorization).Execute()

Get Accounts



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
    search := "DemoUser" // string | List of keywords to search for in accounts, separated by a space.
    searchType := "contains" // string | Get accounts that either contain or start with the value specified in the Search parameter. Valid values: contains (default) or startswith.
    sort := "UserName" // string | Property or properties by which to sort returned accounts, followed by asc (default) or desc to control sort direction. Separate multiple properties with commas, up to a maximum of three properties.
    offset := float32(8.14) // float32 | Offset of the first account that is returned in the collection of results.
    limit := float32(8.14) // float32 | Maximum number of returned accounts. If not specified, the default value is 50. The maximum number that can be specified is 1000.
    filter := "safeName eq {{Safe}}" // string | Get accounts from a specific safe, using the safe name.
    contentType := "application/json" // string | 
    authorization := "{{CyberArkLogonResult}}" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2APIApi.GetAccounts(context.Background()).Search(search).SearchType(searchType).Sort(sort).Offset(offset).Limit(limit).Filter(filter).ContentType(contentType).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2APIApi.GetAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | List of keywords to search for in accounts, separated by a space. | 
 **searchType** | **string** | Get accounts that either contain or start with the value specified in the Search parameter. Valid values: contains (default) or startswith. | 
 **sort** | **string** | Property or properties by which to sort returned accounts, followed by asc (default) or desc to control sort direction. Separate multiple properties with commas, up to a maximum of three properties. | 
 **offset** | **float32** | Offset of the first account that is returned in the collection of results. | 
 **limit** | **float32** | Maximum number of returned accounts. If not specified, the default value is 50. The maximum number that can be specified is 1000. | 
 **filter** | **string** | Get accounts from a specific safe, using the safe name. | 
 **contentType** | **string** |  | 
 **authorization** | **string** |  | 

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


## LinkanAccount

> LinkanAccount(ctx, accountID).Authorization(authorization).ContentType(contentType).Body(body).Execute()

Link an Account



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
    body := "{
    "Name": "ObjectName",
    "ExtraPassID": "1/2/3",
    "Folder": "Root"
}" // string | This method enables a user to associate a linked account to an existing source account. The linked account can be a Reconcile account, Logon account, or other type of linked account that is defined in the platform configuration.     _**Note:** The type of linked accounts that are allowed are determined by the Platform settings. Each platform can support different types of linked accounts._  To run this service, the user must have the following Safe member authorizations:  * List accounts - For both the Safe of the linked account and the Safe of the source account * Update account properties - For the Safe of the source account

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2APIApi.LinkanAccount(context.Background(), accountID).Authorization(authorization).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2APIApi.LinkanAccount``: %v\n", err)
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

Other parameters are passed through a pointer to a apiLinkanAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Session Authorization Token | 
 **contentType** | **string** |  | 

 **body** | **string** | This method enables a user to associate a linked account to an existing source account. The linked account can be a Reconcile account, Logon account, or other type of linked account that is defined in the platform configuration.     _**Note:** The type of linked accounts that are allowed are determined by the Platform settings. Each platform can support different types of linked accounts._  To run this service, the user must have the following Safe member authorizations:  * List accounts - For both the Safe of the linked account and the Safe of the source account * Update account properties - For the Safe of the source account | 

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


## UpdateAccount

> Model200 UpdateAccount(ctx, accountID).Authorization(authorization).ContentType(contentType).Body(body).Execute()

Update Account



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
    body := "[
	{
		"op": "replace",
		"path": "/address",
		"value": "NewAddress"
	},
	{
		"op": "add",
		"path": "/port",
		"value": "3306"
	},
	{
		"op": "remove",
		"path": "/ticketnumber",
		"value": ""
	}
]" // string | This method updates an existing account's details. It is not mandatory to send all the account’s details. Any changed values sent in the request will be updated. All other properties values will remain the same.  On each property, the following operations can be performed:  * Replace - replace the existing value of a property * Remove (to remove the property from the account) * Add (to add that property to the account)  It is possible to set several properties using the same command using the following structure:  ```json [  {      \"op\": \"replace\",      \"path\": \"/platformaccountproperties\",      \"value\": \"{          \\\"{PropertyID1}\\\":\\\"{Value}\\\",          \\\"{PropertyID2}\\\":\\\"{Value}\\\",          \\\"{PropertyID3}\\\":\\\"{Value}\\\"      }\"  } ] ```  When sending several operations on the same property, only the last operation will affect the property.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2APIApi.UpdateAccount(context.Background(), accountID).Authorization(authorization).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2APIApi.UpdateAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAccount`: Model200
    fmt.Fprintf(os.Stdout, "Response from `V2APIApi.UpdateAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **contentType** | **string** |  | 

 **body** | **string** | This method updates an existing account&#39;s details. It is not mandatory to send all the account’s details. Any changed values sent in the request will be updated. All other properties values will remain the same.  On each property, the following operations can be performed:  * Replace - replace the existing value of a property * Remove (to remove the property from the account) * Add (to add that property to the account)  It is possible to set several properties using the same command using the following structure:  &#x60;&#x60;&#x60;json [  {      \&quot;op\&quot;: \&quot;replace\&quot;,      \&quot;path\&quot;: \&quot;/platformaccountproperties\&quot;,      \&quot;value\&quot;: \&quot;{          \\\&quot;{PropertyID1}\\\&quot;:\\\&quot;{Value}\\\&quot;,          \\\&quot;{PropertyID2}\\\&quot;:\\\&quot;{Value}\\\&quot;,          \\\&quot;{PropertyID3}\\\&quot;:\\\&quot;{Value}\\\&quot;      }\&quot;  } ] &#x60;&#x60;&#x60;  When sending several operations on the same property, only the last operation will affect the property. | 

### Return type

[**Model200**](Model200.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

