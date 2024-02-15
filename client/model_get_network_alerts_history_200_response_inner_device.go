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

// checks if the GetNetworkAlertsHistory200ResponseInnerDevice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkAlertsHistory200ResponseInnerDevice{}

// GetNetworkAlertsHistory200ResponseInnerDevice info related to the device that caused the alert
type GetNetworkAlertsHistory200ResponseInnerDevice struct {
	// device serial
	Serial *string `json:"serial,omitempty"`
}

// NewGetNetworkAlertsHistory200ResponseInnerDevice instantiates a new GetNetworkAlertsHistory200ResponseInnerDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkAlertsHistory200ResponseInnerDevice() *GetNetworkAlertsHistory200ResponseInnerDevice {
	this := GetNetworkAlertsHistory200ResponseInnerDevice{}
	return &this
}

// NewGetNetworkAlertsHistory200ResponseInnerDeviceWithDefaults instantiates a new GetNetworkAlertsHistory200ResponseInnerDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkAlertsHistory200ResponseInnerDeviceWithDefaults() *GetNetworkAlertsHistory200ResponseInnerDevice {
	this := GetNetworkAlertsHistory200ResponseInnerDevice{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *GetNetworkAlertsHistory200ResponseInnerDevice) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkAlertsHistory200ResponseInnerDevice) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *GetNetworkAlertsHistory200ResponseInnerDevice) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *GetNetworkAlertsHistory200ResponseInnerDevice) SetSerial(v string) {
	o.Serial = &v
}

func (o GetNetworkAlertsHistory200ResponseInnerDevice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkAlertsHistory200ResponseInnerDevice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	return toSerialize, nil
}

type NullableGetNetworkAlertsHistory200ResponseInnerDevice struct {
	value *GetNetworkAlertsHistory200ResponseInnerDevice
	isSet bool
}

func (v NullableGetNetworkAlertsHistory200ResponseInnerDevice) Get() *GetNetworkAlertsHistory200ResponseInnerDevice {
	return v.value
}

func (v *NullableGetNetworkAlertsHistory200ResponseInnerDevice) Set(val *GetNetworkAlertsHistory200ResponseInnerDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkAlertsHistory200ResponseInnerDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkAlertsHistory200ResponseInnerDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkAlertsHistory200ResponseInnerDevice(val *GetNetworkAlertsHistory200ResponseInnerDevice) *NullableGetNetworkAlertsHistory200ResponseInnerDevice {
	return &NullableGetNetworkAlertsHistory200ResponseInnerDevice{value: val, isSet: true}
}

func (v NullableGetNetworkAlertsHistory200ResponseInnerDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkAlertsHistory200ResponseInnerDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


