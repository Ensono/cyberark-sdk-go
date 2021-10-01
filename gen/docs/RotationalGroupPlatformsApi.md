# \RotationalGroupPlatformsApi

All URIs are relative to *http://example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateRotationalGroupPlatform**](RotationalGroupPlatformsApi.md#ActivateRotationalGroupPlatform) | **Post** /passwordvault/api/platforms/rotationalGroups/{PlatformName}/activate | Activate Rotational Group Platform
[**DeactivateRotationalGroupPlatform**](RotationalGroupPlatformsApi.md#DeactivateRotationalGroupPlatform) | **Post** /passwordvault/api/platforms/rotationalGroups/{PlatformName}/deactivate | Deactivate Rotational Group Platform
[**DeleteRotationalGroupPlatform**](RotationalGroupPlatformsApi.md#DeleteRotationalGroupPlatform) | **Delete** /passwordvault/api/platforms/rotationalGroups/{PlatformName} | Delete Rotational Group Platform
[**DuplicateRotationalGroupPlatforms**](RotationalGroupPlatformsApi.md#DuplicateRotationalGroupPlatforms) | **Post** /passwordvault/api/platforms/rotationalGroups/{PlatformName}/duplicate | Duplicate Rotational Group Platforms
[**ExportPlatform**](RotationalGroupPlatformsApi.md#ExportPlatform) | **Post** /PasswordVault/API/Platforms/{PlatformName}/Export | Export Platform
[**GetPlatformDetails**](RotationalGroupPlatformsApi.md#GetPlatformDetails) | **Get** /PasswordVault/API/Platforms/{PlatformName} | Get Platform Details
[**GetPlatforms**](RotationalGroupPlatformsApi.md#GetPlatforms) | **Get** /PasswordVault/API/Platforms | Get Platforms
[**GetRotationalGroupPlatforms**](RotationalGroupPlatformsApi.md#GetRotationalGroupPlatforms) | **Get** /passwordvault/api/platforms/rotationalGroups | Get Rotational Group Platforms
[**ImportPlatform**](RotationalGroupPlatformsApi.md#ImportPlatform) | **Post** /API/Platforms/Import | Import Platform



## ActivateRotationalGroupPlatform

> ActivateRotationalGroupPlatform(ctx, platformName).Authorization(authorization).ContentType(contentType).Execute()

Activate Rotational Group Platform



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
    resp, r, err := api_client.RotationalGroupPlatformsApi.ActivateRotationalGroupPlatform(context.Background(), platformName).Authorization(authorization).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RotationalGroupPlatformsApi.ActivateRotationalGroupPlatform``: %v\n", err)
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

Other parameters are passed through a pointer to a apiActivateRotationalGroupPlatformRequest struct via the builder pattern


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


## DeactivateRotationalGroupPlatform

> DeactivateRotationalGroupPlatform(ctx, platformName).Authorization(authorization).ContentType(contentType).Execute()

Deactivate Rotational Group Platform



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
    resp, r, err := api_client.RotationalGroupPlatformsApi.DeactivateRotationalGroupPlatform(context.Background(), platformName).Authorization(authorization).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RotationalGroupPlatformsApi.DeactivateRotationalGroupPlatform``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeactivateRotationalGroupPlatformRequest struct via the builder pattern


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


## DeleteRotationalGroupPlatform

> DeleteRotationalGroupPlatform(ctx, platformName).Authorization(authorization).ContentType(contentType).Execute()

Delete Rotational Group Platform



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
    resp, r, err := api_client.RotationalGroupPlatformsApi.DeleteRotationalGroupPlatform(context.Background(), platformName).Authorization(authorization).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RotationalGroupPlatformsApi.DeleteRotationalGroupPlatform``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteRotationalGroupPlatformRequest struct via the builder pattern


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


## DuplicateRotationalGroupPlatforms

> DuplicateRotationalGroupPlatforms(ctx, platformName).Authorization(authorization).ContentType(contentType).Body(body).Execute()

Duplicate Rotational Group Platforms



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
}" // string | This method allows Vault Admins to duplicate rotational group platforms.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RotationalGroupPlatformsApi.DuplicateRotationalGroupPlatforms(context.Background(), platformName).Authorization(authorization).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RotationalGroupPlatformsApi.DuplicateRotationalGroupPlatforms``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDuplicateRotationalGroupPlatformsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Session Authorization Token | 
 **contentType** | **string** |  | 

 **body** | **string** | This method allows Vault Admins to duplicate rotational group platforms. | 

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


## ExportPlatform

> interface{} ExportPlatform(ctx, platformName).Authorization(authorization).ContentType(contentType).Execute()

Export Platform



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
    resp, r, err := api_client.RotationalGroupPlatformsApi.ExportPlatform(context.Background(), platformName).Authorization(authorization).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RotationalGroupPlatformsApi.ExportPlatform``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportPlatform`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `RotationalGroupPlatformsApi.ExportPlatform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platformName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportPlatformRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **contentType** | **string** |  | 


### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlatformDetails

> Model200 GetPlatformDetails(ctx, platformName).Authorization(authorization).ContentType(contentType).Execute()

Get Platform Details



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
    resp, r, err := api_client.RotationalGroupPlatformsApi.GetPlatformDetails(context.Background(), platformName).Authorization(authorization).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RotationalGroupPlatformsApi.GetPlatformDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPlatformDetails`: Model200
    fmt.Fprintf(os.Stdout, "Response from `RotationalGroupPlatformsApi.GetPlatformDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platformName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlatformDetailsRequest struct via the builder pattern


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


## GetPlatforms

> GetPlatforms(ctx).Active(active).PlatformType(platformType).PlatformName(platformName).Authorization(authorization).ContentType(contentType).Execute()

Get Platforms



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
    active := "true" // string | Filter according to whether the platform is active or not. Valid values: true or false
    platformType := "Regular" // string | Filter according to the platform type. Valid values: Group or Regular
    platformName := "platformName_example" // string | Searching according to the platform name.
    authorization := "{{CyberArkLogonResult}}" // string | 
    contentType := "application/json" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RotationalGroupPlatformsApi.GetPlatforms(context.Background()).Active(active).PlatformType(platformType).PlatformName(platformName).Authorization(authorization).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RotationalGroupPlatformsApi.GetPlatforms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPlatformsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **active** | **string** | Filter according to whether the platform is active or not. Valid values: true or false | 
 **platformType** | **string** | Filter according to the platform type. Valid values: Group or Regular | 
 **platformName** | **string** | Searching according to the platform name. | 
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


## GetRotationalGroupPlatforms

> Model200 GetRotationalGroupPlatforms(ctx).Search(search).Authorization(authorization).ContentType(contentType).Execute()

Get Rotational Group Platforms



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
    search := "MySQL" // string | Platform Name
    authorization := "{{CyberArkLogonResult}}" // string | Session Authorization Token
    contentType := "application/json" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RotationalGroupPlatformsApi.GetRotationalGroupPlatforms(context.Background()).Search(search).Authorization(authorization).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RotationalGroupPlatformsApi.GetRotationalGroupPlatforms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRotationalGroupPlatforms`: Model200
    fmt.Fprintf(os.Stdout, "Response from `RotationalGroupPlatformsApi.GetRotationalGroupPlatforms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRotationalGroupPlatformsRequest struct via the builder pattern


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


## ImportPlatform

> ImportPlatform(ctx).Authorization(authorization).ContentType(contentType).Body(body).Execute()

Import Platform



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
	"ImportFile": {zip file in the format of BASE 64 array}
}" // string | This method enables administrators to import new platforms and dependencies.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RotationalGroupPlatformsApi.ImportPlatform(context.Background()).Authorization(authorization).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RotationalGroupPlatformsApi.ImportPlatform``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportPlatformRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **contentType** | **string** |  | 
 **body** | **string** | This method enables administrators to import new platforms and dependencies. | 

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

