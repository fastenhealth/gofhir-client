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

// EvidenceVariable is documented here http://hl7.org/fhir/StructureDefinition/EvidenceVariable
// The EvidenceVariable resource describes a "PICO" element that knowledge (evidence, assertion, recommendation) is about.
type EvidenceVariable struct {
	Id                *string                          `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                            `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                          `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                          `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                       `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension                      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url               *string                          `bson:"url,omitempty" json:"url,omitempty"`
	Identifier        []Identifier                     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version           *string                          `bson:"version,omitempty" json:"version,omitempty"`
	Name              *string                          `bson:"name,omitempty" json:"name,omitempty"`
	Title             *string                          `bson:"title,omitempty" json:"title,omitempty"`
	ShortTitle        *string                          `bson:"shortTitle,omitempty" json:"shortTitle,omitempty"`
	Subtitle          *string                          `bson:"subtitle,omitempty" json:"subtitle,omitempty"`
	Status            PublicationStatus                `bson:"status" json:"status"`
	Date              *string                          `bson:"date,omitempty" json:"date,omitempty"`
	Publisher         *string                          `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact           []ContactDetail                  `bson:"contact,omitempty" json:"contact,omitempty"`
	Description       *string                          `bson:"description,omitempty" json:"description,omitempty"`
	Note              []Annotation                     `bson:"note,omitempty" json:"note,omitempty"`
	UseContext        []UsageContext                   `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept                `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Copyright         *string                          `bson:"copyright,omitempty" json:"copyright,omitempty"`
	ApprovalDate      *string                          `bson:"approvalDate,omitempty" json:"approvalDate,omitempty"`
	LastReviewDate    *string                          `bson:"lastReviewDate,omitempty" json:"lastReviewDate,omitempty"`
	EffectivePeriod   *Period                          `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	Topic             []CodeableConcept                `bson:"topic,omitempty" json:"topic,omitempty"`
	Author            []ContactDetail                  `bson:"author,omitempty" json:"author,omitempty"`
	Editor            []ContactDetail                  `bson:"editor,omitempty" json:"editor,omitempty"`
	Reviewer          []ContactDetail                  `bson:"reviewer,omitempty" json:"reviewer,omitempty"`
	Endorser          []ContactDetail                  `bson:"endorser,omitempty" json:"endorser,omitempty"`
	RelatedArtifact   []RelatedArtifact                `bson:"relatedArtifact,omitempty" json:"relatedArtifact,omitempty"`
	Type              *EvidenceVariableType            `bson:"type,omitempty" json:"type,omitempty"`
	Characteristic    []EvidenceVariableCharacteristic `bson:"characteristic" json:"characteristic"`
}

// A characteristic that defines the members of the evidence element. Multiple characteristics are applied with "and" semantics.
// Characteristics can be defined flexibly to accommodate different use cases for membership criteria, ranging from simple codes, all the way to using an expression language to express the criteria.
type EvidenceVariableCharacteristic struct {
	Id                *string        `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Description       *string        `bson:"description,omitempty" json:"description,omitempty"`
	UsageContext      []UsageContext `bson:"usageContext,omitempty" json:"usageContext,omitempty"`
	Exclude           *bool          `bson:"exclude,omitempty" json:"exclude,omitempty"`
	TimeFromStart     *Duration      `bson:"timeFromStart,omitempty" json:"timeFromStart,omitempty"`
	GroupMeasure      *GroupMeasure  `bson:"groupMeasure,omitempty" json:"groupMeasure,omitempty"`
}

// This function returns resource reference information
func (r EvidenceVariable) ResourceRef() (string, *string) {
	return "EvidenceVariable", r.Id
}

type OtherEvidenceVariable EvidenceVariable

// MarshalJSON marshals the given EvidenceVariable as JSON into a byte slice
func (r EvidenceVariable) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherEvidenceVariable
		ResourceType string `json:"resourceType"`
	}{
		OtherEvidenceVariable: OtherEvidenceVariable(r),
		ResourceType:          "EvidenceVariable",
	})
}

// UnmarshalEvidenceVariable unmarshals a EvidenceVariable.
func UnmarshalEvidenceVariable(b []byte) (EvidenceVariable, error) {
	var evidenceVariable EvidenceVariable
	if err := json.Unmarshal(b, &evidenceVariable); err != nil {
		return evidenceVariable, err
	}
	return evidenceVariable, nil
}
