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

// checks if the UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFrontImage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFrontImage{}

// UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFrontImage Properties for setting a new image.
type UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFrontImage struct {
	// The format of the encoded contents. Supported formats are 'png', 'gif', and jpg'.
	Format *string `json:"format,omitempty"`
	// The file contents (a base 64 encoded string) of your new prepaid front.
	Contents *string `json:"contents,omitempty"`
}

// NewUpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFrontImage instantiates a new UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFrontImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFrontImage() *UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFrontImage {
	this := UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFrontImage{}
	return &this
}

// NewUpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFrontImageWithDefaults instantiates a new UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFrontImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFrontImageWithDefaults() *UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFrontImage {
	this := UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFrontImage{}
	return &this
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFrontImage) GetFormat() string {
	if o == nil || IsNil(o.Format) {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFrontImage) GetFormatOk() (*string, bool) {
	if o == nil || IsNil(o.Format) {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFrontImage) HasFormat() bool {
	if o != nil && !IsNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFrontImage) SetFormat(v string) {
	o.Format = &v
}

// GetContents returns the Contents field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFrontImage) GetContents() string {
	if o == nil || IsNil(o.Contents) {
		var ret string
		return ret
	}
	return *o.Contents
}

// GetContentsOk returns a tuple with the Contents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFrontImage) GetContentsOk() (*string, bool) {
	if o == nil || IsNil(o.Contents) {
		return nil, false
	}
	return o.Contents, true
}

// HasContents returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFrontImage) HasContents() bool {
	if o != nil && !IsNil(o.Contents) {
		return true
	}

	return false
}

// SetContents gets a reference to the given string and assigns it to the Contents field.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFrontImage) SetContents(v string) {
	o.Contents = &v
}

func (o UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFrontImage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFrontImage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Format) {
		toSerialize["format"] = o.Format
	}
	if !IsNil(o.Contents) {
		toSerialize["contents"] = o.Contents
	}
	return toSerialize, nil
}

type NullableUpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFrontImage struct {
	value *UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFrontImage
	isSet bool
}

func (v NullableUpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFrontImage) Get() *UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFrontImage {
	return v.value
}

func (v *NullableUpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFrontImage) Set(val *UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFrontImage) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFrontImage) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFrontImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFrontImage(val *UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFrontImage) *NullableUpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFrontImage {
	return &NullableUpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFrontImage{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFrontImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFrontImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


