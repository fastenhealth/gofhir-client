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

// CareTeam is documented here http://hl7.org/fhir/StructureDefinition/CareTeam
// The Care Team includes all the people and organizations who plan to participate in the coordination and delivery of care for a patient.
type CareTeam struct {
	Id                   *string               `bson:"id,omitempty" json:"id,omitempty"`
	Meta                 *Meta                 `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules        *string               `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language             *string               `bson:"language,omitempty" json:"language,omitempty"`
	Text                 *Narrative            `bson:"text,omitempty" json:"text,omitempty"`
	Extension            []Extension           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []Extension           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier           []Identifier          `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status               *CareTeamStatus       `bson:"status,omitempty" json:"status,omitempty"`
	Category             []CodeableConcept     `bson:"category,omitempty" json:"category,omitempty"`
	Name                 *string               `bson:"name,omitempty" json:"name,omitempty"`
	Subject              *Reference            `bson:"subject,omitempty" json:"subject,omitempty"`
	Encounter            *Reference            `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Period               *Period               `bson:"period,omitempty" json:"period,omitempty"`
	Participant          []CareTeamParticipant `bson:"participant,omitempty" json:"participant,omitempty"`
	ReasonCode           []CodeableConcept     `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	ReasonReference      []Reference           `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	ManagingOrganization []Reference           `bson:"managingOrganization,omitempty" json:"managingOrganization,omitempty"`
	Telecom              []ContactPoint        `bson:"telecom,omitempty" json:"telecom,omitempty"`
	Note                 []Annotation          `bson:"note,omitempty" json:"note,omitempty"`
}

// Identifies all people and organizations who are expected to be involved in the care team.
type CareTeamParticipant struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Role              []CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
	Member            *Reference        `bson:"member,omitempty" json:"member,omitempty"`
	OnBehalfOf        *Reference        `bson:"onBehalfOf,omitempty" json:"onBehalfOf,omitempty"`
	Period            *Period           `bson:"period,omitempty" json:"period,omitempty"`
}

// This function returns resource reference information
func (r CareTeam) ResourceRef() (string, *string) {
	return "CareTeam", r.Id
}

type OtherCareTeam CareTeam

// MarshalJSON marshals the given CareTeam as JSON into a byte slice
func (r CareTeam) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCareTeam
		ResourceType string `json:"resourceType"`
	}{
		OtherCareTeam: OtherCareTeam(r),
		ResourceType:  "CareTeam",
	})
}

// UnmarshalCareTeam unmarshals a CareTeam.
func UnmarshalCareTeam(b []byte) (CareTeam, error) {
	var careTeam CareTeam
	if err := json.Unmarshal(b, &careTeam); err != nil {
		return careTeam, err
	}
	return careTeam, nil
}
