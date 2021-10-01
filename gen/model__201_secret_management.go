/*
CyberArkIAG

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 18a45ad8-77e8-4ecc-873e-787c6de10a60
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// 201SecretManagement struct for 201SecretManagement
type 201SecretManagement struct {
	AutomaticManagementEnabled interface{} `json:"automaticManagementEnabled,omitempty"`
	ManualManagementReason *string `json:"manualManagementReason,omitempty"`
	LastModifiedTime *int32 `json:"lastModifiedTime,omitempty"`
}

// New201SecretManagement instantiates a new 201SecretManagement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func New201SecretManagement() *201SecretManagement {
	this := 201SecretManagement{}
	var manualManagementReason string = "testing"
	this.ManualManagementReason = &manualManagementReason
	var lastModifiedTime int32 = 1581084295
	this.LastModifiedTime = &lastModifiedTime
	return &this
}

// New201SecretManagementWithDefaults instantiates a new 201SecretManagement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func New201SecretManagementWithDefaults() *201SecretManagement {
	this := 201SecretManagement{}
	var manualManagementReason string = "testing"
	this.ManualManagementReason = &manualManagementReason
	var lastModifiedTime int32 = 1581084295
	this.LastModifiedTime = &lastModifiedTime
	return &this
}

// GetAutomaticManagementEnabled returns the AutomaticManagementEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *201SecretManagement) GetAutomaticManagementEnabled() interface{} {
	if o == nil  {
		var ret interface{}
		return ret
	}
	return o.AutomaticManagementEnabled
}

// GetAutomaticManagementEnabledOk returns a tuple with the AutomaticManagementEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *201SecretManagement) GetAutomaticManagementEnabledOk() (*interface{}, bool) {
	if o == nil || o.AutomaticManagementEnabled == nil {
		return nil, false
	}
	return &o.AutomaticManagementEnabled, true
}

// HasAutomaticManagementEnabled returns a boolean if a field has been set.
func (o *201SecretManagement) HasAutomaticManagementEnabled() bool {
	if o != nil && o.AutomaticManagementEnabled != nil {
		return true
	}

	return false
}

// SetAutomaticManagementEnabled gets a reference to the given interface{} and assigns it to the AutomaticManagementEnabled field.
func (o *201SecretManagement) SetAutomaticManagementEnabled(v interface{}) {
	o.AutomaticManagementEnabled = v
}

// GetManualManagementReason returns the ManualManagementReason field value if set, zero value otherwise.
func (o *201SecretManagement) GetManualManagementReason() string {
	if o == nil || o.ManualManagementReason == nil {
		var ret string
		return ret
	}
	return *o.ManualManagementReason
}

// GetManualManagementReasonOk returns a tuple with the ManualManagementReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *201SecretManagement) GetManualManagementReasonOk() (*string, bool) {
	if o == nil || o.ManualManagementReason == nil {
		return nil, false
	}
	return o.ManualManagementReason, true
}

// HasManualManagementReason returns a boolean if a field has been set.
func (o *201SecretManagement) HasManualManagementReason() bool {
	if o != nil && o.ManualManagementReason != nil {
		return true
	}

	return false
}

// SetManualManagementReason gets a reference to the given string and assigns it to the ManualManagementReason field.
func (o *201SecretManagement) SetManualManagementReason(v string) {
	o.ManualManagementReason = &v
}

// GetLastModifiedTime returns the LastModifiedTime field value if set, zero value otherwise.
func (o *201SecretManagement) GetLastModifiedTime() int32 {
	if o == nil || o.LastModifiedTime == nil {
		var ret int32
		return ret
	}
	return *o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *201SecretManagement) GetLastModifiedTimeOk() (*int32, bool) {
	if o == nil || o.LastModifiedTime == nil {
		return nil, false
	}
	return o.LastModifiedTime, true
}

// HasLastModifiedTime returns a boolean if a field has been set.
func (o *201SecretManagement) HasLastModifiedTime() bool {
	if o != nil && o.LastModifiedTime != nil {
		return true
	}

	return false
}

// SetLastModifiedTime gets a reference to the given int32 and assigns it to the LastModifiedTime field.
func (o *201SecretManagement) SetLastModifiedTime(v int32) {
	o.LastModifiedTime = &v
}

func (o 201SecretManagement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AutomaticManagementEnabled != nil {
		toSerialize["automaticManagementEnabled"] = o.AutomaticManagementEnabled
	}
	if o.ManualManagementReason != nil {
		toSerialize["manualManagementReason"] = o.ManualManagementReason
	}
	if o.LastModifiedTime != nil {
		toSerialize["lastModifiedTime"] = o.LastModifiedTime
	}
	return json.Marshal(toSerialize)
}

type Nullable201SecretManagement struct {
	value *201SecretManagement
	isSet bool
}

func (v Nullable201SecretManagement) Get() *201SecretManagement {
	return v.value
}

func (v *Nullable201SecretManagement) Set(val *201SecretManagement) {
	v.value = val
	v.isSet = true
}

func (v Nullable201SecretManagement) IsSet() bool {
	return v.isSet
}

func (v *Nullable201SecretManagement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullable201SecretManagement(val *201SecretManagement) *Nullable201SecretManagement {
	return &Nullable201SecretManagement{value: val, isSet: true}
}

func (v Nullable201SecretManagement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *Nullable201SecretManagement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

