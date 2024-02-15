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

// checks if the UpdateOrganizationCameraOnboardingStatusesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateOrganizationCameraOnboardingStatusesRequest{}

// UpdateOrganizationCameraOnboardingStatusesRequest struct for UpdateOrganizationCameraOnboardingStatusesRequest
type UpdateOrganizationCameraOnboardingStatusesRequest struct {
	// Serial of camera
	Serial *string `json:"serial,omitempty"`
	// Note whether credentials were sent successfully
	WirelessCredentialsSent *bool `json:"wirelessCredentialsSent,omitempty"`
}

// NewUpdateOrganizationCameraOnboardingStatusesRequest instantiates a new UpdateOrganizationCameraOnboardingStatusesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOrganizationCameraOnboardingStatusesRequest() *UpdateOrganizationCameraOnboardingStatusesRequest {
	this := UpdateOrganizationCameraOnboardingStatusesRequest{}
	return &this
}

// NewUpdateOrganizationCameraOnboardingStatusesRequestWithDefaults instantiates a new UpdateOrganizationCameraOnboardingStatusesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOrganizationCameraOnboardingStatusesRequestWithDefaults() *UpdateOrganizationCameraOnboardingStatusesRequest {
	this := UpdateOrganizationCameraOnboardingStatusesRequest{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *UpdateOrganizationCameraOnboardingStatusesRequest) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationCameraOnboardingStatusesRequest) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *UpdateOrganizationCameraOnboardingStatusesRequest) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *UpdateOrganizationCameraOnboardingStatusesRequest) SetSerial(v string) {
	o.Serial = &v
}

// GetWirelessCredentialsSent returns the WirelessCredentialsSent field value if set, zero value otherwise.
func (o *UpdateOrganizationCameraOnboardingStatusesRequest) GetWirelessCredentialsSent() bool {
	if o == nil || IsNil(o.WirelessCredentialsSent) {
		var ret bool
		return ret
	}
	return *o.WirelessCredentialsSent
}

// GetWirelessCredentialsSentOk returns a tuple with the WirelessCredentialsSent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationCameraOnboardingStatusesRequest) GetWirelessCredentialsSentOk() (*bool, bool) {
	if o == nil || IsNil(o.WirelessCredentialsSent) {
		return nil, false
	}
	return o.WirelessCredentialsSent, true
}

// HasWirelessCredentialsSent returns a boolean if a field has been set.
func (o *UpdateOrganizationCameraOnboardingStatusesRequest) HasWirelessCredentialsSent() bool {
	if o != nil && !IsNil(o.WirelessCredentialsSent) {
		return true
	}

	return false
}

// SetWirelessCredentialsSent gets a reference to the given bool and assigns it to the WirelessCredentialsSent field.
func (o *UpdateOrganizationCameraOnboardingStatusesRequest) SetWirelessCredentialsSent(v bool) {
	o.WirelessCredentialsSent = &v
}

func (o UpdateOrganizationCameraOnboardingStatusesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateOrganizationCameraOnboardingStatusesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !IsNil(o.WirelessCredentialsSent) {
		toSerialize["wirelessCredentialsSent"] = o.WirelessCredentialsSent
	}
	return toSerialize, nil
}

type NullableUpdateOrganizationCameraOnboardingStatusesRequest struct {
	value *UpdateOrganizationCameraOnboardingStatusesRequest
	isSet bool
}

func (v NullableUpdateOrganizationCameraOnboardingStatusesRequest) Get() *UpdateOrganizationCameraOnboardingStatusesRequest {
	return v.value
}

func (v *NullableUpdateOrganizationCameraOnboardingStatusesRequest) Set(val *UpdateOrganizationCameraOnboardingStatusesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOrganizationCameraOnboardingStatusesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOrganizationCameraOnboardingStatusesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOrganizationCameraOnboardingStatusesRequest(val *UpdateOrganizationCameraOnboardingStatusesRequest) *NullableUpdateOrganizationCameraOnboardingStatusesRequest {
	return &NullableUpdateOrganizationCameraOnboardingStatusesRequest{value: val, isSet: true}
}

func (v NullableUpdateOrganizationCameraOnboardingStatusesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOrganizationCameraOnboardingStatusesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


