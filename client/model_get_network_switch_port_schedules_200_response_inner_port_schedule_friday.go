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

// checks if the GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleFriday type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleFriday{}

// GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleFriday Friday schedule
type GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleFriday struct {
	// Whether the schedule is active or inactive
	Active *bool `json:"active,omitempty"`
	// The time, from '00:00' to '24:00'
	From *string `json:"from,omitempty"`
	// The time, from '00:00' to '24:00'
	To *string `json:"to,omitempty"`
}

// NewGetNetworkSwitchPortSchedules200ResponseInnerPortScheduleFriday instantiates a new GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleFriday object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSwitchPortSchedules200ResponseInnerPortScheduleFriday() *GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleFriday {
	this := GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleFriday{}
	return &this
}

// NewGetNetworkSwitchPortSchedules200ResponseInnerPortScheduleFridayWithDefaults instantiates a new GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleFriday object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSwitchPortSchedules200ResponseInnerPortScheduleFridayWithDefaults() *GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleFriday {
	this := GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleFriday{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleFriday) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleFriday) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleFriday) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleFriday) SetActive(v bool) {
	o.Active = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleFriday) GetFrom() string {
	if o == nil || IsNil(o.From) {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleFriday) GetFromOk() (*string, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleFriday) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleFriday) SetFrom(v string) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleFriday) GetTo() string {
	if o == nil || IsNil(o.To) {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleFriday) GetToOk() (*string, bool) {
	if o == nil || IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleFriday) HasTo() bool {
	if o != nil && !IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleFriday) SetTo(v string) {
	o.To = &v
}

func (o GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleFriday) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleFriday) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !IsNil(o.To) {
		toSerialize["to"] = o.To
	}
	return toSerialize, nil
}

type NullableGetNetworkSwitchPortSchedules200ResponseInnerPortScheduleFriday struct {
	value *GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleFriday
	isSet bool
}

func (v NullableGetNetworkSwitchPortSchedules200ResponseInnerPortScheduleFriday) Get() *GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleFriday {
	return v.value
}

func (v *NullableGetNetworkSwitchPortSchedules200ResponseInnerPortScheduleFriday) Set(val *GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleFriday) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSwitchPortSchedules200ResponseInnerPortScheduleFriday) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSwitchPortSchedules200ResponseInnerPortScheduleFriday) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSwitchPortSchedules200ResponseInnerPortScheduleFriday(val *GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleFriday) *NullableGetNetworkSwitchPortSchedules200ResponseInnerPortScheduleFriday {
	return &NullableGetNetworkSwitchPortSchedules200ResponseInnerPortScheduleFriday{value: val, isSet: true}
}

func (v NullableGetNetworkSwitchPortSchedules200ResponseInnerPortScheduleFriday) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSwitchPortSchedules200ResponseInnerPortScheduleFriday) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


