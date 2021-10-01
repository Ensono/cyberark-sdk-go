# \SAMLAuthenticationApi

All URIs are relative to *http://example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Logoff**](SAMLAuthenticationApi.md#Logoff) | **Post** /PasswordVault/api/auth/SAML/Logoff | Logoff
[**Logoff5**](SAMLAuthenticationApi.md#Logoff5) | **Post** /PasswordVault/API/Auth/Logoff | Logoff
[**Logon**](SAMLAuthenticationApi.md#Logon) | **Post** /PasswordVault/API/auth/SAML/Logon | Logon
[**LogonCyberArkAuthentication**](SAMLAuthenticationApi.md#LogonCyberArkAuthentication) | **Post** /PasswordVault/API/Auth/CyberArk/Logon | Logon - CyberArk Authentication
[**LogonLDAPAuthentication**](SAMLAuthenticationApi.md#LogonLDAPAuthentication) | **Post** /PasswordVault/API/Auth/LDAP/Logon | Logon - LDAP Authentication
[**LogonRADIUSAuthentication**](SAMLAuthenticationApi.md#LogonRADIUSAuthentication) | **Post** /PasswordVault/API/Auth/radius/Logon | Logon - RADIUS Authentication
[**LogonWindowsAuthentication**](SAMLAuthenticationApi.md#LogonWindowsAuthentication) | **Post** /PasswordVault/API/auth/Windows/Logon | Logon - Windows Authentication



## Logoff

> Logoff(ctx).ContentType(contentType).Authorization(authorization).Execute()

Logoff



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
    contentType := "application/json" // string | 
    authorization := "{{SAMLToken}}" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SAMLAuthenticationApi.Logoff(context.Background()).ContentType(contentType).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SAMLAuthenticationApi.Logoff``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLogoffRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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


## Logoff5

> Model200 Logoff5(ctx).Authorization(authorization).ContentType(contentType).Execute()

Logoff



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SAMLAuthenticationApi.Logoff5(context.Background()).Authorization(authorization).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SAMLAuthenticationApi.Logoff5``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Logoff5`: Model200
    fmt.Fprintf(os.Stdout, "Response from `SAMLAuthenticationApi.Logoff5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLogoff5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **contentType** | **string** |  | 

### Return type

[**Model200**](Model200.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Logon

> Logon(ctx).ConcurrentSession(concurrentSession).ApiUse(apiUse).SAMLResponse(sAMLResponse).ContentType(contentType).Execute()

Logon



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
    concurrentSession := "true" // string | Boolean value
    apiUse := "true" // string | Never should be false
    sAMLResponse := "{{SAMLToken}}" // string | The SAML Token returned from your IdP
    contentType := "application/x-www-form-urlencoded" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SAMLAuthenticationApi.Logon(context.Background()).ConcurrentSession(concurrentSession).ApiUse(apiUse).SAMLResponse(sAMLResponse).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SAMLAuthenticationApi.Logon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLogonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **concurrentSession** | **string** | Boolean value | 
 **apiUse** | **string** | Never should be false | 
 **sAMLResponse** | **string** | The SAML Token returned from your IdP | 
 **contentType** | **string** |  | 

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


## LogonCyberArkAuthentication

> interface{} LogonCyberArkAuthentication(ctx).ContentType(contentType).Body(body).Execute()

Logon - CyberArk Authentication



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
    contentType := "application/json" // string | 
    body := "{
	"username": "{{apiUsername}}",
	"password": "{{apiPassword}}",
	//"newPassword": "<optional>",
	"concurrentSession": "true"
}" // string | This method authenticates a user to the Vault and returns a token that can be used in subsequent web services calls. In addition, this method enables you to set a new password.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SAMLAuthenticationApi.LogonCyberArkAuthentication(context.Background()).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SAMLAuthenticationApi.LogonCyberArkAuthentication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LogonCyberArkAuthentication`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `SAMLAuthenticationApi.LogonCyberArkAuthentication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLogonCyberArkAuthenticationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string** |  | 
 **body** | **string** | This method authenticates a user to the Vault and returns a token that can be used in subsequent web services calls. In addition, this method enables you to set a new password. | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LogonLDAPAuthentication

> interface{} LogonLDAPAuthentication(ctx).ContentType(contentType).Body(body).Execute()

Logon - LDAP Authentication



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
    contentType := "application/json" // string | 
    body := "{
	"username": "{{apiUsername}}",
	"password": "{{apiPassword}}",
	//"newPassword": "<optional>",
	"concurrentSession": "true"
}" // string | This method authenticates a user to the Vault and returns a token that can be used in subsequent web services calls. In addition, this method enables you to set a new password.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SAMLAuthenticationApi.LogonLDAPAuthentication(context.Background()).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SAMLAuthenticationApi.LogonLDAPAuthentication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LogonLDAPAuthentication`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `SAMLAuthenticationApi.LogonLDAPAuthentication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLogonLDAPAuthenticationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string** |  | 
 **body** | **string** | This method authenticates a user to the Vault and returns a token that can be used in subsequent web services calls. In addition, this method enables you to set a new password. | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LogonRADIUSAuthentication

> interface{} LogonRADIUSAuthentication(ctx).ContentType(contentType).Body(body).Execute()

Logon - RADIUS Authentication



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
    contentType := "application/json" // string | 
    body := "{
	"Username": "{{apiUsername}}",
	"Password": "{{apiPassword}}",
	"concurrentSessions": "true"
}" // string | This method authenticates a user to the Vault and returns a token that can be used in subsequent web services calls. In addition, this method enables you to set a new password.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SAMLAuthenticationApi.LogonRADIUSAuthentication(context.Background()).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SAMLAuthenticationApi.LogonRADIUSAuthentication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LogonRADIUSAuthentication`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `SAMLAuthenticationApi.LogonRADIUSAuthentication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLogonRADIUSAuthenticationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string** |  | 
 **body** | **string** | This method authenticates a user to the Vault and returns a token that can be used in subsequent web services calls. In addition, this method enables you to set a new password. | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LogonWindowsAuthentication

> LogonWindowsAuthentication(ctx).ContentType(contentType).Body(body).Execute()

Logon - Windows Authentication



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
    contentType := "application/json" // string | 
    body := "{
	"username": "{{apiUsername}}",
	"password": "{{apiPassword}}",
	"concurrentSessions": "false"
}" // string | This method authenticates a user to the Vault and returns a token that can be used in subsequent web services calls. In addition, this method enables you to set a new password.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SAMLAuthenticationApi.LogonWindowsAuthentication(context.Background()).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SAMLAuthenticationApi.LogonWindowsAuthentication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLogonWindowsAuthenticationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string** |  | 
 **body** | **string** | This method authenticates a user to the Vault and returns a token that can be used in subsequent web services calls. In addition, this method enables you to set a new password. | 

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

