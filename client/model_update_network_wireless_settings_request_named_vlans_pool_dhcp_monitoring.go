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

// checks if the UpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoring type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoring{}

// UpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoring Named VLAN Pool DHCP Monitoring settings.
type UpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoring struct {
	// Whether or not devices using named VLAN pools should remove dirty VLANs from the pool, thereby preventing clients from being assigned to VLANs where they would be unable to obtain an IP address via DHCP.
	Enabled *bool `json:"enabled,omitempty"`
	// The duration in minutes that devices will refrain from using dirty VLANs before adding them back to the pool.
	Duration *int32 `json:"duration,omitempty"`
}

// NewUpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoring instantiates a new UpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoring object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoring() *UpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoring {
	this := UpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoring{}
	return &this
}

// NewUpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoringWithDefaults instantiates a new UpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoring object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoringWithDefaults() *UpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoring {
	this := UpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoring{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoring) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoring) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoring) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoring) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoring) GetDuration() int32 {
	if o == nil || IsNil(o.Duration) {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoring) GetDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoring) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *UpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoring) SetDuration(v int32) {
	o.Duration = &v
}

func (o UpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoring) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoring) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	return toSerialize, nil
}

type NullableUpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoring struct {
	value *UpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoring
	isSet bool
}

func (v NullableUpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoring) Get() *UpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoring {
	return v.value
}

func (v *NullableUpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoring) Set(val *UpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoring) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoring) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoring) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoring(val *UpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoring) *NullableUpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoring {
	return &NullableUpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoring{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoring) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoring) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


