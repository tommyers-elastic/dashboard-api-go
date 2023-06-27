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

// GetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsage An object containing disk usage details.
type GetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsage struct {
	C *GetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsageC `json:"c,omitempty"`
}

// NewGetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsage instantiates a new GetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsage() *GetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsage {
	this := GetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsage{}
	return &this
}

// NewGetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsageWithDefaults instantiates a new GetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsageWithDefaults() *GetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsage {
	this := GetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsage{}
	return &this
}

// GetC returns the C field value if set, zero value otherwise.
func (o *GetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsage) GetC() GetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsageC {
	if o == nil || isNil(o.C) {
		var ret GetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsageC
		return ret
	}
	return *o.C
}

// GetCOk returns a tuple with the C field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsage) GetCOk() (*GetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsageC, bool) {
	if o == nil || isNil(o.C) {
    return nil, false
	}
	return o.C, true
}

// HasC returns a boolean if a field has been set.
func (o *GetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsage) HasC() bool {
	if o != nil && !isNil(o.C) {
		return true
	}

	return false
}

// SetC gets a reference to the given GetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsageC and assigns it to the C field.
func (o *GetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsage) SetC(v GetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsageC) {
	o.C = &v
}

func (o GetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.C) {
		toSerialize["c"] = o.C
	}
	return json.Marshal(toSerialize)
}

type NullableGetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsage struct {
	value *GetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsage
	isSet bool
}

func (v NullableGetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsage) Get() *GetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsage {
	return v.value
}

func (v *NullableGetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsage) Set(val *GetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsage(val *GetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsage) *NullableGetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsage {
	return &NullableGetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsage{value: val, isSet: true}
}

func (v NullableGetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


