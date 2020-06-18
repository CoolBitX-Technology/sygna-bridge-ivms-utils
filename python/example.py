from __init__ import ivms, normalized_json_output

# Originator
# originator name id
org_name_id = ivms.NaturalPersonNameId()
org_name_id.primary_identifier = "Wu"
org_name_id.secondary_identifier = "Xinli"
org_name_id.name_identifier_type = ivms.NATURAL_PERSON_NAME_TYPE_CODE_LEGL

# another name id for originator
org_name_local_id = ivms.NaturalPersonNameId()
org_name_local_id.primary_identifier = "吳"
org_name_local_id.secondary_identifier = "信利"
org_name_local_id.name_identifier_type = ivms.NATURAL_PERSON_NAME_TYPE_CODE_LEGL

# assign two name id to originator name
org_person_name = ivms.NaturalPersonName()
org_person_name.name_identifiers.append(org_name_id)
org_person_name.name_identifiers.append(org_name_local_id)

# originator national id and data
org_person_national_id = ivms.NationalIdentification()
org_person_national_id.national_identifier = "446005"
org_person_national_id.national_identifier_type = ivms.NATIONAL_IDENTIFIER_TYPE_CODE_RAID
org_person_national_id.registration_authority = "RA000553"

# assign name, national id, country to originator natural person
org_nature_person = ivms.NaturalPerson()
org_nature_person.name.CopyFrom(org_person_name)
org_nature_person.national_identification.CopyFrom(org_person_national_id)
org_nature_person.country_of_residence = "TZ"

# assign originator to person object
org_persion = ivms.Person()
org_persion.natural_person.CopyFrom(org_nature_person)

# assign originator person and account id to originator
originator = ivms.Originator()
originator.originator_persons.append(org_persion)
originator.account_numbers.append("1BvBMSEYstWetqTFn5Au4m4GFg7xJaNVN2")

# Beneficiary
# beneficiary 1 name id
bene_1_name_id = ivms.LegalPersonNameId()
bene_1_name_id.legal_person_name = "ABC Limited"
bene_1_name_id.legal_person_name_identifier_type = ivms.LEGAL_PERSON_NAME_TYPE_CODE_LEGL

# assign beneficiary 1 name with id
bene_1_person_name = ivms.LegalPersonName()
bene_1_person_name.name_identifiers.append(bene_1_name_id)

# beneficiary 1 is a legal person
bene_1_legal_person = ivms.LegalPerson()
bene_1_legal_person.name.CopyFrom(bene_1_person_name)

# assign beneficiary 1 to person object
bene_1_person = ivms.Person()
bene_1_person.legal_person.CopyFrom(bene_1_legal_person)

# beneficiary 2 name id
bene_2_name_id = ivms.LegalPersonNameId()
bene_2_name_id.legal_person_name = "CBA Trading"
bene_2_name_id.legal_person_name_identifier_type = ivms.LEGAL_PERSON_NAME_TYPE_CODE_TRAD

# assign beneficiary 2 name with id
bene_2_person_name = ivms.LegalPersonName()
bene_2_person_name.name_identifiers.append(bene_2_name_id)

# beneficiary 2 is a legal person
bene_2_legal_person = ivms.LegalPerson()
bene_2_legal_person.name.CopyFrom(bene_2_person_name)

# assign beneficiary 2 to person object
bene_2_person = ivms.Person()
bene_2_person.legal_person.CopyFrom(bene_2_legal_person)

# assign both persons to beneficiary object
beneficiary = ivms.Beneficiary()
beneficiary.beneficiary_persons.append(bene_1_person)
beneficiary.beneficiary_persons.append(bene_2_person)

# assign originator and beneficiary data to identity payload
private_info = ivms.IdentityPayload()
private_info.originator.CopyFrom(originator)
private_info.beneficiary.CopyFrom(beneficiary)

# you can call this function to get JSON representation of payload
print(normalized_json_output(private_info))

# you can also get protobuf binary data of payload
print(private_info.SerializeToString())
