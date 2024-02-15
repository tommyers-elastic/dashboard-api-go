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

// checks if the GetOrganizationSummaryTopDevicesByUsage200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationSummaryTopDevicesByUsage200ResponseInner{}

// GetOrganizationSummaryTopDevicesByUsage200ResponseInner struct for GetOrganizationSummaryTopDevicesByUsage200ResponseInner
type GetOrganizationSummaryTopDevicesByUsage200ResponseInner struct {
	// Name of the device
	Name *string `json:"name,omitempty"`
	// Model of the device
	Model *string `json:"model,omitempty"`
	// Serial number of the device
	Serial *string `json:"serial,omitempty"`
	// Mac address of the device
	Mac *string `json:"mac,omitempty"`
	// Product type of the device
	ProductType *string `json:"productType,omitempty"`
	Network *GetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerNetwork `json:"network,omitempty"`
	Usage *GetOrganizationSummaryTopDevicesByUsage200ResponseInnerUsage `json:"usage,omitempty"`
	Clients *GetOrganizationSummaryTopDevicesByUsage200ResponseInnerClients `json:"clients,omitempty"`
}

// NewGetOrganizationSummaryTopDevicesByUsage200ResponseInner instantiates a new GetOrganizationSummaryTopDevicesByUsage200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationSummaryTopDevicesByUsage200ResponseInner() *GetOrganizationSummaryTopDevicesByUsage200ResponseInner {
	this := GetOrganizationSummaryTopDevicesByUsage200ResponseInner{}
	return &this
}

// NewGetOrganizationSummaryTopDevicesByUsage200ResponseInnerWithDefaults instantiates a new GetOrganizationSummaryTopDevicesByUsage200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationSummaryTopDevicesByUsage200ResponseInnerWithDefaults() *GetOrganizationSummaryTopDevicesByUsage200ResponseInner {
	this := GetOrganizationSummaryTopDevicesByUsage200ResponseInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) GetModel() string {
	if o == nil || IsNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) GetModelOk() (*string, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) HasModel() bool {
	if o != nil && !IsNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) SetModel(v string) {
	o.Model = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) SetSerial(v string) {
	o.Serial = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) GetMac() string {
	if o == nil || IsNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) GetMacOk() (*string, bool) {
	if o == nil || IsNil(o.Mac) {
		return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) HasMac() bool {
	if o != nil && !IsNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) SetMac(v string) {
	o.Mac = &v
}

// GetProductType returns the ProductType field value if set, zero value otherwise.
func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) GetProductType() string {
	if o == nil || IsNil(o.ProductType) {
		var ret string
		return ret
	}
	return *o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) GetProductTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ProductType) {
		return nil, false
	}
	return o.ProductType, true
}

// HasProductType returns a boolean if a field has been set.
func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) HasProductType() bool {
	if o != nil && !IsNil(o.ProductType) {
		return true
	}

	return false
}

// SetProductType gets a reference to the given string and assigns it to the ProductType field.
func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) SetProductType(v string) {
	o.ProductType = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) GetNetwork() GetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerNetwork {
	if o == nil || IsNil(o.Network) {
		var ret GetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerNetwork
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) GetNetworkOk() (*GetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerNetwork, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given GetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerNetwork and assigns it to the Network field.
func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) SetNetwork(v GetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerNetwork) {
	o.Network = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) GetUsage() GetOrganizationSummaryTopDevicesByUsage200ResponseInnerUsage {
	if o == nil || IsNil(o.Usage) {
		var ret GetOrganizationSummaryTopDevicesByUsage200ResponseInnerUsage
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) GetUsageOk() (*GetOrganizationSummaryTopDevicesByUsage200ResponseInnerUsage, bool) {
	if o == nil || IsNil(o.Usage) {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) HasUsage() bool {
	if o != nil && !IsNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given GetOrganizationSummaryTopDevicesByUsage200ResponseInnerUsage and assigns it to the Usage field.
func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) SetUsage(v GetOrganizationSummaryTopDevicesByUsage200ResponseInnerUsage) {
	o.Usage = &v
}

// GetClients returns the Clients field value if set, zero value otherwise.
func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) GetClients() GetOrganizationSummaryTopDevicesByUsage200ResponseInnerClients {
	if o == nil || IsNil(o.Clients) {
		var ret GetOrganizationSummaryTopDevicesByUsage200ResponseInnerClients
		return ret
	}
	return *o.Clients
}

// GetClientsOk returns a tuple with the Clients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) GetClientsOk() (*GetOrganizationSummaryTopDevicesByUsage200ResponseInnerClients, bool) {
	if o == nil || IsNil(o.Clients) {
		return nil, false
	}
	return o.Clients, true
}

// HasClients returns a boolean if a field has been set.
func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) HasClients() bool {
	if o != nil && !IsNil(o.Clients) {
		return true
	}

	return false
}

// SetClients gets a reference to the given GetOrganizationSummaryTopDevicesByUsage200ResponseInnerClients and assigns it to the Clients field.
func (o *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) SetClients(v GetOrganizationSummaryTopDevicesByUsage200ResponseInnerClients) {
	o.Clients = &v
}

func (o GetOrganizationSummaryTopDevicesByUsage200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationSummaryTopDevicesByUsage200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !IsNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !IsNil(o.ProductType) {
		toSerialize["productType"] = o.ProductType
	}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}
	if !IsNil(o.Clients) {
		toSerialize["clients"] = o.Clients
	}
	return toSerialize, nil
}

type NullableGetOrganizationSummaryTopDevicesByUsage200ResponseInner struct {
	value *GetOrganizationSummaryTopDevicesByUsage200ResponseInner
	isSet bool
}

func (v NullableGetOrganizationSummaryTopDevicesByUsage200ResponseInner) Get() *GetOrganizationSummaryTopDevicesByUsage200ResponseInner {
	return v.value
}

func (v *NullableGetOrganizationSummaryTopDevicesByUsage200ResponseInner) Set(val *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationSummaryTopDevicesByUsage200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationSummaryTopDevicesByUsage200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationSummaryTopDevicesByUsage200ResponseInner(val *GetOrganizationSummaryTopDevicesByUsage200ResponseInner) *NullableGetOrganizationSummaryTopDevicesByUsage200ResponseInner {
	return &NullableGetOrganizationSummaryTopDevicesByUsage200ResponseInner{value: val, isSet: true}
}

func (v NullableGetOrganizationSummaryTopDevicesByUsage200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationSummaryTopDevicesByUsage200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


