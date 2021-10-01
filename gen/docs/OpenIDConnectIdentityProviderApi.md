# \OpenIDConnectIdentityProviderApi

All URIs are relative to *http://example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAuthenticationMethod**](OpenIDConnectIdentityProviderApi.md#AddAuthenticationMethod) | **Post** /PasswordVault/api/Configuration/AuthenticationMethods | Add Authentication Method
[**AddOpenIDConnectIdentityProvider**](OpenIDConnectIdentityProviderApi.md#AddOpenIDConnectIdentityProvider) | **Post** /PasswordVault/api/Configuration/OIDC/Providers | Add OpenID Connect Identity Provider
[**DeleteAuthenticationMethod**](OpenIDConnectIdentityProviderApi.md#DeleteAuthenticationMethod) | **Delete** /PasswordVault/api/Configuration/AuthenticationMethods/{authID} | Delete Authentication Method
[**DeleteOpenIDConnectIdentityProvider**](OpenIDConnectIdentityProviderApi.md#DeleteOpenIDConnectIdentityProvider) | **Delete** /PasswordVault/api/Configuration/OIDC/Providers/{authID} | Delete OpenID Connect Identity Provider
[**GetAllOpenIDConnectIdentityProviders**](OpenIDConnectIdentityProviderApi.md#GetAllOpenIDConnectIdentityProviders) | **Get** /PasswordVault/api/Configuration/OIDC/Providers | Get All OpenID Connect Identity Providers
[**GetAuthenticationMethods**](OpenIDConnectIdentityProviderApi.md#GetAuthenticationMethods) | **Get** /PasswordVault/api/Configuration/AuthenticationMethods | Get Authentication Methods
[**GetSpecificAuthenticationMethod**](OpenIDConnectIdentityProviderApi.md#GetSpecificAuthenticationMethod) | **Get** /PasswordVault/api/Configuration/AuthenticationMethods/{authID} | Get Specific Authentication Method
[**GetSpecificOpenIDConnectIdentityProvider**](OpenIDConnectIdentityProviderApi.md#GetSpecificOpenIDConnectIdentityProvider) | **Get** /PasswordVault/api/Configuration/OIDC/Providers/{authID} | Get Specific OpenID Connect Identity Provider
[**PTAAuthentication**](OpenIDConnectIdentityProviderApi.md#PTAAuthentication) | **Post** /api/getauthtoken | PTA Authentication
[**UpdateAuthenticationMethod**](OpenIDConnectIdentityProviderApi.md#UpdateAuthenticationMethod) | **Put** /PasswordVault/api/Configuration/AuthenticationMethods/{authID} | Update Authentication Method
[**UpdateOpenIDConnectIdentityProvider**](OpenIDConnectIdentityProviderApi.md#UpdateOpenIDConnectIdentityProvider) | **Put** /PasswordVault/api/Configuration/OIDC/Providers/{authID} | Update OpenID Connect Identity Provider



## AddAuthenticationMethod

> AddAuthenticationMethod(ctx).Authorization(authorization).ContentType(contentType).Body(body).Execute()

Add Authentication Method



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
    "id": "saml",
    "displayName": "IDaptive SAML",
    "enabled": true,
    "mobileEnabled": false,
    "logoffUrl": "https://domain.com/idp/logoff",
    "secondFactorAuth": null,
    "signInLabel": "",
    "usernameFieldLabel": "",
    "passwordFieldLabel": ""
}" // string | This method adds a new authentication method. Any user who is a member of the Vault admins group can run this web service.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OpenIDConnectIdentityProviderApi.AddAuthenticationMethod(context.Background()).Authorization(authorization).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenIDConnectIdentityProviderApi.AddAuthenticationMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddAuthenticationMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Session Authorization Token | 
 **contentType** | **string** |  | 
 **body** | **string** | This method adds a new authentication method. Any user who is a member of the Vault admins group can run this web service. | 

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


## AddOpenIDConnectIdentityProvider

> AddOpenIDConnectIdentityProvider(ctx).Authorization(authorization).ContentType(contentType).Body(body).Execute()

Add OpenID Connect Identity Provider



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
    "id": "opserver",
    "authenticationFlow": "Code",
    "discoveryEndpointUrl": "https://Domain.com/OPServer/.well-known/openid-configuration",
    "clientId": "pvwa",
    "clientSecret": "secret",
    "clientSecretMethod": "basic"
}" // string | This method creates an OpenID Connect (OIDC) Identity Provider in the Vault. Any user who is a member of the Vault admins group can run this web service.  **NOTE:** This API must be used with the Add authentication method API in order to have a complete authentication configuration.  When you add an OIDC Identity Provider, you must also add the provider to the list of authentication methods using the same ID, and add the provider's base URL to the access restriction list by using the following APIs:  * Add authentication method * Add allowed referrer

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OpenIDConnectIdentityProviderApi.AddOpenIDConnectIdentityProvider(context.Background()).Authorization(authorization).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenIDConnectIdentityProviderApi.AddOpenIDConnectIdentityProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddOpenIDConnectIdentityProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Session Authorization Token | 
 **contentType** | **string** |  | 
 **body** | **string** | This method creates an OpenID Connect (OIDC) Identity Provider in the Vault. Any user who is a member of the Vault admins group can run this web service.  **NOTE:** This API must be used with the Add authentication method API in order to have a complete authentication configuration.  When you add an OIDC Identity Provider, you must also add the provider to the list of authentication methods using the same ID, and add the provider&#39;s base URL to the access restriction list by using the following APIs:  * Add authentication method * Add allowed referrer | 

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


## DeleteAuthenticationMethod

> DeleteAuthenticationMethod(ctx, authID).Authorization(authorization).ContentType(contentType).Execute()

Delete Authentication Method



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
    authID := "authID" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OpenIDConnectIdentityProviderApi.DeleteAuthenticationMethod(context.Background(), authID).Authorization(authorization).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenIDConnectIdentityProviderApi.DeleteAuthenticationMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthenticationMethodRequest struct via the builder pattern


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


## DeleteOpenIDConnectIdentityProvider

> DeleteOpenIDConnectIdentityProvider(ctx, authID).Authorization(authorization).ContentType(contentType).Execute()

Delete OpenID Connect Identity Provider



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
    authID := "authID" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OpenIDConnectIdentityProviderApi.DeleteOpenIDConnectIdentityProvider(context.Background(), authID).Authorization(authorization).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenIDConnectIdentityProviderApi.DeleteOpenIDConnectIdentityProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOpenIDConnectIdentityProviderRequest struct via the builder pattern


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


## GetAllOpenIDConnectIdentityProviders

> GetAllOpenIDConnectIdentityProviders(ctx).Authorization(authorization).ContentType(contentType).Execute()

Get All OpenID Connect Identity Providers



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
    resp, r, err := api_client.OpenIDConnectIdentityProviderApi.GetAllOpenIDConnectIdentityProviders(context.Background()).Authorization(authorization).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenIDConnectIdentityProviderApi.GetAllOpenIDConnectIdentityProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllOpenIDConnectIdentityProvidersRequest struct via the builder pattern


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


## GetAuthenticationMethods

> Model200 GetAuthenticationMethods(ctx).Authorization(authorization).ContentType(contentType).Execute()

Get Authentication Methods



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
    resp, r, err := api_client.OpenIDConnectIdentityProviderApi.GetAuthenticationMethods(context.Background()).Authorization(authorization).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenIDConnectIdentityProviderApi.GetAuthenticationMethods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthenticationMethods`: Model200
    fmt.Fprintf(os.Stdout, "Response from `OpenIDConnectIdentityProviderApi.GetAuthenticationMethods`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthenticationMethodsRequest struct via the builder pattern


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


## GetSpecificAuthenticationMethod

> Model200 GetSpecificAuthenticationMethod(ctx, authID).Authorization(authorization).ContentType(contentType).Execute()

Get Specific Authentication Method



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
    authID := "authID" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OpenIDConnectIdentityProviderApi.GetSpecificAuthenticationMethod(context.Background(), authID).Authorization(authorization).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenIDConnectIdentityProviderApi.GetSpecificAuthenticationMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSpecificAuthenticationMethod`: Model200
    fmt.Fprintf(os.Stdout, "Response from `OpenIDConnectIdentityProviderApi.GetSpecificAuthenticationMethod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpecificAuthenticationMethodRequest struct via the builder pattern


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


## GetSpecificOpenIDConnectIdentityProvider

> GetSpecificOpenIDConnectIdentityProvider(ctx, authID).Authorization(authorization).ContentType(contentType).Execute()

Get Specific OpenID Connect Identity Provider



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
    authID := "authID" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OpenIDConnectIdentityProviderApi.GetSpecificOpenIDConnectIdentityProvider(context.Background(), authID).Authorization(authorization).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenIDConnectIdentityProviderApi.GetSpecificOpenIDConnectIdentityProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpecificOpenIDConnectIdentityProviderRequest struct via the builder pattern


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


## PTAAuthentication

> PTAAuthentication(ctx).ContentType(contentType).Username(username).Password(password).Execute()

PTA Authentication



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
    contentType := "application/x-www-form-urlencoded" // string | 
    username := "username_example" // string | 
    password := "password_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OpenIDConnectIdentityProviderApi.PTAAuthentication(context.Background()).ContentType(contentType).Username(username).Password(password).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenIDConnectIdentityProviderApi.PTAAuthentication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPTAAuthenticationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string** |  | 
 **username** | **string** |  | 
 **password** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAuthenticationMethod

> UpdateAuthenticationMethod(ctx, authID).Authorization(authorization).ContentType(contentType).Body(body).Execute()

Update Authentication Method



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
    authID := "authID" // string | 
    body := "{
    "displayName": "",
    "enabled": true,
    "mobileEnabled": false,
    "logoffUrl": "https://domain.com/idp/logoff",
    "secondFactorAuth": null,
    "signInLabel": "",
    "usernameFieldLabel": "",
    "passwordFieldLabel": ""
}" // string | This method updates the properties for a specific authentication method. Any user who is a member of the Vault admins group can run this web service.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OpenIDConnectIdentityProviderApi.UpdateAuthenticationMethod(context.Background(), authID).Authorization(authorization).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenIDConnectIdentityProviderApi.UpdateAuthenticationMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAuthenticationMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Session Authorization Token | 
 **contentType** | **string** |  | 

 **body** | **string** | This method updates the properties for a specific authentication method. Any user who is a member of the Vault admins group can run this web service. | 

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


## UpdateOpenIDConnectIdentityProvider

> UpdateOpenIDConnectIdentityProvider(ctx, authID).Authorization(authorization).ContentType(contentType).Body(body).Execute()

Update OpenID Connect Identity Provider



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
    authID := "authID" // string | 
    body := "{
    "authenticationFlow": "Code",
    "discoveryEndpointUrl": "https://10.10.22.121/OPServer/.well-known/openid-configuration",
    "clientId": "pvwa",
    "clientSecretMethod": "basic",
    "userNameClaim": "given_name"
}" // string | This method updates an existing OIDC Identity Provider. Any user who is a member of the Vault admins group can run this web service.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OpenIDConnectIdentityProviderApi.UpdateOpenIDConnectIdentityProvider(context.Background(), authID).Authorization(authorization).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenIDConnectIdentityProviderApi.UpdateOpenIDConnectIdentityProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOpenIDConnectIdentityProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Session Authorization Token | 
 **contentType** | **string** |  | 

 **body** | **string** | This method updates an existing OIDC Identity Provider. Any user who is a member of the Vault admins group can run this web service. | 

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

