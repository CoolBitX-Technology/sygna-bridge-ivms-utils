/*
Bridge

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ivms101

import (
	"encoding/json"
)

// checks if the LegalPersonNameId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LegalPersonNameId{}

// LegalPersonNameId struct for LegalPersonNameId
type LegalPersonNameId struct {
	LegalPersonName *string `json:"legalPersonName,omitempty"`
	LegalPersonNameIdentifierType *string `json:"legalPersonNameIdentifierType,omitempty"`
}

// NewLegalPersonNameId instantiates a new LegalPersonNameId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLegalPersonNameId() *LegalPersonNameId {
	this := LegalPersonNameId{}
	return &this
}

// NewLegalPersonNameIdWithDefaults instantiates a new LegalPersonNameId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLegalPersonNameIdWithDefaults() *LegalPersonNameId {
	this := LegalPersonNameId{}
	return &this
}

// GetLegalPersonName returns the LegalPersonName field value if set, zero value otherwise.
func (o *LegalPersonNameId) GetLegalPersonName() string {
	if o == nil || IsNil(o.LegalPersonName) {
		var ret string
		return ret
	}
	return *o.LegalPersonName
}

// GetLegalPersonNameOk returns a tuple with the LegalPersonName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegalPersonNameId) GetLegalPersonNameOk() (*string, bool) {
	if o == nil || IsNil(o.LegalPersonName) {
		return nil, false
	}
	return o.LegalPersonName, true
}

// HasLegalPersonName returns a boolean if a field has been set.
func (o *LegalPersonNameId) HasLegalPersonName() bool {
	if o != nil && !IsNil(o.LegalPersonName) {
		return true
	}

	return false
}

// SetLegalPersonName gets a reference to the given string and assigns it to the LegalPersonName field.
func (o *LegalPersonNameId) SetLegalPersonName(v string) {
	o.LegalPersonName = &v
}

// GetLegalPersonNameIdentifierType returns the LegalPersonNameIdentifierType field value if set, zero value otherwise.
func (o *LegalPersonNameId) GetLegalPersonNameIdentifierType() string {
	if o == nil || IsNil(o.LegalPersonNameIdentifierType) {
		var ret string
		return ret
	}
	return *o.LegalPersonNameIdentifierType
}

// GetLegalPersonNameIdentifierTypeOk returns a tuple with the LegalPersonNameIdentifierType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegalPersonNameId) GetLegalPersonNameIdentifierTypeOk() (*string, bool) {
	if o == nil || IsNil(o.LegalPersonNameIdentifierType) {
		return nil, false
	}
	return o.LegalPersonNameIdentifierType, true
}

// HasLegalPersonNameIdentifierType returns a boolean if a field has been set.
func (o *LegalPersonNameId) HasLegalPersonNameIdentifierType() bool {
	if o != nil && !IsNil(o.LegalPersonNameIdentifierType) {
		return true
	}

	return false
}

// SetLegalPersonNameIdentifierType gets a reference to the given string and assigns it to the LegalPersonNameIdentifierType field.
func (o *LegalPersonNameId) SetLegalPersonNameIdentifierType(v string) {
	o.LegalPersonNameIdentifierType = &v
}

func (o LegalPersonNameId) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LegalPersonNameId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LegalPersonName) {
		toSerialize["legalPersonName"] = o.LegalPersonName
	}
	if !IsNil(o.LegalPersonNameIdentifierType) {
		toSerialize["legalPersonNameIdentifierType"] = o.LegalPersonNameIdentifierType
	}
	return toSerialize, nil
}

type NullableLegalPersonNameId struct {
	value *LegalPersonNameId
	isSet bool
}

func (v NullableLegalPersonNameId) Get() *LegalPersonNameId {
	return v.value
}

func (v *NullableLegalPersonNameId) Set(val *LegalPersonNameId) {
	v.value = val
	v.isSet = true
}

func (v NullableLegalPersonNameId) IsSet() bool {
	return v.isSet
}

func (v *NullableLegalPersonNameId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLegalPersonNameId(val *LegalPersonNameId) *NullableLegalPersonNameId {
	return &NullableLegalPersonNameId{value: val, isSet: true}
}

func (v NullableLegalPersonNameId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLegalPersonNameId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

