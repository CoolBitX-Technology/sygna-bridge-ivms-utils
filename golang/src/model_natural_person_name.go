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

// NaturalPersonName struct for NaturalPersonName
type NaturalPersonName struct {
	NameIdentifiers *[]NaturalPersonNameId `json:"name_identifiers,omitempty"`
	LocalNameIdentifiers *[]LocalNaturalPersonNameId `json:"local_name_identifiers,omitempty"`
	PhoneticNameIdentifiers *[]PhoneticNaturalPersonNameId `json:"phonetic_name_identifiers,omitempty"`
}

// NewNaturalPersonName instantiates a new NaturalPersonName object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNaturalPersonName() *NaturalPersonName {
	this := NaturalPersonName{}
	return &this
}

// NewNaturalPersonNameWithDefaults instantiates a new NaturalPersonName object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNaturalPersonNameWithDefaults() *NaturalPersonName {
	this := NaturalPersonName{}
	return &this
}

// GetNameIdentifiers returns the NameIdentifiers field value if set, zero value otherwise.
func (o *NaturalPersonName) GetNameIdentifiers() []NaturalPersonNameId {
	if o == nil || o.NameIdentifiers == nil {
		var ret []NaturalPersonNameId
		return ret
	}
	return *o.NameIdentifiers
}

// GetNameIdentifiersOk returns a tuple with the NameIdentifiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NaturalPersonName) GetNameIdentifiersOk() (*[]NaturalPersonNameId, bool) {
	if o == nil || o.NameIdentifiers == nil {
		return nil, false
	}
	return o.NameIdentifiers, true
}

// HasNameIdentifiers returns a boolean if a field has been set.
func (o *NaturalPersonName) HasNameIdentifiers() bool {
	if o != nil && o.NameIdentifiers != nil {
		return true
	}

	return false
}

// SetNameIdentifiers gets a reference to the given []NaturalPersonNameId and assigns it to the NameIdentifiers field.
func (o *NaturalPersonName) SetNameIdentifiers(v []NaturalPersonNameId) {
	o.NameIdentifiers = &v
}

// GetLocalNameIdentifiers returns the LocalNameIdentifiers field value if set, zero value otherwise.
func (o *NaturalPersonName) GetLocalNameIdentifiers() []LocalNaturalPersonNameId {
	if o == nil || o.LocalNameIdentifiers == nil {
		var ret []LocalNaturalPersonNameId
		return ret
	}
	return *o.LocalNameIdentifiers
}

// GetLocalNameIdentifiersOk returns a tuple with the LocalNameIdentifiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NaturalPersonName) GetLocalNameIdentifiersOk() (*[]LocalNaturalPersonNameId, bool) {
	if o == nil || o.LocalNameIdentifiers == nil {
		return nil, false
	}
	return o.LocalNameIdentifiers, true
}

// HasLocalNameIdentifiers returns a boolean if a field has been set.
func (o *NaturalPersonName) HasLocalNameIdentifiers() bool {
	if o != nil && o.LocalNameIdentifiers != nil {
		return true
	}

	return false
}

// SetLocalNameIdentifiers gets a reference to the given []LocalNaturalPersonNameId and assigns it to the LocalNameIdentifiers field.
func (o *NaturalPersonName) SetLocalNameIdentifiers(v []LocalNaturalPersonNameId) {
	o.LocalNameIdentifiers = &v
}

// GetPhoneticNameIdentifiers returns the PhoneticNameIdentifiers field value if set, zero value otherwise.
func (o *NaturalPersonName) GetPhoneticNameIdentifiers() []PhoneticNaturalPersonNameId {
	if o == nil || o.PhoneticNameIdentifiers == nil {
		var ret []PhoneticNaturalPersonNameId
		return ret
	}
	return *o.PhoneticNameIdentifiers
}

// GetPhoneticNameIdentifiersOk returns a tuple with the PhoneticNameIdentifiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NaturalPersonName) GetPhoneticNameIdentifiersOk() (*[]PhoneticNaturalPersonNameId, bool) {
	if o == nil || o.PhoneticNameIdentifiers == nil {
		return nil, false
	}
	return o.PhoneticNameIdentifiers, true
}

// HasPhoneticNameIdentifiers returns a boolean if a field has been set.
func (o *NaturalPersonName) HasPhoneticNameIdentifiers() bool {
	if o != nil && o.PhoneticNameIdentifiers != nil {
		return true
	}

	return false
}

// SetPhoneticNameIdentifiers gets a reference to the given []PhoneticNaturalPersonNameId and assigns it to the PhoneticNameIdentifiers field.
func (o *NaturalPersonName) SetPhoneticNameIdentifiers(v []PhoneticNaturalPersonNameId) {
	o.PhoneticNameIdentifiers = &v
}

func (o NaturalPersonName) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NameIdentifiers != nil {
		toSerialize["name_identifiers"] = o.NameIdentifiers
	}
	if o.LocalNameIdentifiers != nil {
		toSerialize["local_name_identifiers"] = o.LocalNameIdentifiers
	}
	if o.PhoneticNameIdentifiers != nil {
		toSerialize["phonetic_name_identifiers"] = o.PhoneticNameIdentifiers
	}
	return json.Marshal(toSerialize)
}

type NullableNaturalPersonName struct {
	value *NaturalPersonName
	isSet bool
}

func (v NullableNaturalPersonName) Get() *NaturalPersonName {
	return v.value
}

func (v *NullableNaturalPersonName) Set(val *NaturalPersonName) {
	v.value = val
	v.isSet = true
}

func (v NullableNaturalPersonName) IsSet() bool {
	return v.isSet
}

func (v *NullableNaturalPersonName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNaturalPersonName(val *NaturalPersonName) *NullableNaturalPersonName {
	return &NullableNaturalPersonName{value: val, isSet: true}
}

func (v NullableNaturalPersonName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNaturalPersonName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


