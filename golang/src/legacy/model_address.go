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

// Address struct for Address
type Address struct {
	AddressType *string `json:"address_type,omitempty"`
	Department *string `json:"department,omitempty"`
	SubDepartment *string `json:"sub_department,omitempty"`
	StreetName *string `json:"street_name,omitempty"`
	BuildingNumber *string `json:"building_number,omitempty"`
	BuildingName *string `json:"building_name,omitempty"`
	Floor *string `json:"floor,omitempty"`
	PostBox *string `json:"post_box,omitempty"`
	Room *string `json:"room,omitempty"`
	PostCode *string `json:"post_code,omitempty"`
	TownName *string `json:"town_name,omitempty"`
	TownLocationName *string `json:"town_location_name,omitempty"`
	DistrictName *string `json:"district_name,omitempty"`
	CountrySubDivision *string `json:"country_sub_division,omitempty"`
	AddressLine *[]string `json:"address_line,omitempty"`
	Country *string `json:"country,omitempty"`
}

// NewAddress instantiates a new Address object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddress() *Address {
	this := Address{}
	return &this
}

// NewAddressWithDefaults instantiates a new Address object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressWithDefaults() *Address {
	this := Address{}
	return &this
}

// GetAddressType returns the AddressType field value if set, zero value otherwise.
func (o *Address) GetAddressType() string {
	if o == nil || o.AddressType == nil {
		var ret string
		return ret
	}
	return *o.AddressType
}

// GetAddressTypeOk returns a tuple with the AddressType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetAddressTypeOk() (*string, bool) {
	if o == nil || o.AddressType == nil {
		return nil, false
	}
	return o.AddressType, true
}

// HasAddressType returns a boolean if a field has been set.
func (o *Address) HasAddressType() bool {
	if o != nil && o.AddressType != nil {
		return true
	}

	return false
}

// SetAddressType gets a reference to the given string and assigns it to the AddressType field.
func (o *Address) SetAddressType(v string) {
	o.AddressType = &v
}

// GetDepartment returns the Department field value if set, zero value otherwise.
func (o *Address) GetDepartment() string {
	if o == nil || o.Department == nil {
		var ret string
		return ret
	}
	return *o.Department
}

// GetDepartmentOk returns a tuple with the Department field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetDepartmentOk() (*string, bool) {
	if o == nil || o.Department == nil {
		return nil, false
	}
	return o.Department, true
}

// HasDepartment returns a boolean if a field has been set.
func (o *Address) HasDepartment() bool {
	if o != nil && o.Department != nil {
		return true
	}

	return false
}

// SetDepartment gets a reference to the given string and assigns it to the Department field.
func (o *Address) SetDepartment(v string) {
	o.Department = &v
}

// GetSubDepartment returns the SubDepartment field value if set, zero value otherwise.
func (o *Address) GetSubDepartment() string {
	if o == nil || o.SubDepartment == nil {
		var ret string
		return ret
	}
	return *o.SubDepartment
}

// GetSubDepartmentOk returns a tuple with the SubDepartment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetSubDepartmentOk() (*string, bool) {
	if o == nil || o.SubDepartment == nil {
		return nil, false
	}
	return o.SubDepartment, true
}

// HasSubDepartment returns a boolean if a field has been set.
func (o *Address) HasSubDepartment() bool {
	if o != nil && o.SubDepartment != nil {
		return true
	}

	return false
}

// SetSubDepartment gets a reference to the given string and assigns it to the SubDepartment field.
func (o *Address) SetSubDepartment(v string) {
	o.SubDepartment = &v
}

// GetStreetName returns the StreetName field value if set, zero value otherwise.
func (o *Address) GetStreetName() string {
	if o == nil || o.StreetName == nil {
		var ret string
		return ret
	}
	return *o.StreetName
}

// GetStreetNameOk returns a tuple with the StreetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetStreetNameOk() (*string, bool) {
	if o == nil || o.StreetName == nil {
		return nil, false
	}
	return o.StreetName, true
}

// HasStreetName returns a boolean if a field has been set.
func (o *Address) HasStreetName() bool {
	if o != nil && o.StreetName != nil {
		return true
	}

	return false
}

// SetStreetName gets a reference to the given string and assigns it to the StreetName field.
func (o *Address) SetStreetName(v string) {
	o.StreetName = &v
}

// GetBuildingNumber returns the BuildingNumber field value if set, zero value otherwise.
func (o *Address) GetBuildingNumber() string {
	if o == nil || o.BuildingNumber == nil {
		var ret string
		return ret
	}
	return *o.BuildingNumber
}

// GetBuildingNumberOk returns a tuple with the BuildingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetBuildingNumberOk() (*string, bool) {
	if o == nil || o.BuildingNumber == nil {
		return nil, false
	}
	return o.BuildingNumber, true
}

// HasBuildingNumber returns a boolean if a field has been set.
func (o *Address) HasBuildingNumber() bool {
	if o != nil && o.BuildingNumber != nil {
		return true
	}

	return false
}

// SetBuildingNumber gets a reference to the given string and assigns it to the BuildingNumber field.
func (o *Address) SetBuildingNumber(v string) {
	o.BuildingNumber = &v
}

// GetBuildingName returns the BuildingName field value if set, zero value otherwise.
func (o *Address) GetBuildingName() string {
	if o == nil || o.BuildingName == nil {
		var ret string
		return ret
	}
	return *o.BuildingName
}

// GetBuildingNameOk returns a tuple with the BuildingName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetBuildingNameOk() (*string, bool) {
	if o == nil || o.BuildingName == nil {
		return nil, false
	}
	return o.BuildingName, true
}

// HasBuildingName returns a boolean if a field has been set.
func (o *Address) HasBuildingName() bool {
	if o != nil && o.BuildingName != nil {
		return true
	}

	return false
}

// SetBuildingName gets a reference to the given string and assigns it to the BuildingName field.
func (o *Address) SetBuildingName(v string) {
	o.BuildingName = &v
}

// GetFloor returns the Floor field value if set, zero value otherwise.
func (o *Address) GetFloor() string {
	if o == nil || o.Floor == nil {
		var ret string
		return ret
	}
	return *o.Floor
}

// GetFloorOk returns a tuple with the Floor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetFloorOk() (*string, bool) {
	if o == nil || o.Floor == nil {
		return nil, false
	}
	return o.Floor, true
}

// HasFloor returns a boolean if a field has been set.
func (o *Address) HasFloor() bool {
	if o != nil && o.Floor != nil {
		return true
	}

	return false
}

// SetFloor gets a reference to the given string and assigns it to the Floor field.
func (o *Address) SetFloor(v string) {
	o.Floor = &v
}

// GetPostBox returns the PostBox field value if set, zero value otherwise.
func (o *Address) GetPostBox() string {
	if o == nil || o.PostBox == nil {
		var ret string
		return ret
	}
	return *o.PostBox
}

// GetPostBoxOk returns a tuple with the PostBox field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetPostBoxOk() (*string, bool) {
	if o == nil || o.PostBox == nil {
		return nil, false
	}
	return o.PostBox, true
}

// HasPostBox returns a boolean if a field has been set.
func (o *Address) HasPostBox() bool {
	if o != nil && o.PostBox != nil {
		return true
	}

	return false
}

// SetPostBox gets a reference to the given string and assigns it to the PostBox field.
func (o *Address) SetPostBox(v string) {
	o.PostBox = &v
}

// GetRoom returns the Room field value if set, zero value otherwise.
func (o *Address) GetRoom() string {
	if o == nil || o.Room == nil {
		var ret string
		return ret
	}
	return *o.Room
}

// GetRoomOk returns a tuple with the Room field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetRoomOk() (*string, bool) {
	if o == nil || o.Room == nil {
		return nil, false
	}
	return o.Room, true
}

// HasRoom returns a boolean if a field has been set.
func (o *Address) HasRoom() bool {
	if o != nil && o.Room != nil {
		return true
	}

	return false
}

// SetRoom gets a reference to the given string and assigns it to the Room field.
func (o *Address) SetRoom(v string) {
	o.Room = &v
}

// GetPostCode returns the PostCode field value if set, zero value otherwise.
func (o *Address) GetPostCode() string {
	if o == nil || o.PostCode == nil {
		var ret string
		return ret
	}
	return *o.PostCode
}

// GetPostCodeOk returns a tuple with the PostCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetPostCodeOk() (*string, bool) {
	if o == nil || o.PostCode == nil {
		return nil, false
	}
	return o.PostCode, true
}

// HasPostCode returns a boolean if a field has been set.
func (o *Address) HasPostCode() bool {
	if o != nil && o.PostCode != nil {
		return true
	}

	return false
}

// SetPostCode gets a reference to the given string and assigns it to the PostCode field.
func (o *Address) SetPostCode(v string) {
	o.PostCode = &v
}

// GetTownName returns the TownName field value if set, zero value otherwise.
func (o *Address) GetTownName() string {
	if o == nil || o.TownName == nil {
		var ret string
		return ret
	}
	return *o.TownName
}

// GetTownNameOk returns a tuple with the TownName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetTownNameOk() (*string, bool) {
	if o == nil || o.TownName == nil {
		return nil, false
	}
	return o.TownName, true
}

// HasTownName returns a boolean if a field has been set.
func (o *Address) HasTownName() bool {
	if o != nil && o.TownName != nil {
		return true
	}

	return false
}

// SetTownName gets a reference to the given string and assigns it to the TownName field.
func (o *Address) SetTownName(v string) {
	o.TownName = &v
}

// GetTownLocationName returns the TownLocationName field value if set, zero value otherwise.
func (o *Address) GetTownLocationName() string {
	if o == nil || o.TownLocationName == nil {
		var ret string
		return ret
	}
	return *o.TownLocationName
}

// GetTownLocationNameOk returns a tuple with the TownLocationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetTownLocationNameOk() (*string, bool) {
	if o == nil || o.TownLocationName == nil {
		return nil, false
	}
	return o.TownLocationName, true
}

// HasTownLocationName returns a boolean if a field has been set.
func (o *Address) HasTownLocationName() bool {
	if o != nil && o.TownLocationName != nil {
		return true
	}

	return false
}

// SetTownLocationName gets a reference to the given string and assigns it to the TownLocationName field.
func (o *Address) SetTownLocationName(v string) {
	o.TownLocationName = &v
}

// GetDistrictName returns the DistrictName field value if set, zero value otherwise.
func (o *Address) GetDistrictName() string {
	if o == nil || o.DistrictName == nil {
		var ret string
		return ret
	}
	return *o.DistrictName
}

// GetDistrictNameOk returns a tuple with the DistrictName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetDistrictNameOk() (*string, bool) {
	if o == nil || o.DistrictName == nil {
		return nil, false
	}
	return o.DistrictName, true
}

// HasDistrictName returns a boolean if a field has been set.
func (o *Address) HasDistrictName() bool {
	if o != nil && o.DistrictName != nil {
		return true
	}

	return false
}

// SetDistrictName gets a reference to the given string and assigns it to the DistrictName field.
func (o *Address) SetDistrictName(v string) {
	o.DistrictName = &v
}

// GetCountrySubDivision returns the CountrySubDivision field value if set, zero value otherwise.
func (o *Address) GetCountrySubDivision() string {
	if o == nil || o.CountrySubDivision == nil {
		var ret string
		return ret
	}
	return *o.CountrySubDivision
}

// GetCountrySubDivisionOk returns a tuple with the CountrySubDivision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetCountrySubDivisionOk() (*string, bool) {
	if o == nil || o.CountrySubDivision == nil {
		return nil, false
	}
	return o.CountrySubDivision, true
}

// HasCountrySubDivision returns a boolean if a field has been set.
func (o *Address) HasCountrySubDivision() bool {
	if o != nil && o.CountrySubDivision != nil {
		return true
	}

	return false
}

// SetCountrySubDivision gets a reference to the given string and assigns it to the CountrySubDivision field.
func (o *Address) SetCountrySubDivision(v string) {
	o.CountrySubDivision = &v
}

// GetAddressLine returns the AddressLine field value if set, zero value otherwise.
func (o *Address) GetAddressLine() []string {
	if o == nil || o.AddressLine == nil {
		var ret []string
		return ret
	}
	return *o.AddressLine
}

// GetAddressLineOk returns a tuple with the AddressLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetAddressLineOk() (*[]string, bool) {
	if o == nil || o.AddressLine == nil {
		return nil, false
	}
	return o.AddressLine, true
}

// HasAddressLine returns a boolean if a field has been set.
func (o *Address) HasAddressLine() bool {
	if o != nil && o.AddressLine != nil {
		return true
	}

	return false
}

// SetAddressLine gets a reference to the given []string and assigns it to the AddressLine field.
func (o *Address) SetAddressLine(v []string) {
	o.AddressLine = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *Address) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *Address) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *Address) SetCountry(v string) {
	o.Country = &v
}

func (o Address) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AddressType != nil {
		toSerialize["address_type"] = o.AddressType
	}
	if o.Department != nil {
		toSerialize["department"] = o.Department
	}
	if o.SubDepartment != nil {
		toSerialize["sub_department"] = o.SubDepartment
	}
	if o.StreetName != nil {
		toSerialize["street_name"] = o.StreetName
	}
	if o.BuildingNumber != nil {
		toSerialize["building_number"] = o.BuildingNumber
	}
	if o.BuildingName != nil {
		toSerialize["building_name"] = o.BuildingName
	}
	if o.Floor != nil {
		toSerialize["floor"] = o.Floor
	}
	if o.PostBox != nil {
		toSerialize["post_box"] = o.PostBox
	}
	if o.Room != nil {
		toSerialize["room"] = o.Room
	}
	if o.PostCode != nil {
		toSerialize["post_code"] = o.PostCode
	}
	if o.TownName != nil {
		toSerialize["town_name"] = o.TownName
	}
	if o.TownLocationName != nil {
		toSerialize["town_location_name"] = o.TownLocationName
	}
	if o.DistrictName != nil {
		toSerialize["district_name"] = o.DistrictName
	}
	if o.CountrySubDivision != nil {
		toSerialize["country_sub_division"] = o.CountrySubDivision
	}
	if o.AddressLine != nil {
		toSerialize["address_line"] = o.AddressLine
	}
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}
	return json.Marshal(toSerialize)
}

type NullableAddress struct {
	value *Address
	isSet bool
}

func (v NullableAddress) Get() *Address {
	return v.value
}

func (v *NullableAddress) Set(val *Address) {
	v.value = val
	v.isSet = true
}

func (v NullableAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddress(val *Address) *NullableAddress {
	return &NullableAddress{value: val, isSet: true}
}

func (v NullableAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

