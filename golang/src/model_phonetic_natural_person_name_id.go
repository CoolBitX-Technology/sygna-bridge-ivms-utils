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

// PhoneticNaturalPersonNameId struct for PhoneticNaturalPersonNameId
type PhoneticNaturalPersonNameId struct {
	PrimaryIdentifier   *string `json:"primary_identifier,omitempty"`
	SecondaryIdentifier *string `json:"secondary_identifier,omitempty"`
	NameIdentifierType  *string `json:"name_identifier_type,omitempty"`
}

// NewPhoneticNaturalPersonNameId instantiates a new PhoneticNaturalPersonNameId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPhoneticNaturalPersonNameId() *PhoneticNaturalPersonNameId {
	this := PhoneticNaturalPersonNameId{}
	return &this
}

// NewPhoneticNaturalPersonNameIdWithDefaults instantiates a new PhoneticNaturalPersonNameId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPhoneticNaturalPersonNameIdWithDefaults() *PhoneticNaturalPersonNameId {
	this := PhoneticNaturalPersonNameId{}
	return &this
}

// GetPrimaryIdentifier returns the PrimaryIdentifier field value if set, zero value otherwise.
func (o *PhoneticNaturalPersonNameId) GetPrimaryIdentifier() string {
	if o == nil || o.PrimaryIdentifier == nil {
		var ret string
		return ret
	}
	return *o.PrimaryIdentifier
}

// GetPrimaryIdentifierOk returns a tuple with the PrimaryIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhoneticNaturalPersonNameId) GetPrimaryIdentifierOk() (*string, bool) {
	if o == nil || o.PrimaryIdentifier == nil {
		return nil, false
	}
	return o.PrimaryIdentifier, true
}

// HasPrimaryIdentifier returns a boolean if a field has been set.
func (o *PhoneticNaturalPersonNameId) HasPrimaryIdentifier() bool {
	if o != nil && o.PrimaryIdentifier != nil {
		return true
	}

	return false
}

// SetPrimaryIdentifier gets a reference to the given string and assigns it to the PrimaryIdentifier field.
func (o *PhoneticNaturalPersonNameId) SetPrimaryIdentifier(v string) {
	o.PrimaryIdentifier = &v
}

// GetSecondaryIdentifier returns the SecondaryIdentifier field value if set, zero value otherwise.
func (o *PhoneticNaturalPersonNameId) GetSecondaryIdentifier() string {
	if o == nil || o.SecondaryIdentifier == nil {
		var ret string
		return ret
	}
	return *o.SecondaryIdentifier
}

// GetSecondaryIdentifierOk returns a tuple with the SecondaryIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhoneticNaturalPersonNameId) GetSecondaryIdentifierOk() (*string, bool) {
	if o == nil || o.SecondaryIdentifier == nil {
		return nil, false
	}
	return o.SecondaryIdentifier, true
}

// HasSecondaryIdentifier returns a boolean if a field has been set.
func (o *PhoneticNaturalPersonNameId) HasSecondaryIdentifier() bool {
	if o != nil && o.SecondaryIdentifier != nil {
		return true
	}

	return false
}

// SetSecondaryIdentifier gets a reference to the given string and assigns it to the SecondaryIdentifier field.
func (o *PhoneticNaturalPersonNameId) SetSecondaryIdentifier(v string) {
	o.SecondaryIdentifier = &v
}

// GetNameIdentifierType returns the NameIdentifierType field value if set, zero value otherwise.
func (o *PhoneticNaturalPersonNameId) GetNameIdentifierType() string {
	if o == nil || o.NameIdentifierType == nil {
		var ret string
		return ret
	}
	return *o.NameIdentifierType
}

// GetNameIdentifierTypeOk returns a tuple with the NameIdentifierType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhoneticNaturalPersonNameId) GetNameIdentifierTypeOk() (*string, bool) {
	if o == nil || o.NameIdentifierType == nil {
		return nil, false
	}
	return o.NameIdentifierType, true
}

// HasNameIdentifierType returns a boolean if a field has been set.
func (o *PhoneticNaturalPersonNameId) HasNameIdentifierType() bool {
	if o != nil && o.NameIdentifierType != nil {
		return true
	}

	return false
}

// SetNameIdentifierType gets a reference to the given string and assigns it to the NameIdentifierType field.
func (o *PhoneticNaturalPersonNameId) SetNameIdentifierType(v string) {
	o.NameIdentifierType = &v
}

func (o PhoneticNaturalPersonNameId) MarshalJSON() ([]byte, error) {
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

type NullablePhoneticNaturalPersonNameId struct {
	value *PhoneticNaturalPersonNameId
	isSet bool
}

func (v NullablePhoneticNaturalPersonNameId) Get() *PhoneticNaturalPersonNameId {
	return v.value
}

func (v *NullablePhoneticNaturalPersonNameId) Set(val *PhoneticNaturalPersonNameId) {
	v.value = val
	v.isSet = true
}

func (v NullablePhoneticNaturalPersonNameId) IsSet() bool {
	return v.isSet
}

func (v *NullablePhoneticNaturalPersonNameId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePhoneticNaturalPersonNameId(val *PhoneticNaturalPersonNameId) *NullablePhoneticNaturalPersonNameId {
	return &NullablePhoneticNaturalPersonNameId{value: val, isSet: true}
}

func (v NullablePhoneticNaturalPersonNameId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePhoneticNaturalPersonNameId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
