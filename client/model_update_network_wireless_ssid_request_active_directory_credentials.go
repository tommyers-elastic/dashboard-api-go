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

// checks if the UpdateNetworkWirelessSsidRequestActiveDirectoryCredentials type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkWirelessSsidRequestActiveDirectoryCredentials{}

// UpdateNetworkWirelessSsidRequestActiveDirectoryCredentials (Optional) The credentials of the user account to be used by the AP to bind to your Active Directory server. The Active Directory account should have permissions on all your Active Directory servers. Only valid if the splashPage is 'Password-protected with Active Directory'.
type UpdateNetworkWirelessSsidRequestActiveDirectoryCredentials struct {
	// The logon name of the Active Directory account.
	LogonName *string `json:"logonName,omitempty"`
	// The password to the Active Directory user account.
	Password *string `json:"password,omitempty"`
}

// NewUpdateNetworkWirelessSsidRequestActiveDirectoryCredentials instantiates a new UpdateNetworkWirelessSsidRequestActiveDirectoryCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessSsidRequestActiveDirectoryCredentials() *UpdateNetworkWirelessSsidRequestActiveDirectoryCredentials {
	this := UpdateNetworkWirelessSsidRequestActiveDirectoryCredentials{}
	return &this
}

// NewUpdateNetworkWirelessSsidRequestActiveDirectoryCredentialsWithDefaults instantiates a new UpdateNetworkWirelessSsidRequestActiveDirectoryCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessSsidRequestActiveDirectoryCredentialsWithDefaults() *UpdateNetworkWirelessSsidRequestActiveDirectoryCredentials {
	this := UpdateNetworkWirelessSsidRequestActiveDirectoryCredentials{}
	return &this
}

// GetLogonName returns the LogonName field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidRequestActiveDirectoryCredentials) GetLogonName() string {
	if o == nil || IsNil(o.LogonName) {
		var ret string
		return ret
	}
	return *o.LogonName
}

// GetLogonNameOk returns a tuple with the LogonName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidRequestActiveDirectoryCredentials) GetLogonNameOk() (*string, bool) {
	if o == nil || IsNil(o.LogonName) {
		return nil, false
	}
	return o.LogonName, true
}

// HasLogonName returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidRequestActiveDirectoryCredentials) HasLogonName() bool {
	if o != nil && !IsNil(o.LogonName) {
		return true
	}

	return false
}

// SetLogonName gets a reference to the given string and assigns it to the LogonName field.
func (o *UpdateNetworkWirelessSsidRequestActiveDirectoryCredentials) SetLogonName(v string) {
	o.LogonName = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidRequestActiveDirectoryCredentials) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidRequestActiveDirectoryCredentials) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidRequestActiveDirectoryCredentials) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *UpdateNetworkWirelessSsidRequestActiveDirectoryCredentials) SetPassword(v string) {
	o.Password = &v
}

func (o UpdateNetworkWirelessSsidRequestActiveDirectoryCredentials) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkWirelessSsidRequestActiveDirectoryCredentials) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LogonName) {
		toSerialize["logonName"] = o.LogonName
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	return toSerialize, nil
}

type NullableUpdateNetworkWirelessSsidRequestActiveDirectoryCredentials struct {
	value *UpdateNetworkWirelessSsidRequestActiveDirectoryCredentials
	isSet bool
}

func (v NullableUpdateNetworkWirelessSsidRequestActiveDirectoryCredentials) Get() *UpdateNetworkWirelessSsidRequestActiveDirectoryCredentials {
	return v.value
}

func (v *NullableUpdateNetworkWirelessSsidRequestActiveDirectoryCredentials) Set(val *UpdateNetworkWirelessSsidRequestActiveDirectoryCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessSsidRequestActiveDirectoryCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessSsidRequestActiveDirectoryCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessSsidRequestActiveDirectoryCredentials(val *UpdateNetworkWirelessSsidRequestActiveDirectoryCredentials) *NullableUpdateNetworkWirelessSsidRequestActiveDirectoryCredentials {
	return &NullableUpdateNetworkWirelessSsidRequestActiveDirectoryCredentials{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessSsidRequestActiveDirectoryCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessSsidRequestActiveDirectoryCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


