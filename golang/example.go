package main

import (
	"fmt"

	ivms101 "github.com/CoolBitX-Technology/sygna-bridge-ivms-utils/golang/src"
)

func main() {
	// Originator
	// originator name id
	orgNameID := ivms101.NewNaturalPersonNameId()
	orgNameID.SetPrimaryIdentifier("Wu")
	orgNameID.SetSecondaryIdentifier("Xinli")
	orgNameID.SetNameIdentifierType(string(ivms101.NATURALPERSONNAMETYPECODE_LEGL))

	// another name id for originator
	orgNameIDLocal := ivms101.NewLocalNaturalPersonNameId()
	orgNameIDLocal.SetPrimaryIdentifier("吳")
	orgNameIDLocal.SetSecondaryIdentifier("信利")
	orgNameIDLocal.SetNameIdentifierType(string(ivms101.NATURALPERSONNAMETYPECODE_LEGL))

	// assign two name id to originator name
	orgPersonName := ivms101.NewNaturalPersonName()
	orgPersonName.SetNameIdentifiers([]ivms101.NaturalPersonNameId{*orgNameID})
	orgPersonName.SetLocalNameIdentifiers([]ivms101.LocalNaturalPersonNameId{*orgNameIDLocal})

	// originator national id and data
	orgPersonNationalID := ivms101.NewNationalIdentification()
	orgPersonNationalID.SetNationalIdentifier("446005")
	orgPersonNationalID.SetNationalIdentifierType(string(ivms101.NATIONALIDENTIFIERTYPECODE_RAID))
	orgPersonNationalID.SetRegistrationAuthority("RA000553")

	// assign name, national id, country to originator natural person
	originatingNaturalPerson := ivms101.NewNaturalPerson()
	originatingNaturalPerson.SetName(*orgPersonName)
	originatingNaturalPerson.SetNationalIdentification(*orgPersonNationalID)
	originatingNaturalPerson.SetCountryOfResidence("TZ")

	// assign originator to person object
	originatingPerson := ivms101.NewPerson()
	originatingPerson.SetNaturalPerson(*originatingNaturalPerson)

	// assign originator person and account id to originator
	originator := ivms101.NewOriginator()
	originator.SetOriginatorPersons([]ivms101.Person{*originatingPerson})
	originator.SetAccountNumbers([]string{"1BvBMSEYstWetqTFn5Au4m4GFg7xJaNVN2"})

	// Beneficiary
	// beneficiary 1 name id
	bene1NameID := ivms101.NewLegalPersonNameId()
	bene1NameID.SetLegalPersonName("ABC Limited")
	bene1NameID.SetLegalPersonNameIdentifierType(string(ivms101.LEGALPERSONNAMETYPECODE_LEGL))

	// assign beneficiary 1 name with id
	bene1PersonName := ivms101.NewLegalPersonName()
	bene1PersonName.SetNameIdentifiers([]ivms101.LegalPersonNameId{*bene1NameID})

	// beneficiary 1 is a legal person
	beneficiary1LegalPerson := ivms101.NewLegalPerson()
	beneficiary1LegalPerson.SetName(*bene1PersonName)

	// assign beneficiary 1 to person object
	beneficiary1Person := ivms101.NewPerson()
	beneficiary1Person.SetLegalPerson(*beneficiary1LegalPerson)

	// beneficiary 2 name id
	bene2NameID := ivms101.NewLegalPersonNameId()
	bene2NameID.SetLegalPersonName("CBA Trading")
	bene2NameID.SetLegalPersonNameIdentifierType(string(ivms101.LEGALPERSONNAMETYPECODE_TRAD))

	// assign beneficiary 2 name with id
	bene2PersonName := ivms101.NewLegalPersonName()
	bene2PersonName.SetNameIdentifiers([]ivms101.LegalPersonNameId{*bene2NameID})

	// beneficiary 2 is a legal person
	beneficiary2LegalPerson := ivms101.NewLegalPerson()
	beneficiary2LegalPerson.SetName(*bene2PersonName)

	// assign beneficiary 2 to person object
	beneficiary2Person := ivms101.NewPerson()
	beneficiary2Person.SetLegalPerson(*beneficiary2LegalPerson)

	// assign both persons to beneficiary object
	beneficiary := ivms101.NewBeneficiary()
	beneficiary.SetBeneficiaryPersons([]ivms101.Person{*beneficiary1Person, *beneficiary2Person})

	// assign originator and beneficiary data to identity payload
	privateInfo := ivms101.NewIdentityPayload()
	privateInfo.SetOriginator(*originator)
	privateInfo.SetBeneficiary(*beneficiary)

	// to json
	jsonb, _ := privateInfo.MarshalJSON()
	fmt.Print(string(jsonb))

	// from json
	decodedPayload := ivms101.NewNullableIdentityPayload(nil)
	decodedPayload.UnmarshalJSON(jsonb)

	jsonb2, _ := decodedPayload.MarshalJSON()
	fmt.Print("\n\n" + string(jsonb2))
}
