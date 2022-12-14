/*
 * CAMPUSOnline Webservice proxy
 *
 * This is the proxy server for CAMPUSOnline Webservices, enabling a nicely documented and uniform rest access to CAMPUSOnline resources.
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// RoutesPersonName struct for RoutesPersonName
type RoutesPersonName struct {
	Family *string `json:"family,omitempty"`
	Given *string `json:"given,omitempty"`
}

// NewRoutesPersonName instantiates a new RoutesPersonName object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoutesPersonName() *RoutesPersonName {
	this := RoutesPersonName{}
	return &this
}

// NewRoutesPersonNameWithDefaults instantiates a new RoutesPersonName object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoutesPersonNameWithDefaults() *RoutesPersonName {
	this := RoutesPersonName{}
	return &this
}

// GetFamily returns the Family field value if set, zero value otherwise.
func (o *RoutesPersonName) GetFamily() string {
	if o == nil || o.Family == nil {
		var ret string
		return ret
	}
	return *o.Family
}

// GetFamilyOk returns a tuple with the Family field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutesPersonName) GetFamilyOk() (*string, bool) {
	if o == nil || o.Family == nil {
		return nil, false
	}
	return o.Family, true
}

// HasFamily returns a boolean if a field has been set.
func (o *RoutesPersonName) HasFamily() bool {
	if o != nil && o.Family != nil {
		return true
	}

	return false
}

// SetFamily gets a reference to the given string and assigns it to the Family field.
func (o *RoutesPersonName) SetFamily(v string) {
	o.Family = &v
}

// GetGiven returns the Given field value if set, zero value otherwise.
func (o *RoutesPersonName) GetGiven() string {
	if o == nil || o.Given == nil {
		var ret string
		return ret
	}
	return *o.Given
}

// GetGivenOk returns a tuple with the Given field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutesPersonName) GetGivenOk() (*string, bool) {
	if o == nil || o.Given == nil {
		return nil, false
	}
	return o.Given, true
}

// HasGiven returns a boolean if a field has been set.
func (o *RoutesPersonName) HasGiven() bool {
	if o != nil && o.Given != nil {
		return true
	}

	return false
}

// SetGiven gets a reference to the given string and assigns it to the Given field.
func (o *RoutesPersonName) SetGiven(v string) {
	o.Given = &v
}

func (o RoutesPersonName) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Family != nil {
		toSerialize["family"] = o.Family
	}
	if o.Given != nil {
		toSerialize["given"] = o.Given
	}
	return json.Marshal(toSerialize)
}

type NullableRoutesPersonName struct {
	value *RoutesPersonName
	isSet bool
}

func (v NullableRoutesPersonName) Get() *RoutesPersonName {
	return v.value
}

func (v *NullableRoutesPersonName) Set(val *RoutesPersonName) {
	v.value = val
	v.isSet = true
}

func (v NullableRoutesPersonName) IsSet() bool {
	return v.isSet
}

func (v *NullableRoutesPersonName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoutesPersonName(val *RoutesPersonName) *NullableRoutesPersonName {
	return &NullableRoutesPersonName{value: val, isSet: true}
}

func (v NullableRoutesPersonName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoutesPersonName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


