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

// checks if the GetNetworkFirmwareUpgradesStagedStages200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkFirmwareUpgradesStagedStages200ResponseInner{}

// GetNetworkFirmwareUpgradesStagedStages200ResponseInner struct for GetNetworkFirmwareUpgradesStagedStages200ResponseInner
type GetNetworkFirmwareUpgradesStagedStages200ResponseInner struct {
	Group *GetNetworkFirmwareUpgradesStagedStages200ResponseInnerGroup `json:"group,omitempty"`
}

// NewGetNetworkFirmwareUpgradesStagedStages200ResponseInner instantiates a new GetNetworkFirmwareUpgradesStagedStages200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkFirmwareUpgradesStagedStages200ResponseInner() *GetNetworkFirmwareUpgradesStagedStages200ResponseInner {
	this := GetNetworkFirmwareUpgradesStagedStages200ResponseInner{}
	return &this
}

// NewGetNetworkFirmwareUpgradesStagedStages200ResponseInnerWithDefaults instantiates a new GetNetworkFirmwareUpgradesStagedStages200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkFirmwareUpgradesStagedStages200ResponseInnerWithDefaults() *GetNetworkFirmwareUpgradesStagedStages200ResponseInner {
	this := GetNetworkFirmwareUpgradesStagedStages200ResponseInner{}
	return &this
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *GetNetworkFirmwareUpgradesStagedStages200ResponseInner) GetGroup() GetNetworkFirmwareUpgradesStagedStages200ResponseInnerGroup {
	if o == nil || IsNil(o.Group) {
		var ret GetNetworkFirmwareUpgradesStagedStages200ResponseInnerGroup
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFirmwareUpgradesStagedStages200ResponseInner) GetGroupOk() (*GetNetworkFirmwareUpgradesStagedStages200ResponseInnerGroup, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *GetNetworkFirmwareUpgradesStagedStages200ResponseInner) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given GetNetworkFirmwareUpgradesStagedStages200ResponseInnerGroup and assigns it to the Group field.
func (o *GetNetworkFirmwareUpgradesStagedStages200ResponseInner) SetGroup(v GetNetworkFirmwareUpgradesStagedStages200ResponseInnerGroup) {
	o.Group = &v
}

func (o GetNetworkFirmwareUpgradesStagedStages200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkFirmwareUpgradesStagedStages200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	return toSerialize, nil
}

type NullableGetNetworkFirmwareUpgradesStagedStages200ResponseInner struct {
	value *GetNetworkFirmwareUpgradesStagedStages200ResponseInner
	isSet bool
}

func (v NullableGetNetworkFirmwareUpgradesStagedStages200ResponseInner) Get() *GetNetworkFirmwareUpgradesStagedStages200ResponseInner {
	return v.value
}

func (v *NullableGetNetworkFirmwareUpgradesStagedStages200ResponseInner) Set(val *GetNetworkFirmwareUpgradesStagedStages200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkFirmwareUpgradesStagedStages200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkFirmwareUpgradesStagedStages200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkFirmwareUpgradesStagedStages200ResponseInner(val *GetNetworkFirmwareUpgradesStagedStages200ResponseInner) *NullableGetNetworkFirmwareUpgradesStagedStages200ResponseInner {
	return &NullableGetNetworkFirmwareUpgradesStagedStages200ResponseInner{value: val, isSet: true}
}

func (v NullableGetNetworkFirmwareUpgradesStagedStages200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkFirmwareUpgradesStagedStages200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


