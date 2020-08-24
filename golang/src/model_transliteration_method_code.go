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
	"fmt"
)

// TransliterationMethodCode the model 'TransliterationMethodCode'
type TransliterationMethodCode string

// List of TransliterationMethodCode
const (
	TRANSLITERATIONMETHODCODE_ARAB TransliterationMethodCode = "ARAB"
	TRANSLITERATIONMETHODCODE_ARAN TransliterationMethodCode = "ARAN"
	TRANSLITERATIONMETHODCODE_ARMN TransliterationMethodCode = "ARMN"
	TRANSLITERATIONMETHODCODE_CYRL TransliterationMethodCode = "CYRL"
	TRANSLITERATIONMETHODCODE_DEVA TransliterationMethodCode = "DEVA"
	TRANSLITERATIONMETHODCODE_GEOR TransliterationMethodCode = "GEOR"
	TRANSLITERATIONMETHODCODE_GREK TransliterationMethodCode = "GREK"
	TRANSLITERATIONMETHODCODE_HANI TransliterationMethodCode = "HANI"
	TRANSLITERATIONMETHODCODE_HEBR TransliterationMethodCode = "HEBR"
	TRANSLITERATIONMETHODCODE_KANA TransliterationMethodCode = "KANA"
	TRANSLITERATIONMETHODCODE_KORE TransliterationMethodCode = "KORE"
	TRANSLITERATIONMETHODCODE_THAI TransliterationMethodCode = "THAI"
	TRANSLITERATIONMETHODCODE_OTHR TransliterationMethodCode = "OTHR"
)

func (v *TransliterationMethodCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TransliterationMethodCode(value)
	for _, existing := range []TransliterationMethodCode{ "ARAB", "ARAN", "ARMN", "CYRL", "DEVA", "GEOR", "GREK", "HANI", "HEBR", "KANA", "KORE", "THAI", "OTHR",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TransliterationMethodCode", value)
}

// Ptr returns reference to TransliterationMethodCode value
func (v TransliterationMethodCode) Ptr() *TransliterationMethodCode {
	return &v
}

type NullableTransliterationMethodCode struct {
	value *TransliterationMethodCode
	isSet bool
}

func (v NullableTransliterationMethodCode) Get() *TransliterationMethodCode {
	return v.value
}

func (v *NullableTransliterationMethodCode) Set(val *TransliterationMethodCode) {
	v.value = val
	v.isSet = true
}

func (v NullableTransliterationMethodCode) IsSet() bool {
	return v.isSet
}

func (v *NullableTransliterationMethodCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransliterationMethodCode(val *TransliterationMethodCode) *NullableTransliterationMethodCode {
	return &NullableTransliterationMethodCode{value: val, isSet: true}
}

func (v NullableTransliterationMethodCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransliterationMethodCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

