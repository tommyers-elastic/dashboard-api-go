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

// checks if the UpdateNetworkApplianceSecurityIntrusionRequestProtectedNetworks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkApplianceSecurityIntrusionRequestProtectedNetworks{}

// UpdateNetworkApplianceSecurityIntrusionRequestProtectedNetworks Set the included/excluded networks from the intrusion engine (optional - omitting will leave current config unchanged). This is available only in 'passthrough' mode
type UpdateNetworkApplianceSecurityIntrusionRequestProtectedNetworks struct {
	// true/false whether to use special IPv4 addresses: https://tools.ietf.org/html/rfc5735 (required). Default value is true if none currently saved
	UseDefault *bool `json:"useDefault,omitempty"`
	// list of IP addresses or subnets being protected (required if 'useDefault' is false)
	IncludedCidr []string `json:"includedCidr,omitempty"`
	// list of IP addresses or subnets being excluded from protection (required if 'useDefault' is false)
	ExcludedCidr []string `json:"excludedCidr,omitempty"`
}

// NewUpdateNetworkApplianceSecurityIntrusionRequestProtectedNetworks instantiates a new UpdateNetworkApplianceSecurityIntrusionRequestProtectedNetworks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceSecurityIntrusionRequestProtectedNetworks() *UpdateNetworkApplianceSecurityIntrusionRequestProtectedNetworks {
	this := UpdateNetworkApplianceSecurityIntrusionRequestProtectedNetworks{}
	return &this
}

// NewUpdateNetworkApplianceSecurityIntrusionRequestProtectedNetworksWithDefaults instantiates a new UpdateNetworkApplianceSecurityIntrusionRequestProtectedNetworks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceSecurityIntrusionRequestProtectedNetworksWithDefaults() *UpdateNetworkApplianceSecurityIntrusionRequestProtectedNetworks {
	this := UpdateNetworkApplianceSecurityIntrusionRequestProtectedNetworks{}
	return &this
}

// GetUseDefault returns the UseDefault field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceSecurityIntrusionRequestProtectedNetworks) GetUseDefault() bool {
	if o == nil || IsNil(o.UseDefault) {
		var ret bool
		return ret
	}
	return *o.UseDefault
}

// GetUseDefaultOk returns a tuple with the UseDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceSecurityIntrusionRequestProtectedNetworks) GetUseDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.UseDefault) {
		return nil, false
	}
	return o.UseDefault, true
}

// HasUseDefault returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceSecurityIntrusionRequestProtectedNetworks) HasUseDefault() bool {
	if o != nil && !IsNil(o.UseDefault) {
		return true
	}

	return false
}

// SetUseDefault gets a reference to the given bool and assigns it to the UseDefault field.
func (o *UpdateNetworkApplianceSecurityIntrusionRequestProtectedNetworks) SetUseDefault(v bool) {
	o.UseDefault = &v
}

// GetIncludedCidr returns the IncludedCidr field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceSecurityIntrusionRequestProtectedNetworks) GetIncludedCidr() []string {
	if o == nil || IsNil(o.IncludedCidr) {
		var ret []string
		return ret
	}
	return o.IncludedCidr
}

// GetIncludedCidrOk returns a tuple with the IncludedCidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceSecurityIntrusionRequestProtectedNetworks) GetIncludedCidrOk() ([]string, bool) {
	if o == nil || IsNil(o.IncludedCidr) {
		return nil, false
	}
	return o.IncludedCidr, true
}

// HasIncludedCidr returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceSecurityIntrusionRequestProtectedNetworks) HasIncludedCidr() bool {
	if o != nil && !IsNil(o.IncludedCidr) {
		return true
	}

	return false
}

// SetIncludedCidr gets a reference to the given []string and assigns it to the IncludedCidr field.
func (o *UpdateNetworkApplianceSecurityIntrusionRequestProtectedNetworks) SetIncludedCidr(v []string) {
	o.IncludedCidr = v
}

// GetExcludedCidr returns the ExcludedCidr field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceSecurityIntrusionRequestProtectedNetworks) GetExcludedCidr() []string {
	if o == nil || IsNil(o.ExcludedCidr) {
		var ret []string
		return ret
	}
	return o.ExcludedCidr
}

// GetExcludedCidrOk returns a tuple with the ExcludedCidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceSecurityIntrusionRequestProtectedNetworks) GetExcludedCidrOk() ([]string, bool) {
	if o == nil || IsNil(o.ExcludedCidr) {
		return nil, false
	}
	return o.ExcludedCidr, true
}

// HasExcludedCidr returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceSecurityIntrusionRequestProtectedNetworks) HasExcludedCidr() bool {
	if o != nil && !IsNil(o.ExcludedCidr) {
		return true
	}

	return false
}

// SetExcludedCidr gets a reference to the given []string and assigns it to the ExcludedCidr field.
func (o *UpdateNetworkApplianceSecurityIntrusionRequestProtectedNetworks) SetExcludedCidr(v []string) {
	o.ExcludedCidr = v
}

func (o UpdateNetworkApplianceSecurityIntrusionRequestProtectedNetworks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkApplianceSecurityIntrusionRequestProtectedNetworks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UseDefault) {
		toSerialize["useDefault"] = o.UseDefault
	}
	if !IsNil(o.IncludedCidr) {
		toSerialize["includedCidr"] = o.IncludedCidr
	}
	if !IsNil(o.ExcludedCidr) {
		toSerialize["excludedCidr"] = o.ExcludedCidr
	}
	return toSerialize, nil
}

type NullableUpdateNetworkApplianceSecurityIntrusionRequestProtectedNetworks struct {
	value *UpdateNetworkApplianceSecurityIntrusionRequestProtectedNetworks
	isSet bool
}

func (v NullableUpdateNetworkApplianceSecurityIntrusionRequestProtectedNetworks) Get() *UpdateNetworkApplianceSecurityIntrusionRequestProtectedNetworks {
	return v.value
}

func (v *NullableUpdateNetworkApplianceSecurityIntrusionRequestProtectedNetworks) Set(val *UpdateNetworkApplianceSecurityIntrusionRequestProtectedNetworks) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceSecurityIntrusionRequestProtectedNetworks) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceSecurityIntrusionRequestProtectedNetworks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceSecurityIntrusionRequestProtectedNetworks(val *UpdateNetworkApplianceSecurityIntrusionRequestProtectedNetworks) *NullableUpdateNetworkApplianceSecurityIntrusionRequestProtectedNetworks {
	return &NullableUpdateNetworkApplianceSecurityIntrusionRequestProtectedNetworks{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceSecurityIntrusionRequestProtectedNetworks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceSecurityIntrusionRequestProtectedNetworks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


