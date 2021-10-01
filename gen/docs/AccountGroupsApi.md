# \AccountGroupsApi

All URIs are relative to *http://example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAccountGroup**](AccountGroupsApi.md#AddAccountGroup) | **Post** /API/AccountGroups/ | Add Account Group
[**AddAccounttoAccountGroup**](AccountGroupsApi.md#AddAccounttoAccountGroup) | **Post** /API/AccountGroups/{GroupName}/Members | Add Account to Account Group
[**DeleteMemberfromAccountGroup**](AccountGroupsApi.md#DeleteMemberfromAccountGroup) | **Delete** /PasswordVault/API/AccountGroups/{GroupName}/Members/{AccountID} | Delete Member from Account Group
[**GetAccountGroupMembers**](AccountGroupsApi.md#GetAccountGroupMembers) | **Get** /PasswordVault/API/AccountGroups/{GroupName}/Members | Get Account Group Members
[**GetAccountGroupbySafe**](AccountGroupsApi.md#GetAccountGroupbySafe) | **Get** /PasswordVault/API/AccountGroups | Get Account Group by Safe



## AddAccountGroup

> AddAccountGroup(ctx).Authorization(authorization).ContentType(contentType).Body(body).Execute()

Add Account Group



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
  "GroupName":"{{GroupName}}",
  "GroupPlatform":"{{PlatformID}}",
  "Safe":"{{Safe}}"
}" // string | This method enables application managers to define a new account group automatically, and manage accounts as part of a group.  To create an account group, users require the following permissions in the Safe where the group is created: * Add accounts * Update account content * Update account properties * Create folders  **Note:** _The following characters are not supported in URL values in the Body:_ **+ & % #**

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountGroupsApi.AddAccountGroup(context.Background()).Authorization(authorization).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountGroupsApi.AddAccountGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddAccountGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **contentType** | **string** |  | 
 **body** | **string** | This method enables application managers to define a new account group automatically, and manage accounts as part of a group.  To create an account group, users require the following permissions in the Safe where the group is created: * Add accounts * Update account content * Update account properties * Create folders  **Note:** _The following characters are not supported in URL values in the Body:_ **+ &amp; % #** | 

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


## AddAccounttoAccountGroup

> AddAccounttoAccountGroup(ctx, groupName).Authorization(authorization).ContentType(contentType).Body(body).Execute()

Add Account to Account Group



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
    groupName := "GroupName" // string | 
    body := "{
  "AccountID":"{{AccountID}}"
}" // string | This method adds an account as a member to an existing account group.  The account can contain either a password or SSH key.  All members of an account group must be stored in the same Safe as the group itself.  To add an account as a member to an account group, users require the following permissions to the Safe where the group is created: * Add accounts * Update account content * Update account properties  **Note:** _The following characters are not support in URL values in the Body:_ **+ & % #**

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountGroupsApi.AddAccounttoAccountGroup(context.Background(), groupName).Authorization(authorization).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountGroupsApi.AddAccounttoAccountGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddAccounttoAccountGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **contentType** | **string** |  | 

 **body** | **string** | This method adds an account as a member to an existing account group.  The account can contain either a password or SSH key.  All members of an account group must be stored in the same Safe as the group itself.  To add an account as a member to an account group, users require the following permissions to the Safe where the group is created: * Add accounts * Update account content * Update account properties  **Note:** _The following characters are not support in URL values in the Body:_ **+ &amp; % #** | 

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


## DeleteMemberfromAccountGroup

> DeleteMemberfromAccountGroup(ctx, groupName, accountID).Authorization(authorization).ContentType(contentType).Execute()

Delete Member from Account Group



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
    groupName := "GroupName" // string | 
    accountID := "AccountID" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountGroupsApi.DeleteMemberfromAccountGroup(context.Background(), groupName, accountID).Authorization(authorization).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountGroupsApi.DeleteMemberfromAccountGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupName** | **string** |  | 
**accountID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMemberfromAccountGroupRequest struct via the builder pattern


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


## GetAccountGroupMembers

> GetAccountGroupMembers(ctx, groupName).Authorization(authorization).ContentType(contentType).Execute()

Get Account Group Members



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
    groupName := "GroupName" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountGroupsApi.GetAccountGroupMembers(context.Background(), groupName).Authorization(authorization).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountGroupsApi.GetAccountGroupMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountGroupMembersRequest struct via the builder pattern


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


## GetAccountGroupbySafe

> GetAccountGroupbySafe(ctx).Safe(safe).Authorization(authorization).ContentType(contentType).Execute()

Get Account Group by Safe



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
    safe := "{{Safe}}" // string | 
    authorization := "{{CyberArkLogonResult}}" // string | 
    contentType := "application/json" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountGroupsApi.GetAccountGroupbySafe(context.Background()).Safe(safe).Authorization(authorization).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountGroupsApi.GetAccountGroupbySafe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountGroupbySafeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **safe** | **string** |  | 
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

