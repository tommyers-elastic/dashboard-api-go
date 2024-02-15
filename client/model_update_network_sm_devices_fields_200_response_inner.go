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

// checks if the UpdateNetworkSmDevicesFields200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkSmDevicesFields200ResponseInner{}

// UpdateNetworkSmDevicesFields200ResponseInner struct for UpdateNetworkSmDevicesFields200ResponseInner
type UpdateNetworkSmDevicesFields200ResponseInner struct {
	// The Meraki Id of the device record.
	Id *string `json:"id,omitempty"`
	// The name of the device.
	Name *string `json:"name,omitempty"`
	// The MAC of the device.
	WifiMac *string `json:"wifiMac,omitempty"`
	// The device serial.
	Serial *string `json:"serial,omitempty"`
	// Notes associated with the device.
	Notes *string `json:"notes,omitempty"`
}

// NewUpdateNetworkSmDevicesFields200ResponseInner instantiates a new UpdateNetworkSmDevicesFields200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkSmDevicesFields200ResponseInner() *UpdateNetworkSmDevicesFields200ResponseInner {
	this := UpdateNetworkSmDevicesFields200ResponseInner{}
	return &this
}

// NewUpdateNetworkSmDevicesFields200ResponseInnerWithDefaults instantiates a new UpdateNetworkSmDevicesFields200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkSmDevicesFields200ResponseInnerWithDefaults() *UpdateNetworkSmDevicesFields200ResponseInner {
	this := UpdateNetworkSmDevicesFields200ResponseInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UpdateNetworkSmDevicesFields200ResponseInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSmDevicesFields200ResponseInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UpdateNetworkSmDevicesFields200ResponseInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UpdateNetworkSmDevicesFields200ResponseInner) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateNetworkSmDevicesFields200ResponseInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSmDevicesFields200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateNetworkSmDevicesFields200ResponseInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateNetworkSmDevicesFields200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetWifiMac returns the WifiMac field value if set, zero value otherwise.
func (o *UpdateNetworkSmDevicesFields200ResponseInner) GetWifiMac() string {
	if o == nil || IsNil(o.WifiMac) {
		var ret string
		return ret
	}
	return *o.WifiMac
}

// GetWifiMacOk returns a tuple with the WifiMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSmDevicesFields200ResponseInner) GetWifiMacOk() (*string, bool) {
	if o == nil || IsNil(o.WifiMac) {
		return nil, false
	}
	return o.WifiMac, true
}

// HasWifiMac returns a boolean if a field has been set.
func (o *UpdateNetworkSmDevicesFields200ResponseInner) HasWifiMac() bool {
	if o != nil && !IsNil(o.WifiMac) {
		return true
	}

	return false
}

// SetWifiMac gets a reference to the given string and assigns it to the WifiMac field.
func (o *UpdateNetworkSmDevicesFields200ResponseInner) SetWifiMac(v string) {
	o.WifiMac = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *UpdateNetworkSmDevicesFields200ResponseInner) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSmDevicesFields200ResponseInner) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *UpdateNetworkSmDevicesFields200ResponseInner) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *UpdateNetworkSmDevicesFields200ResponseInner) SetSerial(v string) {
	o.Serial = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *UpdateNetworkSmDevicesFields200ResponseInner) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSmDevicesFields200ResponseInner) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *UpdateNetworkSmDevicesFields200ResponseInner) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *UpdateNetworkSmDevicesFields200ResponseInner) SetNotes(v string) {
	o.Notes = &v
}

func (o UpdateNetworkSmDevicesFields200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkSmDevicesFields200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.WifiMac) {
		toSerialize["wifiMac"] = o.WifiMac
	}
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	return toSerialize, nil
}

type NullableUpdateNetworkSmDevicesFields200ResponseInner struct {
	value *UpdateNetworkSmDevicesFields200ResponseInner
	isSet bool
}

func (v NullableUpdateNetworkSmDevicesFields200ResponseInner) Get() *UpdateNetworkSmDevicesFields200ResponseInner {
	return v.value
}

func (v *NullableUpdateNetworkSmDevicesFields200ResponseInner) Set(val *UpdateNetworkSmDevicesFields200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkSmDevicesFields200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkSmDevicesFields200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkSmDevicesFields200ResponseInner(val *UpdateNetworkSmDevicesFields200ResponseInner) *NullableUpdateNetworkSmDevicesFields200ResponseInner {
	return &NullableUpdateNetworkSmDevicesFields200ResponseInner{value: val, isSet: true}
}

func (v NullableUpdateNetworkSmDevicesFields200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkSmDevicesFields200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


