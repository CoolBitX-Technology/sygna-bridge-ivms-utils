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

// checks if the DateAndPlaceOfBirth type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DateAndPlaceOfBirth{}

// DateAndPlaceOfBirth struct for DateAndPlaceOfBirth
type DateAndPlaceOfBirth struct {
	DateOfBirth *string `json:"dateOfBirth,omitempty"`
	PlaceOfBirth *string `json:"placeOfBirth,omitempty"`
}

// NewDateAndPlaceOfBirth instantiates a new DateAndPlaceOfBirth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDateAndPlaceOfBirth() *DateAndPlaceOfBirth {
	this := DateAndPlaceOfBirth{}
	return &this
}

// NewDateAndPlaceOfBirthWithDefaults instantiates a new DateAndPlaceOfBirth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDateAndPlaceOfBirthWithDefaults() *DateAndPlaceOfBirth {
	this := DateAndPlaceOfBirth{}
	return &this
}

// GetDateOfBirth returns the DateOfBirth field value if set, zero value otherwise.
func (o *DateAndPlaceOfBirth) GetDateOfBirth() string {
	if o == nil || IsNil(o.DateOfBirth) {
		var ret string
		return ret
	}
	return *o.DateOfBirth
}

// GetDateOfBirthOk returns a tuple with the DateOfBirth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateAndPlaceOfBirth) GetDateOfBirthOk() (*string, bool) {
	if o == nil || IsNil(o.DateOfBirth) {
		return nil, false
	}
	return o.DateOfBirth, true
}

// HasDateOfBirth returns a boolean if a field has been set.
func (o *DateAndPlaceOfBirth) HasDateOfBirth() bool {
	if o != nil && !IsNil(o.DateOfBirth) {
		return true
	}

	return false
}

// SetDateOfBirth gets a reference to the given string and assigns it to the DateOfBirth field.
func (o *DateAndPlaceOfBirth) SetDateOfBirth(v string) {
	o.DateOfBirth = &v
}

// GetPlaceOfBirth returns the PlaceOfBirth field value if set, zero value otherwise.
func (o *DateAndPlaceOfBirth) GetPlaceOfBirth() string {
	if o == nil || IsNil(o.PlaceOfBirth) {
		var ret string
		return ret
	}
	return *o.PlaceOfBirth
}

// GetPlaceOfBirthOk returns a tuple with the PlaceOfBirth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateAndPlaceOfBirth) GetPlaceOfBirthOk() (*string, bool) {
	if o == nil || IsNil(o.PlaceOfBirth) {
		return nil, false
	}
	return o.PlaceOfBirth, true
}

// HasPlaceOfBirth returns a boolean if a field has been set.
func (o *DateAndPlaceOfBirth) HasPlaceOfBirth() bool {
	if o != nil && !IsNil(o.PlaceOfBirth) {
		return true
	}

	return false
}

// SetPlaceOfBirth gets a reference to the given string and assigns it to the PlaceOfBirth field.
func (o *DateAndPlaceOfBirth) SetPlaceOfBirth(v string) {
	o.PlaceOfBirth = &v
}

func (o DateAndPlaceOfBirth) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DateAndPlaceOfBirth) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DateOfBirth) {
		toSerialize["dateOfBirth"] = o.DateOfBirth
	}
	if !IsNil(o.PlaceOfBirth) {
		toSerialize["placeOfBirth"] = o.PlaceOfBirth
	}
	return toSerialize, nil
}

type NullableDateAndPlaceOfBirth struct {
	value *DateAndPlaceOfBirth
	isSet bool
}

func (v NullableDateAndPlaceOfBirth) Get() *DateAndPlaceOfBirth {
	return v.value
}

func (v *NullableDateAndPlaceOfBirth) Set(val *DateAndPlaceOfBirth) {
	v.value = val
	v.isSet = true
}

func (v NullableDateAndPlaceOfBirth) IsSet() bool {
	return v.isSet
}

func (v *NullableDateAndPlaceOfBirth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDateAndPlaceOfBirth(val *DateAndPlaceOfBirth) *NullableDateAndPlaceOfBirth {
	return &NullableDateAndPlaceOfBirth{value: val, isSet: true}
}

func (v NullableDateAndPlaceOfBirth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDateAndPlaceOfBirth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

