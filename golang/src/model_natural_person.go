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

// NaturalPerson struct for NaturalPerson
type NaturalPerson struct {
	Name *NaturalPersonName `json:"name,omitempty"`
	GeographicAddresses *[]Address `json:"geographic_addresses,omitempty"`
	NationalIdentification *NationalIdentification `json:"national_identification,omitempty"`
	CustomerIdentification *string `json:"customer_identification,omitempty"`
	DateAndPlaceOfBirth *DateAndPlaceOfBirth `json:"date_and_place_of_birth,omitempty"`
	CountryOfResidence *string `json:"country_of_residence,omitempty"`
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
	if o == nil || o.Name == nil {
		var ret NaturalPersonName
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NaturalPerson) GetNameOk() (*NaturalPersonName, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NaturalPerson) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given NaturalPersonName and assigns it to the Name field.
func (o *NaturalPerson) SetName(v NaturalPersonName) {
	o.Name = &v
}

// GetGeographicAddresses returns the GeographicAddresses field value if set, zero value otherwise.
func (o *NaturalPerson) GetGeographicAddresses() []Address {
	if o == nil || o.GeographicAddresses == nil {
		var ret []Address
		return ret
	}
	return *o.GeographicAddresses
}

// GetGeographicAddressesOk returns a tuple with the GeographicAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NaturalPerson) GetGeographicAddressesOk() (*[]Address, bool) {
	if o == nil || o.GeographicAddresses == nil {
		return nil, false
	}
	return o.GeographicAddresses, true
}

// HasGeographicAddresses returns a boolean if a field has been set.
func (o *NaturalPerson) HasGeographicAddresses() bool {
	if o != nil && o.GeographicAddresses != nil {
		return true
	}

	return false
}

// SetGeographicAddresses gets a reference to the given []Address and assigns it to the GeographicAddresses field.
func (o *NaturalPerson) SetGeographicAddresses(v []Address) {
	o.GeographicAddresses = &v
}

// GetNationalIdentification returns the NationalIdentification field value if set, zero value otherwise.
func (o *NaturalPerson) GetNationalIdentification() NationalIdentification {
	if o == nil || o.NationalIdentification == nil {
		var ret NationalIdentification
		return ret
	}
	return *o.NationalIdentification
}

// GetNationalIdentificationOk returns a tuple with the NationalIdentification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NaturalPerson) GetNationalIdentificationOk() (*NationalIdentification, bool) {
	if o == nil || o.NationalIdentification == nil {
		return nil, false
	}
	return o.NationalIdentification, true
}

// HasNationalIdentification returns a boolean if a field has been set.
func (o *NaturalPerson) HasNationalIdentification() bool {
	if o != nil && o.NationalIdentification != nil {
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
	if o == nil || o.CustomerIdentification == nil {
		var ret string
		return ret
	}
	return *o.CustomerIdentification
}

// GetCustomerIdentificationOk returns a tuple with the CustomerIdentification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NaturalPerson) GetCustomerIdentificationOk() (*string, bool) {
	if o == nil || o.CustomerIdentification == nil {
		return nil, false
	}
	return o.CustomerIdentification, true
}

// HasCustomerIdentification returns a boolean if a field has been set.
func (o *NaturalPerson) HasCustomerIdentification() bool {
	if o != nil && o.CustomerIdentification != nil {
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
	if o == nil || o.DateAndPlaceOfBirth == nil {
		var ret DateAndPlaceOfBirth
		return ret
	}
	return *o.DateAndPlaceOfBirth
}

// GetDateAndPlaceOfBirthOk returns a tuple with the DateAndPlaceOfBirth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NaturalPerson) GetDateAndPlaceOfBirthOk() (*DateAndPlaceOfBirth, bool) {
	if o == nil || o.DateAndPlaceOfBirth == nil {
		return nil, false
	}
	return o.DateAndPlaceOfBirth, true
}

// HasDateAndPlaceOfBirth returns a boolean if a field has been set.
func (o *NaturalPerson) HasDateAndPlaceOfBirth() bool {
	if o != nil && o.DateAndPlaceOfBirth != nil {
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
	if o == nil || o.CountryOfResidence == nil {
		var ret string
		return ret
	}
	return *o.CountryOfResidence
}

// GetCountryOfResidenceOk returns a tuple with the CountryOfResidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NaturalPerson) GetCountryOfResidenceOk() (*string, bool) {
	if o == nil || o.CountryOfResidence == nil {
		return nil, false
	}
	return o.CountryOfResidence, true
}

// HasCountryOfResidence returns a boolean if a field has been set.
func (o *NaturalPerson) HasCountryOfResidence() bool {
	if o != nil && o.CountryOfResidence != nil {
		return true
	}

	return false
}

// SetCountryOfResidence gets a reference to the given string and assigns it to the CountryOfResidence field.
func (o *NaturalPerson) SetCountryOfResidence(v string) {
	o.CountryOfResidence = &v
}

func (o NaturalPerson) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.GeographicAddresses != nil {
		toSerialize["geographic_addresses"] = o.GeographicAddresses
	}
	if o.NationalIdentification != nil {
		toSerialize["national_identification"] = o.NationalIdentification
	}
	if o.CustomerIdentification != nil {
		toSerialize["customer_identification"] = o.CustomerIdentification
	}
	if o.DateAndPlaceOfBirth != nil {
		toSerialize["date_and_place_of_birth"] = o.DateAndPlaceOfBirth
	}
	if o.CountryOfResidence != nil {
		toSerialize["country_of_residence"] = o.CountryOfResidence
	}
	return json.Marshal(toSerialize)
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


