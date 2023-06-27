/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 07 June, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.34.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// UpdateDeviceApplianceUplinksSettingsRequestInterfaces Interface settings.
type UpdateDeviceApplianceUplinksSettingsRequestInterfaces struct {
	Wan1 *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1 `json:"wan1,omitempty"`
	Wan2 *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2 `json:"wan2,omitempty"`
}

// NewUpdateDeviceApplianceUplinksSettingsRequestInterfaces instantiates a new UpdateDeviceApplianceUplinksSettingsRequestInterfaces object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDeviceApplianceUplinksSettingsRequestInterfaces() *UpdateDeviceApplianceUplinksSettingsRequestInterfaces {
	this := UpdateDeviceApplianceUplinksSettingsRequestInterfaces{}
	return &this
}

// NewUpdateDeviceApplianceUplinksSettingsRequestInterfacesWithDefaults instantiates a new UpdateDeviceApplianceUplinksSettingsRequestInterfaces object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDeviceApplianceUplinksSettingsRequestInterfacesWithDefaults() *UpdateDeviceApplianceUplinksSettingsRequestInterfaces {
	this := UpdateDeviceApplianceUplinksSettingsRequestInterfaces{}
	return &this
}

// GetWan1 returns the Wan1 field value if set, zero value otherwise.
func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfaces) GetWan1() UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1 {
	if o == nil || isNil(o.Wan1) {
		var ret UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1
		return ret
	}
	return *o.Wan1
}

// GetWan1Ok returns a tuple with the Wan1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfaces) GetWan1Ok() (*UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1, bool) {
	if o == nil || isNil(o.Wan1) {
    return nil, false
	}
	return o.Wan1, true
}

// HasWan1 returns a boolean if a field has been set.
func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfaces) HasWan1() bool {
	if o != nil && !isNil(o.Wan1) {
		return true
	}

	return false
}

// SetWan1 gets a reference to the given UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1 and assigns it to the Wan1 field.
func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfaces) SetWan1(v UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1) {
	o.Wan1 = &v
}

// GetWan2 returns the Wan2 field value if set, zero value otherwise.
func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfaces) GetWan2() UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2 {
	if o == nil || isNil(o.Wan2) {
		var ret UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2
		return ret
	}
	return *o.Wan2
}

// GetWan2Ok returns a tuple with the Wan2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfaces) GetWan2Ok() (*UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2, bool) {
	if o == nil || isNil(o.Wan2) {
    return nil, false
	}
	return o.Wan2, true
}

// HasWan2 returns a boolean if a field has been set.
func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfaces) HasWan2() bool {
	if o != nil && !isNil(o.Wan2) {
		return true
	}

	return false
}

// SetWan2 gets a reference to the given UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2 and assigns it to the Wan2 field.
func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfaces) SetWan2(v UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan2) {
	o.Wan2 = &v
}

func (o UpdateDeviceApplianceUplinksSettingsRequestInterfaces) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Wan1) {
		toSerialize["wan1"] = o.Wan1
	}
	if !isNil(o.Wan2) {
		toSerialize["wan2"] = o.Wan2
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateDeviceApplianceUplinksSettingsRequestInterfaces struct {
	value *UpdateDeviceApplianceUplinksSettingsRequestInterfaces
	isSet bool
}

func (v NullableUpdateDeviceApplianceUplinksSettingsRequestInterfaces) Get() *UpdateDeviceApplianceUplinksSettingsRequestInterfaces {
	return v.value
}

func (v *NullableUpdateDeviceApplianceUplinksSettingsRequestInterfaces) Set(val *UpdateDeviceApplianceUplinksSettingsRequestInterfaces) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDeviceApplianceUplinksSettingsRequestInterfaces) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDeviceApplianceUplinksSettingsRequestInterfaces) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDeviceApplianceUplinksSettingsRequestInterfaces(val *UpdateDeviceApplianceUplinksSettingsRequestInterfaces) *NullableUpdateDeviceApplianceUplinksSettingsRequestInterfaces {
	return &NullableUpdateDeviceApplianceUplinksSettingsRequestInterfaces{value: val, isSet: true}
}

func (v NullableUpdateDeviceApplianceUplinksSettingsRequestInterfaces) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDeviceApplianceUplinksSettingsRequestInterfaces) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


