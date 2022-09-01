// Copyright 2022 - Fasten Health
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fhir430

import "encoding/json"

// THIS FILE IS GENERATED BY https://github.com/fastenhealth/gofhir-models-gen
// PLEASE DO NOT EDIT BY HAND

// Patient is documented here http://hl7.org/fhir/StructureDefinition/Patient
// Demographics and other administrative information about an individual or animal receiving care or other health-related services.
type Patient struct {
	Id                   *string                `bson:"id,omitempty" json:"id,omitempty"`
	Meta                 *Meta                  `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules        *string                `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language             *string                `bson:"language,omitempty" json:"language,omitempty"`
	Text                 *Narrative             `bson:"text,omitempty" json:"text,omitempty"`
	Extension            []Extension            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []Extension            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier           []Identifier           `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Active               *bool                  `bson:"active,omitempty" json:"active,omitempty"`
	Name                 []HumanName            `bson:"name,omitempty" json:"name,omitempty"`
	Telecom              []ContactPoint         `bson:"telecom,omitempty" json:"telecom,omitempty"`
	Gender               *AdministrativeGender  `bson:"gender,omitempty" json:"gender,omitempty"`
	BirthDate            *string                `bson:"birthDate,omitempty" json:"birthDate,omitempty"`
	Address              []Address              `bson:"address,omitempty" json:"address,omitempty"`
	MaritalStatus        *CodeableConcept       `bson:"maritalStatus,omitempty" json:"maritalStatus,omitempty"`
	Photo                []Attachment           `bson:"photo,omitempty" json:"photo,omitempty"`
	Contact              []PatientContact       `bson:"contact,omitempty" json:"contact,omitempty"`
	Communication        []PatientCommunication `bson:"communication,omitempty" json:"communication,omitempty"`
	GeneralPractitioner  []Reference            `bson:"generalPractitioner,omitempty" json:"generalPractitioner,omitempty"`
	ManagingOrganization *Reference             `bson:"managingOrganization,omitempty" json:"managingOrganization,omitempty"`
	Link                 []PatientLink          `bson:"link,omitempty" json:"link,omitempty"`
}

// A contact party (e.g. guardian, partner, friend) for the patient.
// Contact covers all kinds of contact parties: family members, business contacts, guardians, caregivers. Not applicable to register pedigree and family ties beyond use of having contact.
type PatientContact struct {
	Id                *string               `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Relationship      []CodeableConcept     `bson:"relationship,omitempty" json:"relationship,omitempty"`
	Name              *HumanName            `bson:"name,omitempty" json:"name,omitempty"`
	Telecom           []ContactPoint        `bson:"telecom,omitempty" json:"telecom,omitempty"`
	Address           *Address              `bson:"address,omitempty" json:"address,omitempty"`
	Gender            *AdministrativeGender `bson:"gender,omitempty" json:"gender,omitempty"`
	Organization      *Reference            `bson:"organization,omitempty" json:"organization,omitempty"`
	Period            *Period               `bson:"period,omitempty" json:"period,omitempty"`
}

// A language which may be used to communicate with the patient about his or her health.
// If no language is specified, this *implies* that the default local language is spoken.  If you need to convey proficiency for multiple modes, then you need multiple Patient.Communication associations.   For animals, language is not a relevant field, and should be absent from the instance. If the Patient does not speak the default local language, then the Interpreter Required Standard can be used to explicitly declare that an interpreter is required.
type PatientCommunication struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Language          CodeableConcept `bson:"language" json:"language"`
	Preferred         *bool           `bson:"preferred,omitempty" json:"preferred,omitempty"`
}

// Link to another patient resource that concerns the same actual patient.
// There is no assumption that linked patient records have mutual links.
type PatientLink struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Other             Reference   `bson:"other" json:"other"`
	Type              LinkType    `bson:"type" json:"type"`
}

// This function returns resource reference information
func (r Patient) ResourceRef() (string, *string) {
	return "Patient", r.Id
}

type OtherPatient Patient

// MarshalJSON marshals the given Patient as JSON into a byte slice
func (r Patient) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherPatient
		ResourceType string `json:"resourceType"`
	}{
		OtherPatient: OtherPatient(r),
		ResourceType: "Patient",
	})
}

// UnmarshalPatient unmarshals a Patient.
func UnmarshalPatient(b []byte) (Patient, error) {
	var patient Patient
	if err := json.Unmarshal(b, &patient); err != nil {
		return patient, err
	}
	return patient, nil
}
