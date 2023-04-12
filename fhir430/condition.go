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

// Condition is documented here http://hl7.org/fhir/StructureDefinition/Condition
// A clinical condition, problem, diagnosis, or other event, situation, issue, or clinical concept that has risen to a level of concern.
type Condition struct {
	Id                 *string             `bson:"id,omitempty" json:"id,omitempty"`
	Meta               *Meta               `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules      *string             `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language           *string             `bson:"language,omitempty" json:"language,omitempty"`
	Text               *Narrative          `bson:"text,omitempty" json:"text,omitempty"`
	Extension          []Extension         `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension  []Extension         `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier         []Identifier        `bson:"identifier,omitempty" json:"identifier,omitempty"`
	ClinicalStatus     *CodeableConcept    `bson:"clinicalStatus,omitempty" json:"clinicalStatus,omitempty"`
	VerificationStatus *CodeableConcept    `bson:"verificationStatus,omitempty" json:"verificationStatus,omitempty"`
	Category           []CodeableConcept   `bson:"category,omitempty" json:"category,omitempty"`
	Severity           *CodeableConcept    `bson:"severity,omitempty" json:"severity,omitempty"`
	Code               *CodeableConcept    `bson:"code,omitempty" json:"code,omitempty"`
	BodySite           []CodeableConcept   `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	Subject            Reference           `bson:"subject" json:"subject"`
	Encounter          *Reference          `bson:"encounter,omitempty" json:"encounter,omitempty"`
	OnsetDateTime      *string             `bson:"onsetDateTime,omitempty" json:"onsetDateTime,omitempty"`
	OnsetAge           *Age                `bson:"onsetAge,omitempty" json:"onsetAge,omitempty"`
	OnsetPeriod        *Period             `bson:"onsetPeriod,omitempty" json:"onsetPeriod,omitempty"`
	OnsetRange         *Range              `bson:"onsetRange,omitempty" json:"onsetRange,omitempty"`
	OnsetString        *string             `bson:"onsetString,omitempty" json:"onsetString,omitempty"`
	AbatementDateTime  *string             `bson:"abatementDateTime,omitempty" json:"abatementDateTime,omitempty"`
	AbatementAge       *Age                `bson:"abatementAge,omitempty" json:"abatementAge,omitempty"`
	AbatementPeriod    *Period             `bson:"abatementPeriod,omitempty" json:"abatementPeriod,omitempty"`
	AbatementRange     *Range              `bson:"abatementRange,omitempty" json:"abatementRange,omitempty"`
	AbatementString    *string             `bson:"abatementString,omitempty" json:"abatementString,omitempty"`
	RecordedDate       *string             `bson:"recordedDate,omitempty" json:"recordedDate,omitempty"`
	Recorder           *Reference          `bson:"recorder,omitempty" json:"recorder,omitempty"`
	Asserter           *Reference          `bson:"asserter,omitempty" json:"asserter,omitempty"`
	Stage              []ConditionStage    `bson:"stage,omitempty" json:"stage,omitempty"`
	Evidence           []ConditionEvidence `bson:"evidence,omitempty" json:"evidence,omitempty"`
	Note               []Annotation        `bson:"note,omitempty" json:"note,omitempty"`
}

// Clinical stage or grade of a condition. May include formal severity assessments.
type ConditionStage struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Summary           *CodeableConcept `bson:"summary,omitempty" json:"summary,omitempty"`
	Assessment        []Reference      `bson:"assessment,omitempty" json:"assessment,omitempty"`
	Type              *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
}

// Supporting evidence / manifestations that are the basis of the Condition's verification status, such as evidence that confirmed or refuted the condition.
// The evidence may be a simple list of coded symptoms/manifestations, or references to observations or formal assessments, or both.
type ConditionEvidence struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              []CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	Detail            []Reference       `bson:"detail,omitempty" json:"detail,omitempty"`
}

// This function returns resource reference information
func (r Condition) ResourceRef() (string, *string) {
	return "Condition", r.Id
}

type OtherCondition Condition

// MarshalJSON marshals the given Condition as JSON into a byte slice
func (r Condition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCondition
		ResourceType string `json:"resourceType"`
	}{
		OtherCondition: OtherCondition(r),
		ResourceType:   "Condition",
	})
}

// UnmarshalCondition unmarshals a Condition.
func UnmarshalCondition(b []byte) (Condition, error) {
	var condition Condition
	if err := json.Unmarshal(b, &condition); err != nil {
		return condition, err
	}
	return condition, nil
}
