/*
Bridge

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ivms101

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the Originator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Originator{}

// Originator struct for Originator
type Originator struct {
	OriginatorPerson []Person `json:"originatorPerson"`
}

type _Originator Originator

// NewOriginator instantiates a new Originator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOriginator(originatorPerson []Person) *Originator {
	this := Originator{}
	this.OriginatorPerson = originatorPerson
	return &this
}

// NewOriginatorWithDefaults instantiates a new Originator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOriginatorWithDefaults() *Originator {
	this := Originator{}
	return &this
}

// GetOriginatorPerson returns the OriginatorPerson field value
func (o *Originator) GetOriginatorPerson() []Person {
	if o == nil {
		var ret []Person
		return ret
	}

	return o.OriginatorPerson
}

// GetOriginatorPersonOk returns a tuple with the OriginatorPerson field value
// and a boolean to check if the value has been set.
func (o *Originator) GetOriginatorPersonOk() ([]Person, bool) {
	if o == nil {
		return nil, false
	}
	return o.OriginatorPerson, true
}

// SetOriginatorPerson sets field value
func (o *Originator) SetOriginatorPerson(v []Person) {
	o.OriginatorPerson = v
}

func (o Originator) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Originator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["originatorPerson"] = o.OriginatorPerson
	return toSerialize, nil
}

func (o *Originator) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"originatorPerson",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varOriginator := _Originator{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOriginator)

	if err != nil {
		return err
	}

	*o = Originator(varOriginator)

	return err
}

type NullableOriginator struct {
	value *Originator
	isSet bool
}

func (v NullableOriginator) Get() *Originator {
	return v.value
}

func (v *NullableOriginator) Set(val *Originator) {
	v.value = val
	v.isSet = true
}

func (v NullableOriginator) IsSet() bool {
	return v.isSet
}

func (v *NullableOriginator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOriginator(val *Originator) *NullableOriginator {
	return &NullableOriginator{value: val, isSet: true}
}

func (v NullableOriginator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOriginator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

