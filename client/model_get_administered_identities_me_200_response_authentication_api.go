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

// checks if the GetAdministeredIdentitiesMe200ResponseAuthenticationApi type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAdministeredIdentitiesMe200ResponseAuthenticationApi{}

// GetAdministeredIdentitiesMe200ResponseAuthenticationApi API authentication
type GetAdministeredIdentitiesMe200ResponseAuthenticationApi struct {
	Key *GetAdministeredIdentitiesMe200ResponseAuthenticationApiKey `json:"key,omitempty"`
}

// NewGetAdministeredIdentitiesMe200ResponseAuthenticationApi instantiates a new GetAdministeredIdentitiesMe200ResponseAuthenticationApi object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAdministeredIdentitiesMe200ResponseAuthenticationApi() *GetAdministeredIdentitiesMe200ResponseAuthenticationApi {
	this := GetAdministeredIdentitiesMe200ResponseAuthenticationApi{}
	return &this
}

// NewGetAdministeredIdentitiesMe200ResponseAuthenticationApiWithDefaults instantiates a new GetAdministeredIdentitiesMe200ResponseAuthenticationApi object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAdministeredIdentitiesMe200ResponseAuthenticationApiWithDefaults() *GetAdministeredIdentitiesMe200ResponseAuthenticationApi {
	this := GetAdministeredIdentitiesMe200ResponseAuthenticationApi{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *GetAdministeredIdentitiesMe200ResponseAuthenticationApi) GetKey() GetAdministeredIdentitiesMe200ResponseAuthenticationApiKey {
	if o == nil || IsNil(o.Key) {
		var ret GetAdministeredIdentitiesMe200ResponseAuthenticationApiKey
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAdministeredIdentitiesMe200ResponseAuthenticationApi) GetKeyOk() (*GetAdministeredIdentitiesMe200ResponseAuthenticationApiKey, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *GetAdministeredIdentitiesMe200ResponseAuthenticationApi) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given GetAdministeredIdentitiesMe200ResponseAuthenticationApiKey and assigns it to the Key field.
func (o *GetAdministeredIdentitiesMe200ResponseAuthenticationApi) SetKey(v GetAdministeredIdentitiesMe200ResponseAuthenticationApiKey) {
	o.Key = &v
}

func (o GetAdministeredIdentitiesMe200ResponseAuthenticationApi) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAdministeredIdentitiesMe200ResponseAuthenticationApi) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	return toSerialize, nil
}

type NullableGetAdministeredIdentitiesMe200ResponseAuthenticationApi struct {
	value *GetAdministeredIdentitiesMe200ResponseAuthenticationApi
	isSet bool
}

func (v NullableGetAdministeredIdentitiesMe200ResponseAuthenticationApi) Get() *GetAdministeredIdentitiesMe200ResponseAuthenticationApi {
	return v.value
}

func (v *NullableGetAdministeredIdentitiesMe200ResponseAuthenticationApi) Set(val *GetAdministeredIdentitiesMe200ResponseAuthenticationApi) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAdministeredIdentitiesMe200ResponseAuthenticationApi) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAdministeredIdentitiesMe200ResponseAuthenticationApi) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAdministeredIdentitiesMe200ResponseAuthenticationApi(val *GetAdministeredIdentitiesMe200ResponseAuthenticationApi) *NullableGetAdministeredIdentitiesMe200ResponseAuthenticationApi {
	return &NullableGetAdministeredIdentitiesMe200ResponseAuthenticationApi{value: val, isSet: true}
}

func (v NullableGetAdministeredIdentitiesMe200ResponseAuthenticationApi) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAdministeredIdentitiesMe200ResponseAuthenticationApi) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


