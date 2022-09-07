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

// RoutesPersonRole struct for RoutesPersonRole
type RoutesPersonRole struct {
	RoleId *string `json:"role_id,omitempty"`
	Text *string `json:"text,omitempty"`
}

// NewRoutesPersonRole instantiates a new RoutesPersonRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoutesPersonRole() *RoutesPersonRole {
	this := RoutesPersonRole{}
	return &this
}

// NewRoutesPersonRoleWithDefaults instantiates a new RoutesPersonRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoutesPersonRoleWithDefaults() *RoutesPersonRole {
	this := RoutesPersonRole{}
	return &this
}

// GetRoleId returns the RoleId field value if set, zero value otherwise.
func (o *RoutesPersonRole) GetRoleId() string {
	if o == nil || o.RoleId == nil {
		var ret string
		return ret
	}
	return *o.RoleId
}

// GetRoleIdOk returns a tuple with the RoleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutesPersonRole) GetRoleIdOk() (*string, bool) {
	if o == nil || o.RoleId == nil {
		return nil, false
	}
	return o.RoleId, true
}

// HasRoleId returns a boolean if a field has been set.
func (o *RoutesPersonRole) HasRoleId() bool {
	if o != nil && o.RoleId != nil {
		return true
	}

	return false
}

// SetRoleId gets a reference to the given string and assigns it to the RoleId field.
func (o *RoutesPersonRole) SetRoleId(v string) {
	o.RoleId = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *RoutesPersonRole) GetText() string {
	if o == nil || o.Text == nil {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutesPersonRole) GetTextOk() (*string, bool) {
	if o == nil || o.Text == nil {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *RoutesPersonRole) HasText() bool {
	if o != nil && o.Text != nil {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *RoutesPersonRole) SetText(v string) {
	o.Text = &v
}

func (o RoutesPersonRole) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RoleId != nil {
		toSerialize["role_id"] = o.RoleId
	}
	if o.Text != nil {
		toSerialize["text"] = o.Text
	}
	return json.Marshal(toSerialize)
}

type NullableRoutesPersonRole struct {
	value *RoutesPersonRole
	isSet bool
}

func (v NullableRoutesPersonRole) Get() *RoutesPersonRole {
	return v.value
}

func (v *NullableRoutesPersonRole) Set(val *RoutesPersonRole) {
	v.value = val
	v.isSet = true
}

func (v NullableRoutesPersonRole) IsSet() bool {
	return v.isSet
}

func (v *NullableRoutesPersonRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoutesPersonRole(val *RoutesPersonRole) *NullableRoutesPersonRole {
	return &NullableRoutesPersonRole{value: val, isSet: true}
}

func (v NullableRoutesPersonRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoutesPersonRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


