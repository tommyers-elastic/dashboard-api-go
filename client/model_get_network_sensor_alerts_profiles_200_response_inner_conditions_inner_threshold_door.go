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

// checks if the GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdDoor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdDoor{}

// GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdDoor Door open threshold. 'open' must be provided and set to true.
type GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdDoor struct {
	// Alerting threshold for a door open event. Must be set to true.
	Open bool `json:"open"`
}

// NewGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdDoor instantiates a new GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdDoor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdDoor(open bool) *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdDoor {
	this := GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdDoor{}
	this.Open = open
	return &this
}

// NewGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdDoorWithDefaults instantiates a new GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdDoor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdDoorWithDefaults() *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdDoor {
	this := GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdDoor{}
	return &this
}

// GetOpen returns the Open field value
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdDoor) GetOpen() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Open
}

// GetOpenOk returns a tuple with the Open field value
// and a boolean to check if the value has been set.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdDoor) GetOpenOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Open, true
}

// SetOpen sets field value
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdDoor) SetOpen(v bool) {
	o.Open = v
}

func (o GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdDoor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdDoor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["open"] = o.Open
	return toSerialize, nil
}

type NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdDoor struct {
	value *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdDoor
	isSet bool
}

func (v NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdDoor) Get() *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdDoor {
	return v.value
}

func (v *NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdDoor) Set(val *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdDoor) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdDoor) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdDoor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdDoor(val *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdDoor) *NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdDoor {
	return &NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdDoor{value: val, isSet: true}
}

func (v NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdDoor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdDoor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


