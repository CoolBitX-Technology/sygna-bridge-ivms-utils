package org.openapitools.client.model;

import org.openapitools.client.model.*;
import com.google.gson.Gson;
import com.google.gson.GsonBuilder;

public class Example {
  public static void main(String[] args) {
    // Originator
    // originator name id
    NaturalPersonNameId orgNameId = new NaturalPersonNameId();
    orgNameId.setPrimaryIdentifier("Wu");
    orgNameId.setSecondaryIdentifier("Xinli");
    orgNameId.setNameIdentifierType(NaturalPersonNameId.NameIdentifierTypeEnum.LEGL);

    // another name id for originator
    LocalNaturalPersonNameId orgNameIdLocal = new LocalNaturalPersonNameId();
    orgNameIdLocal.setPrimaryIdentifier("吳");
    orgNameIdLocal.setSecondaryIdentifier("信利");
    orgNameIdLocal.setNameIdentifierType(LocalNaturalPersonNameId.NameIdentifierTypeEnum.LEGL);

    // assign two name id to originator name
    NaturalPersonName orgPersonName = new NaturalPersonName();
    orgPersonName.addNameIdentifiersItem(orgNameId);
    orgPersonName.addLocalNameIdentifiersItem(orgNameIdLocal);

    // originator national id and data
    NationalIdentification orgPersonNationalId = new NationalIdentification();
    orgPersonNationalId.setNationalIdentifier("446005");
    orgPersonNationalId.setNationalIdentifierType(NationalIdentification.NationalIdentifierTypeEnum.RAID);
    orgPersonNationalId.setRegistrationAuthority("RA000553");

    // assign name, national id, country to originator natural person
    NaturalPerson originatingNaturalPerson = new NaturalPerson();
    originatingNaturalPerson.setName(orgPersonName);
    originatingNaturalPerson.setNationalIdentification(orgPersonNationalId);
    originatingNaturalPerson.setCountryOfResidence("TZ");

    // assign originator to person object
    Person originatingPerson = new Person();
    originatingPerson.setNaturalPerson(originatingNaturalPerson);

    // assign originator person and account id to originator
    Originator originator = new Originator();
    originator.addOriginatorPersonsItem(originatingPerson);
    originator.addAccountNumbersItem("1BvBMSEYstWetqTFn5Au4m4GFg7xJaNVN2");

    // Beneficiary
    // beneficiary 1 name id
    LegalPersonNameId bene1NameId = new LegalPersonNameId();
    bene1NameId.setLegalPersonName("ABC Limited");
    bene1NameId.setLegalPersonNameIdentifierType(LegalPersonNameId.LegalPersonNameIdentifierTypeEnum.LEGL);

    // assign beneficiary 1 name with id
    LegalPersonName bene1PersonName = new LegalPersonName();
    bene1PersonName.addNameIdentifiersItem(bene1NameId);

    // beneficiary 1 is a legal person
    LegalPerson beneficiary1LegalPerson = new LegalPerson();
    beneficiary1LegalPerson.setName(bene1PersonName);

    // assign beneficiary 1 to person object
    Person beneficiary1Person = new Person();
    beneficiary1Person.setLegalPerson(beneficiary1LegalPerson);

    // beneficiary 2 name id
    LegalPersonNameId bene2NameId = new LegalPersonNameId();
    bene2NameId.setLegalPersonName("CBA Trading");
    bene2NameId.setLegalPersonNameIdentifierType(LegalPersonNameId.LegalPersonNameIdentifierTypeEnum.TRAD);

    // assign beneficiary 2 name with id
    LegalPersonName bene2PersonName = new LegalPersonName();
    bene2PersonName.addNameIdentifiersItem(bene2NameId);

    // beneficiary 2 is a legal person
    LegalPerson beneficiary2LegalPerson = new LegalPerson();
    beneficiary2LegalPerson.setName(bene2PersonName);

    // assign beneficiary 2 to person object
    Person beneficiary2Person = new Person();
    beneficiary2Person.setLegalPerson(beneficiary2LegalPerson);

    // assign both persons to beneficiary object
    Beneficiary beneficiary = new Beneficiary();
    beneficiary.addBeneficiaryPersonsItem(beneficiary1Person);
    beneficiary.addBeneficiaryPersonsItem(beneficiary2Person);

    // assign originator and beneficiary data to identity payload
    IdentityPayload privateInfo = new IdentityPayload();
    privateInfo.setOriginator(originator);
    privateInfo.setBeneficiary(beneficiary);

    Gson gson = new GsonBuilder().setPrettyPrinting().create();
    System.out.println(gson.toJson(privateInfo));
  }
}