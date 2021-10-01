# Model200

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [default to "24_3"]
**Name** | Pointer to **string** |  | [optional] [default to "ANSIBLE-RESTAPI-USER"]
**Address** | Pointer to **string** |  | [optional] [default to "192.168.3.101"]
**UserName** | Pointer to **string** |  | [optional] [default to "Svc_AnsibleREST"]
**PlatformId** | Pointer to **string** |  | [optional] [default to "JG-CyberArkVault"]
**SafeName** | Pointer to **string** |  | [optional] [default to "D-CYBR-RESTAPI-ACCTS"]
**SecretType** | Pointer to **string** |  | [optional] [default to "password"]
**SecretManagement** | Pointer to [**Model200SecretManagement**](200SecretManagement.md) |  | [optional] 
**CreatedTime** | Pointer to **int32** |  | [optional] [default to 1541876282]

## Methods

### NewModel200

`func NewModel200() *Model200`

NewModel200 instantiates a new Model200 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModel200WithDefaults

`func NewModel200WithDefaults() *Model200`

NewModel200WithDefaults instantiates a new Model200 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Model200) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Model200) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Model200) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Model200) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Model200) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Model200) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Model200) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Model200) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAddress

`func (o *Model200) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Model200) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Model200) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Model200) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetUserName

`func (o *Model200) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *Model200) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *Model200) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *Model200) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetPlatformId

`func (o *Model200) GetPlatformId() string`

GetPlatformId returns the PlatformId field if non-nil, zero value otherwise.

### GetPlatformIdOk

`func (o *Model200) GetPlatformIdOk() (*string, bool)`

GetPlatformIdOk returns a tuple with the PlatformId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformId

`func (o *Model200) SetPlatformId(v string)`

SetPlatformId sets PlatformId field to given value.

### HasPlatformId

`func (o *Model200) HasPlatformId() bool`

HasPlatformId returns a boolean if a field has been set.

### GetSafeName

`func (o *Model200) GetSafeName() string`

GetSafeName returns the SafeName field if non-nil, zero value otherwise.

### GetSafeNameOk

`func (o *Model200) GetSafeNameOk() (*string, bool)`

GetSafeNameOk returns a tuple with the SafeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeName

`func (o *Model200) SetSafeName(v string)`

SetSafeName sets SafeName field to given value.

### HasSafeName

`func (o *Model200) HasSafeName() bool`

HasSafeName returns a boolean if a field has been set.

### GetSecretType

`func (o *Model200) GetSecretType() string`

GetSecretType returns the SecretType field if non-nil, zero value otherwise.

### GetSecretTypeOk

`func (o *Model200) GetSecretTypeOk() (*string, bool)`

GetSecretTypeOk returns a tuple with the SecretType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretType

`func (o *Model200) SetSecretType(v string)`

SetSecretType sets SecretType field to given value.

### HasSecretType

`func (o *Model200) HasSecretType() bool`

HasSecretType returns a boolean if a field has been set.

### GetSecretManagement

`func (o *Model200) GetSecretManagement() Model200SecretManagement`

GetSecretManagement returns the SecretManagement field if non-nil, zero value otherwise.

### GetSecretManagementOk

`func (o *Model200) GetSecretManagementOk() (*Model200SecretManagement, bool)`

GetSecretManagementOk returns a tuple with the SecretManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretManagement

`func (o *Model200) SetSecretManagement(v Model200SecretManagement)`

SetSecretManagement sets SecretManagement field to given value.

### HasSecretManagement

`func (o *Model200) HasSecretManagement() bool`

HasSecretManagement returns a boolean if a field has been set.

### GetCreatedTime

`func (o *Model200) GetCreatedTime() int32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Model200) GetCreatedTimeOk() (*int32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Model200) SetCreatedTime(v int32)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *Model200) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


