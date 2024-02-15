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

// checks if the GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPowerFactor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPowerFactor{}

// GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPowerFactor Power factor threshold. 'percentage' must be provided.
type GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPowerFactor struct {
	// Alerting threshold as the ratio of active power to apparent power. Must be between 0 and 100.
	Percentage int32 `json:"percentage"`
}

// NewGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPowerFactor instantiates a new GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPowerFactor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPowerFactor(percentage int32) *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPowerFactor {
	this := GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPowerFactor{}
	this.Percentage = percentage
	return &this
}

// NewGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPowerFactorWithDefaults instantiates a new GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPowerFactor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPowerFactorWithDefaults() *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPowerFactor {
	this := GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPowerFactor{}
	return &this
}

// GetPercentage returns the Percentage field value
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPowerFactor) GetPercentage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Percentage
}

// GetPercentageOk returns a tuple with the Percentage field value
// and a boolean to check if the value has been set.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPowerFactor) GetPercentageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Percentage, true
}

// SetPercentage sets field value
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPowerFactor) SetPercentage(v int32) {
	o.Percentage = v
}

func (o GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPowerFactor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPowerFactor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["percentage"] = o.Percentage
	return toSerialize, nil
}

type NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPowerFactor struct {
	value *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPowerFactor
	isSet bool
}

func (v NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPowerFactor) Get() *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPowerFactor {
	return v.value
}

func (v *NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPowerFactor) Set(val *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPowerFactor) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPowerFactor) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPowerFactor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPowerFactor(val *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPowerFactor) *NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPowerFactor {
	return &NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPowerFactor{value: val, isSet: true}
}

func (v NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPowerFactor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdPowerFactor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


