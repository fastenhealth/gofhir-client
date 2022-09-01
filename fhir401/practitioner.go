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

package fhir401

import "encoding/json"

// THIS FILE IS GENERATED BY https://github.com/fastenhealth/gofhir-models-gen
// PLEASE DO NOT EDIT BY HAND

// Practitioner is documented here http://hl7.org/fhir/StructureDefinition/Practitioner
// A person who is directly or indirectly involved in the provisioning of healthcare.
type Practitioner struct {
	Id                *string                     `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                       `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                     `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                     `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                  `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension                 `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                 `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier                `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Active            *bool                       `bson:"active,omitempty" json:"active,omitempty"`
	Name              []HumanName                 `bson:"name,omitempty" json:"name,omitempty"`
	Telecom           []ContactPoint              `bson:"telecom,omitempty" json:"telecom,omitempty"`
	Address           []Address                   `bson:"address,omitempty" json:"address,omitempty"`
	Gender            *AdministrativeGender       `bson:"gender,omitempty" json:"gender,omitempty"`
	BirthDate         *string                     `bson:"birthDate,omitempty" json:"birthDate,omitempty"`
	Photo             []Attachment                `bson:"photo,omitempty" json:"photo,omitempty"`
	Qualification     []PractitionerQualification `bson:"qualification,omitempty" json:"qualification,omitempty"`
	Communication     []CodeableConcept           `bson:"communication,omitempty" json:"communication,omitempty"`
}

// The official certifications, training, and licenses that authorize or otherwise pertain to the provision of care by the practitioner.  For example, a medical license issued by a medical board authorizing the practitioner to practice medicine within a certian locality.
type PractitionerQualification struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier    `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Code              CodeableConcept `bson:"code" json:"code"`
	Period            *Period         `bson:"period,omitempty" json:"period,omitempty"`
	Issuer            *Reference      `bson:"issuer,omitempty" json:"issuer,omitempty"`
}

// This function returns resource reference information
func (r Practitioner) ResourceRef() (string, *string) {
	return "Practitioner", r.Id
}

type OtherPractitioner Practitioner

// MarshalJSON marshals the given Practitioner as JSON into a byte slice
func (r Practitioner) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherPractitioner
		ResourceType string `json:"resourceType"`
	}{
		OtherPractitioner: OtherPractitioner(r),
		ResourceType:      "Practitioner",
	})
}

// UnmarshalPractitioner unmarshals a Practitioner.
func UnmarshalPractitioner(b []byte) (Practitioner, error) {
	var practitioner Practitioner
	if err := json.Unmarshal(b, &practitioner); err != nil {
		return practitioner, err
	}
	return practitioner, nil
}
