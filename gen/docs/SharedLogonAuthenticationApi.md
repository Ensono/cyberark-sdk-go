# \SharedLogonAuthenticationApi

All URIs are relative to *http://example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Logoff1**](SharedLogonAuthenticationApi.md#Logoff1) | **Post** /PasswordVault/WebServices/auth/Shared/RestfulAuthenticationService.svc/Logoff | Logoff
[**Logon0**](SharedLogonAuthenticationApi.md#Logon0) | **Post** /PasswordVault/WebServices/auth/Shared/RestfulAuthenticationService.svc/Logon | Logon



## Logoff1

> Logoff1(ctx).ContentType(contentType).Authorization(authorization).Execute()

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
    authorization := "{{CyberArkLogonResult}}" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharedLogonAuthenticationApi.Logoff1(context.Background()).ContentType(contentType).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharedLogonAuthenticationApi.Logoff1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLogoff1Request struct via the builder pattern


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


## Logon0

> Logon0(ctx).ContentType(contentType).Body(body).Execute()

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
    contentType := "application/json" // string | 
    body := "body_example" // string | Shared authentication is based on a user credential file that is stored in the PVWA web server. During shared authentication, only the user defined in the credential file can log on to the PVWA, but multiple users can use the logon token.  This type of authentication **requires** the application using the REST services to manage the users as the Vault can't identify which specific user performs each action.  Multiple concurrent connections can be created using the same token, without affecting each other.  The shared user is defined in a user credential file, whose location is specified in the WSCredentialFile parameter, in the appsettings section of the PVWAweb.config file:  ``` <add key=\"WSCredentialFile\" value=\"C:\\CyberArk\\Password Vault Web Access\\CredFiles\\WSUser.ini\"/> ```  * Make sure that this user can access the PVWA interface. * Make sure the user only has the permissions in the Vault that they require.  For information about securing communication when using the SDK, refer to the following:  * [Securing application-to-REST communication](https://docs.cyberark.com/Product-Doc/OnlineHelp/PAS/Latest/en/Content/SDK/Securing%20Communication.htm) * [Configuring client authentication via certificates](https://docs.cyberark.com/Product-Doc/OnlineHelp/PAS/Latest/en/Content/SDK/Configuring%20Client%20Authentication%20via%20Client%20Certificates.htm)  This method authenticates to the Vault with a shared webservices user and returns a token that will be used in subsequent web services calls.  This is supported for CyberArk authentication only, and not for third party authentication.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharedLogonAuthenticationApi.Logon0(context.Background()).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharedLogonAuthenticationApi.Logon0``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLogon0Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string** |  | 
 **body** | **string** | Shared authentication is based on a user credential file that is stored in the PVWA web server. During shared authentication, only the user defined in the credential file can log on to the PVWA, but multiple users can use the logon token.  This type of authentication **requires** the application using the REST services to manage the users as the Vault can&#39;t identify which specific user performs each action.  Multiple concurrent connections can be created using the same token, without affecting each other.  The shared user is defined in a user credential file, whose location is specified in the WSCredentialFile parameter, in the appsettings section of the PVWAweb.config file:  &#x60;&#x60;&#x60; &lt;add key&#x3D;\&quot;WSCredentialFile\&quot; value&#x3D;\&quot;C:\\CyberArk\\Password Vault Web Access\\CredFiles\\WSUser.ini\&quot;/&gt; &#x60;&#x60;&#x60;  * Make sure that this user can access the PVWA interface. * Make sure the user only has the permissions in the Vault that they require.  For information about securing communication when using the SDK, refer to the following:  * [Securing application-to-REST communication](https://docs.cyberark.com/Product-Doc/OnlineHelp/PAS/Latest/en/Content/SDK/Securing%20Communication.htm) * [Configuring client authentication via certificates](https://docs.cyberark.com/Product-Doc/OnlineHelp/PAS/Latest/en/Content/SDK/Configuring%20Client%20Authentication%20via%20Client%20Certificates.htm)  This method authenticates to the Vault with a shared webservices user and returns a token that will be used in subsequent web services calls.  This is supported for CyberArk authentication only, and not for third party authentication. | 

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

