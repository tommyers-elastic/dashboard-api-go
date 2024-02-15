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

// checks if the UpdateNetworkWirelessBillingRequestPlansInnerBandwidthLimits type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkWirelessBillingRequestPlansInnerBandwidthLimits{}

// UpdateNetworkWirelessBillingRequestPlansInnerBandwidthLimits The uplink bandwidth settings for the pricing plan.
type UpdateNetworkWirelessBillingRequestPlansInnerBandwidthLimits struct {
	// The maximum upload limit (integer, in Kbps). null indicates no limit
	LimitUp *int32 `json:"limitUp,omitempty"`
	// The maximum download limit (integer, in Kbps). null indicates no limit
	LimitDown *int32 `json:"limitDown,omitempty"`
}

// NewUpdateNetworkWirelessBillingRequestPlansInnerBandwidthLimits instantiates a new UpdateNetworkWirelessBillingRequestPlansInnerBandwidthLimits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessBillingRequestPlansInnerBandwidthLimits() *UpdateNetworkWirelessBillingRequestPlansInnerBandwidthLimits {
	this := UpdateNetworkWirelessBillingRequestPlansInnerBandwidthLimits{}
	return &this
}

// NewUpdateNetworkWirelessBillingRequestPlansInnerBandwidthLimitsWithDefaults instantiates a new UpdateNetworkWirelessBillingRequestPlansInnerBandwidthLimits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessBillingRequestPlansInnerBandwidthLimitsWithDefaults() *UpdateNetworkWirelessBillingRequestPlansInnerBandwidthLimits {
	this := UpdateNetworkWirelessBillingRequestPlansInnerBandwidthLimits{}
	return &this
}

// GetLimitUp returns the LimitUp field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessBillingRequestPlansInnerBandwidthLimits) GetLimitUp() int32 {
	if o == nil || IsNil(o.LimitUp) {
		var ret int32
		return ret
	}
	return *o.LimitUp
}

// GetLimitUpOk returns a tuple with the LimitUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessBillingRequestPlansInnerBandwidthLimits) GetLimitUpOk() (*int32, bool) {
	if o == nil || IsNil(o.LimitUp) {
		return nil, false
	}
	return o.LimitUp, true
}

// HasLimitUp returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessBillingRequestPlansInnerBandwidthLimits) HasLimitUp() bool {
	if o != nil && !IsNil(o.LimitUp) {
		return true
	}

	return false
}

// SetLimitUp gets a reference to the given int32 and assigns it to the LimitUp field.
func (o *UpdateNetworkWirelessBillingRequestPlansInnerBandwidthLimits) SetLimitUp(v int32) {
	o.LimitUp = &v
}

// GetLimitDown returns the LimitDown field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessBillingRequestPlansInnerBandwidthLimits) GetLimitDown() int32 {
	if o == nil || IsNil(o.LimitDown) {
		var ret int32
		return ret
	}
	return *o.LimitDown
}

// GetLimitDownOk returns a tuple with the LimitDown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessBillingRequestPlansInnerBandwidthLimits) GetLimitDownOk() (*int32, bool) {
	if o == nil || IsNil(o.LimitDown) {
		return nil, false
	}
	return o.LimitDown, true
}

// HasLimitDown returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessBillingRequestPlansInnerBandwidthLimits) HasLimitDown() bool {
	if o != nil && !IsNil(o.LimitDown) {
		return true
	}

	return false
}

// SetLimitDown gets a reference to the given int32 and assigns it to the LimitDown field.
func (o *UpdateNetworkWirelessBillingRequestPlansInnerBandwidthLimits) SetLimitDown(v int32) {
	o.LimitDown = &v
}

func (o UpdateNetworkWirelessBillingRequestPlansInnerBandwidthLimits) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkWirelessBillingRequestPlansInnerBandwidthLimits) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LimitUp) {
		toSerialize["limitUp"] = o.LimitUp
	}
	if !IsNil(o.LimitDown) {
		toSerialize["limitDown"] = o.LimitDown
	}
	return toSerialize, nil
}

type NullableUpdateNetworkWirelessBillingRequestPlansInnerBandwidthLimits struct {
	value *UpdateNetworkWirelessBillingRequestPlansInnerBandwidthLimits
	isSet bool
}

func (v NullableUpdateNetworkWirelessBillingRequestPlansInnerBandwidthLimits) Get() *UpdateNetworkWirelessBillingRequestPlansInnerBandwidthLimits {
	return v.value
}

func (v *NullableUpdateNetworkWirelessBillingRequestPlansInnerBandwidthLimits) Set(val *UpdateNetworkWirelessBillingRequestPlansInnerBandwidthLimits) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessBillingRequestPlansInnerBandwidthLimits) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessBillingRequestPlansInnerBandwidthLimits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessBillingRequestPlansInnerBandwidthLimits(val *UpdateNetworkWirelessBillingRequestPlansInnerBandwidthLimits) *NullableUpdateNetworkWirelessBillingRequestPlansInnerBandwidthLimits {
	return &NullableUpdateNetworkWirelessBillingRequestPlansInnerBandwidthLimits{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessBillingRequestPlansInnerBandwidthLimits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessBillingRequestPlansInnerBandwidthLimits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


