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

// GetOrganizationSensorReadingsHistory200ResponseInnerButton Reading for the 'button' metric. This will only be present if the 'metric' property equals 'button'.
type GetOrganizationSensorReadingsHistory200ResponseInnerButton struct {
	// Type of button press that occurred.
	PressType *string `json:"pressType,omitempty"`
}

// NewGetOrganizationSensorReadingsHistory200ResponseInnerButton instantiates a new GetOrganizationSensorReadingsHistory200ResponseInnerButton object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationSensorReadingsHistory200ResponseInnerButton() *GetOrganizationSensorReadingsHistory200ResponseInnerButton {
	this := GetOrganizationSensorReadingsHistory200ResponseInnerButton{}
	return &this
}

// NewGetOrganizationSensorReadingsHistory200ResponseInnerButtonWithDefaults instantiates a new GetOrganizationSensorReadingsHistory200ResponseInnerButton object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationSensorReadingsHistory200ResponseInnerButtonWithDefaults() *GetOrganizationSensorReadingsHistory200ResponseInnerButton {
	this := GetOrganizationSensorReadingsHistory200ResponseInnerButton{}
	return &this
}

// GetPressType returns the PressType field value if set, zero value otherwise.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerButton) GetPressType() string {
	if o == nil || isNil(o.PressType) {
		var ret string
		return ret
	}
	return *o.PressType
}

// GetPressTypeOk returns a tuple with the PressType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerButton) GetPressTypeOk() (*string, bool) {
	if o == nil || isNil(o.PressType) {
    return nil, false
	}
	return o.PressType, true
}

// HasPressType returns a boolean if a field has been set.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerButton) HasPressType() bool {
	if o != nil && !isNil(o.PressType) {
		return true
	}

	return false
}

// SetPressType gets a reference to the given string and assigns it to the PressType field.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerButton) SetPressType(v string) {
	o.PressType = &v
}

func (o GetOrganizationSensorReadingsHistory200ResponseInnerButton) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PressType) {
		toSerialize["pressType"] = o.PressType
	}
	return json.Marshal(toSerialize)
}

type NullableGetOrganizationSensorReadingsHistory200ResponseInnerButton struct {
	value *GetOrganizationSensorReadingsHistory200ResponseInnerButton
	isSet bool
}

func (v NullableGetOrganizationSensorReadingsHistory200ResponseInnerButton) Get() *GetOrganizationSensorReadingsHistory200ResponseInnerButton {
	return v.value
}

func (v *NullableGetOrganizationSensorReadingsHistory200ResponseInnerButton) Set(val *GetOrganizationSensorReadingsHistory200ResponseInnerButton) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationSensorReadingsHistory200ResponseInnerButton) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationSensorReadingsHistory200ResponseInnerButton) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationSensorReadingsHistory200ResponseInnerButton(val *GetOrganizationSensorReadingsHistory200ResponseInnerButton) *NullableGetOrganizationSensorReadingsHistory200ResponseInnerButton {
	return &NullableGetOrganizationSensorReadingsHistory200ResponseInnerButton{value: val, isSet: true}
}

func (v NullableGetOrganizationSensorReadingsHistory200ResponseInnerButton) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationSensorReadingsHistory200ResponseInnerButton) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


