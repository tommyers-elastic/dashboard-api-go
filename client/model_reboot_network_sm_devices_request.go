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

// checks if the RebootNetworkSmDevicesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RebootNetworkSmDevicesRequest{}

// RebootNetworkSmDevicesRequest struct for RebootNetworkSmDevicesRequest
type RebootNetworkSmDevicesRequest struct {
	// The wifiMacs of the endpoints to be rebooted.
	WifiMacs []string `json:"wifiMacs,omitempty"`
	// The ids of the endpoints to be rebooted.
	Ids []string `json:"ids,omitempty"`
	// The serials of the endpoints to be rebooted.
	Serials []string `json:"serials,omitempty"`
	// The scope (one of all, none, withAny, withAll, withoutAny, or withoutAll) and a set of tags of the endpoints to be rebooted.
	Scope []string `json:"scope,omitempty"`
	// The KextPaths of the endpoints to be rebooted. Available for macOS 11+
	KextPaths []string `json:"kextPaths,omitempty"`
	// Whether or not to notify the user before rebooting the endpoint. Available for macOS 11.3+
	NotifyUser *bool `json:"notifyUser,omitempty"`
	// Whether or not to rebuild the kernel cache when rebooting the endpoint. Available for macOS 11+
	RebuildKernelCache *bool `json:"rebuildKernelCache,omitempty"`
	// Whether or not the request requires network tethering. Available for macOS and supervised iOS or tvOS
	RequestRequiresNetworkTether *bool `json:"requestRequiresNetworkTether,omitempty"`
}

// NewRebootNetworkSmDevicesRequest instantiates a new RebootNetworkSmDevicesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRebootNetworkSmDevicesRequest() *RebootNetworkSmDevicesRequest {
	this := RebootNetworkSmDevicesRequest{}
	return &this
}

// NewRebootNetworkSmDevicesRequestWithDefaults instantiates a new RebootNetworkSmDevicesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRebootNetworkSmDevicesRequestWithDefaults() *RebootNetworkSmDevicesRequest {
	this := RebootNetworkSmDevicesRequest{}
	return &this
}

// GetWifiMacs returns the WifiMacs field value if set, zero value otherwise.
func (o *RebootNetworkSmDevicesRequest) GetWifiMacs() []string {
	if o == nil || IsNil(o.WifiMacs) {
		var ret []string
		return ret
	}
	return o.WifiMacs
}

// GetWifiMacsOk returns a tuple with the WifiMacs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RebootNetworkSmDevicesRequest) GetWifiMacsOk() ([]string, bool) {
	if o == nil || IsNil(o.WifiMacs) {
		return nil, false
	}
	return o.WifiMacs, true
}

// HasWifiMacs returns a boolean if a field has been set.
func (o *RebootNetworkSmDevicesRequest) HasWifiMacs() bool {
	if o != nil && !IsNil(o.WifiMacs) {
		return true
	}

	return false
}

// SetWifiMacs gets a reference to the given []string and assigns it to the WifiMacs field.
func (o *RebootNetworkSmDevicesRequest) SetWifiMacs(v []string) {
	o.WifiMacs = v
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *RebootNetworkSmDevicesRequest) GetIds() []string {
	if o == nil || IsNil(o.Ids) {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RebootNetworkSmDevicesRequest) GetIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *RebootNetworkSmDevicesRequest) HasIds() bool {
	if o != nil && !IsNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *RebootNetworkSmDevicesRequest) SetIds(v []string) {
	o.Ids = v
}

// GetSerials returns the Serials field value if set, zero value otherwise.
func (o *RebootNetworkSmDevicesRequest) GetSerials() []string {
	if o == nil || IsNil(o.Serials) {
		var ret []string
		return ret
	}
	return o.Serials
}

// GetSerialsOk returns a tuple with the Serials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RebootNetworkSmDevicesRequest) GetSerialsOk() ([]string, bool) {
	if o == nil || IsNil(o.Serials) {
		return nil, false
	}
	return o.Serials, true
}

// HasSerials returns a boolean if a field has been set.
func (o *RebootNetworkSmDevicesRequest) HasSerials() bool {
	if o != nil && !IsNil(o.Serials) {
		return true
	}

	return false
}

// SetSerials gets a reference to the given []string and assigns it to the Serials field.
func (o *RebootNetworkSmDevicesRequest) SetSerials(v []string) {
	o.Serials = v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *RebootNetworkSmDevicesRequest) GetScope() []string {
	if o == nil || IsNil(o.Scope) {
		var ret []string
		return ret
	}
	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RebootNetworkSmDevicesRequest) GetScopeOk() ([]string, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *RebootNetworkSmDevicesRequest) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given []string and assigns it to the Scope field.
func (o *RebootNetworkSmDevicesRequest) SetScope(v []string) {
	o.Scope = v
}

// GetKextPaths returns the KextPaths field value if set, zero value otherwise.
func (o *RebootNetworkSmDevicesRequest) GetKextPaths() []string {
	if o == nil || IsNil(o.KextPaths) {
		var ret []string
		return ret
	}
	return o.KextPaths
}

// GetKextPathsOk returns a tuple with the KextPaths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RebootNetworkSmDevicesRequest) GetKextPathsOk() ([]string, bool) {
	if o == nil || IsNil(o.KextPaths) {
		return nil, false
	}
	return o.KextPaths, true
}

// HasKextPaths returns a boolean if a field has been set.
func (o *RebootNetworkSmDevicesRequest) HasKextPaths() bool {
	if o != nil && !IsNil(o.KextPaths) {
		return true
	}

	return false
}

// SetKextPaths gets a reference to the given []string and assigns it to the KextPaths field.
func (o *RebootNetworkSmDevicesRequest) SetKextPaths(v []string) {
	o.KextPaths = v
}

// GetNotifyUser returns the NotifyUser field value if set, zero value otherwise.
func (o *RebootNetworkSmDevicesRequest) GetNotifyUser() bool {
	if o == nil || IsNil(o.NotifyUser) {
		var ret bool
		return ret
	}
	return *o.NotifyUser
}

// GetNotifyUserOk returns a tuple with the NotifyUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RebootNetworkSmDevicesRequest) GetNotifyUserOk() (*bool, bool) {
	if o == nil || IsNil(o.NotifyUser) {
		return nil, false
	}
	return o.NotifyUser, true
}

// HasNotifyUser returns a boolean if a field has been set.
func (o *RebootNetworkSmDevicesRequest) HasNotifyUser() bool {
	if o != nil && !IsNil(o.NotifyUser) {
		return true
	}

	return false
}

// SetNotifyUser gets a reference to the given bool and assigns it to the NotifyUser field.
func (o *RebootNetworkSmDevicesRequest) SetNotifyUser(v bool) {
	o.NotifyUser = &v
}

// GetRebuildKernelCache returns the RebuildKernelCache field value if set, zero value otherwise.
func (o *RebootNetworkSmDevicesRequest) GetRebuildKernelCache() bool {
	if o == nil || IsNil(o.RebuildKernelCache) {
		var ret bool
		return ret
	}
	return *o.RebuildKernelCache
}

// GetRebuildKernelCacheOk returns a tuple with the RebuildKernelCache field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RebootNetworkSmDevicesRequest) GetRebuildKernelCacheOk() (*bool, bool) {
	if o == nil || IsNil(o.RebuildKernelCache) {
		return nil, false
	}
	return o.RebuildKernelCache, true
}

// HasRebuildKernelCache returns a boolean if a field has been set.
func (o *RebootNetworkSmDevicesRequest) HasRebuildKernelCache() bool {
	if o != nil && !IsNil(o.RebuildKernelCache) {
		return true
	}

	return false
}

// SetRebuildKernelCache gets a reference to the given bool and assigns it to the RebuildKernelCache field.
func (o *RebootNetworkSmDevicesRequest) SetRebuildKernelCache(v bool) {
	o.RebuildKernelCache = &v
}

// GetRequestRequiresNetworkTether returns the RequestRequiresNetworkTether field value if set, zero value otherwise.
func (o *RebootNetworkSmDevicesRequest) GetRequestRequiresNetworkTether() bool {
	if o == nil || IsNil(o.RequestRequiresNetworkTether) {
		var ret bool
		return ret
	}
	return *o.RequestRequiresNetworkTether
}

// GetRequestRequiresNetworkTetherOk returns a tuple with the RequestRequiresNetworkTether field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RebootNetworkSmDevicesRequest) GetRequestRequiresNetworkTetherOk() (*bool, bool) {
	if o == nil || IsNil(o.RequestRequiresNetworkTether) {
		return nil, false
	}
	return o.RequestRequiresNetworkTether, true
}

// HasRequestRequiresNetworkTether returns a boolean if a field has been set.
func (o *RebootNetworkSmDevicesRequest) HasRequestRequiresNetworkTether() bool {
	if o != nil && !IsNil(o.RequestRequiresNetworkTether) {
		return true
	}

	return false
}

// SetRequestRequiresNetworkTether gets a reference to the given bool and assigns it to the RequestRequiresNetworkTether field.
func (o *RebootNetworkSmDevicesRequest) SetRequestRequiresNetworkTether(v bool) {
	o.RequestRequiresNetworkTether = &v
}

func (o RebootNetworkSmDevicesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RebootNetworkSmDevicesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.WifiMacs) {
		toSerialize["wifiMacs"] = o.WifiMacs
	}
	if !IsNil(o.Ids) {
		toSerialize["ids"] = o.Ids
	}
	if !IsNil(o.Serials) {
		toSerialize["serials"] = o.Serials
	}
	if !IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	if !IsNil(o.KextPaths) {
		toSerialize["kextPaths"] = o.KextPaths
	}
	if !IsNil(o.NotifyUser) {
		toSerialize["notifyUser"] = o.NotifyUser
	}
	if !IsNil(o.RebuildKernelCache) {
		toSerialize["rebuildKernelCache"] = o.RebuildKernelCache
	}
	if !IsNil(o.RequestRequiresNetworkTether) {
		toSerialize["requestRequiresNetworkTether"] = o.RequestRequiresNetworkTether
	}
	return toSerialize, nil
}

type NullableRebootNetworkSmDevicesRequest struct {
	value *RebootNetworkSmDevicesRequest
	isSet bool
}

func (v NullableRebootNetworkSmDevicesRequest) Get() *RebootNetworkSmDevicesRequest {
	return v.value
}

func (v *NullableRebootNetworkSmDevicesRequest) Set(val *RebootNetworkSmDevicesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRebootNetworkSmDevicesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRebootNetworkSmDevicesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRebootNetworkSmDevicesRequest(val *RebootNetworkSmDevicesRequest) *NullableRebootNetworkSmDevicesRequest {
	return &NullableRebootNetworkSmDevicesRequest{value: val, isSet: true}
}

func (v NullableRebootNetworkSmDevicesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRebootNetworkSmDevicesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


