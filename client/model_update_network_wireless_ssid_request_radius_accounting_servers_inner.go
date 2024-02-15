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

// checks if the UpdateNetworkWirelessSsidRequestRadiusAccountingServersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkWirelessSsidRequestRadiusAccountingServersInner{}

// UpdateNetworkWirelessSsidRequestRadiusAccountingServersInner struct for UpdateNetworkWirelessSsidRequestRadiusAccountingServersInner
type UpdateNetworkWirelessSsidRequestRadiusAccountingServersInner struct {
	// IP address (or FQDN) to which the APs will send RADIUS accounting messages
	Host string `json:"host"`
	// Port on the RADIUS server that is listening for accounting messages
	Port *int32 `json:"port,omitempty"`
	// Shared key used to authenticate messages between the APs and RADIUS server
	Secret *string `json:"secret,omitempty"`
	// Use RADSEC (TLS over TCP) to connect to this RADIUS accounting server. Requires radiusProxyEnabled.
	RadsecEnabled *bool `json:"radsecEnabled,omitempty"`
	// Certificate used for authorization for the RADSEC Server
	CaCertificate *string `json:"caCertificate,omitempty"`
}

// NewUpdateNetworkWirelessSsidRequestRadiusAccountingServersInner instantiates a new UpdateNetworkWirelessSsidRequestRadiusAccountingServersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessSsidRequestRadiusAccountingServersInner(host string) *UpdateNetworkWirelessSsidRequestRadiusAccountingServersInner {
	this := UpdateNetworkWirelessSsidRequestRadiusAccountingServersInner{}
	this.Host = host
	return &this
}

// NewUpdateNetworkWirelessSsidRequestRadiusAccountingServersInnerWithDefaults instantiates a new UpdateNetworkWirelessSsidRequestRadiusAccountingServersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessSsidRequestRadiusAccountingServersInnerWithDefaults() *UpdateNetworkWirelessSsidRequestRadiusAccountingServersInner {
	this := UpdateNetworkWirelessSsidRequestRadiusAccountingServersInner{}
	return &this
}

// GetHost returns the Host field value
func (o *UpdateNetworkWirelessSsidRequestRadiusAccountingServersInner) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidRequestRadiusAccountingServersInner) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *UpdateNetworkWirelessSsidRequestRadiusAccountingServersInner) SetHost(v string) {
	o.Host = v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidRequestRadiusAccountingServersInner) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidRequestRadiusAccountingServersInner) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidRequestRadiusAccountingServersInner) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *UpdateNetworkWirelessSsidRequestRadiusAccountingServersInner) SetPort(v int32) {
	o.Port = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidRequestRadiusAccountingServersInner) GetSecret() string {
	if o == nil || IsNil(o.Secret) {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidRequestRadiusAccountingServersInner) GetSecretOk() (*string, bool) {
	if o == nil || IsNil(o.Secret) {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidRequestRadiusAccountingServersInner) HasSecret() bool {
	if o != nil && !IsNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *UpdateNetworkWirelessSsidRequestRadiusAccountingServersInner) SetSecret(v string) {
	o.Secret = &v
}

// GetRadsecEnabled returns the RadsecEnabled field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidRequestRadiusAccountingServersInner) GetRadsecEnabled() bool {
	if o == nil || IsNil(o.RadsecEnabled) {
		var ret bool
		return ret
	}
	return *o.RadsecEnabled
}

// GetRadsecEnabledOk returns a tuple with the RadsecEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidRequestRadiusAccountingServersInner) GetRadsecEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.RadsecEnabled) {
		return nil, false
	}
	return o.RadsecEnabled, true
}

// HasRadsecEnabled returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidRequestRadiusAccountingServersInner) HasRadsecEnabled() bool {
	if o != nil && !IsNil(o.RadsecEnabled) {
		return true
	}

	return false
}

// SetRadsecEnabled gets a reference to the given bool and assigns it to the RadsecEnabled field.
func (o *UpdateNetworkWirelessSsidRequestRadiusAccountingServersInner) SetRadsecEnabled(v bool) {
	o.RadsecEnabled = &v
}

// GetCaCertificate returns the CaCertificate field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidRequestRadiusAccountingServersInner) GetCaCertificate() string {
	if o == nil || IsNil(o.CaCertificate) {
		var ret string
		return ret
	}
	return *o.CaCertificate
}

// GetCaCertificateOk returns a tuple with the CaCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidRequestRadiusAccountingServersInner) GetCaCertificateOk() (*string, bool) {
	if o == nil || IsNil(o.CaCertificate) {
		return nil, false
	}
	return o.CaCertificate, true
}

// HasCaCertificate returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidRequestRadiusAccountingServersInner) HasCaCertificate() bool {
	if o != nil && !IsNil(o.CaCertificate) {
		return true
	}

	return false
}

// SetCaCertificate gets a reference to the given string and assigns it to the CaCertificate field.
func (o *UpdateNetworkWirelessSsidRequestRadiusAccountingServersInner) SetCaCertificate(v string) {
	o.CaCertificate = &v
}

func (o UpdateNetworkWirelessSsidRequestRadiusAccountingServersInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkWirelessSsidRequestRadiusAccountingServersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["host"] = o.Host
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.Secret) {
		toSerialize["secret"] = o.Secret
	}
	if !IsNil(o.RadsecEnabled) {
		toSerialize["radsecEnabled"] = o.RadsecEnabled
	}
	if !IsNil(o.CaCertificate) {
		toSerialize["caCertificate"] = o.CaCertificate
	}
	return toSerialize, nil
}

type NullableUpdateNetworkWirelessSsidRequestRadiusAccountingServersInner struct {
	value *UpdateNetworkWirelessSsidRequestRadiusAccountingServersInner
	isSet bool
}

func (v NullableUpdateNetworkWirelessSsidRequestRadiusAccountingServersInner) Get() *UpdateNetworkWirelessSsidRequestRadiusAccountingServersInner {
	return v.value
}

func (v *NullableUpdateNetworkWirelessSsidRequestRadiusAccountingServersInner) Set(val *UpdateNetworkWirelessSsidRequestRadiusAccountingServersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessSsidRequestRadiusAccountingServersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessSsidRequestRadiusAccountingServersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessSsidRequestRadiusAccountingServersInner(val *UpdateNetworkWirelessSsidRequestRadiusAccountingServersInner) *NullableUpdateNetworkWirelessSsidRequestRadiusAccountingServersInner {
	return &NullableUpdateNetworkWirelessSsidRequestRadiusAccountingServersInner{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessSsidRequestRadiusAccountingServersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessSsidRequestRadiusAccountingServersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


