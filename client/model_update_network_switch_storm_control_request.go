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

// checks if the UpdateNetworkSwitchStormControlRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkSwitchStormControlRequest{}

// UpdateNetworkSwitchStormControlRequest struct for UpdateNetworkSwitchStormControlRequest
type UpdateNetworkSwitchStormControlRequest struct {
	// Percentage (1 to 99) of total available port bandwidth for broadcast traffic type. Default value 100 percent rate is to clear the configuration.
	BroadcastThreshold *int32 `json:"broadcastThreshold,omitempty"`
	// Percentage (1 to 99) of total available port bandwidth for multicast traffic type. Default value 100 percent rate is to clear the configuration.
	MulticastThreshold *int32 `json:"multicastThreshold,omitempty"`
	// Percentage (1 to 99) of total available port bandwidth for unknown unicast (dlf-destination lookup failure) traffic type. Default value 100 percent rate is to clear the configuration.
	UnknownUnicastThreshold *int32 `json:"unknownUnicastThreshold,omitempty"`
}

// NewUpdateNetworkSwitchStormControlRequest instantiates a new UpdateNetworkSwitchStormControlRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkSwitchStormControlRequest() *UpdateNetworkSwitchStormControlRequest {
	this := UpdateNetworkSwitchStormControlRequest{}
	return &this
}

// NewUpdateNetworkSwitchStormControlRequestWithDefaults instantiates a new UpdateNetworkSwitchStormControlRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkSwitchStormControlRequestWithDefaults() *UpdateNetworkSwitchStormControlRequest {
	this := UpdateNetworkSwitchStormControlRequest{}
	return &this
}

// GetBroadcastThreshold returns the BroadcastThreshold field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchStormControlRequest) GetBroadcastThreshold() int32 {
	if o == nil || IsNil(o.BroadcastThreshold) {
		var ret int32
		return ret
	}
	return *o.BroadcastThreshold
}

// GetBroadcastThresholdOk returns a tuple with the BroadcastThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchStormControlRequest) GetBroadcastThresholdOk() (*int32, bool) {
	if o == nil || IsNil(o.BroadcastThreshold) {
		return nil, false
	}
	return o.BroadcastThreshold, true
}

// HasBroadcastThreshold returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchStormControlRequest) HasBroadcastThreshold() bool {
	if o != nil && !IsNil(o.BroadcastThreshold) {
		return true
	}

	return false
}

// SetBroadcastThreshold gets a reference to the given int32 and assigns it to the BroadcastThreshold field.
func (o *UpdateNetworkSwitchStormControlRequest) SetBroadcastThreshold(v int32) {
	o.BroadcastThreshold = &v
}

// GetMulticastThreshold returns the MulticastThreshold field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchStormControlRequest) GetMulticastThreshold() int32 {
	if o == nil || IsNil(o.MulticastThreshold) {
		var ret int32
		return ret
	}
	return *o.MulticastThreshold
}

// GetMulticastThresholdOk returns a tuple with the MulticastThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchStormControlRequest) GetMulticastThresholdOk() (*int32, bool) {
	if o == nil || IsNil(o.MulticastThreshold) {
		return nil, false
	}
	return o.MulticastThreshold, true
}

// HasMulticastThreshold returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchStormControlRequest) HasMulticastThreshold() bool {
	if o != nil && !IsNil(o.MulticastThreshold) {
		return true
	}

	return false
}

// SetMulticastThreshold gets a reference to the given int32 and assigns it to the MulticastThreshold field.
func (o *UpdateNetworkSwitchStormControlRequest) SetMulticastThreshold(v int32) {
	o.MulticastThreshold = &v
}

// GetUnknownUnicastThreshold returns the UnknownUnicastThreshold field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchStormControlRequest) GetUnknownUnicastThreshold() int32 {
	if o == nil || IsNil(o.UnknownUnicastThreshold) {
		var ret int32
		return ret
	}
	return *o.UnknownUnicastThreshold
}

// GetUnknownUnicastThresholdOk returns a tuple with the UnknownUnicastThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchStormControlRequest) GetUnknownUnicastThresholdOk() (*int32, bool) {
	if o == nil || IsNil(o.UnknownUnicastThreshold) {
		return nil, false
	}
	return o.UnknownUnicastThreshold, true
}

// HasUnknownUnicastThreshold returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchStormControlRequest) HasUnknownUnicastThreshold() bool {
	if o != nil && !IsNil(o.UnknownUnicastThreshold) {
		return true
	}

	return false
}

// SetUnknownUnicastThreshold gets a reference to the given int32 and assigns it to the UnknownUnicastThreshold field.
func (o *UpdateNetworkSwitchStormControlRequest) SetUnknownUnicastThreshold(v int32) {
	o.UnknownUnicastThreshold = &v
}

func (o UpdateNetworkSwitchStormControlRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkSwitchStormControlRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BroadcastThreshold) {
		toSerialize["broadcastThreshold"] = o.BroadcastThreshold
	}
	if !IsNil(o.MulticastThreshold) {
		toSerialize["multicastThreshold"] = o.MulticastThreshold
	}
	if !IsNil(o.UnknownUnicastThreshold) {
		toSerialize["unknownUnicastThreshold"] = o.UnknownUnicastThreshold
	}
	return toSerialize, nil
}

type NullableUpdateNetworkSwitchStormControlRequest struct {
	value *UpdateNetworkSwitchStormControlRequest
	isSet bool
}

func (v NullableUpdateNetworkSwitchStormControlRequest) Get() *UpdateNetworkSwitchStormControlRequest {
	return v.value
}

func (v *NullableUpdateNetworkSwitchStormControlRequest) Set(val *UpdateNetworkSwitchStormControlRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkSwitchStormControlRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkSwitchStormControlRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkSwitchStormControlRequest(val *UpdateNetworkSwitchStormControlRequest) *NullableUpdateNetworkSwitchStormControlRequest {
	return &NullableUpdateNetworkSwitchStormControlRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkSwitchStormControlRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkSwitchStormControlRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


