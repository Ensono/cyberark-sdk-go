# \BulkUploadOfAccountsApi

All URIs are relative to *http://example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBulkUploadofAccounts**](BulkUploadOfAccountsApi.md#CreateBulkUploadofAccounts) | **Post** /passwordvault/api/bulkactions/accounts | Create Bulk Upload of Accounts
[**GetAllBulkAccountUploadsforUser**](BulkUploadOfAccountsApi.md#GetAllBulkAccountUploadsforUser) | **Get** /passwordvault/api/bulkactions/accounts | Get All Bulk Account Uploads for User
[**GetBulkAccountUploadResult**](BulkUploadOfAccountsApi.md#GetBulkAccountUploadResult) | **Get** /passwordvault/api/bulkactions/accounts/{BulkID} | Get Bulk Account Upload Result



## CreateBulkUploadofAccounts

> CreateBulkUploadofAccounts(ctx).Authorization(authorization).ContentType(contentType).Body(body).Execute()

Create Bulk Upload of Accounts



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
    body := "{
    "source": "filename.csv",
    "accountsList": [
        {
            "uploadIndex": "1",
            "username": "JohnDoe",
            "address": "192.0.2.0",
            "platformId": "WinDomain",
            "safeName": "WinDomainSafe",
            "secret": "123456",
            "platformAccountProperties": {
                "port": "111"
            },
            "secretManagement": {
                "automaticManagementEnabled": true,
                "manualManagementReason": ""
            },
            "remoteMachinesAccess": {
                "accessRestrictedToRemoteMachines": true,
                "remoteMachines": "example.com"
            },
            "groupName": "DomainGroup"
        },
        {
            "uploadIndex": "2",
            "username": "JaneDoe",
            "address": "198.51.100.0",
            "platformId": "WinDesktopLocal",
            "safeName": "WinUsersSafe",
            "secret": "123456",
            "platformAccountProperties": {
                "port": "222"
            },
            "secretManagement": {
                "automaticManagementEnabled": true,
                "manualManagementReason": ""
            },
            "remoteMachinesAccess": {
                "accessRestrictedToRemoteMachines": true,
                "remoteMachines": "example.net"
            },
            "groupName": "WinGroup"
        }
    ]
}" // string | This method allows a developer to add multiple accounts to existing Safes. The response contains the ID of the bulk account upload that was performed.     **Note:** This option is only available if you have **Add accounts**, **Update account content**, and **Update account properties** authorization in at least one Safe.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BulkUploadOfAccountsApi.CreateBulkUploadofAccounts(context.Background()).Authorization(authorization).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BulkUploadOfAccountsApi.CreateBulkUploadofAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBulkUploadofAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Session Authorization Token | 
 **contentType** | **string** |  | 
 **body** | **string** | This method allows a developer to add multiple accounts to existing Safes. The response contains the ID of the bulk account upload that was performed.     **Note:** This option is only available if you have **Add accounts**, **Update account content**, and **Update account properties** authorization in at least one Safe. | 

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


## GetAllBulkAccountUploadsforUser

> GetAllBulkAccountUploadsforUser(ctx).Filter(filter).Limit(limit).Authorization(authorization).ContentType(contentType).Execute()

Get All Bulk Account Uploads for User



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
    filter := float32(8.14) // float32 | status - returns all bulk account uploads that meet the required status
    limit := float32(8.14) // float32 | number of accounts to return, starting from first account
    authorization := "{{CyberArkLogonResult}}" // string | Session Authorization Token
    contentType := "application/json" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BulkUploadOfAccountsApi.GetAllBulkAccountUploadsforUser(context.Background()).Filter(filter).Limit(limit).Authorization(authorization).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BulkUploadOfAccountsApi.GetAllBulkAccountUploadsforUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllBulkAccountUploadsforUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **float32** | status - returns all bulk account uploads that meet the required status | 
 **limit** | **float32** | number of accounts to return, starting from first account | 
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


## GetBulkAccountUploadResult

> GetBulkAccountUploadResult(ctx, bulkID).Authorization(authorization).ContentType(contentType).Execute()

Get Bulk Account Upload Result



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
    bulkID := "BulkID" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BulkUploadOfAccountsApi.GetBulkAccountUploadResult(context.Background(), bulkID).Authorization(authorization).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BulkUploadOfAccountsApi.GetBulkAccountUploadResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bulkID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBulkAccountUploadResultRequest struct via the builder pattern


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

