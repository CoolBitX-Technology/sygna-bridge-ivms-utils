from sygna_bridge_ivms_util import ivms
import json

# Originator
# originator name id
org_name_id = ivms.NaturalPersonNameId()
org_name_id.primary_identifier = "Wu"
org_name_id.secondary_identifier = "Xinli"
org_name_id.name_identifier_type = ivms.NaturalPersonNameTypeCode.LEGL

# another name id for originator
org_name_local_id = ivms.LocalNaturalPersonNameId()
org_name_local_id.primary_identifier = "吳"
org_name_local_id.secondary_identifier = "信利"
org_name_local_id.name_identifier_type = ivms.NaturalPersonNameTypeCode.LEGL

# assign two name id to originator name
org_person_name = ivms.NaturalPersonName()
org_person_name.name_identifiers = [org_name_id]
org_person_name.local_name_identifiers = [org_name_local_id]

# originator national id and data
org_person_national_id = ivms.NationalIdentification()
org_person_national_id.national_identifier = "446005"
org_person_national_id.national_identifier_type = ivms.NationalIdentifierTypeCode.RAID
org_person_national_id.registration_authority = "RA000553"

# assign name, national id, country to originator natural person
org_nature_person = ivms.NaturalPerson()
org_nature_person.name = org_person_name
org_nature_person.national_identification = org_person_national_id
org_nature_person.country_of_residence = "TZ"

# assign originator to person object
org_persion = ivms.Person()
org_persion.natural_person = org_nature_person

# assign originator person and account id to originator
originator = ivms.Originator()
originator.originator_persons = [org_persion]
originator.account_numbers = ["1BvBMSEYstWetqTFn5Au4m4GFg7xJaNVN2"]

# Beneficiary
# beneficiary 1 name id
bene_1_name_id = ivms.LegalPersonNameId()
bene_1_name_id.legal_person_name = "ABC Limited"
bene_1_name_id.legal_person_name_identifier_type = ivms.LegalPersonNameTypeCode.LEGL

# assign beneficiary 1 name with id
bene_1_person_name = ivms.LegalPersonName()
bene_1_person_name.name_identifiers = [bene_1_name_id]

# beneficiary 1 is a legal person
bene_1_legal_person = ivms.LegalPerson()
bene_1_legal_person.name = bene_1_person_name

# assign beneficiary 1 to person object
bene_1_person = ivms.Person()
bene_1_person.legal_person = bene_1_legal_person

# beneficiary 2 name id
bene_2_name_id = ivms.LegalPersonNameId()
bene_2_name_id.legal_person_name = "CBA Trading"
bene_2_name_id.legal_person_name_identifier_type = ivms.LegalPersonNameTypeCode.TRAD

# assign beneficiary 2 name with id
bene_2_person_name = ivms.LegalPersonName()
bene_2_person_name.name_identifiers = [bene_2_name_id]

# beneficiary 2 is a legal person
bene_2_legal_person = ivms.LegalPerson()
bene_2_legal_person.name = bene_2_person_name

# assign beneficiary 2 to person object
bene_2_person = ivms.Person()
bene_2_person.legal_person = bene_2_legal_person

# assign both persons to beneficiary object
beneficiary = ivms.Beneficiary()
beneficiary.beneficiary_persons = [bene_1_person, bene_2_person]

# assign originator and beneficiary data to identity payload
private_info = ivms.IdentityPayload()
private_info.originator = originator
private_info.beneficiary = beneficiary

# pretty print json
print(json.dumps(private_info.to_dict(), indent=2, ensure_ascii=False))
