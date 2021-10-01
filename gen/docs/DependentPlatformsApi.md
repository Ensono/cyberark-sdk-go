# \DependentPlatformsApi

All URIs are relative to *http://example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDependentPlatform**](DependentPlatformsApi.md#DeleteDependentPlatform) | **Delete** /passwordvault/api/platforms/dependents/{PlatformName} | Delete Dependent Platform
[**DuplicateDependentPlatforms**](DependentPlatformsApi.md#DuplicateDependentPlatforms) | **Post** /passwordvault/api/platforms/dependent/{PlatformName}/duplicate | Duplicate Dependent Platforms
[**GetDependentPlatforms**](DependentPlatformsApi.md#GetDependentPlatforms) | **Get** /passwordvault/api/platforms/dependents | Get Dependent Platforms



## DeleteDependentPlatform

> DeleteDependentPlatform(ctx, platformName).Authorization(authorization).ContentType(contentType).Execute()

Delete Dependent Platform



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
    resp, r, err := api_client.DependentPlatformsApi.DeleteDependentPlatform(context.Background(), platformName).Authorization(authorization).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DependentPlatformsApi.DeleteDependentPlatform``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteDependentPlatformRequest struct via the builder pattern


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


## DuplicateDependentPlatforms

> DuplicateDependentPlatforms(ctx, platformName).Authorization(authorization).ContentType(contentType).Body(body).Execute()

Duplicate Dependent Platforms



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
}" // string | This method allows Vault Admins to duplicate dependent platforms.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DependentPlatformsApi.DuplicateDependentPlatforms(context.Background(), platformName).Authorization(authorization).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DependentPlatformsApi.DuplicateDependentPlatforms``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDuplicateDependentPlatformsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Session Authorization Token | 
 **contentType** | **string** |  | 

 **body** | **string** | This method allows Vault Admins to duplicate dependent platforms. | 

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


## GetDependentPlatforms

> GetDependentPlatforms(ctx).Search(search).Authorization(authorization).ContentType(contentType).Execute()

Get Dependent Platforms



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
    search := float32(8.14) // float32 | Platform Name
    authorization := "{{CyberArkLogonResult}}" // string | Session Authorization Token
    contentType := "application/json" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DependentPlatformsApi.GetDependentPlatforms(context.Background()).Search(search).Authorization(authorization).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DependentPlatformsApi.GetDependentPlatforms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDependentPlatformsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **float32** | Platform Name | 
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

