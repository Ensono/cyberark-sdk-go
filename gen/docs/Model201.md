# Model201

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [default to "29_7"]
**Name** | Pointer to **string** |  | [optional] [default to "Database-MySQL-db2.joegarcia.dev-cluster02sqluser01"]
**Address** | Pointer to **string** |  | [optional] [default to "db2.joegarcia.dev"]
**UserName** | Pointer to **string** |  | [optional] [default to "cluster02sqluser01"]
**PlatformId** | Pointer to **string** |  | [optional] [default to "MySQLServer-DualAccounts"]
**SafeName** | Pointer to **string** |  | [optional] [default to "D-MySQL-Users"]
**SecretType** | Pointer to **string** |  | [optional] [default to "password"]
**PlatformAccountProperties** | Pointer to [**Model201PlatformAccountProperties**](201PlatformAccountProperties.md) |  | [optional] 
**SecretManagement** | Pointer to [**Model201SecretManagement**](201SecretManagement.md) |  | [optional] 
**CreatedTime** | Pointer to **int32** |  | [optional] [default to 1581084295]

## Methods

### NewModel201

`func NewModel201() *Model201`

NewModel201 instantiates a new Model201 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModel201WithDefaults

`func NewModel201WithDefaults() *Model201`

NewModel201WithDefaults instantiates a new Model201 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Model201) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Model201) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Model201) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Model201) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Model201) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Model201) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Model201) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Model201) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAddress

`func (o *Model201) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Model201) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Model201) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Model201) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetUserName

`func (o *Model201) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *Model201) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *Model201) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *Model201) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetPlatformId

`func (o *Model201) GetPlatformId() string`

GetPlatformId returns the PlatformId field if non-nil, zero value otherwise.

### GetPlatformIdOk

`func (o *Model201) GetPlatformIdOk() (*string, bool)`

GetPlatformIdOk returns a tuple with the PlatformId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformId

`func (o *Model201) SetPlatformId(v string)`

SetPlatformId sets PlatformId field to given value.

### HasPlatformId

`func (o *Model201) HasPlatformId() bool`

HasPlatformId returns a boolean if a field has been set.

### GetSafeName

`func (o *Model201) GetSafeName() string`

GetSafeName returns the SafeName field if non-nil, zero value otherwise.

### GetSafeNameOk

`func (o *Model201) GetSafeNameOk() (*string, bool)`

GetSafeNameOk returns a tuple with the SafeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeName

`func (o *Model201) SetSafeName(v string)`

SetSafeName sets SafeName field to given value.

### HasSafeName

`func (o *Model201) HasSafeName() bool`

HasSafeName returns a boolean if a field has been set.

### GetSecretType

`func (o *Model201) GetSecretType() string`

GetSecretType returns the SecretType field if non-nil, zero value otherwise.

### GetSecretTypeOk

`func (o *Model201) GetSecretTypeOk() (*string, bool)`

GetSecretTypeOk returns a tuple with the SecretType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretType

`func (o *Model201) SetSecretType(v string)`

SetSecretType sets SecretType field to given value.

### HasSecretType

`func (o *Model201) HasSecretType() bool`

HasSecretType returns a boolean if a field has been set.

### GetPlatformAccountProperties

`func (o *Model201) GetPlatformAccountProperties() Model201PlatformAccountProperties`

GetPlatformAccountProperties returns the PlatformAccountProperties field if non-nil, zero value otherwise.

### GetPlatformAccountPropertiesOk

`func (o *Model201) GetPlatformAccountPropertiesOk() (*Model201PlatformAccountProperties, bool)`

GetPlatformAccountPropertiesOk returns a tuple with the PlatformAccountProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformAccountProperties

`func (o *Model201) SetPlatformAccountProperties(v Model201PlatformAccountProperties)`

SetPlatformAccountProperties sets PlatformAccountProperties field to given value.

### HasPlatformAccountProperties

`func (o *Model201) HasPlatformAccountProperties() bool`

HasPlatformAccountProperties returns a boolean if a field has been set.

### GetSecretManagement

`func (o *Model201) GetSecretManagement() Model201SecretManagement`

GetSecretManagement returns the SecretManagement field if non-nil, zero value otherwise.

### GetSecretManagementOk

`func (o *Model201) GetSecretManagementOk() (*Model201SecretManagement, bool)`

GetSecretManagementOk returns a tuple with the SecretManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretManagement

`func (o *Model201) SetSecretManagement(v Model201SecretManagement)`

SetSecretManagement sets SecretManagement field to given value.

### HasSecretManagement

`func (o *Model201) HasSecretManagement() bool`

HasSecretManagement returns a boolean if a field has been set.

### GetCreatedTime

`func (o *Model201) GetCreatedTime() int32`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Model201) GetCreatedTimeOk() (*int32, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Model201) SetCreatedTime(v int32)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *Model201) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


