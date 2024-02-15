/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 07 February, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.43.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// checks if the GetNetworkWirelessSsidIdentityPsks200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkWirelessSsidIdentityPsks200ResponseInner{}

// GetNetworkWirelessSsidIdentityPsks200ResponseInner struct for GetNetworkWirelessSsidIdentityPsks200ResponseInner
type GetNetworkWirelessSsidIdentityPsks200ResponseInner struct {
	// The name of the Identity PSK
	Name *string `json:"name,omitempty"`
	// The unique identifier of the Identity PSK
	Id *string `json:"id,omitempty"`
	// The group policy to be applied to clients
	GroupPolicyId *string `json:"groupPolicyId,omitempty"`
	// The passphrase for client authentication
	Passphrase *string `json:"passphrase,omitempty"`
	// The WiFi Personal Network unique identifier
	WifiPersonalNetworkId *string `json:"wifiPersonalNetworkId,omitempty"`
	// The email associated with the System's Manager User
	Email *string `json:"email,omitempty"`
	// Timestamp for when the Identity PSK expires, or 'null' to never expire
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
}

// NewGetNetworkWirelessSsidIdentityPsks200ResponseInner instantiates a new GetNetworkWirelessSsidIdentityPsks200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkWirelessSsidIdentityPsks200ResponseInner() *GetNetworkWirelessSsidIdentityPsks200ResponseInner {
	this := GetNetworkWirelessSsidIdentityPsks200ResponseInner{}
	return &this
}

// NewGetNetworkWirelessSsidIdentityPsks200ResponseInnerWithDefaults instantiates a new GetNetworkWirelessSsidIdentityPsks200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkWirelessSsidIdentityPsks200ResponseInnerWithDefaults() *GetNetworkWirelessSsidIdentityPsks200ResponseInner {
	this := GetNetworkWirelessSsidIdentityPsks200ResponseInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetNetworkWirelessSsidIdentityPsks200ResponseInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSsidIdentityPsks200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetNetworkWirelessSsidIdentityPsks200ResponseInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetNetworkWirelessSsidIdentityPsks200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetNetworkWirelessSsidIdentityPsks200ResponseInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSsidIdentityPsks200ResponseInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetNetworkWirelessSsidIdentityPsks200ResponseInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetNetworkWirelessSsidIdentityPsks200ResponseInner) SetId(v string) {
	o.Id = &v
}

// GetGroupPolicyId returns the GroupPolicyId field value if set, zero value otherwise.
func (o *GetNetworkWirelessSsidIdentityPsks200ResponseInner) GetGroupPolicyId() string {
	if o == nil || IsNil(o.GroupPolicyId) {
		var ret string
		return ret
	}
	return *o.GroupPolicyId
}

// GetGroupPolicyIdOk returns a tuple with the GroupPolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSsidIdentityPsks200ResponseInner) GetGroupPolicyIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupPolicyId) {
		return nil, false
	}
	return o.GroupPolicyId, true
}

// HasGroupPolicyId returns a boolean if a field has been set.
func (o *GetNetworkWirelessSsidIdentityPsks200ResponseInner) HasGroupPolicyId() bool {
	if o != nil && !IsNil(o.GroupPolicyId) {
		return true
	}

	return false
}

// SetGroupPolicyId gets a reference to the given string and assigns it to the GroupPolicyId field.
func (o *GetNetworkWirelessSsidIdentityPsks200ResponseInner) SetGroupPolicyId(v string) {
	o.GroupPolicyId = &v
}

// GetPassphrase returns the Passphrase field value if set, zero value otherwise.
func (o *GetNetworkWirelessSsidIdentityPsks200ResponseInner) GetPassphrase() string {
	if o == nil || IsNil(o.Passphrase) {
		var ret string
		return ret
	}
	return *o.Passphrase
}

// GetPassphraseOk returns a tuple with the Passphrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSsidIdentityPsks200ResponseInner) GetPassphraseOk() (*string, bool) {
	if o == nil || IsNil(o.Passphrase) {
		return nil, false
	}
	return o.Passphrase, true
}

// HasPassphrase returns a boolean if a field has been set.
func (o *GetNetworkWirelessSsidIdentityPsks200ResponseInner) HasPassphrase() bool {
	if o != nil && !IsNil(o.Passphrase) {
		return true
	}

	return false
}

// SetPassphrase gets a reference to the given string and assigns it to the Passphrase field.
func (o *GetNetworkWirelessSsidIdentityPsks200ResponseInner) SetPassphrase(v string) {
	o.Passphrase = &v
}

// GetWifiPersonalNetworkId returns the WifiPersonalNetworkId field value if set, zero value otherwise.
func (o *GetNetworkWirelessSsidIdentityPsks200ResponseInner) GetWifiPersonalNetworkId() string {
	if o == nil || IsNil(o.WifiPersonalNetworkId) {
		var ret string
		return ret
	}
	return *o.WifiPersonalNetworkId
}

// GetWifiPersonalNetworkIdOk returns a tuple with the WifiPersonalNetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSsidIdentityPsks200ResponseInner) GetWifiPersonalNetworkIdOk() (*string, bool) {
	if o == nil || IsNil(o.WifiPersonalNetworkId) {
		return nil, false
	}
	return o.WifiPersonalNetworkId, true
}

// HasWifiPersonalNetworkId returns a boolean if a field has been set.
func (o *GetNetworkWirelessSsidIdentityPsks200ResponseInner) HasWifiPersonalNetworkId() bool {
	if o != nil && !IsNil(o.WifiPersonalNetworkId) {
		return true
	}

	return false
}

// SetWifiPersonalNetworkId gets a reference to the given string and assigns it to the WifiPersonalNetworkId field.
func (o *GetNetworkWirelessSsidIdentityPsks200ResponseInner) SetWifiPersonalNetworkId(v string) {
	o.WifiPersonalNetworkId = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *GetNetworkWirelessSsidIdentityPsks200ResponseInner) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSsidIdentityPsks200ResponseInner) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *GetNetworkWirelessSsidIdentityPsks200ResponseInner) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *GetNetworkWirelessSsidIdentityPsks200ResponseInner) SetEmail(v string) {
	o.Email = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *GetNetworkWirelessSsidIdentityPsks200ResponseInner) GetExpiresAt() time.Time {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSsidIdentityPsks200ResponseInner) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *GetNetworkWirelessSsidIdentityPsks200ResponseInner) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *GetNetworkWirelessSsidIdentityPsks200ResponseInner) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

func (o GetNetworkWirelessSsidIdentityPsks200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkWirelessSsidIdentityPsks200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.GroupPolicyId) {
		toSerialize["groupPolicyId"] = o.GroupPolicyId
	}
	if !IsNil(o.Passphrase) {
		toSerialize["passphrase"] = o.Passphrase
	}
	if !IsNil(o.WifiPersonalNetworkId) {
		toSerialize["wifiPersonalNetworkId"] = o.WifiPersonalNetworkId
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.ExpiresAt) {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	return toSerialize, nil
}

type NullableGetNetworkWirelessSsidIdentityPsks200ResponseInner struct {
	value *GetNetworkWirelessSsidIdentityPsks200ResponseInner
	isSet bool
}

func (v NullableGetNetworkWirelessSsidIdentityPsks200ResponseInner) Get() *GetNetworkWirelessSsidIdentityPsks200ResponseInner {
	return v.value
}

func (v *NullableGetNetworkWirelessSsidIdentityPsks200ResponseInner) Set(val *GetNetworkWirelessSsidIdentityPsks200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkWirelessSsidIdentityPsks200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkWirelessSsidIdentityPsks200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkWirelessSsidIdentityPsks200ResponseInner(val *GetNetworkWirelessSsidIdentityPsks200ResponseInner) *NullableGetNetworkWirelessSsidIdentityPsks200ResponseInner {
	return &NullableGetNetworkWirelessSsidIdentityPsks200ResponseInner{value: val, isSet: true}
}

func (v NullableGetNetworkWirelessSsidIdentityPsks200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkWirelessSsidIdentityPsks200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


