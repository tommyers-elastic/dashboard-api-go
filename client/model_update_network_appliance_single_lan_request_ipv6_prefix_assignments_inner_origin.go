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

// checks if the UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInnerOrigin type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInnerOrigin{}

// UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInnerOrigin The origin of the prefix
type UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInnerOrigin struct {
	// Type of the origin
	Type string `json:"type"`
	// Interfaces associated with the prefix
	Interfaces []string `json:"interfaces,omitempty"`
}

// NewUpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInnerOrigin instantiates a new UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInnerOrigin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInnerOrigin(type_ string) *UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInnerOrigin {
	this := UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInnerOrigin{}
	this.Type = type_
	return &this
}

// NewUpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInnerOriginWithDefaults instantiates a new UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInnerOrigin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInnerOriginWithDefaults() *UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInnerOrigin {
	this := UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInnerOrigin{}
	return &this
}

// GetType returns the Type field value
func (o *UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInnerOrigin) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInnerOrigin) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInnerOrigin) SetType(v string) {
	o.Type = v
}

// GetInterfaces returns the Interfaces field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInnerOrigin) GetInterfaces() []string {
	if o == nil || IsNil(o.Interfaces) {
		var ret []string
		return ret
	}
	return o.Interfaces
}

// GetInterfacesOk returns a tuple with the Interfaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInnerOrigin) GetInterfacesOk() ([]string, bool) {
	if o == nil || IsNil(o.Interfaces) {
		return nil, false
	}
	return o.Interfaces, true
}

// HasInterfaces returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInnerOrigin) HasInterfaces() bool {
	if o != nil && !IsNil(o.Interfaces) {
		return true
	}

	return false
}

// SetInterfaces gets a reference to the given []string and assigns it to the Interfaces field.
func (o *UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInnerOrigin) SetInterfaces(v []string) {
	o.Interfaces = v
}

func (o UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInnerOrigin) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInnerOrigin) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Interfaces) {
		toSerialize["interfaces"] = o.Interfaces
	}
	return toSerialize, nil
}

type NullableUpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInnerOrigin struct {
	value *UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInnerOrigin
	isSet bool
}

func (v NullableUpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInnerOrigin) Get() *UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInnerOrigin {
	return v.value
}

func (v *NullableUpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInnerOrigin) Set(val *UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInnerOrigin) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInnerOrigin) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInnerOrigin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInnerOrigin(val *UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInnerOrigin) *NullableUpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInnerOrigin {
	return &NullableUpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInnerOrigin{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInnerOrigin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInnerOrigin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


