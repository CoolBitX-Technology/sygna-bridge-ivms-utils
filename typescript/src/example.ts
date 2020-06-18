import { ivms, normalizedJsonOutput } from ".";

// Originator
// originator name id
const orgNameId = new ivms.NaturalPersonNameId();
orgNameId.setPrimaryIdentifier("Wu");
orgNameId.setSecondaryIdentifier("Xinli");
orgNameId.setNameIdentifierType(
  ivms.NaturalPersonNameTypeCode.NATURAL_PERSON_NAME_TYPE_CODE_LEGL
);

// another name id for originator
const orgNameIdLocal = new ivms.NaturalPersonNameId();
orgNameIdLocal.setPrimaryIdentifier("吳");
orgNameIdLocal.setSecondaryIdentifier("信利");
orgNameIdLocal.setNameIdentifierType(
  ivms.NaturalPersonNameTypeCode.NATURAL_PERSON_NAME_TYPE_CODE_LEGL
);

// assign two name id to originator name
const orgPersonName = new ivms.NaturalPersonName();
orgPersonName.addNameIdentifiers(orgNameId);
orgPersonName.addNameIdentifiers(orgNameIdLocal);

// originator national id and data
const orgPersonNationalId = new ivms.NationalIdentification();
orgPersonNationalId.setNationalIdentifier("446005");
orgPersonNationalId.setNationalIdentifierType(
  ivms.NationalIdentifierTypeCode.NATIONAL_IDENTIFIER_TYPE_CODE_RAID
);
orgPersonNationalId.setRegistrationAuthority("RA000553");

// assign name, national id, country to originator natural person
const originatingNaturalPerson = new ivms.NaturalPerson();
originatingNaturalPerson.setName(orgPersonName);
originatingNaturalPerson.setNationalIdentification(orgPersonNationalId);
originatingNaturalPerson.setCountryOfResidence("TZ");

// assign originator to person object
const originatingPerson = new ivms.Person();
originatingPerson.setNaturalPerson(originatingNaturalPerson);

// assign originator person and account id to originator
const originator = new ivms.Originator();
originator.addOriginatorPersons(originatingPerson);
originator.addAccountNumbers("1BvBMSEYstWetqTFn5Au4m4GFg7xJaNVN2");

// Beneficiary
// beneficiary 1 name id
const bene1NameId = new ivms.LegalPersonNameId();
bene1NameId.setLegalPersonName("ABC Limited");
bene1NameId.setLegalPersonNameIdentifierType(
  ivms.LegalPersonNameTypeCode.LEGAL_PERSON_NAME_TYPE_CODE_LEGL
);

// assign beneficiary 1 name with id
const bene1PersonName = new ivms.LegalPersonName();
bene1PersonName.addNameIdentifiers(bene1NameId);

// beneficiary 1 is a legal person
const beneficiary1LegalPerson = new ivms.LegalPerson();
beneficiary1LegalPerson.setName(bene1PersonName);

// assign beneficiary 1 to person object
const beneficiary1Person = new ivms.Person();
beneficiary1Person.setLegalPerson(beneficiary1LegalPerson);

// beneficiary 2 name id
const bene2NameId = new ivms.LegalPersonNameId();
bene2NameId.setLegalPersonName("CBA Trading");
bene2NameId.setLegalPersonNameIdentifierType(
  ivms.LegalPersonNameTypeCode.LEGAL_PERSON_NAME_TYPE_CODE_TRAD
);

// assign beneficiary 2 name with id
const bene2PersonName = new ivms.LegalPersonName();
bene2PersonName.addNameIdentifiers(bene2NameId);

// beneficiary 2 is a legal person
const beneficiary2LegalPerson = new ivms.LegalPerson();
beneficiary2LegalPerson.setName(bene2PersonName);

// assign beneficiary 2 to person object
const beneficiary2Person = new ivms.Person();
beneficiary2Person.setLegalPerson(beneficiary2LegalPerson);

// assign both persons to beneficiary object
const beneficiary = new ivms.Beneficiary();
beneficiary.addBeneficiaryPersons(beneficiary1Person);
beneficiary.addBeneficiaryPersons(beneficiary2Person);

// assign originator and beneficiary data to identity payload
const privateInfo = new ivms.IdentityPayload();
privateInfo.setOriginator(originator);
privateInfo.setBeneficiary(beneficiary);

// you can call this function to get JSON representation of payload
normalizedJsonOutput(privateInfo).then((d) =>
  console.log(JSON.stringify(d, null, 2))
);

// you can also get protobuf binary data of payload
console.log(privateInfo.serializeBinary());
