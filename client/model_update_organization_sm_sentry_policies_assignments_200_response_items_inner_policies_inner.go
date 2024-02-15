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

// checks if the UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner{}

// UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner struct for UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner
type UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner struct {
	// The Id of the Sentry Policy
	PolicyId *string `json:"policyId,omitempty"`
	// The Id of the Network the Sentry Policy is associated with. In a locale, this should be the Wireless Group if present, otherwise the Wired Group.
	NetworkId *string `json:"networkId,omitempty"`
	// The Id of the Systems Manager Network the Sentry Policy is assigned to
	SmNetworkId *string `json:"smNetworkId,omitempty"`
	// The tags of the Sentry Policy
	Tags []string `json:"tags,omitempty"`
	// The scope of the Sentry Policy
	Scope *string `json:"scope,omitempty"`
	// The number of the Group Policy
	GroupNumber *string `json:"groupNumber,omitempty"`
	// The Id of the Group Policy. This is associated with the network specified by the networkId.
	GroupPolicyId *string `json:"groupPolicyId,omitempty"`
	// The priority of the Sentry Policy
	Priority *string `json:"priority,omitempty"`
	// The creation time of the Sentry Policy
	CreatedAt *string `json:"createdAt,omitempty"`
	// The last update time of the Sentry Policy
	LastUpdatedAt *string `json:"lastUpdatedAt,omitempty"`
}

// NewUpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner instantiates a new UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner() *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner {
	this := UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner{}
	return &this
}

// NewUpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInnerWithDefaults instantiates a new UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInnerWithDefaults() *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner {
	this := UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner{}
	return &this
}

// GetPolicyId returns the PolicyId field value if set, zero value otherwise.
func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) GetPolicyId() string {
	if o == nil || IsNil(o.PolicyId) {
		var ret string
		return ret
	}
	return *o.PolicyId
}

// GetPolicyIdOk returns a tuple with the PolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) GetPolicyIdOk() (*string, bool) {
	if o == nil || IsNil(o.PolicyId) {
		return nil, false
	}
	return o.PolicyId, true
}

// HasPolicyId returns a boolean if a field has been set.
func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) HasPolicyId() bool {
	if o != nil && !IsNil(o.PolicyId) {
		return true
	}

	return false
}

// SetPolicyId gets a reference to the given string and assigns it to the PolicyId field.
func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) SetPolicyId(v string) {
	o.PolicyId = &v
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) GetNetworkId() string {
	if o == nil || IsNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) GetNetworkIdOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkId) {
		return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) HasNetworkId() bool {
	if o != nil && !IsNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetSmNetworkId returns the SmNetworkId field value if set, zero value otherwise.
func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) GetSmNetworkId() string {
	if o == nil || IsNil(o.SmNetworkId) {
		var ret string
		return ret
	}
	return *o.SmNetworkId
}

// GetSmNetworkIdOk returns a tuple with the SmNetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) GetSmNetworkIdOk() (*string, bool) {
	if o == nil || IsNil(o.SmNetworkId) {
		return nil, false
	}
	return o.SmNetworkId, true
}

// HasSmNetworkId returns a boolean if a field has been set.
func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) HasSmNetworkId() bool {
	if o != nil && !IsNil(o.SmNetworkId) {
		return true
	}

	return false
}

// SetSmNetworkId gets a reference to the given string and assigns it to the SmNetworkId field.
func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) SetSmNetworkId(v string) {
	o.SmNetworkId = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) SetTags(v []string) {
	o.Tags = v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) GetScope() string {
	if o == nil || IsNil(o.Scope) {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) GetScopeOk() (*string, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) SetScope(v string) {
	o.Scope = &v
}

// GetGroupNumber returns the GroupNumber field value if set, zero value otherwise.
func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) GetGroupNumber() string {
	if o == nil || IsNil(o.GroupNumber) {
		var ret string
		return ret
	}
	return *o.GroupNumber
}

// GetGroupNumberOk returns a tuple with the GroupNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) GetGroupNumberOk() (*string, bool) {
	if o == nil || IsNil(o.GroupNumber) {
		return nil, false
	}
	return o.GroupNumber, true
}

// HasGroupNumber returns a boolean if a field has been set.
func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) HasGroupNumber() bool {
	if o != nil && !IsNil(o.GroupNumber) {
		return true
	}

	return false
}

// SetGroupNumber gets a reference to the given string and assigns it to the GroupNumber field.
func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) SetGroupNumber(v string) {
	o.GroupNumber = &v
}

// GetGroupPolicyId returns the GroupPolicyId field value if set, zero value otherwise.
func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) GetGroupPolicyId() string {
	if o == nil || IsNil(o.GroupPolicyId) {
		var ret string
		return ret
	}
	return *o.GroupPolicyId
}

// GetGroupPolicyIdOk returns a tuple with the GroupPolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) GetGroupPolicyIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupPolicyId) {
		return nil, false
	}
	return o.GroupPolicyId, true
}

// HasGroupPolicyId returns a boolean if a field has been set.
func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) HasGroupPolicyId() bool {
	if o != nil && !IsNil(o.GroupPolicyId) {
		return true
	}

	return false
}

// SetGroupPolicyId gets a reference to the given string and assigns it to the GroupPolicyId field.
func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) SetGroupPolicyId(v string) {
	o.GroupPolicyId = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) GetPriority() string {
	if o == nil || IsNil(o.Priority) {
		var ret string
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) GetPriorityOk() (*string, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given string and assigns it to the Priority field.
func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) SetPriority(v string) {
	o.Priority = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetLastUpdatedAt returns the LastUpdatedAt field value if set, zero value otherwise.
func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) GetLastUpdatedAt() string {
	if o == nil || IsNil(o.LastUpdatedAt) {
		var ret string
		return ret
	}
	return *o.LastUpdatedAt
}

// GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) GetLastUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.LastUpdatedAt) {
		return nil, false
	}
	return o.LastUpdatedAt, true
}

// HasLastUpdatedAt returns a boolean if a field has been set.
func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) HasLastUpdatedAt() bool {
	if o != nil && !IsNil(o.LastUpdatedAt) {
		return true
	}

	return false
}

// SetLastUpdatedAt gets a reference to the given string and assigns it to the LastUpdatedAt field.
func (o *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) SetLastUpdatedAt(v string) {
	o.LastUpdatedAt = &v
}

func (o UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PolicyId) {
		toSerialize["policyId"] = o.PolicyId
	}
	if !IsNil(o.NetworkId) {
		toSerialize["networkId"] = o.NetworkId
	}
	if !IsNil(o.SmNetworkId) {
		toSerialize["smNetworkId"] = o.SmNetworkId
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	if !IsNil(o.GroupNumber) {
		toSerialize["groupNumber"] = o.GroupNumber
	}
	if !IsNil(o.GroupPolicyId) {
		toSerialize["groupPolicyId"] = o.GroupPolicyId
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.LastUpdatedAt) {
		toSerialize["lastUpdatedAt"] = o.LastUpdatedAt
	}
	return toSerialize, nil
}

type NullableUpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner struct {
	value *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner
	isSet bool
}

func (v NullableUpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) Get() *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner {
	return v.value
}

func (v *NullableUpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) Set(val *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner(val *UpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) *NullableUpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner {
	return &NullableUpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner{value: val, isSet: true}
}

func (v NullableUpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOrganizationSmSentryPoliciesAssignments200ResponseItemsInnerPoliciesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


