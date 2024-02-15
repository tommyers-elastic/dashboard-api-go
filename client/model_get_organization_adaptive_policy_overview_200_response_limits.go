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

// checks if the GetOrganizationAdaptivePolicyOverview200ResponseLimits type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationAdaptivePolicyOverview200ResponseLimits{}

// GetOrganizationAdaptivePolicyOverview200ResponseLimits The current limits of various adaptive policy objects.
type GetOrganizationAdaptivePolicyOverview200ResponseLimits struct {
	// Maximum number of user-created adaptive policy groups allowed in the organization.
	CustomGroups *int32 `json:"customGroups,omitempty"`
	// Maximum number of rules allowed in an adaptive policy ACL in the organization.
	RulesInAnAcl *int32 `json:"rulesInAnAcl,omitempty"`
	// Maximum number of adaptive policy ACLs that can be assigned to an adaptive policy in the organization.
	AclsInAPolicy *int32 `json:"aclsInAPolicy,omitempty"`
	// Maximum number of policy objects (with the adaptive policy type) allowed in the organization.
	PolicyObjects *int32 `json:"policyObjects,omitempty"`
}

// NewGetOrganizationAdaptivePolicyOverview200ResponseLimits instantiates a new GetOrganizationAdaptivePolicyOverview200ResponseLimits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationAdaptivePolicyOverview200ResponseLimits() *GetOrganizationAdaptivePolicyOverview200ResponseLimits {
	this := GetOrganizationAdaptivePolicyOverview200ResponseLimits{}
	return &this
}

// NewGetOrganizationAdaptivePolicyOverview200ResponseLimitsWithDefaults instantiates a new GetOrganizationAdaptivePolicyOverview200ResponseLimits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationAdaptivePolicyOverview200ResponseLimitsWithDefaults() *GetOrganizationAdaptivePolicyOverview200ResponseLimits {
	this := GetOrganizationAdaptivePolicyOverview200ResponseLimits{}
	return &this
}

// GetCustomGroups returns the CustomGroups field value if set, zero value otherwise.
func (o *GetOrganizationAdaptivePolicyOverview200ResponseLimits) GetCustomGroups() int32 {
	if o == nil || IsNil(o.CustomGroups) {
		var ret int32
		return ret
	}
	return *o.CustomGroups
}

// GetCustomGroupsOk returns a tuple with the CustomGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAdaptivePolicyOverview200ResponseLimits) GetCustomGroupsOk() (*int32, bool) {
	if o == nil || IsNil(o.CustomGroups) {
		return nil, false
	}
	return o.CustomGroups, true
}

// HasCustomGroups returns a boolean if a field has been set.
func (o *GetOrganizationAdaptivePolicyOverview200ResponseLimits) HasCustomGroups() bool {
	if o != nil && !IsNil(o.CustomGroups) {
		return true
	}

	return false
}

// SetCustomGroups gets a reference to the given int32 and assigns it to the CustomGroups field.
func (o *GetOrganizationAdaptivePolicyOverview200ResponseLimits) SetCustomGroups(v int32) {
	o.CustomGroups = &v
}

// GetRulesInAnAcl returns the RulesInAnAcl field value if set, zero value otherwise.
func (o *GetOrganizationAdaptivePolicyOverview200ResponseLimits) GetRulesInAnAcl() int32 {
	if o == nil || IsNil(o.RulesInAnAcl) {
		var ret int32
		return ret
	}
	return *o.RulesInAnAcl
}

// GetRulesInAnAclOk returns a tuple with the RulesInAnAcl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAdaptivePolicyOverview200ResponseLimits) GetRulesInAnAclOk() (*int32, bool) {
	if o == nil || IsNil(o.RulesInAnAcl) {
		return nil, false
	}
	return o.RulesInAnAcl, true
}

// HasRulesInAnAcl returns a boolean if a field has been set.
func (o *GetOrganizationAdaptivePolicyOverview200ResponseLimits) HasRulesInAnAcl() bool {
	if o != nil && !IsNil(o.RulesInAnAcl) {
		return true
	}

	return false
}

// SetRulesInAnAcl gets a reference to the given int32 and assigns it to the RulesInAnAcl field.
func (o *GetOrganizationAdaptivePolicyOverview200ResponseLimits) SetRulesInAnAcl(v int32) {
	o.RulesInAnAcl = &v
}

// GetAclsInAPolicy returns the AclsInAPolicy field value if set, zero value otherwise.
func (o *GetOrganizationAdaptivePolicyOverview200ResponseLimits) GetAclsInAPolicy() int32 {
	if o == nil || IsNil(o.AclsInAPolicy) {
		var ret int32
		return ret
	}
	return *o.AclsInAPolicy
}

// GetAclsInAPolicyOk returns a tuple with the AclsInAPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAdaptivePolicyOverview200ResponseLimits) GetAclsInAPolicyOk() (*int32, bool) {
	if o == nil || IsNil(o.AclsInAPolicy) {
		return nil, false
	}
	return o.AclsInAPolicy, true
}

// HasAclsInAPolicy returns a boolean if a field has been set.
func (o *GetOrganizationAdaptivePolicyOverview200ResponseLimits) HasAclsInAPolicy() bool {
	if o != nil && !IsNil(o.AclsInAPolicy) {
		return true
	}

	return false
}

// SetAclsInAPolicy gets a reference to the given int32 and assigns it to the AclsInAPolicy field.
func (o *GetOrganizationAdaptivePolicyOverview200ResponseLimits) SetAclsInAPolicy(v int32) {
	o.AclsInAPolicy = &v
}

// GetPolicyObjects returns the PolicyObjects field value if set, zero value otherwise.
func (o *GetOrganizationAdaptivePolicyOverview200ResponseLimits) GetPolicyObjects() int32 {
	if o == nil || IsNil(o.PolicyObjects) {
		var ret int32
		return ret
	}
	return *o.PolicyObjects
}

// GetPolicyObjectsOk returns a tuple with the PolicyObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAdaptivePolicyOverview200ResponseLimits) GetPolicyObjectsOk() (*int32, bool) {
	if o == nil || IsNil(o.PolicyObjects) {
		return nil, false
	}
	return o.PolicyObjects, true
}

// HasPolicyObjects returns a boolean if a field has been set.
func (o *GetOrganizationAdaptivePolicyOverview200ResponseLimits) HasPolicyObjects() bool {
	if o != nil && !IsNil(o.PolicyObjects) {
		return true
	}

	return false
}

// SetPolicyObjects gets a reference to the given int32 and assigns it to the PolicyObjects field.
func (o *GetOrganizationAdaptivePolicyOverview200ResponseLimits) SetPolicyObjects(v int32) {
	o.PolicyObjects = &v
}

func (o GetOrganizationAdaptivePolicyOverview200ResponseLimits) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationAdaptivePolicyOverview200ResponseLimits) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomGroups) {
		toSerialize["customGroups"] = o.CustomGroups
	}
	if !IsNil(o.RulesInAnAcl) {
		toSerialize["rulesInAnAcl"] = o.RulesInAnAcl
	}
	if !IsNil(o.AclsInAPolicy) {
		toSerialize["aclsInAPolicy"] = o.AclsInAPolicy
	}
	if !IsNil(o.PolicyObjects) {
		toSerialize["policyObjects"] = o.PolicyObjects
	}
	return toSerialize, nil
}

type NullableGetOrganizationAdaptivePolicyOverview200ResponseLimits struct {
	value *GetOrganizationAdaptivePolicyOverview200ResponseLimits
	isSet bool
}

func (v NullableGetOrganizationAdaptivePolicyOverview200ResponseLimits) Get() *GetOrganizationAdaptivePolicyOverview200ResponseLimits {
	return v.value
}

func (v *NullableGetOrganizationAdaptivePolicyOverview200ResponseLimits) Set(val *GetOrganizationAdaptivePolicyOverview200ResponseLimits) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationAdaptivePolicyOverview200ResponseLimits) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationAdaptivePolicyOverview200ResponseLimits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationAdaptivePolicyOverview200ResponseLimits(val *GetOrganizationAdaptivePolicyOverview200ResponseLimits) *NullableGetOrganizationAdaptivePolicyOverview200ResponseLimits {
	return &NullableGetOrganizationAdaptivePolicyOverview200ResponseLimits{value: val, isSet: true}
}

func (v NullableGetOrganizationAdaptivePolicyOverview200ResponseLimits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationAdaptivePolicyOverview200ResponseLimits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


