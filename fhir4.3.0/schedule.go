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

package fhir

import "encoding/json"

// THIS FILE IS GENERATED BY https://github.com/fastenhealth/gofhir-models-gen
// PLEASE DO NOT EDIT BY HAND

// Schedule is documented here http://hl7.org/fhir/StructureDefinition/Schedule
type Schedule struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta             `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string           `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string           `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative        `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Active            *bool             `bson:"active,omitempty" json:"active,omitempty"`
	ServiceCategory   []CodeableConcept `bson:"serviceCategory,omitempty" json:"serviceCategory,omitempty"`
	ServiceType       []CodeableConcept `bson:"serviceType,omitempty" json:"serviceType,omitempty"`
	Specialty         []CodeableConcept `bson:"specialty,omitempty" json:"specialty,omitempty"`
	Actor             []Reference       `bson:"actor" json:"actor"`
	PlanningHorizon   *Period           `bson:"planningHorizon,omitempty" json:"planningHorizon,omitempty"`
	Comment           *string           `bson:"comment,omitempty" json:"comment,omitempty"`
}
type OtherSchedule Schedule

// MarshalJSON marshals the given Schedule as JSON into a byte slice
func (r Schedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSchedule
		ResourceType string `json:"resourceType"`
	}{
		OtherSchedule: OtherSchedule(r),
		ResourceType:  "Schedule",
	})
}

// UnmarshalSchedule unmarshals a Schedule.
func UnmarshalSchedule(b []byte) (Schedule, error) {
	var schedule Schedule
	if err := json.Unmarshal(b, &schedule); err != nil {
		return schedule, err
	}
	return schedule, nil
}
