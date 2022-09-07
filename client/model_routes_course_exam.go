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

// RoutesCourseExam struct for RoutesCourseExam
type RoutesCourseExam struct {
	InfoBlock *RoutesInfoBlock `json:"info_block,omitempty"`
}

// NewRoutesCourseExam instantiates a new RoutesCourseExam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoutesCourseExam() *RoutesCourseExam {
	this := RoutesCourseExam{}
	return &this
}

// NewRoutesCourseExamWithDefaults instantiates a new RoutesCourseExam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoutesCourseExamWithDefaults() *RoutesCourseExam {
	this := RoutesCourseExam{}
	return &this
}

// GetInfoBlock returns the InfoBlock field value if set, zero value otherwise.
func (o *RoutesCourseExam) GetInfoBlock() RoutesInfoBlock {
	if o == nil || o.InfoBlock == nil {
		var ret RoutesInfoBlock
		return ret
	}
	return *o.InfoBlock
}

// GetInfoBlockOk returns a tuple with the InfoBlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutesCourseExam) GetInfoBlockOk() (*RoutesInfoBlock, bool) {
	if o == nil || o.InfoBlock == nil {
		return nil, false
	}
	return o.InfoBlock, true
}

// HasInfoBlock returns a boolean if a field has been set.
func (o *RoutesCourseExam) HasInfoBlock() bool {
	if o != nil && o.InfoBlock != nil {
		return true
	}

	return false
}

// SetInfoBlock gets a reference to the given RoutesInfoBlock and assigns it to the InfoBlock field.
func (o *RoutesCourseExam) SetInfoBlock(v RoutesInfoBlock) {
	o.InfoBlock = &v
}

func (o RoutesCourseExam) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.InfoBlock != nil {
		toSerialize["info_block"] = o.InfoBlock
	}
	return json.Marshal(toSerialize)
}

type NullableRoutesCourseExam struct {
	value *RoutesCourseExam
	isSet bool
}

func (v NullableRoutesCourseExam) Get() *RoutesCourseExam {
	return v.value
}

func (v *NullableRoutesCourseExam) Set(val *RoutesCourseExam) {
	v.value = val
	v.isSet = true
}

func (v NullableRoutesCourseExam) IsSet() bool {
	return v.isSet
}

func (v *NullableRoutesCourseExam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoutesCourseExam(val *RoutesCourseExam) *NullableRoutesCourseExam {
	return &NullableRoutesCourseExam{value: val, isSet: true}
}

func (v NullableRoutesCourseExam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoutesCourseExam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


