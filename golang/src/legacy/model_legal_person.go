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

// LegalPerson struct for LegalPerson
type LegalPerson struct {
	Name                   *LegalPersonName        `json:"name,omitempty"`
	GeographicAddresses    *[]Address              `json:"geographic_addresses,omitempty"`
	CustomerIdentification *string                 `json:"customer_identification,omitempty"`
	CustomerNumber         *string                 `json:"customer_number,omitempty"`
	NationalIdentification *NationalIdentification `json:"national_identification,omitempty"`
	CountryOfRegistration *string `json:"country_of_registration,omitempty"`
}

// NewLegalPerson instantiates a new LegalPerson object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLegalPerson() *LegalPerson {
	this := LegalPerson{}
	return &this
}

// NewLegalPersonWithDefaults instantiates a new LegalPerson object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLegalPersonWithDefaults() *LegalPerson {
	this := LegalPerson{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LegalPerson) GetName() LegalPersonName {
	if o == nil || o.Name == nil {
		var ret LegalPersonName
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegalPerson) GetNameOk() (*LegalPersonName, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LegalPerson) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given LegalPersonName and assigns it to the Name field.
func (o *LegalPerson) SetName(v LegalPersonName) {
	o.Name = &v
}

// GetGeographicAddresses returns the GeographicAddresses field value if set, zero value otherwise.
func (o *LegalPerson) GetGeographicAddresses() []Address {
	if o == nil || o.GeographicAddresses == nil {
		var ret []Address
		return ret
	}
	return *o.GeographicAddresses
}

// GetGeographicAddressesOk returns a tuple with the GeographicAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegalPerson) GetGeographicAddressesOk() (*[]Address, bool) {
	if o == nil || o.GeographicAddresses == nil {
		return nil, false
	}
	return o.GeographicAddresses, true
}

// HasGeographicAddresses returns a boolean if a field has been set.
func (o *LegalPerson) HasGeographicAddresses() bool {
	if o != nil && o.GeographicAddresses != nil {
		return true
	}

	return false
}

// SetGeographicAddresses gets a reference to the given []Address and assigns it to the GeographicAddresses field.
func (o *LegalPerson) SetGeographicAddresses(v []Address) {
	o.GeographicAddresses = &v
}

// GetCustomerIdentification returns the CustomerIdentification field value if set, zero value otherwise.
func (o *LegalPerson) GetCustomerIdentification() string {
	if o == nil || o.CustomerIdentification == nil {
		var ret string
		return ret
	}
	return *o.CustomerIdentification
}

// GetCustomerIdentificationOk returns a tuple with the CustomerIdentification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegalPerson) GetCustomerIdentificationOk() (*string, bool) {
	if o == nil || o.CustomerIdentification == nil {
		return nil, false
	}
	return o.CustomerIdentification, true
}

// HasCustomerIdentification returns a boolean if a field has been set.
func (o *LegalPerson) HasCustomerIdentification() bool {
	if o != nil && o.CustomerIdentification != nil {
		return true
	}

	return false
}

// SetCustomerIdentification gets a reference to the given string and assigns it to the CustomerIdentification field.
func (o *LegalPerson) SetCustomerIdentification(v string) {
	o.CustomerIdentification = &v
}

// GetNationalIdentification returns the NationalIdentification field value if set, zero value otherwise.
func (o *LegalPerson) GetNationalIdentification() NationalIdentification {
	if o == nil || o.NationalIdentification == nil {
		var ret NationalIdentification
		return ret
	}
	return *o.NationalIdentification
}

// GetNationalIdentificationOk returns a tuple with the NationalIdentification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegalPerson) GetNationalIdentificationOk() (*NationalIdentification, bool) {
	if o == nil || o.NationalIdentification == nil {
		return nil, false
	}
	return o.NationalIdentification, true
}

// HasNationalIdentification returns a boolean if a field has been set.
func (o *LegalPerson) HasNationalIdentification() bool {
	if o != nil && o.NationalIdentification != nil {
		return true
	}

	return false
}

// SetNationalIdentification gets a reference to the given NationalIdentification and assigns it to the NationalIdentification field.
func (o *LegalPerson) SetNationalIdentification(v NationalIdentification) {
	o.NationalIdentification = &v
}

// GetCountryOfRegistration returns the CountryOfRegistration field value if set, zero value otherwise.
func (o *LegalPerson) GetCountryOfRegistration() string {
	if o == nil || o.CountryOfRegistration == nil {
		var ret string
		return ret
	}
	return *o.CountryOfRegistration
}

// GetCountryOfRegistrationOk returns a tuple with the CountryOfRegistration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegalPerson) GetCountryOfRegistrationOk() (*string, bool) {
	if o == nil || o.CountryOfRegistration == nil {
		return nil, false
	}
	return o.CountryOfRegistration, true
}

// HasCountryOfRegistration returns a boolean if a field has been set.
func (o *LegalPerson) HasCountryOfRegistration() bool {
	if o != nil && o.CountryOfRegistration != nil {
		return true
	}

	return false
}

// SetCountryOfRegistration gets a reference to the given string and assigns it to the CountryOfRegistration field.
func (o *LegalPerson) SetCountryOfRegistration(v string) {
	o.CountryOfRegistration = &v
}

func (o LegalPerson) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.GeographicAddresses != nil {
		toSerialize["geographic_addresses"] = o.GeographicAddresses
	}
	if o.CustomerIdentification != nil {
		toSerialize["customer_identification"] = o.CustomerIdentification
		toSerialize["customer_number"] = o.CustomerIdentification
	}
	if o.CustomerNumber != nil {
		toSerialize["customer_identification"] = o.CustomerNumber
	}
	if o.NationalIdentification != nil {
		toSerialize["national_identification"] = o.NationalIdentification
	}
	if o.CountryOfRegistration != nil {
		toSerialize["country_of_registration"] = o.CountryOfRegistration
	}
	return json.Marshal(toSerialize)
}

type NullableLegalPerson struct {
	value *LegalPerson
	isSet bool
}

func (v NullableLegalPerson) Get() *LegalPerson {
	return v.value
}

func (v *NullableLegalPerson) Set(val *LegalPerson) {
	v.value = val
	v.isSet = true
}

func (v NullableLegalPerson) IsSet() bool {
	return v.isSet
}

func (v *NullableLegalPerson) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLegalPerson(val *LegalPerson) *NullableLegalPerson {
	return &NullableLegalPerson{value: val, isSet: true}
}

func (v NullableLegalPerson) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLegalPerson) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

