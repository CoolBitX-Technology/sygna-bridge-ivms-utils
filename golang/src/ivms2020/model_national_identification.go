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

// checks if the NationalIdentification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NationalIdentification{}

// NationalIdentification struct for NationalIdentification
type NationalIdentification struct {
	NationalIdentifier *string `json:"nationalIdentifier,omitempty"`
	NationalIdentifierType *string `json:"nationalIdentifierType,omitempty"`
	CountryOfIssue *string `json:"countryOfIssue,omitempty"`
	RegistrationAuthority *string `json:"registrationAuthority,omitempty"`
}

// NewNationalIdentification instantiates a new NationalIdentification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNationalIdentification() *NationalIdentification {
	this := NationalIdentification{}
	return &this
}

// NewNationalIdentificationWithDefaults instantiates a new NationalIdentification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNationalIdentificationWithDefaults() *NationalIdentification {
	this := NationalIdentification{}
	return &this
}

// GetNationalIdentifier returns the NationalIdentifier field value if set, zero value otherwise.
func (o *NationalIdentification) GetNationalIdentifier() string {
	if o == nil || IsNil(o.NationalIdentifier) {
		var ret string
		return ret
	}
	return *o.NationalIdentifier
}

// GetNationalIdentifierOk returns a tuple with the NationalIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NationalIdentification) GetNationalIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.NationalIdentifier) {
		return nil, false
	}
	return o.NationalIdentifier, true
}

// HasNationalIdentifier returns a boolean if a field has been set.
func (o *NationalIdentification) HasNationalIdentifier() bool {
	if o != nil && !IsNil(o.NationalIdentifier) {
		return true
	}

	return false
}

// SetNationalIdentifier gets a reference to the given string and assigns it to the NationalIdentifier field.
func (o *NationalIdentification) SetNationalIdentifier(v string) {
	o.NationalIdentifier = &v
}

// GetNationalIdentifierType returns the NationalIdentifierType field value if set, zero value otherwise.
func (o *NationalIdentification) GetNationalIdentifierType() string {
	if o == nil || IsNil(o.NationalIdentifierType) {
		var ret string
		return ret
	}
	return *o.NationalIdentifierType
}

// GetNationalIdentifierTypeOk returns a tuple with the NationalIdentifierType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NationalIdentification) GetNationalIdentifierTypeOk() (*string, bool) {
	if o == nil || IsNil(o.NationalIdentifierType) {
		return nil, false
	}
	return o.NationalIdentifierType, true
}

// HasNationalIdentifierType returns a boolean if a field has been set.
func (o *NationalIdentification) HasNationalIdentifierType() bool {
	if o != nil && !IsNil(o.NationalIdentifierType) {
		return true
	}

	return false
}

// SetNationalIdentifierType gets a reference to the given string and assigns it to the NationalIdentifierType field.
func (o *NationalIdentification) SetNationalIdentifierType(v string) {
	o.NationalIdentifierType = &v
}

// GetCountryOfIssue returns the CountryOfIssue field value if set, zero value otherwise.
func (o *NationalIdentification) GetCountryOfIssue() string {
	if o == nil || IsNil(o.CountryOfIssue) {
		var ret string
		return ret
	}
	return *o.CountryOfIssue
}

// GetCountryOfIssueOk returns a tuple with the CountryOfIssue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NationalIdentification) GetCountryOfIssueOk() (*string, bool) {
	if o == nil || IsNil(o.CountryOfIssue) {
		return nil, false
	}
	return o.CountryOfIssue, true
}

// HasCountryOfIssue returns a boolean if a field has been set.
func (o *NationalIdentification) HasCountryOfIssue() bool {
	if o != nil && !IsNil(o.CountryOfIssue) {
		return true
	}

	return false
}

// SetCountryOfIssue gets a reference to the given string and assigns it to the CountryOfIssue field.
func (o *NationalIdentification) SetCountryOfIssue(v string) {
	o.CountryOfIssue = &v
}

// GetRegistrationAuthority returns the RegistrationAuthority field value if set, zero value otherwise.
func (o *NationalIdentification) GetRegistrationAuthority() string {
	if o == nil || IsNil(o.RegistrationAuthority) {
		var ret string
		return ret
	}
	return *o.RegistrationAuthority
}

// GetRegistrationAuthorityOk returns a tuple with the RegistrationAuthority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NationalIdentification) GetRegistrationAuthorityOk() (*string, bool) {
	if o == nil || IsNil(o.RegistrationAuthority) {
		return nil, false
	}
	return o.RegistrationAuthority, true
}

// HasRegistrationAuthority returns a boolean if a field has been set.
func (o *NationalIdentification) HasRegistrationAuthority() bool {
	if o != nil && !IsNil(o.RegistrationAuthority) {
		return true
	}

	return false
}

// SetRegistrationAuthority gets a reference to the given string and assigns it to the RegistrationAuthority field.
func (o *NationalIdentification) SetRegistrationAuthority(v string) {
	o.RegistrationAuthority = &v
}

func (o NationalIdentification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NationalIdentification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NationalIdentifier) {
		toSerialize["nationalIdentifier"] = o.NationalIdentifier
	}
	if !IsNil(o.NationalIdentifierType) {
		toSerialize["nationalIdentifierType"] = o.NationalIdentifierType
	}
	if !IsNil(o.CountryOfIssue) {
		toSerialize["countryOfIssue"] = o.CountryOfIssue
	}
	if !IsNil(o.RegistrationAuthority) {
		toSerialize["registrationAuthority"] = o.RegistrationAuthority
	}
	return toSerialize, nil
}

type NullableNationalIdentification struct {
	value *NationalIdentification
	isSet bool
}

func (v NullableNationalIdentification) Get() *NationalIdentification {
	return v.value
}

func (v *NullableNationalIdentification) Set(val *NationalIdentification) {
	v.value = val
	v.isSet = true
}

func (v NullableNationalIdentification) IsSet() bool {
	return v.isSet
}

func (v *NullableNationalIdentification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNationalIdentification(val *NationalIdentification) *NullableNationalIdentification {
	return &NullableNationalIdentification{value: val, isSet: true}
}

func (v NullableNationalIdentification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNationalIdentification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

