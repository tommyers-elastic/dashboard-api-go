/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 07 June, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.34.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// GetNetworkSwitchDhcpV4ServersSeen200ResponseInner struct for GetNetworkSwitchDhcpV4ServersSeen200ResponseInner
type GetNetworkSwitchDhcpV4ServersSeen200ResponseInner struct {
	// Mac address of the server.
	Mac *string `json:"mac,omitempty"`
	// Vlan id of the server.
	Vlan *int32 `json:"vlan,omitempty"`
	// Client id of the server if available.
	ClientId *string `json:"clientId,omitempty"`
	// Whether the server is allowed or blocked. Always true for configured servers.
	IsAllowed *bool `json:"isAllowed,omitempty"`
	// Last time the server was seen.
	LastSeenAt *time.Time `json:"lastSeenAt,omitempty"`
	// Devices that saw the server.
	SeenBy []GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerSeenByInner `json:"seenBy,omitempty"`
	// server type. Can be a 'device', 'stack', or 'discovered' (i.e client).
	Type *string `json:"type,omitempty"`
	Device *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDevice `json:"device,omitempty"`
	Ipv4 *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerIpv4 `json:"ipv4,omitempty"`
	// Whether the server is configured.
	IsConfigured *bool `json:"isConfigured,omitempty"`
	LastAck *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastAck `json:"lastAck,omitempty"`
	LastPacket *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket `json:"lastPacket,omitempty"`
}

// NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInner instantiates a new GetNetworkSwitchDhcpV4ServersSeen200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInner() *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner {
	this := GetNetworkSwitchDhcpV4ServersSeen200ResponseInner{}
	return &this
}

// NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerWithDefaults instantiates a new GetNetworkSwitchDhcpV4ServersSeen200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerWithDefaults() *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner {
	this := GetNetworkSwitchDhcpV4ServersSeen200ResponseInner{}
	return &this
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) GetMac() string {
	if o == nil || isNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) GetMacOk() (*string, bool) {
	if o == nil || isNil(o.Mac) {
    return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) HasMac() bool {
	if o != nil && !isNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) SetMac(v string) {
	o.Mac = &v
}

// GetVlan returns the Vlan field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) GetVlan() int32 {
	if o == nil || isNil(o.Vlan) {
		var ret int32
		return ret
	}
	return *o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) GetVlanOk() (*int32, bool) {
	if o == nil || isNil(o.Vlan) {
    return nil, false
	}
	return o.Vlan, true
}

// HasVlan returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) HasVlan() bool {
	if o != nil && !isNil(o.Vlan) {
		return true
	}

	return false
}

// SetVlan gets a reference to the given int32 and assigns it to the Vlan field.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) SetVlan(v int32) {
	o.Vlan = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) GetClientId() string {
	if o == nil || isNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) GetClientIdOk() (*string, bool) {
	if o == nil || isNil(o.ClientId) {
    return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) HasClientId() bool {
	if o != nil && !isNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) SetClientId(v string) {
	o.ClientId = &v
}

// GetIsAllowed returns the IsAllowed field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) GetIsAllowed() bool {
	if o == nil || isNil(o.IsAllowed) {
		var ret bool
		return ret
	}
	return *o.IsAllowed
}

// GetIsAllowedOk returns a tuple with the IsAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) GetIsAllowedOk() (*bool, bool) {
	if o == nil || isNil(o.IsAllowed) {
    return nil, false
	}
	return o.IsAllowed, true
}

// HasIsAllowed returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) HasIsAllowed() bool {
	if o != nil && !isNil(o.IsAllowed) {
		return true
	}

	return false
}

// SetIsAllowed gets a reference to the given bool and assigns it to the IsAllowed field.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) SetIsAllowed(v bool) {
	o.IsAllowed = &v
}

// GetLastSeenAt returns the LastSeenAt field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) GetLastSeenAt() time.Time {
	if o == nil || isNil(o.LastSeenAt) {
		var ret time.Time
		return ret
	}
	return *o.LastSeenAt
}

// GetLastSeenAtOk returns a tuple with the LastSeenAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) GetLastSeenAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.LastSeenAt) {
    return nil, false
	}
	return o.LastSeenAt, true
}

// HasLastSeenAt returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) HasLastSeenAt() bool {
	if o != nil && !isNil(o.LastSeenAt) {
		return true
	}

	return false
}

// SetLastSeenAt gets a reference to the given time.Time and assigns it to the LastSeenAt field.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) SetLastSeenAt(v time.Time) {
	o.LastSeenAt = &v
}

// GetSeenBy returns the SeenBy field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) GetSeenBy() []GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerSeenByInner {
	if o == nil || isNil(o.SeenBy) {
		var ret []GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerSeenByInner
		return ret
	}
	return o.SeenBy
}

// GetSeenByOk returns a tuple with the SeenBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) GetSeenByOk() ([]GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerSeenByInner, bool) {
	if o == nil || isNil(o.SeenBy) {
    return nil, false
	}
	return o.SeenBy, true
}

// HasSeenBy returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) HasSeenBy() bool {
	if o != nil && !isNil(o.SeenBy) {
		return true
	}

	return false
}

// SetSeenBy gets a reference to the given []GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerSeenByInner and assigns it to the SeenBy field.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) SetSeenBy(v []GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerSeenByInner) {
	o.SeenBy = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) SetType(v string) {
	o.Type = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) GetDevice() GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDevice {
	if o == nil || isNil(o.Device) {
		var ret GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDevice
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) GetDeviceOk() (*GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDevice, bool) {
	if o == nil || isNil(o.Device) {
    return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) HasDevice() bool {
	if o != nil && !isNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDevice and assigns it to the Device field.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) SetDevice(v GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDevice) {
	o.Device = &v
}

// GetIpv4 returns the Ipv4 field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) GetIpv4() GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerIpv4 {
	if o == nil || isNil(o.Ipv4) {
		var ret GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerIpv4
		return ret
	}
	return *o.Ipv4
}

// GetIpv4Ok returns a tuple with the Ipv4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) GetIpv4Ok() (*GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerIpv4, bool) {
	if o == nil || isNil(o.Ipv4) {
    return nil, false
	}
	return o.Ipv4, true
}

// HasIpv4 returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) HasIpv4() bool {
	if o != nil && !isNil(o.Ipv4) {
		return true
	}

	return false
}

// SetIpv4 gets a reference to the given GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerIpv4 and assigns it to the Ipv4 field.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) SetIpv4(v GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerIpv4) {
	o.Ipv4 = &v
}

// GetIsConfigured returns the IsConfigured field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) GetIsConfigured() bool {
	if o == nil || isNil(o.IsConfigured) {
		var ret bool
		return ret
	}
	return *o.IsConfigured
}

// GetIsConfiguredOk returns a tuple with the IsConfigured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) GetIsConfiguredOk() (*bool, bool) {
	if o == nil || isNil(o.IsConfigured) {
    return nil, false
	}
	return o.IsConfigured, true
}

// HasIsConfigured returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) HasIsConfigured() bool {
	if o != nil && !isNil(o.IsConfigured) {
		return true
	}

	return false
}

// SetIsConfigured gets a reference to the given bool and assigns it to the IsConfigured field.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) SetIsConfigured(v bool) {
	o.IsConfigured = &v
}

// GetLastAck returns the LastAck field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) GetLastAck() GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastAck {
	if o == nil || isNil(o.LastAck) {
		var ret GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastAck
		return ret
	}
	return *o.LastAck
}

// GetLastAckOk returns a tuple with the LastAck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) GetLastAckOk() (*GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastAck, bool) {
	if o == nil || isNil(o.LastAck) {
    return nil, false
	}
	return o.LastAck, true
}

// HasLastAck returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) HasLastAck() bool {
	if o != nil && !isNil(o.LastAck) {
		return true
	}

	return false
}

// SetLastAck gets a reference to the given GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastAck and assigns it to the LastAck field.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) SetLastAck(v GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastAck) {
	o.LastAck = &v
}

// GetLastPacket returns the LastPacket field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) GetLastPacket() GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket {
	if o == nil || isNil(o.LastPacket) {
		var ret GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket
		return ret
	}
	return *o.LastPacket
}

// GetLastPacketOk returns a tuple with the LastPacket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) GetLastPacketOk() (*GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket, bool) {
	if o == nil || isNil(o.LastPacket) {
    return nil, false
	}
	return o.LastPacket, true
}

// HasLastPacket returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) HasLastPacket() bool {
	if o != nil && !isNil(o.LastPacket) {
		return true
	}

	return false
}

// SetLastPacket gets a reference to the given GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket and assigns it to the LastPacket field.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) SetLastPacket(v GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacket) {
	o.LastPacket = &v
}

func (o GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !isNil(o.Vlan) {
		toSerialize["vlan"] = o.Vlan
	}
	if !isNil(o.ClientId) {
		toSerialize["clientId"] = o.ClientId
	}
	if !isNil(o.IsAllowed) {
		toSerialize["isAllowed"] = o.IsAllowed
	}
	if !isNil(o.LastSeenAt) {
		toSerialize["lastSeenAt"] = o.LastSeenAt
	}
	if !isNil(o.SeenBy) {
		toSerialize["seenBy"] = o.SeenBy
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Device) {
		toSerialize["device"] = o.Device
	}
	if !isNil(o.Ipv4) {
		toSerialize["ipv4"] = o.Ipv4
	}
	if !isNil(o.IsConfigured) {
		toSerialize["isConfigured"] = o.IsConfigured
	}
	if !isNil(o.LastAck) {
		toSerialize["lastAck"] = o.LastAck
	}
	if !isNil(o.LastPacket) {
		toSerialize["lastPacket"] = o.LastPacket
	}
	return json.Marshal(toSerialize)
}

type NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInner struct {
	value *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner
	isSet bool
}

func (v NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInner) Get() *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner {
	return v.value
}

func (v *NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInner) Set(val *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInner(val *GetNetworkSwitchDhcpV4ServersSeen200ResponseInner) *NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInner {
	return &NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInner{value: val, isSet: true}
}

func (v NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


