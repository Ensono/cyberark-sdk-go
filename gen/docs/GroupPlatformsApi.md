# \GroupPlatformsApi

All URIs are relative to *http://example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateGroupPlatform**](GroupPlatformsApi.md#ActivateGroupPlatform) | **Post** /passwordvault/api/platforms/groups/{PlatformName}/activate | Activate Group Platform
[**DeactivateGroupPlatform**](GroupPlatformsApi.md#DeactivateGroupPlatform) | **Post** /passwordvault/api/platforms/groups/{PlatformName}/deactivate | Deactivate Group Platform
[**DeleteGroupPlatform**](GroupPlatformsApi.md#DeleteGroupPlatform) | **Delete** /passwordvault/api/platforms/groups/{PlatformName} | Delete Group Platform
[**DuplicateGroupPlatforms**](GroupPlatformsApi.md#DuplicateGroupPlatforms) | **Post** /passwordvault/api/platforms/groups/{PlatformName}/duplicate | Duplicate Group Platforms
[**GetGroupPlatforms**](GroupPlatformsApi.md#GetGroupPlatforms) | **Get** /passwordvault/api/platforms/groups | Get Group Platforms



## ActivateGroupPlatform

> ActivateGroupPlatform(ctx, platformName).Authorization(authorization).ContentType(contentType).Execute()

Activate Group Platform



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
    platformName := "PlatformName" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupPlatformsApi.ActivateGroupPlatform(context.Background(), platformName).Authorization(authorization).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupPlatformsApi.ActivateGroupPlatform``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platformName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateGroupPlatformRequest struct via the builder pattern


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


## DeactivateGroupPlatform

> DeactivateGroupPlatform(ctx, platformName).Authorization(authorization).ContentType(contentType).Execute()

Deactivate Group Platform



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
    platformName := "PlatformName" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupPlatformsApi.DeactivateGroupPlatform(context.Background(), platformName).Authorization(authorization).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupPlatformsApi.DeactivateGroupPlatform``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platformName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateGroupPlatformRequest struct via the builder pattern


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


## DeleteGroupPlatform

> DeleteGroupPlatform(ctx, platformName).Authorization(authorization).ContentType(contentType).Execute()

Delete Group Platform



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
    platformName := "PlatformName" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupPlatformsApi.DeleteGroupPlatform(context.Background(), platformName).Authorization(authorization).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupPlatformsApi.DeleteGroupPlatform``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platformName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupPlatformRequest struct via the builder pattern


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


## DuplicateGroupPlatforms

> DuplicateGroupPlatforms(ctx, platformName).Authorization(authorization).ContentType(contentType).Body(body).Execute()

Duplicate Group Platforms



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
    platformName := "PlatformName" // string | 
    body := "{
    "Name": "test Platform",
    "Description": ""
}" // string | This method allows Vault Admins to duplicate group platforms.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupPlatformsApi.DuplicateGroupPlatforms(context.Background(), platformName).Authorization(authorization).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupPlatformsApi.DuplicateGroupPlatforms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platformName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDuplicateGroupPlatformsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Session Authorization Token | 
 **contentType** | **string** |  | 

 **body** | **string** | This method allows Vault Admins to duplicate group platforms. | 

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


## GetGroupPlatforms

> Model200 GetGroupPlatforms(ctx).Search(search).Authorization(authorization).ContentType(contentType).Execute()

Get Group Platforms



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
    search := "SSH" // string | Platform Name
    authorization := "{{CyberArkLogonResult}}" // string | Session Authorization Token
    contentType := "application/json" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupPlatformsApi.GetGroupPlatforms(context.Background()).Search(search).Authorization(authorization).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupPlatformsApi.GetGroupPlatforms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroupPlatforms`: Model200
    fmt.Fprintf(os.Stdout, "Response from `GroupPlatformsApi.GetGroupPlatforms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupPlatformsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | Platform Name | 
 **authorization** | **string** | Session Authorization Token | 
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

