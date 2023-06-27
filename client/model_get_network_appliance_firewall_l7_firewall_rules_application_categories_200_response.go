/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 07 June, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.34.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200Response struct for GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200Response
type GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200Response struct {
	//  The L7 firewall application categories and their associated applications for an MX network
	ApplicationCategories []GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInner `json:"applicationCategories,omitempty"`
}

// NewGetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200Response instantiates a new GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200Response() *GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200Response {
	this := GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200Response{}
	return &this
}

// NewGetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseWithDefaults instantiates a new GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseWithDefaults() *GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200Response {
	this := GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200Response{}
	return &this
}

// GetApplicationCategories returns the ApplicationCategories field value if set, zero value otherwise.
func (o *GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200Response) GetApplicationCategories() []GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInner {
	if o == nil || isNil(o.ApplicationCategories) {
		var ret []GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInner
		return ret
	}
	return o.ApplicationCategories
}

// GetApplicationCategoriesOk returns a tuple with the ApplicationCategories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200Response) GetApplicationCategoriesOk() ([]GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInner, bool) {
	if o == nil || isNil(o.ApplicationCategories) {
    return nil, false
	}
	return o.ApplicationCategories, true
}

// HasApplicationCategories returns a boolean if a field has been set.
func (o *GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200Response) HasApplicationCategories() bool {
	if o != nil && !isNil(o.ApplicationCategories) {
		return true
	}

	return false
}

// SetApplicationCategories gets a reference to the given []GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInner and assigns it to the ApplicationCategories field.
func (o *GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200Response) SetApplicationCategories(v []GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInner) {
	o.ApplicationCategories = v
}

func (o GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ApplicationCategories) {
		toSerialize["applicationCategories"] = o.ApplicationCategories
	}
	return json.Marshal(toSerialize)
}

type NullableGetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200Response struct {
	value *GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200Response
	isSet bool
}

func (v NullableGetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200Response) Get() *GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200Response {
	return v.value
}

func (v *NullableGetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200Response) Set(val *GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200Response(val *GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200Response) *NullableGetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200Response {
	return &NullableGetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200Response{value: val, isSet: true}
}

func (v NullableGetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


