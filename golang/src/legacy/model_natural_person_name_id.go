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

// NaturalPersonNameId struct for NaturalPersonNameId
type NaturalPersonNameId struct {
	PrimaryIdentifier *string `json:"primary_identifier,omitempty"`
	SecondaryIdentifier *string `json:"secondary_identifier,omitempty"`
	NameIdentifierType *string `json:"name_identifier_type,omitempty"`
}

// NewNaturalPersonNameId instantiates a new NaturalPersonNameId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNaturalPersonNameId() *NaturalPersonNameId {
	this := NaturalPersonNameId{}
	return &this
}

// NewNaturalPersonNameIdWithDefaults instantiates a new NaturalPersonNameId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNaturalPersonNameIdWithDefaults() *NaturalPersonNameId {
	this := NaturalPersonNameId{}
	return &this
}

// GetPrimaryIdentifier returns the PrimaryIdentifier field value if set, zero value otherwise.
func (o *NaturalPersonNameId) GetPrimaryIdentifier() string {
	if o == nil || o.PrimaryIdentifier == nil {
		var ret string
		return ret
	}
	return *o.PrimaryIdentifier
}

// GetPrimaryIdentifierOk returns a tuple with the PrimaryIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NaturalPersonNameId) GetPrimaryIdentifierOk() (*string, bool) {
	if o == nil || o.PrimaryIdentifier == nil {
		return nil, false
	}
	return o.PrimaryIdentifier, true
}

// HasPrimaryIdentifier returns a boolean if a field has been set.
func (o *NaturalPersonNameId) HasPrimaryIdentifier() bool {
	if o != nil && o.PrimaryIdentifier != nil {
		return true
	}

	return false
}

// SetPrimaryIdentifier gets a reference to the given string and assigns it to the PrimaryIdentifier field.
func (o *NaturalPersonNameId) SetPrimaryIdentifier(v string) {
	o.PrimaryIdentifier = &v
}

// GetSecondaryIdentifier returns the SecondaryIdentifier field value if set, zero value otherwise.
func (o *NaturalPersonNameId) GetSecondaryIdentifier() string {
	if o == nil || o.SecondaryIdentifier == nil {
		var ret string
		return ret
	}
	return *o.SecondaryIdentifier
}

// GetSecondaryIdentifierOk returns a tuple with the SecondaryIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NaturalPersonNameId) GetSecondaryIdentifierOk() (*string, bool) {
	if o == nil || o.SecondaryIdentifier == nil {
		return nil, false
	}
	return o.SecondaryIdentifier, true
}

// HasSecondaryIdentifier returns a boolean if a field has been set.
func (o *NaturalPersonNameId) HasSecondaryIdentifier() bool {
	if o != nil && o.SecondaryIdentifier != nil {
		return true
	}

	return false
}

// SetSecondaryIdentifier gets a reference to the given string and assigns it to the SecondaryIdentifier field.
func (o *NaturalPersonNameId) SetSecondaryIdentifier(v string) {
	o.SecondaryIdentifier = &v
}

// GetNameIdentifierType returns the NameIdentifierType field value if set, zero value otherwise.
func (o *NaturalPersonNameId) GetNameIdentifierType() string {
	if o == nil || o.NameIdentifierType == nil {
		var ret string
		return ret
	}
	return *o.NameIdentifierType
}

// GetNameIdentifierTypeOk returns a tuple with the NameIdentifierType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NaturalPersonNameId) GetNameIdentifierTypeOk() (*string, bool) {
	if o == nil || o.NameIdentifierType == nil {
		return nil, false
	}
	return o.NameIdentifierType, true
}

// HasNameIdentifierType returns a boolean if a field has been set.
func (o *NaturalPersonNameId) HasNameIdentifierType() bool {
	if o != nil && o.NameIdentifierType != nil {
		return true
	}

	return false
}

// SetNameIdentifierType gets a reference to the given string and assigns it to the NameIdentifierType field.
func (o *NaturalPersonNameId) SetNameIdentifierType(v string) {
	o.NameIdentifierType = &v
}

func (o NaturalPersonNameId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PrimaryIdentifier != nil {
		toSerialize["primary_identifier"] = o.PrimaryIdentifier
	}
	if o.SecondaryIdentifier != nil {
		toSerialize["secondary_identifier"] = o.SecondaryIdentifier
	}
	if o.NameIdentifierType != nil {
		toSerialize["name_identifier_type"] = o.NameIdentifierType
	}
	return json.Marshal(toSerialize)
}

type NullableNaturalPersonNameId struct {
	value *NaturalPersonNameId
	isSet bool
}

func (v NullableNaturalPersonNameId) Get() *NaturalPersonNameId {
	return v.value
}

func (v *NullableNaturalPersonNameId) Set(val *NaturalPersonNameId) {
	v.value = val
	v.isSet = true
}

func (v NullableNaturalPersonNameId) IsSet() bool {
	return v.isSet
}

func (v *NullableNaturalPersonNameId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNaturalPersonNameId(val *NaturalPersonNameId) *NullableNaturalPersonNameId {
	return &NullableNaturalPersonNameId{value: val, isSet: true}
}

func (v NullableNaturalPersonNameId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNaturalPersonNameId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

