/*
Bridge

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ivms101

import (
	"encoding/json"
	"fmt"
)

// LegalPersonNameTypeCode the model 'LegalPersonNameTypeCode'
type LegalPersonNameTypeCode string

// List of LegalPersonNameTypeCode
const (
	LEGALPERSONNAMETYPECODE_LEGL LegalPersonNameTypeCode = "LEGL"
	LEGALPERSONNAMETYPECODE_SHRT LegalPersonNameTypeCode = "SHRT"
	LEGALPERSONNAMETYPECODE_TRAD LegalPersonNameTypeCode = "TRAD"
)

// All allowed values of LegalPersonNameTypeCode enum
var AllowedLegalPersonNameTypeCodeEnumValues = []LegalPersonNameTypeCode{
	"LEGL",
	"SHRT",
	"TRAD",
}

func (v *LegalPersonNameTypeCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LegalPersonNameTypeCode(value)
	for _, existing := range AllowedLegalPersonNameTypeCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LegalPersonNameTypeCode", value)
}

// NewLegalPersonNameTypeCodeFromValue returns a pointer to a valid LegalPersonNameTypeCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLegalPersonNameTypeCodeFromValue(v string) (*LegalPersonNameTypeCode, error) {
	ev := LegalPersonNameTypeCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LegalPersonNameTypeCode: valid values are %v", v, AllowedLegalPersonNameTypeCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LegalPersonNameTypeCode) IsValid() bool {
	for _, existing := range AllowedLegalPersonNameTypeCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LegalPersonNameTypeCode value
func (v LegalPersonNameTypeCode) Ptr() *LegalPersonNameTypeCode {
	return &v
}

type NullableLegalPersonNameTypeCode struct {
	value *LegalPersonNameTypeCode
	isSet bool
}

func (v NullableLegalPersonNameTypeCode) Get() *LegalPersonNameTypeCode {
	return v.value
}

func (v *NullableLegalPersonNameTypeCode) Set(val *LegalPersonNameTypeCode) {
	v.value = val
	v.isSet = true
}

func (v NullableLegalPersonNameTypeCode) IsSet() bool {
	return v.isSet
}

func (v *NullableLegalPersonNameTypeCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLegalPersonNameTypeCode(val *LegalPersonNameTypeCode) *NullableLegalPersonNameTypeCode {
	return &NullableLegalPersonNameTypeCode{value: val, isSet: true}
}

func (v NullableLegalPersonNameTypeCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLegalPersonNameTypeCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

