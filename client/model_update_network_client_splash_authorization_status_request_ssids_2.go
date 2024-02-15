/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 07 February, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.43.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the UpdateNetworkClientSplashAuthorizationStatusRequestSsids2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkClientSplashAuthorizationStatusRequestSsids2{}

// UpdateNetworkClientSplashAuthorizationStatusRequestSsids2 Splash authorization for SSID 2
type UpdateNetworkClientSplashAuthorizationStatusRequestSsids2 struct {
	// New authorization status for the SSID (true, false).
	IsAuthorized *bool `json:"isAuthorized,omitempty"`
}

// NewUpdateNetworkClientSplashAuthorizationStatusRequestSsids2 instantiates a new UpdateNetworkClientSplashAuthorizationStatusRequestSsids2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkClientSplashAuthorizationStatusRequestSsids2() *UpdateNetworkClientSplashAuthorizationStatusRequestSsids2 {
	this := UpdateNetworkClientSplashAuthorizationStatusRequestSsids2{}
	return &this
}

// NewUpdateNetworkClientSplashAuthorizationStatusRequestSsids2WithDefaults instantiates a new UpdateNetworkClientSplashAuthorizationStatusRequestSsids2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkClientSplashAuthorizationStatusRequestSsids2WithDefaults() *UpdateNetworkClientSplashAuthorizationStatusRequestSsids2 {
	this := UpdateNetworkClientSplashAuthorizationStatusRequestSsids2{}
	return &this
}

// GetIsAuthorized returns the IsAuthorized field value if set, zero value otherwise.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids2) GetIsAuthorized() bool {
	if o == nil || IsNil(o.IsAuthorized) {
		var ret bool
		return ret
	}
	return *o.IsAuthorized
}

// GetIsAuthorizedOk returns a tuple with the IsAuthorized field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids2) GetIsAuthorizedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAuthorized) {
		return nil, false
	}
	return o.IsAuthorized, true
}

// HasIsAuthorized returns a boolean if a field has been set.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids2) HasIsAuthorized() bool {
	if o != nil && !IsNil(o.IsAuthorized) {
		return true
	}

	return false
}

// SetIsAuthorized gets a reference to the given bool and assigns it to the IsAuthorized field.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids2) SetIsAuthorized(v bool) {
	o.IsAuthorized = &v
}

func (o UpdateNetworkClientSplashAuthorizationStatusRequestSsids2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkClientSplashAuthorizationStatusRequestSsids2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsAuthorized) {
		toSerialize["isAuthorized"] = o.IsAuthorized
	}
	return toSerialize, nil
}

type NullableUpdateNetworkClientSplashAuthorizationStatusRequestSsids2 struct {
	value *UpdateNetworkClientSplashAuthorizationStatusRequestSsids2
	isSet bool
}

func (v NullableUpdateNetworkClientSplashAuthorizationStatusRequestSsids2) Get() *UpdateNetworkClientSplashAuthorizationStatusRequestSsids2 {
	return v.value
}

func (v *NullableUpdateNetworkClientSplashAuthorizationStatusRequestSsids2) Set(val *UpdateNetworkClientSplashAuthorizationStatusRequestSsids2) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkClientSplashAuthorizationStatusRequestSsids2) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkClientSplashAuthorizationStatusRequestSsids2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkClientSplashAuthorizationStatusRequestSsids2(val *UpdateNetworkClientSplashAuthorizationStatusRequestSsids2) *NullableUpdateNetworkClientSplashAuthorizationStatusRequestSsids2 {
	return &NullableUpdateNetworkClientSplashAuthorizationStatusRequestSsids2{value: val, isSet: true}
}

func (v NullableUpdateNetworkClientSplashAuthorizationStatusRequestSsids2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkClientSplashAuthorizationStatusRequestSsids2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


