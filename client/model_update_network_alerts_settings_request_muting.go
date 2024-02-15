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

// checks if the UpdateNetworkAlertsSettingsRequestMuting type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkAlertsSettingsRequestMuting{}

// UpdateNetworkAlertsSettingsRequestMuting Mute alerts under certain conditions
type UpdateNetworkAlertsSettingsRequestMuting struct {
	ByPortSchedules *UpdateNetworkAlertsSettingsRequestMutingByPortSchedules `json:"byPortSchedules,omitempty"`
}

// NewUpdateNetworkAlertsSettingsRequestMuting instantiates a new UpdateNetworkAlertsSettingsRequestMuting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkAlertsSettingsRequestMuting() *UpdateNetworkAlertsSettingsRequestMuting {
	this := UpdateNetworkAlertsSettingsRequestMuting{}
	return &this
}

// NewUpdateNetworkAlertsSettingsRequestMutingWithDefaults instantiates a new UpdateNetworkAlertsSettingsRequestMuting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkAlertsSettingsRequestMutingWithDefaults() *UpdateNetworkAlertsSettingsRequestMuting {
	this := UpdateNetworkAlertsSettingsRequestMuting{}
	return &this
}

// GetByPortSchedules returns the ByPortSchedules field value if set, zero value otherwise.
func (o *UpdateNetworkAlertsSettingsRequestMuting) GetByPortSchedules() UpdateNetworkAlertsSettingsRequestMutingByPortSchedules {
	if o == nil || IsNil(o.ByPortSchedules) {
		var ret UpdateNetworkAlertsSettingsRequestMutingByPortSchedules
		return ret
	}
	return *o.ByPortSchedules
}

// GetByPortSchedulesOk returns a tuple with the ByPortSchedules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkAlertsSettingsRequestMuting) GetByPortSchedulesOk() (*UpdateNetworkAlertsSettingsRequestMutingByPortSchedules, bool) {
	if o == nil || IsNil(o.ByPortSchedules) {
		return nil, false
	}
	return o.ByPortSchedules, true
}

// HasByPortSchedules returns a boolean if a field has been set.
func (o *UpdateNetworkAlertsSettingsRequestMuting) HasByPortSchedules() bool {
	if o != nil && !IsNil(o.ByPortSchedules) {
		return true
	}

	return false
}

// SetByPortSchedules gets a reference to the given UpdateNetworkAlertsSettingsRequestMutingByPortSchedules and assigns it to the ByPortSchedules field.
func (o *UpdateNetworkAlertsSettingsRequestMuting) SetByPortSchedules(v UpdateNetworkAlertsSettingsRequestMutingByPortSchedules) {
	o.ByPortSchedules = &v
}

func (o UpdateNetworkAlertsSettingsRequestMuting) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkAlertsSettingsRequestMuting) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ByPortSchedules) {
		toSerialize["byPortSchedules"] = o.ByPortSchedules
	}
	return toSerialize, nil
}

type NullableUpdateNetworkAlertsSettingsRequestMuting struct {
	value *UpdateNetworkAlertsSettingsRequestMuting
	isSet bool
}

func (v NullableUpdateNetworkAlertsSettingsRequestMuting) Get() *UpdateNetworkAlertsSettingsRequestMuting {
	return v.value
}

func (v *NullableUpdateNetworkAlertsSettingsRequestMuting) Set(val *UpdateNetworkAlertsSettingsRequestMuting) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkAlertsSettingsRequestMuting) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkAlertsSettingsRequestMuting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkAlertsSettingsRequestMuting(val *UpdateNetworkAlertsSettingsRequestMuting) *NullableUpdateNetworkAlertsSettingsRequestMuting {
	return &NullableUpdateNetworkAlertsSettingsRequestMuting{value: val, isSet: true}
}

func (v NullableUpdateNetworkAlertsSettingsRequestMuting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkAlertsSettingsRequestMuting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


