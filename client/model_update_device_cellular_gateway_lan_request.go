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

// checks if the UpdateDeviceCellularGatewayLanRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDeviceCellularGatewayLanRequest{}

// UpdateDeviceCellularGatewayLanRequest struct for UpdateDeviceCellularGatewayLanRequest
type UpdateDeviceCellularGatewayLanRequest struct {
	// list of all reserved IP ranges for a single MG
	ReservedIpRanges []UpdateDeviceCellularGatewayLanRequestReservedIpRangesInner `json:"reservedIpRanges,omitempty"`
	// list of all fixed IP assignments for a single MG
	FixedIpAssignments []UpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner `json:"fixedIpAssignments,omitempty"`
}

// NewUpdateDeviceCellularGatewayLanRequest instantiates a new UpdateDeviceCellularGatewayLanRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDeviceCellularGatewayLanRequest() *UpdateDeviceCellularGatewayLanRequest {
	this := UpdateDeviceCellularGatewayLanRequest{}
	return &this
}

// NewUpdateDeviceCellularGatewayLanRequestWithDefaults instantiates a new UpdateDeviceCellularGatewayLanRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDeviceCellularGatewayLanRequestWithDefaults() *UpdateDeviceCellularGatewayLanRequest {
	this := UpdateDeviceCellularGatewayLanRequest{}
	return &this
}

// GetReservedIpRanges returns the ReservedIpRanges field value if set, zero value otherwise.
func (o *UpdateDeviceCellularGatewayLanRequest) GetReservedIpRanges() []UpdateDeviceCellularGatewayLanRequestReservedIpRangesInner {
	if o == nil || IsNil(o.ReservedIpRanges) {
		var ret []UpdateDeviceCellularGatewayLanRequestReservedIpRangesInner
		return ret
	}
	return o.ReservedIpRanges
}

// GetReservedIpRangesOk returns a tuple with the ReservedIpRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceCellularGatewayLanRequest) GetReservedIpRangesOk() ([]UpdateDeviceCellularGatewayLanRequestReservedIpRangesInner, bool) {
	if o == nil || IsNil(o.ReservedIpRanges) {
		return nil, false
	}
	return o.ReservedIpRanges, true
}

// HasReservedIpRanges returns a boolean if a field has been set.
func (o *UpdateDeviceCellularGatewayLanRequest) HasReservedIpRanges() bool {
	if o != nil && !IsNil(o.ReservedIpRanges) {
		return true
	}

	return false
}

// SetReservedIpRanges gets a reference to the given []UpdateDeviceCellularGatewayLanRequestReservedIpRangesInner and assigns it to the ReservedIpRanges field.
func (o *UpdateDeviceCellularGatewayLanRequest) SetReservedIpRanges(v []UpdateDeviceCellularGatewayLanRequestReservedIpRangesInner) {
	o.ReservedIpRanges = v
}

// GetFixedIpAssignments returns the FixedIpAssignments field value if set, zero value otherwise.
func (o *UpdateDeviceCellularGatewayLanRequest) GetFixedIpAssignments() []UpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner {
	if o == nil || IsNil(o.FixedIpAssignments) {
		var ret []UpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner
		return ret
	}
	return o.FixedIpAssignments
}

// GetFixedIpAssignmentsOk returns a tuple with the FixedIpAssignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceCellularGatewayLanRequest) GetFixedIpAssignmentsOk() ([]UpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner, bool) {
	if o == nil || IsNil(o.FixedIpAssignments) {
		return nil, false
	}
	return o.FixedIpAssignments, true
}

// HasFixedIpAssignments returns a boolean if a field has been set.
func (o *UpdateDeviceCellularGatewayLanRequest) HasFixedIpAssignments() bool {
	if o != nil && !IsNil(o.FixedIpAssignments) {
		return true
	}

	return false
}

// SetFixedIpAssignments gets a reference to the given []UpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner and assigns it to the FixedIpAssignments field.
func (o *UpdateDeviceCellularGatewayLanRequest) SetFixedIpAssignments(v []UpdateDeviceCellularGatewayLanRequestFixedIpAssignmentsInner) {
	o.FixedIpAssignments = v
}

func (o UpdateDeviceCellularGatewayLanRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateDeviceCellularGatewayLanRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReservedIpRanges) {
		toSerialize["reservedIpRanges"] = o.ReservedIpRanges
	}
	if !IsNil(o.FixedIpAssignments) {
		toSerialize["fixedIpAssignments"] = o.FixedIpAssignments
	}
	return toSerialize, nil
}

type NullableUpdateDeviceCellularGatewayLanRequest struct {
	value *UpdateDeviceCellularGatewayLanRequest
	isSet bool
}

func (v NullableUpdateDeviceCellularGatewayLanRequest) Get() *UpdateDeviceCellularGatewayLanRequest {
	return v.value
}

func (v *NullableUpdateDeviceCellularGatewayLanRequest) Set(val *UpdateDeviceCellularGatewayLanRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDeviceCellularGatewayLanRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDeviceCellularGatewayLanRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDeviceCellularGatewayLanRequest(val *UpdateDeviceCellularGatewayLanRequest) *NullableUpdateDeviceCellularGatewayLanRequest {
	return &NullableUpdateDeviceCellularGatewayLanRequest{value: val, isSet: true}
}

func (v NullableUpdateDeviceCellularGatewayLanRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDeviceCellularGatewayLanRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


