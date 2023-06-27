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

// UpdateNetworkWebhooksHttpServerRequest struct for UpdateNetworkWebhooksHttpServerRequest
type UpdateNetworkWebhooksHttpServerRequest struct {
	// A name for easy reference to the HTTP server
	Name *string `json:"name,omitempty"`
	// A shared secret that will be included in POSTs sent to the HTTP server. This secret can be used to verify that the request was sent by Meraki.
	SharedSecret *string `json:"sharedSecret,omitempty"`
	PayloadTemplate *UpdateNetworkWebhooksHttpServerRequestPayloadTemplate `json:"payloadTemplate,omitempty"`
}

// NewUpdateNetworkWebhooksHttpServerRequest instantiates a new UpdateNetworkWebhooksHttpServerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWebhooksHttpServerRequest() *UpdateNetworkWebhooksHttpServerRequest {
	this := UpdateNetworkWebhooksHttpServerRequest{}
	return &this
}

// NewUpdateNetworkWebhooksHttpServerRequestWithDefaults instantiates a new UpdateNetworkWebhooksHttpServerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWebhooksHttpServerRequestWithDefaults() *UpdateNetworkWebhooksHttpServerRequest {
	this := UpdateNetworkWebhooksHttpServerRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateNetworkWebhooksHttpServerRequest) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWebhooksHttpServerRequest) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateNetworkWebhooksHttpServerRequest) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateNetworkWebhooksHttpServerRequest) SetName(v string) {
	o.Name = &v
}

// GetSharedSecret returns the SharedSecret field value if set, zero value otherwise.
func (o *UpdateNetworkWebhooksHttpServerRequest) GetSharedSecret() string {
	if o == nil || isNil(o.SharedSecret) {
		var ret string
		return ret
	}
	return *o.SharedSecret
}

// GetSharedSecretOk returns a tuple with the SharedSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWebhooksHttpServerRequest) GetSharedSecretOk() (*string, bool) {
	if o == nil || isNil(o.SharedSecret) {
    return nil, false
	}
	return o.SharedSecret, true
}

// HasSharedSecret returns a boolean if a field has been set.
func (o *UpdateNetworkWebhooksHttpServerRequest) HasSharedSecret() bool {
	if o != nil && !isNil(o.SharedSecret) {
		return true
	}

	return false
}

// SetSharedSecret gets a reference to the given string and assigns it to the SharedSecret field.
func (o *UpdateNetworkWebhooksHttpServerRequest) SetSharedSecret(v string) {
	o.SharedSecret = &v
}

// GetPayloadTemplate returns the PayloadTemplate field value if set, zero value otherwise.
func (o *UpdateNetworkWebhooksHttpServerRequest) GetPayloadTemplate() UpdateNetworkWebhooksHttpServerRequestPayloadTemplate {
	if o == nil || isNil(o.PayloadTemplate) {
		var ret UpdateNetworkWebhooksHttpServerRequestPayloadTemplate
		return ret
	}
	return *o.PayloadTemplate
}

// GetPayloadTemplateOk returns a tuple with the PayloadTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWebhooksHttpServerRequest) GetPayloadTemplateOk() (*UpdateNetworkWebhooksHttpServerRequestPayloadTemplate, bool) {
	if o == nil || isNil(o.PayloadTemplate) {
    return nil, false
	}
	return o.PayloadTemplate, true
}

// HasPayloadTemplate returns a boolean if a field has been set.
func (o *UpdateNetworkWebhooksHttpServerRequest) HasPayloadTemplate() bool {
	if o != nil && !isNil(o.PayloadTemplate) {
		return true
	}

	return false
}

// SetPayloadTemplate gets a reference to the given UpdateNetworkWebhooksHttpServerRequestPayloadTemplate and assigns it to the PayloadTemplate field.
func (o *UpdateNetworkWebhooksHttpServerRequest) SetPayloadTemplate(v UpdateNetworkWebhooksHttpServerRequestPayloadTemplate) {
	o.PayloadTemplate = &v
}

func (o UpdateNetworkWebhooksHttpServerRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.SharedSecret) {
		toSerialize["sharedSecret"] = o.SharedSecret
	}
	if !isNil(o.PayloadTemplate) {
		toSerialize["payloadTemplate"] = o.PayloadTemplate
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkWebhooksHttpServerRequest struct {
	value *UpdateNetworkWebhooksHttpServerRequest
	isSet bool
}

func (v NullableUpdateNetworkWebhooksHttpServerRequest) Get() *UpdateNetworkWebhooksHttpServerRequest {
	return v.value
}

func (v *NullableUpdateNetworkWebhooksHttpServerRequest) Set(val *UpdateNetworkWebhooksHttpServerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWebhooksHttpServerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWebhooksHttpServerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWebhooksHttpServerRequest(val *UpdateNetworkWebhooksHttpServerRequest) *NullableUpdateNetworkWebhooksHttpServerRequest {
	return &NullableUpdateNetworkWebhooksHttpServerRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkWebhooksHttpServerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWebhooksHttpServerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


