# 201SecretManagement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutomaticManagementEnabled** | Pointer to **interface{}** |  | [optional] 
**ManualManagementReason** | Pointer to **string** |  | [optional] [default to "testing"]
**LastModifiedTime** | Pointer to **int32** |  | [optional] [default to 1581084295]

## Methods

### New201SecretManagement

`func New201SecretManagement() *201SecretManagement`

New201SecretManagement instantiates a new 201SecretManagement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### New201SecretManagementWithDefaults

`func New201SecretManagementWithDefaults() *201SecretManagement`

New201SecretManagementWithDefaults instantiates a new 201SecretManagement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutomaticManagementEnabled

`func (o *201SecretManagement) GetAutomaticManagementEnabled() interface{}`

GetAutomaticManagementEnabled returns the AutomaticManagementEnabled field if non-nil, zero value otherwise.

### GetAutomaticManagementEnabledOk

`func (o *201SecretManagement) GetAutomaticManagementEnabledOk() (*interface{}, bool)`

GetAutomaticManagementEnabledOk returns a tuple with the AutomaticManagementEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticManagementEnabled

`func (o *201SecretManagement) SetAutomaticManagementEnabled(v interface{})`

SetAutomaticManagementEnabled sets AutomaticManagementEnabled field to given value.

### HasAutomaticManagementEnabled

`func (o *201SecretManagement) HasAutomaticManagementEnabled() bool`

HasAutomaticManagementEnabled returns a boolean if a field has been set.

### SetAutomaticManagementEnabledNil

`func (o *201SecretManagement) SetAutomaticManagementEnabledNil(b bool)`

 SetAutomaticManagementEnabledNil sets the value for AutomaticManagementEnabled to be an explicit nil

### UnsetAutomaticManagementEnabled
`func (o *201SecretManagement) UnsetAutomaticManagementEnabled()`

UnsetAutomaticManagementEnabled ensures that no value is present for AutomaticManagementEnabled, not even an explicit nil
### GetManualManagementReason

`func (o *201SecretManagement) GetManualManagementReason() string`

GetManualManagementReason returns the ManualManagementReason field if non-nil, zero value otherwise.

### GetManualManagementReasonOk

`func (o *201SecretManagement) GetManualManagementReasonOk() (*string, bool)`

GetManualManagementReasonOk returns a tuple with the ManualManagementReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualManagementReason

`func (o *201SecretManagement) SetManualManagementReason(v string)`

SetManualManagementReason sets ManualManagementReason field to given value.

### HasManualManagementReason

`func (o *201SecretManagement) HasManualManagementReason() bool`

HasManualManagementReason returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *201SecretManagement) GetLastModifiedTime() int32`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *201SecretManagement) GetLastModifiedTimeOk() (*int32, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *201SecretManagement) SetLastModifiedTime(v int32)`

SetLastModifiedTime sets LastModifiedTime field to given value.

### HasLastModifiedTime

`func (o *201SecretManagement) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


