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

// RoutesOrgUnitOrgUnitKind struct for RoutesOrgUnitOrgUnitKind
type RoutesOrgUnitOrgUnitKind struct {
	SubBlock *[]RoutesOrgUnitOrgUnitKindSubBlock `json:"subBlock,omitempty"`
	Text *string `json:"text,omitempty"`
}

// NewRoutesOrgUnitOrgUnitKind instantiates a new RoutesOrgUnitOrgUnitKind object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoutesOrgUnitOrgUnitKind() *RoutesOrgUnitOrgUnitKind {
	this := RoutesOrgUnitOrgUnitKind{}
	return &this
}

// NewRoutesOrgUnitOrgUnitKindWithDefaults instantiates a new RoutesOrgUnitOrgUnitKind object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoutesOrgUnitOrgUnitKindWithDefaults() *RoutesOrgUnitOrgUnitKind {
	this := RoutesOrgUnitOrgUnitKind{}
	return &this
}

// GetSubBlock returns the SubBlock field value if set, zero value otherwise.
func (o *RoutesOrgUnitOrgUnitKind) GetSubBlock() []RoutesOrgUnitOrgUnitKindSubBlock {
	if o == nil || o.SubBlock == nil {
		var ret []RoutesOrgUnitOrgUnitKindSubBlock
		return ret
	}
	return *o.SubBlock
}

// GetSubBlockOk returns a tuple with the SubBlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutesOrgUnitOrgUnitKind) GetSubBlockOk() (*[]RoutesOrgUnitOrgUnitKindSubBlock, bool) {
	if o == nil || o.SubBlock == nil {
		return nil, false
	}
	return o.SubBlock, true
}

// HasSubBlock returns a boolean if a field has been set.
func (o *RoutesOrgUnitOrgUnitKind) HasSubBlock() bool {
	if o != nil && o.SubBlock != nil {
		return true
	}

	return false
}

// SetSubBlock gets a reference to the given []RoutesOrgUnitOrgUnitKindSubBlock and assigns it to the SubBlock field.
func (o *RoutesOrgUnitOrgUnitKind) SetSubBlock(v []RoutesOrgUnitOrgUnitKindSubBlock) {
	o.SubBlock = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *RoutesOrgUnitOrgUnitKind) GetText() string {
	if o == nil || o.Text == nil {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutesOrgUnitOrgUnitKind) GetTextOk() (*string, bool) {
	if o == nil || o.Text == nil {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *RoutesOrgUnitOrgUnitKind) HasText() bool {
	if o != nil && o.Text != nil {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *RoutesOrgUnitOrgUnitKind) SetText(v string) {
	o.Text = &v
}

func (o RoutesOrgUnitOrgUnitKind) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SubBlock != nil {
		toSerialize["subBlock"] = o.SubBlock
	}
	if o.Text != nil {
		toSerialize["text"] = o.Text
	}
	return json.Marshal(toSerialize)
}

type NullableRoutesOrgUnitOrgUnitKind struct {
	value *RoutesOrgUnitOrgUnitKind
	isSet bool
}

func (v NullableRoutesOrgUnitOrgUnitKind) Get() *RoutesOrgUnitOrgUnitKind {
	return v.value
}

func (v *NullableRoutesOrgUnitOrgUnitKind) Set(val *RoutesOrgUnitOrgUnitKind) {
	v.value = val
	v.isSet = true
}

func (v NullableRoutesOrgUnitOrgUnitKind) IsSet() bool {
	return v.isSet
}

func (v *NullableRoutesOrgUnitOrgUnitKind) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoutesOrgUnitOrgUnitKind(val *RoutesOrgUnitOrgUnitKind) *NullableRoutesOrgUnitOrgUnitKind {
	return &NullableRoutesOrgUnitOrgUnitKind{value: val, isSet: true}
}

func (v NullableRoutesOrgUnitOrgUnitKind) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoutesOrgUnitOrgUnitKind) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


