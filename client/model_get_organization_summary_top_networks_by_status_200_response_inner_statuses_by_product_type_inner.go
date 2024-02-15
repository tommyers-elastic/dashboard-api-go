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

// checks if the GetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatusesByProductTypeInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatusesByProductTypeInner{}

// GetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatusesByProductTypeInner struct for GetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatusesByProductTypeInner
type GetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatusesByProductTypeInner struct {
	// Product type
	ProductType *string `json:"productType,omitempty"`
	Counts *GetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatusesByProductTypeInnerCounts `json:"counts,omitempty"`
}

// NewGetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatusesByProductTypeInner instantiates a new GetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatusesByProductTypeInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatusesByProductTypeInner() *GetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatusesByProductTypeInner {
	this := GetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatusesByProductTypeInner{}
	return &this
}

// NewGetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatusesByProductTypeInnerWithDefaults instantiates a new GetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatusesByProductTypeInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatusesByProductTypeInnerWithDefaults() *GetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatusesByProductTypeInner {
	this := GetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatusesByProductTypeInner{}
	return &this
}

// GetProductType returns the ProductType field value if set, zero value otherwise.
func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatusesByProductTypeInner) GetProductType() string {
	if o == nil || IsNil(o.ProductType) {
		var ret string
		return ret
	}
	return *o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatusesByProductTypeInner) GetProductTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ProductType) {
		return nil, false
	}
	return o.ProductType, true
}

// HasProductType returns a boolean if a field has been set.
func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatusesByProductTypeInner) HasProductType() bool {
	if o != nil && !IsNil(o.ProductType) {
		return true
	}

	return false
}

// SetProductType gets a reference to the given string and assigns it to the ProductType field.
func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatusesByProductTypeInner) SetProductType(v string) {
	o.ProductType = &v
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatusesByProductTypeInner) GetCounts() GetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatusesByProductTypeInnerCounts {
	if o == nil || IsNil(o.Counts) {
		var ret GetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatusesByProductTypeInnerCounts
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatusesByProductTypeInner) GetCountsOk() (*GetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatusesByProductTypeInnerCounts, bool) {
	if o == nil || IsNil(o.Counts) {
		return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatusesByProductTypeInner) HasCounts() bool {
	if o != nil && !IsNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given GetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatusesByProductTypeInnerCounts and assigns it to the Counts field.
func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatusesByProductTypeInner) SetCounts(v GetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatusesByProductTypeInnerCounts) {
	o.Counts = &v
}

func (o GetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatusesByProductTypeInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatusesByProductTypeInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProductType) {
		toSerialize["productType"] = o.ProductType
	}
	if !IsNil(o.Counts) {
		toSerialize["counts"] = o.Counts
	}
	return toSerialize, nil
}

type NullableGetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatusesByProductTypeInner struct {
	value *GetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatusesByProductTypeInner
	isSet bool
}

func (v NullableGetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatusesByProductTypeInner) Get() *GetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatusesByProductTypeInner {
	return v.value
}

func (v *NullableGetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatusesByProductTypeInner) Set(val *GetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatusesByProductTypeInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatusesByProductTypeInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatusesByProductTypeInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatusesByProductTypeInner(val *GetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatusesByProductTypeInner) *NullableGetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatusesByProductTypeInner {
	return &NullableGetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatusesByProductTypeInner{value: val, isSet: true}
}

func (v NullableGetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatusesByProductTypeInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatusesByProductTypeInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


