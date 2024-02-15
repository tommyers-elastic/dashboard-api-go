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

// checks if the GetNetworkApplianceSsids200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkApplianceSsids200ResponseInner{}

// GetNetworkApplianceSsids200ResponseInner struct for GetNetworkApplianceSsids200ResponseInner
type GetNetworkApplianceSsids200ResponseInner struct {
	// The number of the SSID.
	Number *int32 `json:"number,omitempty"`
	// The name of the SSID.
	Name *string `json:"name,omitempty"`
	// Whether or not the SSID is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// The VLAN ID of the VLAN associated to this SSID.
	DefaultVlanId *int32 `json:"defaultVlanId,omitempty"`
	// The association control method for the SSID.
	AuthMode *string `json:"authMode,omitempty"`
	// The RADIUS 802.1x servers to be used for authentication.
	RadiusServers []GetNetworkApplianceSsids200ResponseInnerRadiusServersInner `json:"radiusServers,omitempty"`
	// The psk encryption mode for the SSID.
	EncryptionMode *string `json:"encryptionMode,omitempty"`
	// WPA encryption mode for the SSID.
	WpaEncryptionMode *string `json:"wpaEncryptionMode,omitempty"`
	// Boolean indicating whether the MX should advertise or hide this SSID.
	Visible *bool `json:"visible,omitempty"`
}

// NewGetNetworkApplianceSsids200ResponseInner instantiates a new GetNetworkApplianceSsids200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkApplianceSsids200ResponseInner() *GetNetworkApplianceSsids200ResponseInner {
	this := GetNetworkApplianceSsids200ResponseInner{}
	return &this
}

// NewGetNetworkApplianceSsids200ResponseInnerWithDefaults instantiates a new GetNetworkApplianceSsids200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkApplianceSsids200ResponseInnerWithDefaults() *GetNetworkApplianceSsids200ResponseInner {
	this := GetNetworkApplianceSsids200ResponseInner{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *GetNetworkApplianceSsids200ResponseInner) GetNumber() int32 {
	if o == nil || IsNil(o.Number) {
		var ret int32
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceSsids200ResponseInner) GetNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *GetNetworkApplianceSsids200ResponseInner) HasNumber() bool {
	if o != nil && !IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given int32 and assigns it to the Number field.
func (o *GetNetworkApplianceSsids200ResponseInner) SetNumber(v int32) {
	o.Number = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetNetworkApplianceSsids200ResponseInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceSsids200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetNetworkApplianceSsids200ResponseInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetNetworkApplianceSsids200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *GetNetworkApplianceSsids200ResponseInner) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceSsids200ResponseInner) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *GetNetworkApplianceSsids200ResponseInner) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *GetNetworkApplianceSsids200ResponseInner) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetDefaultVlanId returns the DefaultVlanId field value if set, zero value otherwise.
func (o *GetNetworkApplianceSsids200ResponseInner) GetDefaultVlanId() int32 {
	if o == nil || IsNil(o.DefaultVlanId) {
		var ret int32
		return ret
	}
	return *o.DefaultVlanId
}

// GetDefaultVlanIdOk returns a tuple with the DefaultVlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceSsids200ResponseInner) GetDefaultVlanIdOk() (*int32, bool) {
	if o == nil || IsNil(o.DefaultVlanId) {
		return nil, false
	}
	return o.DefaultVlanId, true
}

// HasDefaultVlanId returns a boolean if a field has been set.
func (o *GetNetworkApplianceSsids200ResponseInner) HasDefaultVlanId() bool {
	if o != nil && !IsNil(o.DefaultVlanId) {
		return true
	}

	return false
}

// SetDefaultVlanId gets a reference to the given int32 and assigns it to the DefaultVlanId field.
func (o *GetNetworkApplianceSsids200ResponseInner) SetDefaultVlanId(v int32) {
	o.DefaultVlanId = &v
}

// GetAuthMode returns the AuthMode field value if set, zero value otherwise.
func (o *GetNetworkApplianceSsids200ResponseInner) GetAuthMode() string {
	if o == nil || IsNil(o.AuthMode) {
		var ret string
		return ret
	}
	return *o.AuthMode
}

// GetAuthModeOk returns a tuple with the AuthMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceSsids200ResponseInner) GetAuthModeOk() (*string, bool) {
	if o == nil || IsNil(o.AuthMode) {
		return nil, false
	}
	return o.AuthMode, true
}

// HasAuthMode returns a boolean if a field has been set.
func (o *GetNetworkApplianceSsids200ResponseInner) HasAuthMode() bool {
	if o != nil && !IsNil(o.AuthMode) {
		return true
	}

	return false
}

// SetAuthMode gets a reference to the given string and assigns it to the AuthMode field.
func (o *GetNetworkApplianceSsids200ResponseInner) SetAuthMode(v string) {
	o.AuthMode = &v
}

// GetRadiusServers returns the RadiusServers field value if set, zero value otherwise.
func (o *GetNetworkApplianceSsids200ResponseInner) GetRadiusServers() []GetNetworkApplianceSsids200ResponseInnerRadiusServersInner {
	if o == nil || IsNil(o.RadiusServers) {
		var ret []GetNetworkApplianceSsids200ResponseInnerRadiusServersInner
		return ret
	}
	return o.RadiusServers
}

// GetRadiusServersOk returns a tuple with the RadiusServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceSsids200ResponseInner) GetRadiusServersOk() ([]GetNetworkApplianceSsids200ResponseInnerRadiusServersInner, bool) {
	if o == nil || IsNil(o.RadiusServers) {
		return nil, false
	}
	return o.RadiusServers, true
}

// HasRadiusServers returns a boolean if a field has been set.
func (o *GetNetworkApplianceSsids200ResponseInner) HasRadiusServers() bool {
	if o != nil && !IsNil(o.RadiusServers) {
		return true
	}

	return false
}

// SetRadiusServers gets a reference to the given []GetNetworkApplianceSsids200ResponseInnerRadiusServersInner and assigns it to the RadiusServers field.
func (o *GetNetworkApplianceSsids200ResponseInner) SetRadiusServers(v []GetNetworkApplianceSsids200ResponseInnerRadiusServersInner) {
	o.RadiusServers = v
}

// GetEncryptionMode returns the EncryptionMode field value if set, zero value otherwise.
func (o *GetNetworkApplianceSsids200ResponseInner) GetEncryptionMode() string {
	if o == nil || IsNil(o.EncryptionMode) {
		var ret string
		return ret
	}
	return *o.EncryptionMode
}

// GetEncryptionModeOk returns a tuple with the EncryptionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceSsids200ResponseInner) GetEncryptionModeOk() (*string, bool) {
	if o == nil || IsNil(o.EncryptionMode) {
		return nil, false
	}
	return o.EncryptionMode, true
}

// HasEncryptionMode returns a boolean if a field has been set.
func (o *GetNetworkApplianceSsids200ResponseInner) HasEncryptionMode() bool {
	if o != nil && !IsNil(o.EncryptionMode) {
		return true
	}

	return false
}

// SetEncryptionMode gets a reference to the given string and assigns it to the EncryptionMode field.
func (o *GetNetworkApplianceSsids200ResponseInner) SetEncryptionMode(v string) {
	o.EncryptionMode = &v
}

// GetWpaEncryptionMode returns the WpaEncryptionMode field value if set, zero value otherwise.
func (o *GetNetworkApplianceSsids200ResponseInner) GetWpaEncryptionMode() string {
	if o == nil || IsNil(o.WpaEncryptionMode) {
		var ret string
		return ret
	}
	return *o.WpaEncryptionMode
}

// GetWpaEncryptionModeOk returns a tuple with the WpaEncryptionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceSsids200ResponseInner) GetWpaEncryptionModeOk() (*string, bool) {
	if o == nil || IsNil(o.WpaEncryptionMode) {
		return nil, false
	}
	return o.WpaEncryptionMode, true
}

// HasWpaEncryptionMode returns a boolean if a field has been set.
func (o *GetNetworkApplianceSsids200ResponseInner) HasWpaEncryptionMode() bool {
	if o != nil && !IsNil(o.WpaEncryptionMode) {
		return true
	}

	return false
}

// SetWpaEncryptionMode gets a reference to the given string and assigns it to the WpaEncryptionMode field.
func (o *GetNetworkApplianceSsids200ResponseInner) SetWpaEncryptionMode(v string) {
	o.WpaEncryptionMode = &v
}

// GetVisible returns the Visible field value if set, zero value otherwise.
func (o *GetNetworkApplianceSsids200ResponseInner) GetVisible() bool {
	if o == nil || IsNil(o.Visible) {
		var ret bool
		return ret
	}
	return *o.Visible
}

// GetVisibleOk returns a tuple with the Visible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceSsids200ResponseInner) GetVisibleOk() (*bool, bool) {
	if o == nil || IsNil(o.Visible) {
		return nil, false
	}
	return o.Visible, true
}

// HasVisible returns a boolean if a field has been set.
func (o *GetNetworkApplianceSsids200ResponseInner) HasVisible() bool {
	if o != nil && !IsNil(o.Visible) {
		return true
	}

	return false
}

// SetVisible gets a reference to the given bool and assigns it to the Visible field.
func (o *GetNetworkApplianceSsids200ResponseInner) SetVisible(v bool) {
	o.Visible = &v
}

func (o GetNetworkApplianceSsids200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkApplianceSsids200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Number) {
		toSerialize["number"] = o.Number
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.DefaultVlanId) {
		toSerialize["defaultVlanId"] = o.DefaultVlanId
	}
	if !IsNil(o.AuthMode) {
		toSerialize["authMode"] = o.AuthMode
	}
	if !IsNil(o.RadiusServers) {
		toSerialize["radiusServers"] = o.RadiusServers
	}
	if !IsNil(o.EncryptionMode) {
		toSerialize["encryptionMode"] = o.EncryptionMode
	}
	if !IsNil(o.WpaEncryptionMode) {
		toSerialize["wpaEncryptionMode"] = o.WpaEncryptionMode
	}
	if !IsNil(o.Visible) {
		toSerialize["visible"] = o.Visible
	}
	return toSerialize, nil
}

type NullableGetNetworkApplianceSsids200ResponseInner struct {
	value *GetNetworkApplianceSsids200ResponseInner
	isSet bool
}

func (v NullableGetNetworkApplianceSsids200ResponseInner) Get() *GetNetworkApplianceSsids200ResponseInner {
	return v.value
}

func (v *NullableGetNetworkApplianceSsids200ResponseInner) Set(val *GetNetworkApplianceSsids200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkApplianceSsids200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkApplianceSsids200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkApplianceSsids200ResponseInner(val *GetNetworkApplianceSsids200ResponseInner) *NullableGetNetworkApplianceSsids200ResponseInner {
	return &NullableGetNetworkApplianceSsids200ResponseInner{value: val, isSet: true}
}

func (v NullableGetNetworkApplianceSsids200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkApplianceSsids200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


