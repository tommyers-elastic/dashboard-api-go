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

// checks if the UpdateDeviceApplianceRadioSettingsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDeviceApplianceRadioSettingsRequest{}

// UpdateDeviceApplianceRadioSettingsRequest struct for UpdateDeviceApplianceRadioSettingsRequest
type UpdateDeviceApplianceRadioSettingsRequest struct {
	// The ID of an RF profile to assign to the device. If the value of this parameter is null, the appropriate basic RF profile (indoor or outdoor) will be assigned to the device. Assigning an RF profile will clear ALL manually configured overrides on the device (channel width, channel, power).
	RfProfileId *string `json:"rfProfileId,omitempty"`
	TwoFourGhzSettings *UpdateDeviceApplianceRadioSettingsRequestTwoFourGhzSettings `json:"twoFourGhzSettings,omitempty"`
	FiveGhzSettings *UpdateDeviceApplianceRadioSettingsRequestFiveGhzSettings `json:"fiveGhzSettings,omitempty"`
}

// NewUpdateDeviceApplianceRadioSettingsRequest instantiates a new UpdateDeviceApplianceRadioSettingsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDeviceApplianceRadioSettingsRequest() *UpdateDeviceApplianceRadioSettingsRequest {
	this := UpdateDeviceApplianceRadioSettingsRequest{}
	return &this
}

// NewUpdateDeviceApplianceRadioSettingsRequestWithDefaults instantiates a new UpdateDeviceApplianceRadioSettingsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDeviceApplianceRadioSettingsRequestWithDefaults() *UpdateDeviceApplianceRadioSettingsRequest {
	this := UpdateDeviceApplianceRadioSettingsRequest{}
	return &this
}

// GetRfProfileId returns the RfProfileId field value if set, zero value otherwise.
func (o *UpdateDeviceApplianceRadioSettingsRequest) GetRfProfileId() string {
	if o == nil || IsNil(o.RfProfileId) {
		var ret string
		return ret
	}
	return *o.RfProfileId
}

// GetRfProfileIdOk returns a tuple with the RfProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceApplianceRadioSettingsRequest) GetRfProfileIdOk() (*string, bool) {
	if o == nil || IsNil(o.RfProfileId) {
		return nil, false
	}
	return o.RfProfileId, true
}

// HasRfProfileId returns a boolean if a field has been set.
func (o *UpdateDeviceApplianceRadioSettingsRequest) HasRfProfileId() bool {
	if o != nil && !IsNil(o.RfProfileId) {
		return true
	}

	return false
}

// SetRfProfileId gets a reference to the given string and assigns it to the RfProfileId field.
func (o *UpdateDeviceApplianceRadioSettingsRequest) SetRfProfileId(v string) {
	o.RfProfileId = &v
}

// GetTwoFourGhzSettings returns the TwoFourGhzSettings field value if set, zero value otherwise.
func (o *UpdateDeviceApplianceRadioSettingsRequest) GetTwoFourGhzSettings() UpdateDeviceApplianceRadioSettingsRequestTwoFourGhzSettings {
	if o == nil || IsNil(o.TwoFourGhzSettings) {
		var ret UpdateDeviceApplianceRadioSettingsRequestTwoFourGhzSettings
		return ret
	}
	return *o.TwoFourGhzSettings
}

// GetTwoFourGhzSettingsOk returns a tuple with the TwoFourGhzSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceApplianceRadioSettingsRequest) GetTwoFourGhzSettingsOk() (*UpdateDeviceApplianceRadioSettingsRequestTwoFourGhzSettings, bool) {
	if o == nil || IsNil(o.TwoFourGhzSettings) {
		return nil, false
	}
	return o.TwoFourGhzSettings, true
}

// HasTwoFourGhzSettings returns a boolean if a field has been set.
func (o *UpdateDeviceApplianceRadioSettingsRequest) HasTwoFourGhzSettings() bool {
	if o != nil && !IsNil(o.TwoFourGhzSettings) {
		return true
	}

	return false
}

// SetTwoFourGhzSettings gets a reference to the given UpdateDeviceApplianceRadioSettingsRequestTwoFourGhzSettings and assigns it to the TwoFourGhzSettings field.
func (o *UpdateDeviceApplianceRadioSettingsRequest) SetTwoFourGhzSettings(v UpdateDeviceApplianceRadioSettingsRequestTwoFourGhzSettings) {
	o.TwoFourGhzSettings = &v
}

// GetFiveGhzSettings returns the FiveGhzSettings field value if set, zero value otherwise.
func (o *UpdateDeviceApplianceRadioSettingsRequest) GetFiveGhzSettings() UpdateDeviceApplianceRadioSettingsRequestFiveGhzSettings {
	if o == nil || IsNil(o.FiveGhzSettings) {
		var ret UpdateDeviceApplianceRadioSettingsRequestFiveGhzSettings
		return ret
	}
	return *o.FiveGhzSettings
}

// GetFiveGhzSettingsOk returns a tuple with the FiveGhzSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceApplianceRadioSettingsRequest) GetFiveGhzSettingsOk() (*UpdateDeviceApplianceRadioSettingsRequestFiveGhzSettings, bool) {
	if o == nil || IsNil(o.FiveGhzSettings) {
		return nil, false
	}
	return o.FiveGhzSettings, true
}

// HasFiveGhzSettings returns a boolean if a field has been set.
func (o *UpdateDeviceApplianceRadioSettingsRequest) HasFiveGhzSettings() bool {
	if o != nil && !IsNil(o.FiveGhzSettings) {
		return true
	}

	return false
}

// SetFiveGhzSettings gets a reference to the given UpdateDeviceApplianceRadioSettingsRequestFiveGhzSettings and assigns it to the FiveGhzSettings field.
func (o *UpdateDeviceApplianceRadioSettingsRequest) SetFiveGhzSettings(v UpdateDeviceApplianceRadioSettingsRequestFiveGhzSettings) {
	o.FiveGhzSettings = &v
}

func (o UpdateDeviceApplianceRadioSettingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateDeviceApplianceRadioSettingsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RfProfileId) {
		toSerialize["rfProfileId"] = o.RfProfileId
	}
	if !IsNil(o.TwoFourGhzSettings) {
		toSerialize["twoFourGhzSettings"] = o.TwoFourGhzSettings
	}
	if !IsNil(o.FiveGhzSettings) {
		toSerialize["fiveGhzSettings"] = o.FiveGhzSettings
	}
	return toSerialize, nil
}

type NullableUpdateDeviceApplianceRadioSettingsRequest struct {
	value *UpdateDeviceApplianceRadioSettingsRequest
	isSet bool
}

func (v NullableUpdateDeviceApplianceRadioSettingsRequest) Get() *UpdateDeviceApplianceRadioSettingsRequest {
	return v.value
}

func (v *NullableUpdateDeviceApplianceRadioSettingsRequest) Set(val *UpdateDeviceApplianceRadioSettingsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDeviceApplianceRadioSettingsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDeviceApplianceRadioSettingsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDeviceApplianceRadioSettingsRequest(val *UpdateDeviceApplianceRadioSettingsRequest) *NullableUpdateDeviceApplianceRadioSettingsRequest {
	return &NullableUpdateDeviceApplianceRadioSettingsRequest{value: val, isSet: true}
}

func (v NullableUpdateDeviceApplianceRadioSettingsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDeviceApplianceRadioSettingsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


