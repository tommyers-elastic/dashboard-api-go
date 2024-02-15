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

// checks if the ClaimIntoOrganizationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClaimIntoOrganizationRequest{}

// ClaimIntoOrganizationRequest struct for ClaimIntoOrganizationRequest
type ClaimIntoOrganizationRequest struct {
	// The numbers of the orders that should be claimed
	Orders []string `json:"orders,omitempty"`
	// The serials of the devices that should be claimed
	Serials []string `json:"serials,omitempty"`
	// The licenses that should be claimed
	Licenses []ClaimIntoOrganizationRequestLicensesInner `json:"licenses,omitempty"`
}

// NewClaimIntoOrganizationRequest instantiates a new ClaimIntoOrganizationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClaimIntoOrganizationRequest() *ClaimIntoOrganizationRequest {
	this := ClaimIntoOrganizationRequest{}
	return &this
}

// NewClaimIntoOrganizationRequestWithDefaults instantiates a new ClaimIntoOrganizationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClaimIntoOrganizationRequestWithDefaults() *ClaimIntoOrganizationRequest {
	this := ClaimIntoOrganizationRequest{}
	return &this
}

// GetOrders returns the Orders field value if set, zero value otherwise.
func (o *ClaimIntoOrganizationRequest) GetOrders() []string {
	if o == nil || IsNil(o.Orders) {
		var ret []string
		return ret
	}
	return o.Orders
}

// GetOrdersOk returns a tuple with the Orders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClaimIntoOrganizationRequest) GetOrdersOk() ([]string, bool) {
	if o == nil || IsNil(o.Orders) {
		return nil, false
	}
	return o.Orders, true
}

// HasOrders returns a boolean if a field has been set.
func (o *ClaimIntoOrganizationRequest) HasOrders() bool {
	if o != nil && !IsNil(o.Orders) {
		return true
	}

	return false
}

// SetOrders gets a reference to the given []string and assigns it to the Orders field.
func (o *ClaimIntoOrganizationRequest) SetOrders(v []string) {
	o.Orders = v
}

// GetSerials returns the Serials field value if set, zero value otherwise.
func (o *ClaimIntoOrganizationRequest) GetSerials() []string {
	if o == nil || IsNil(o.Serials) {
		var ret []string
		return ret
	}
	return o.Serials
}

// GetSerialsOk returns a tuple with the Serials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClaimIntoOrganizationRequest) GetSerialsOk() ([]string, bool) {
	if o == nil || IsNil(o.Serials) {
		return nil, false
	}
	return o.Serials, true
}

// HasSerials returns a boolean if a field has been set.
func (o *ClaimIntoOrganizationRequest) HasSerials() bool {
	if o != nil && !IsNil(o.Serials) {
		return true
	}

	return false
}

// SetSerials gets a reference to the given []string and assigns it to the Serials field.
func (o *ClaimIntoOrganizationRequest) SetSerials(v []string) {
	o.Serials = v
}

// GetLicenses returns the Licenses field value if set, zero value otherwise.
func (o *ClaimIntoOrganizationRequest) GetLicenses() []ClaimIntoOrganizationRequestLicensesInner {
	if o == nil || IsNil(o.Licenses) {
		var ret []ClaimIntoOrganizationRequestLicensesInner
		return ret
	}
	return o.Licenses
}

// GetLicensesOk returns a tuple with the Licenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClaimIntoOrganizationRequest) GetLicensesOk() ([]ClaimIntoOrganizationRequestLicensesInner, bool) {
	if o == nil || IsNil(o.Licenses) {
		return nil, false
	}
	return o.Licenses, true
}

// HasLicenses returns a boolean if a field has been set.
func (o *ClaimIntoOrganizationRequest) HasLicenses() bool {
	if o != nil && !IsNil(o.Licenses) {
		return true
	}

	return false
}

// SetLicenses gets a reference to the given []ClaimIntoOrganizationRequestLicensesInner and assigns it to the Licenses field.
func (o *ClaimIntoOrganizationRequest) SetLicenses(v []ClaimIntoOrganizationRequestLicensesInner) {
	o.Licenses = v
}

func (o ClaimIntoOrganizationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClaimIntoOrganizationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Orders) {
		toSerialize["orders"] = o.Orders
	}
	if !IsNil(o.Serials) {
		toSerialize["serials"] = o.Serials
	}
	if !IsNil(o.Licenses) {
		toSerialize["licenses"] = o.Licenses
	}
	return toSerialize, nil
}

type NullableClaimIntoOrganizationRequest struct {
	value *ClaimIntoOrganizationRequest
	isSet bool
}

func (v NullableClaimIntoOrganizationRequest) Get() *ClaimIntoOrganizationRequest {
	return v.value
}

func (v *NullableClaimIntoOrganizationRequest) Set(val *ClaimIntoOrganizationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableClaimIntoOrganizationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableClaimIntoOrganizationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClaimIntoOrganizationRequest(val *ClaimIntoOrganizationRequest) *NullableClaimIntoOrganizationRequest {
	return &NullableClaimIntoOrganizationRequest{value: val, isSet: true}
}

func (v NullableClaimIntoOrganizationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClaimIntoOrganizationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


