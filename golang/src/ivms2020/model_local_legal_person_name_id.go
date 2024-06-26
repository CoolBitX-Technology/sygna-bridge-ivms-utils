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

// checks if the LocalLegalPersonNameId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LocalLegalPersonNameId{}

// LocalLegalPersonNameId struct for LocalLegalPersonNameId
type LocalLegalPersonNameId struct {
	LegalPersonName *string `json:"legalPersonName,omitempty"`
	LegalPersonNameIdentifierType *string `json:"legalPersonNameIdentifierType,omitempty"`
}

// NewLocalLegalPersonNameId instantiates a new LocalLegalPersonNameId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocalLegalPersonNameId() *LocalLegalPersonNameId {
	this := LocalLegalPersonNameId{}
	return &this
}

// NewLocalLegalPersonNameIdWithDefaults instantiates a new LocalLegalPersonNameId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocalLegalPersonNameIdWithDefaults() *LocalLegalPersonNameId {
	this := LocalLegalPersonNameId{}
	return &this
}

// GetLegalPersonName returns the LegalPersonName field value if set, zero value otherwise.
func (o *LocalLegalPersonNameId) GetLegalPersonName() string {
	if o == nil || IsNil(o.LegalPersonName) {
		var ret string
		return ret
	}
	return *o.LegalPersonName
}

// GetLegalPersonNameOk returns a tuple with the LegalPersonName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalLegalPersonNameId) GetLegalPersonNameOk() (*string, bool) {
	if o == nil || IsNil(o.LegalPersonName) {
		return nil, false
	}
	return o.LegalPersonName, true
}

// HasLegalPersonName returns a boolean if a field has been set.
func (o *LocalLegalPersonNameId) HasLegalPersonName() bool {
	if o != nil && !IsNil(o.LegalPersonName) {
		return true
	}

	return false
}

// SetLegalPersonName gets a reference to the given string and assigns it to the LegalPersonName field.
func (o *LocalLegalPersonNameId) SetLegalPersonName(v string) {
	o.LegalPersonName = &v
}

// GetLegalPersonNameIdentifierType returns the LegalPersonNameIdentifierType field value if set, zero value otherwise.
func (o *LocalLegalPersonNameId) GetLegalPersonNameIdentifierType() string {
	if o == nil || IsNil(o.LegalPersonNameIdentifierType) {
		var ret string
		return ret
	}
	return *o.LegalPersonNameIdentifierType
}

// GetLegalPersonNameIdentifierTypeOk returns a tuple with the LegalPersonNameIdentifierType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalLegalPersonNameId) GetLegalPersonNameIdentifierTypeOk() (*string, bool) {
	if o == nil || IsNil(o.LegalPersonNameIdentifierType) {
		return nil, false
	}
	return o.LegalPersonNameIdentifierType, true
}

// HasLegalPersonNameIdentifierType returns a boolean if a field has been set.
func (o *LocalLegalPersonNameId) HasLegalPersonNameIdentifierType() bool {
	if o != nil && !IsNil(o.LegalPersonNameIdentifierType) {
		return true
	}

	return false
}

// SetLegalPersonNameIdentifierType gets a reference to the given string and assigns it to the LegalPersonNameIdentifierType field.
func (o *LocalLegalPersonNameId) SetLegalPersonNameIdentifierType(v string) {
	o.LegalPersonNameIdentifierType = &v
}

func (o LocalLegalPersonNameId) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LocalLegalPersonNameId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LegalPersonName) {
		toSerialize["legalPersonName"] = o.LegalPersonName
	}
	if !IsNil(o.LegalPersonNameIdentifierType) {
		toSerialize["legalPersonNameIdentifierType"] = o.LegalPersonNameIdentifierType
	}
	return toSerialize, nil
}

type NullableLocalLegalPersonNameId struct {
	value *LocalLegalPersonNameId
	isSet bool
}

func (v NullableLocalLegalPersonNameId) Get() *LocalLegalPersonNameId {
	return v.value
}

func (v *NullableLocalLegalPersonNameId) Set(val *LocalLegalPersonNameId) {
	v.value = val
	v.isSet = true
}

func (v NullableLocalLegalPersonNameId) IsSet() bool {
	return v.isSet
}

func (v *NullableLocalLegalPersonNameId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocalLegalPersonNameId(val *LocalLegalPersonNameId) *NullableLocalLegalPersonNameId {
	return &NullableLocalLegalPersonNameId{value: val, isSet: true}
}

func (v NullableLocalLegalPersonNameId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocalLegalPersonNameId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


