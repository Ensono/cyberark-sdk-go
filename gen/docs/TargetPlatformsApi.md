# \TargetPlatformsApi

All URIs are relative to *http://example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateTargetPlatform**](TargetPlatformsApi.md#ActivateTargetPlatform) | **Post** /passwordvault/api/platforms/targets/{PlatformName}/activate | Activate Target Platform
[**DeactivateTargetPlatform**](TargetPlatformsApi.md#DeactivateTargetPlatform) | **Post** /passwordvault/api/platforms/targets/{PlatformName}/deactivate | Deactivate Target Platform
[**DeleteTargetPlatform**](TargetPlatformsApi.md#DeleteTargetPlatform) | **Delete** /passwordvault/api/platforms/targets/{PlatformName} | Delete Target Platform
[**DuplicateTargetPlatforms**](TargetPlatformsApi.md#DuplicateTargetPlatforms) | **Post** /passwordvault/api/platforms/targets/{PlatformName}/duplicate | Duplicate Target Platforms
[**GetTargetPlatforms**](TargetPlatformsApi.md#GetTargetPlatforms) | **Get** /passwordvault/api/platforms/targets | Get Target Platforms



## ActivateTargetPlatform

> ActivateTargetPlatform(ctx, platformName).Authorization(authorization).ContentType(contentType).Execute()

Activate Target Platform



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
    resp, r, err := api_client.TargetPlatformsApi.ActivateTargetPlatform(context.Background(), platformName).Authorization(authorization).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TargetPlatformsApi.ActivateTargetPlatform``: %v\n", err)
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

Other parameters are passed through a pointer to a apiActivateTargetPlatformRequest struct via the builder pattern


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


## DeactivateTargetPlatform

> DeactivateTargetPlatform(ctx, platformName).Authorization(authorization).ContentType(contentType).Execute()

Deactivate Target Platform



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
    resp, r, err := api_client.TargetPlatformsApi.DeactivateTargetPlatform(context.Background(), platformName).Authorization(authorization).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TargetPlatformsApi.DeactivateTargetPlatform``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeactivateTargetPlatformRequest struct via the builder pattern


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


## DeleteTargetPlatform

> DeleteTargetPlatform(ctx, platformName).Authorization(authorization).ContentType(contentType).Execute()

Delete Target Platform



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
    resp, r, err := api_client.TargetPlatformsApi.DeleteTargetPlatform(context.Background(), platformName).Authorization(authorization).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TargetPlatformsApi.DeleteTargetPlatform``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteTargetPlatformRequest struct via the builder pattern


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


## DuplicateTargetPlatforms

> DuplicateTargetPlatforms(ctx, platformName).Authorization(authorization).ContentType(contentType).Body(body).Execute()

Duplicate Target Platforms



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
}" // string | This method allows Vault Admins to duplicate target platforms.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TargetPlatformsApi.DuplicateTargetPlatforms(context.Background(), platformName).Authorization(authorization).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TargetPlatformsApi.DuplicateTargetPlatforms``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDuplicateTargetPlatformsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Session Authorization Token | 
 **contentType** | **string** |  | 

 **body** | **string** | This method allows Vault Admins to duplicate target platforms. | 

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


## GetTargetPlatforms

> Model200 GetTargetPlatforms(ctx).Authorization(authorization).ContentType(contentType).Execute()

Get Target Platforms



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TargetPlatformsApi.GetTargetPlatforms(context.Background()).Authorization(authorization).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TargetPlatformsApi.GetTargetPlatforms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTargetPlatforms`: Model200
    fmt.Fprintf(os.Stdout, "Response from `TargetPlatformsApi.GetTargetPlatforms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTargetPlatformsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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

