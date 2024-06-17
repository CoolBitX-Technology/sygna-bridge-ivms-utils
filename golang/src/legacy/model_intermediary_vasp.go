/*
 * Bridge
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ivms101

import (
	"encoding/json"
)

// IntermediaryVasp struct for IntermediaryVasp
type IntermediaryVasp struct {
	IntermediaryVasp *Person `json:"intermediary_vasp,omitempty"`
	Sequence *int32 `json:"sequence,omitempty"`
}

// NewIntermediaryVasp instantiates a new IntermediaryVasp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntermediaryVasp() *IntermediaryVasp {
	this := IntermediaryVasp{}
	return &this
}

// NewIntermediaryVaspWithDefaults instantiates a new IntermediaryVasp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntermediaryVaspWithDefaults() *IntermediaryVasp {
	this := IntermediaryVasp{}
	return &this
}

// GetIntermediaryVasp returns the IntermediaryVasp field value if set, zero value otherwise.
func (o *IntermediaryVasp) GetIntermediaryVasp() Person {
	if o == nil || o.IntermediaryVasp == nil {
		var ret Person
		return ret
	}
	return *o.IntermediaryVasp
}

// GetIntermediaryVaspOk returns a tuple with the IntermediaryVasp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntermediaryVasp) GetIntermediaryVaspOk() (*Person, bool) {
	if o == nil || o.IntermediaryVasp == nil {
		return nil, false
	}
	return o.IntermediaryVasp, true
}

// HasIntermediaryVasp returns a boolean if a field has been set.
func (o *IntermediaryVasp) HasIntermediaryVasp() bool {
	if o != nil && o.IntermediaryVasp != nil {
		return true
	}

	return false
}

// SetIntermediaryVasp gets a reference to the given Person and assigns it to the IntermediaryVasp field.
func (o *IntermediaryVasp) SetIntermediaryVasp(v Person) {
	o.IntermediaryVasp = &v
}

// GetSequence returns the Sequence field value if set, zero value otherwise.
func (o *IntermediaryVasp) GetSequence() int32 {
	if o == nil || o.Sequence == nil {
		var ret int32
		return ret
	}
	return *o.Sequence
}

// GetSequenceOk returns a tuple with the Sequence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntermediaryVasp) GetSequenceOk() (*int32, bool) {
	if o == nil || o.Sequence == nil {
		return nil, false
	}
	return o.Sequence, true
}

// HasSequence returns a boolean if a field has been set.
func (o *IntermediaryVasp) HasSequence() bool {
	if o != nil && o.Sequence != nil {
		return true
	}

	return false
}

// SetSequence gets a reference to the given int32 and assigns it to the Sequence field.
func (o *IntermediaryVasp) SetSequence(v int32) {
	o.Sequence = &v
}

func (o IntermediaryVasp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IntermediaryVasp != nil {
		toSerialize["intermediary_vasp"] = o.IntermediaryVasp
	}
	if o.Sequence != nil {
		toSerialize["sequence"] = o.Sequence
	}
	return json.Marshal(toSerialize)
}

type NullableIntermediaryVasp struct {
	value *IntermediaryVasp
	isSet bool
}

func (v NullableIntermediaryVasp) Get() *IntermediaryVasp {
	return v.value
}

func (v *NullableIntermediaryVasp) Set(val *IntermediaryVasp) {
	v.value = val
	v.isSet = true
}

func (v NullableIntermediaryVasp) IsSet() bool {
	return v.isSet
}

func (v *NullableIntermediaryVasp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntermediaryVasp(val *IntermediaryVasp) *NullableIntermediaryVasp {
	return &NullableIntermediaryVasp{value: val, isSet: true}
}

func (v NullableIntermediaryVasp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntermediaryVasp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

