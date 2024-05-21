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

// checks if the NaturalPerson type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NaturalPerson{}

// NaturalPerson struct for NaturalPerson
type NaturalPerson struct {
	Name *NaturalPersonName `json:"name,omitempty"`
	GeographicAddress []Address `json:"geographicAddress,omitempty"`
	NationalIdentification *NationalIdentification `json:"nationalIdentification,omitempty"`
	CustomerIdentification *string `json:"customerIdentification,omitempty"`
	DateAndPlaceOfBirth *DateAndPlaceOfBirth `json:"dateAndPlaceOfBirth,omitempty"`
	CountryOfResidence *string `json:"countryOfResidence,omitempty"`
}

// NewNaturalPerson instantiates a new NaturalPerson object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNaturalPerson() *NaturalPerson {
	this := NaturalPerson{}
	return &this
}

// NewNaturalPersonWithDefaults instantiates a new NaturalPerson object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNaturalPersonWithDefaults() *NaturalPerson {
	this := NaturalPerson{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NaturalPerson) GetName() NaturalPersonName {
	if o == nil || IsNil(o.Name) {
		var ret NaturalPersonName
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NaturalPerson) GetNameOk() (*NaturalPersonName, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NaturalPerson) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given NaturalPersonName and assigns it to the Name field.
func (o *NaturalPerson) SetName(v NaturalPersonName) {
	o.Name = &v
}

// GetGeographicAddress returns the GeographicAddress field value if set, zero value otherwise.
func (o *NaturalPerson) GetGeographicAddress() []Address {
	if o == nil || IsNil(o.GeographicAddress) {
		var ret []Address
		return ret
	}
	return o.GeographicAddress
}

// GetGeographicAddressOk returns a tuple with the GeographicAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NaturalPerson) GetGeographicAddressOk() ([]Address, bool) {
	if o == nil || IsNil(o.GeographicAddress) {
		return nil, false
	}
	return o.GeographicAddress, true
}

// HasGeographicAddress returns a boolean if a field has been set.
func (o *NaturalPerson) HasGeographicAddress() bool {
	if o != nil && !IsNil(o.GeographicAddress) {
		return true
	}

	return false
}

// SetGeographicAddress gets a reference to the given []Address and assigns it to the GeographicAddress field.
func (o *NaturalPerson) SetGeographicAddress(v []Address) {
	o.GeographicAddress = v
}

// GetNationalIdentification returns the NationalIdentification field value if set, zero value otherwise.
func (o *NaturalPerson) GetNationalIdentification() NationalIdentification {
	if o == nil || IsNil(o.NationalIdentification) {
		var ret NationalIdentification
		return ret
	}
	return *o.NationalIdentification
}

// GetNationalIdentificationOk returns a tuple with the NationalIdentification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NaturalPerson) GetNationalIdentificationOk() (*NationalIdentification, bool) {
	if o == nil || IsNil(o.NationalIdentification) {
		return nil, false
	}
	return o.NationalIdentification, true
}

// HasNationalIdentification returns a boolean if a field has been set.
func (o *NaturalPerson) HasNationalIdentification() bool {
	if o != nil && !IsNil(o.NationalIdentification) {
		return true
	}

	return false
}

// SetNationalIdentification gets a reference to the given NationalIdentification and assigns it to the NationalIdentification field.
func (o *NaturalPerson) SetNationalIdentification(v NationalIdentification) {
	o.NationalIdentification = &v
}

// GetCustomerIdentification returns the CustomerIdentification field value if set, zero value otherwise.
func (o *NaturalPerson) GetCustomerIdentification() string {
	if o == nil || IsNil(o.CustomerIdentification) {
		var ret string
		return ret
	}
	return *o.CustomerIdentification
}

// GetCustomerIdentificationOk returns a tuple with the CustomerIdentification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NaturalPerson) GetCustomerIdentificationOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerIdentification) {
		return nil, false
	}
	return o.CustomerIdentification, true
}

// HasCustomerIdentification returns a boolean if a field has been set.
func (o *NaturalPerson) HasCustomerIdentification() bool {
	if o != nil && !IsNil(o.CustomerIdentification) {
		return true
	}

	return false
}

// SetCustomerIdentification gets a reference to the given string and assigns it to the CustomerIdentification field.
func (o *NaturalPerson) SetCustomerIdentification(v string) {
	o.CustomerIdentification = &v
}

// GetDateAndPlaceOfBirth returns the DateAndPlaceOfBirth field value if set, zero value otherwise.
func (o *NaturalPerson) GetDateAndPlaceOfBirth() DateAndPlaceOfBirth {
	if o == nil || IsNil(o.DateAndPlaceOfBirth) {
		var ret DateAndPlaceOfBirth
		return ret
	}
	return *o.DateAndPlaceOfBirth
}

// GetDateAndPlaceOfBirthOk returns a tuple with the DateAndPlaceOfBirth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NaturalPerson) GetDateAndPlaceOfBirthOk() (*DateAndPlaceOfBirth, bool) {
	if o == nil || IsNil(o.DateAndPlaceOfBirth) {
		return nil, false
	}
	return o.DateAndPlaceOfBirth, true
}

// HasDateAndPlaceOfBirth returns a boolean if a field has been set.
func (o *NaturalPerson) HasDateAndPlaceOfBirth() bool {
	if o != nil && !IsNil(o.DateAndPlaceOfBirth) {
		return true
	}

	return false
}

// SetDateAndPlaceOfBirth gets a reference to the given DateAndPlaceOfBirth and assigns it to the DateAndPlaceOfBirth field.
func (o *NaturalPerson) SetDateAndPlaceOfBirth(v DateAndPlaceOfBirth) {
	o.DateAndPlaceOfBirth = &v
}

// GetCountryOfResidence returns the CountryOfResidence field value if set, zero value otherwise.
func (o *NaturalPerson) GetCountryOfResidence() string {
	if o == nil || IsNil(o.CountryOfResidence) {
		var ret string
		return ret
	}
	return *o.CountryOfResidence
}

// GetCountryOfResidenceOk returns a tuple with the CountryOfResidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NaturalPerson) GetCountryOfResidenceOk() (*string, bool) {
	if o == nil || IsNil(o.CountryOfResidence) {
		return nil, false
	}
	return o.CountryOfResidence, true
}

// HasCountryOfResidence returns a boolean if a field has been set.
func (o *NaturalPerson) HasCountryOfResidence() bool {
	if o != nil && !IsNil(o.CountryOfResidence) {
		return true
	}

	return false
}

// SetCountryOfResidence gets a reference to the given string and assigns it to the CountryOfResidence field.
func (o *NaturalPerson) SetCountryOfResidence(v string) {
	o.CountryOfResidence = &v
}

func (o NaturalPerson) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NaturalPerson) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.GeographicAddress) {
		toSerialize["geographicAddress"] = o.GeographicAddress
	}
	if !IsNil(o.NationalIdentification) {
		toSerialize["nationalIdentification"] = o.NationalIdentification
	}
	if !IsNil(o.CustomerIdentification) {
		toSerialize["customerIdentification"] = o.CustomerIdentification
	}
	if !IsNil(o.DateAndPlaceOfBirth) {
		toSerialize["dateAndPlaceOfBirth"] = o.DateAndPlaceOfBirth
	}
	if !IsNil(o.CountryOfResidence) {
		toSerialize["countryOfResidence"] = o.CountryOfResidence
	}
	return toSerialize, nil
}

type NullableNaturalPerson struct {
	value *NaturalPerson
	isSet bool
}

func (v NullableNaturalPerson) Get() *NaturalPerson {
	return v.value
}

func (v *NullableNaturalPerson) Set(val *NaturalPerson) {
	v.value = val
	v.isSet = true
}

func (v NullableNaturalPerson) IsSet() bool {
	return v.isSet
}

func (v *NullableNaturalPerson) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNaturalPerson(val *NaturalPerson) *NullableNaturalPerson {
	return &NullableNaturalPerson{value: val, isSet: true}
}

func (v NullableNaturalPerson) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNaturalPerson) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


