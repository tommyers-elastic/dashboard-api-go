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

// checks if the GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdVoltage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdVoltage{}

// GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdVoltage Voltage threshold. 'level' must be provided.
type GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdVoltage struct {
	// Alerting threshold in volts. Must be between 0 and 250.
	Level float32 `json:"level"`
}

// NewGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdVoltage instantiates a new GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdVoltage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdVoltage(level float32) *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdVoltage {
	this := GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdVoltage{}
	this.Level = level
	return &this
}

// NewGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdVoltageWithDefaults instantiates a new GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdVoltage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdVoltageWithDefaults() *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdVoltage {
	this := GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdVoltage{}
	return &this
}

// GetLevel returns the Level field value
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdVoltage) GetLevel() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Level
}

// GetLevelOk returns a tuple with the Level field value
// and a boolean to check if the value has been set.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdVoltage) GetLevelOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Level, true
}

// SetLevel sets field value
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdVoltage) SetLevel(v float32) {
	o.Level = v
}

func (o GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdVoltage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdVoltage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["level"] = o.Level
	return toSerialize, nil
}

type NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdVoltage struct {
	value *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdVoltage
	isSet bool
}

func (v NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdVoltage) Get() *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdVoltage {
	return v.value
}

func (v *NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdVoltage) Set(val *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdVoltage) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdVoltage) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdVoltage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdVoltage(val *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdVoltage) *NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdVoltage {
	return &NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdVoltage{value: val, isSet: true}
}

func (v NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdVoltage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdVoltage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


