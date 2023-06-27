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

// GetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailbackImmediate Immediate WAN failover and failback
type GetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailbackImmediate struct {
	// Whether immediate WAN failover and failback is enabled
	Enabled bool `json:"enabled"`
}

// NewGetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailbackImmediate instantiates a new GetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailbackImmediate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailbackImmediate(enabled bool) *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailbackImmediate {
	this := GetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailbackImmediate{}
	this.Enabled = enabled
	return &this
}

// NewGetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailbackImmediateWithDefaults instantiates a new GetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailbackImmediate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailbackImmediateWithDefaults() *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailbackImmediate {
	this := GetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailbackImmediate{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailbackImmediate) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailbackImmediate) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailbackImmediate) SetEnabled(v bool) {
	o.Enabled = v
}

func (o GetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailbackImmediate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableGetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailbackImmediate struct {
	value *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailbackImmediate
	isSet bool
}

func (v NullableGetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailbackImmediate) Get() *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailbackImmediate {
	return v.value
}

func (v *NullableGetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailbackImmediate) Set(val *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailbackImmediate) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailbackImmediate) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailbackImmediate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailbackImmediate(val *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailbackImmediate) *NullableGetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailbackImmediate {
	return &NullableGetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailbackImmediate{value: val, isSet: true}
}

func (v NullableGetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailbackImmediate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailbackImmediate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


