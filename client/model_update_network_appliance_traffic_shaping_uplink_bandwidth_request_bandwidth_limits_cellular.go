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

// checks if the UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsCellular type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsCellular{}

// UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsCellular The bandwidth settings for the 'cellular' uplink
type UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsCellular struct {
	// The maximum upload limit (integer, in Kbps). null indicates no limit
	LimitUp *int32 `json:"limitUp,omitempty"`
	// The maximum download limit (integer, in Kbps). null indicates no limit
	LimitDown *int32 `json:"limitDown,omitempty"`
}

// NewUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsCellular instantiates a new UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsCellular object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsCellular() *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsCellular {
	this := UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsCellular{}
	return &this
}

// NewUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsCellularWithDefaults instantiates a new UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsCellular object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsCellularWithDefaults() *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsCellular {
	this := UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsCellular{}
	return &this
}

// GetLimitUp returns the LimitUp field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsCellular) GetLimitUp() int32 {
	if o == nil || IsNil(o.LimitUp) {
		var ret int32
		return ret
	}
	return *o.LimitUp
}

// GetLimitUpOk returns a tuple with the LimitUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsCellular) GetLimitUpOk() (*int32, bool) {
	if o == nil || IsNil(o.LimitUp) {
		return nil, false
	}
	return o.LimitUp, true
}

// HasLimitUp returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsCellular) HasLimitUp() bool {
	if o != nil && !IsNil(o.LimitUp) {
		return true
	}

	return false
}

// SetLimitUp gets a reference to the given int32 and assigns it to the LimitUp field.
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsCellular) SetLimitUp(v int32) {
	o.LimitUp = &v
}

// GetLimitDown returns the LimitDown field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsCellular) GetLimitDown() int32 {
	if o == nil || IsNil(o.LimitDown) {
		var ret int32
		return ret
	}
	return *o.LimitDown
}

// GetLimitDownOk returns a tuple with the LimitDown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsCellular) GetLimitDownOk() (*int32, bool) {
	if o == nil || IsNil(o.LimitDown) {
		return nil, false
	}
	return o.LimitDown, true
}

// HasLimitDown returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsCellular) HasLimitDown() bool {
	if o != nil && !IsNil(o.LimitDown) {
		return true
	}

	return false
}

// SetLimitDown gets a reference to the given int32 and assigns it to the LimitDown field.
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsCellular) SetLimitDown(v int32) {
	o.LimitDown = &v
}

func (o UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsCellular) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsCellular) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LimitUp) {
		toSerialize["limitUp"] = o.LimitUp
	}
	if !IsNil(o.LimitDown) {
		toSerialize["limitDown"] = o.LimitDown
	}
	return toSerialize, nil
}

type NullableUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsCellular struct {
	value *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsCellular
	isSet bool
}

func (v NullableUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsCellular) Get() *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsCellular {
	return v.value
}

func (v *NullableUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsCellular) Set(val *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsCellular) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsCellular) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsCellular) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsCellular(val *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsCellular) *NullableUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsCellular {
	return &NullableUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsCellular{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsCellular) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsCellular) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


