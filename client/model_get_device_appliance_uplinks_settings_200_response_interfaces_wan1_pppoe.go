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

// GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1Pppoe Configuration options for PPPoE.
type GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1Pppoe struct {
	// Whether PPPoE is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	Authentication *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1PppoeAuthentication `json:"authentication,omitempty"`
}

// NewGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1Pppoe instantiates a new GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1Pppoe object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1Pppoe() *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1Pppoe {
	this := GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1Pppoe{}
	return &this
}

// NewGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1PppoeWithDefaults instantiates a new GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1Pppoe object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1PppoeWithDefaults() *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1Pppoe {
	this := GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1Pppoe{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1Pppoe) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1Pppoe) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1Pppoe) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1Pppoe) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1Pppoe) GetAuthentication() GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1PppoeAuthentication {
	if o == nil || isNil(o.Authentication) {
		var ret GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1PppoeAuthentication
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1Pppoe) GetAuthenticationOk() (*GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1PppoeAuthentication, bool) {
	if o == nil || isNil(o.Authentication) {
    return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1Pppoe) HasAuthentication() bool {
	if o != nil && !isNil(o.Authentication) {
		return true
	}

	return false
}

// SetAuthentication gets a reference to the given GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1PppoeAuthentication and assigns it to the Authentication field.
func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1Pppoe) SetAuthentication(v GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1PppoeAuthentication) {
	o.Authentication = &v
}

func (o GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1Pppoe) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.Authentication) {
		toSerialize["authentication"] = o.Authentication
	}
	return json.Marshal(toSerialize)
}

type NullableGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1Pppoe struct {
	value *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1Pppoe
	isSet bool
}

func (v NullableGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1Pppoe) Get() *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1Pppoe {
	return v.value
}

func (v *NullableGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1Pppoe) Set(val *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1Pppoe) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1Pppoe) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1Pppoe) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1Pppoe(val *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1Pppoe) *NullableGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1Pppoe {
	return &NullableGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1Pppoe{value: val, isSet: true}
}

func (v NullableGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1Pppoe) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1Pppoe) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


