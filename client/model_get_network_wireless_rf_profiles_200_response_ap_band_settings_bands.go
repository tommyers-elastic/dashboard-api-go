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

// checks if the GetNetworkWirelessRfProfiles200ResponseApBandSettingsBands type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkWirelessRfProfiles200ResponseApBandSettingsBands{}

// GetNetworkWirelessRfProfiles200ResponseApBandSettingsBands Settings related to all bands
type GetNetworkWirelessRfProfiles200ResponseApBandSettingsBands struct {
	// List of enabled bands. Can include [\"2.4\", \"5\", \"6\"]
	Enabled []string `json:"enabled,omitempty"`
}

// NewGetNetworkWirelessRfProfiles200ResponseApBandSettingsBands instantiates a new GetNetworkWirelessRfProfiles200ResponseApBandSettingsBands object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkWirelessRfProfiles200ResponseApBandSettingsBands() *GetNetworkWirelessRfProfiles200ResponseApBandSettingsBands {
	this := GetNetworkWirelessRfProfiles200ResponseApBandSettingsBands{}
	return &this
}

// NewGetNetworkWirelessRfProfiles200ResponseApBandSettingsBandsWithDefaults instantiates a new GetNetworkWirelessRfProfiles200ResponseApBandSettingsBands object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkWirelessRfProfiles200ResponseApBandSettingsBandsWithDefaults() *GetNetworkWirelessRfProfiles200ResponseApBandSettingsBands {
	this := GetNetworkWirelessRfProfiles200ResponseApBandSettingsBands{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *GetNetworkWirelessRfProfiles200ResponseApBandSettingsBands) GetEnabled() []string {
	if o == nil || IsNil(o.Enabled) {
		var ret []string
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessRfProfiles200ResponseApBandSettingsBands) GetEnabledOk() ([]string, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *GetNetworkWirelessRfProfiles200ResponseApBandSettingsBands) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given []string and assigns it to the Enabled field.
func (o *GetNetworkWirelessRfProfiles200ResponseApBandSettingsBands) SetEnabled(v []string) {
	o.Enabled = v
}

func (o GetNetworkWirelessRfProfiles200ResponseApBandSettingsBands) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkWirelessRfProfiles200ResponseApBandSettingsBands) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableGetNetworkWirelessRfProfiles200ResponseApBandSettingsBands struct {
	value *GetNetworkWirelessRfProfiles200ResponseApBandSettingsBands
	isSet bool
}

func (v NullableGetNetworkWirelessRfProfiles200ResponseApBandSettingsBands) Get() *GetNetworkWirelessRfProfiles200ResponseApBandSettingsBands {
	return v.value
}

func (v *NullableGetNetworkWirelessRfProfiles200ResponseApBandSettingsBands) Set(val *GetNetworkWirelessRfProfiles200ResponseApBandSettingsBands) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkWirelessRfProfiles200ResponseApBandSettingsBands) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkWirelessRfProfiles200ResponseApBandSettingsBands) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkWirelessRfProfiles200ResponseApBandSettingsBands(val *GetNetworkWirelessRfProfiles200ResponseApBandSettingsBands) *NullableGetNetworkWirelessRfProfiles200ResponseApBandSettingsBands {
	return &NullableGetNetworkWirelessRfProfiles200ResponseApBandSettingsBands{value: val, isSet: true}
}

func (v NullableGetNetworkWirelessRfProfiles200ResponseApBandSettingsBands) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkWirelessRfProfiles200ResponseApBandSettingsBands) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


