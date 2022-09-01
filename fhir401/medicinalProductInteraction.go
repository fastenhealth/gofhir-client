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

// MedicinalProductInteraction is documented here http://hl7.org/fhir/StructureDefinition/MedicinalProductInteraction
// The interactions of the medicinal product with other medicinal products, or other forms of interactions.
type MedicinalProductInteraction struct {
	Id                *string                                  `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                                    `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                                  `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                                  `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                               `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension                              `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                              `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Subject           []Reference                              `bson:"subject,omitempty" json:"subject,omitempty"`
	Description       *string                                  `bson:"description,omitempty" json:"description,omitempty"`
	Interactant       []MedicinalProductInteractionInteractant `bson:"interactant,omitempty" json:"interactant,omitempty"`
	Type              *CodeableConcept                         `bson:"type,omitempty" json:"type,omitempty"`
	Effect            *CodeableConcept                         `bson:"effect,omitempty" json:"effect,omitempty"`
	Incidence         *CodeableConcept                         `bson:"incidence,omitempty" json:"incidence,omitempty"`
	Management        *CodeableConcept                         `bson:"management,omitempty" json:"management,omitempty"`
}

// The specific medication, food or laboratory test that interacts.
type MedicinalProductInteractionInteractant struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
}

// This function returns resource reference information
func (r MedicinalProductInteraction) ResourceRef() (string, *string) {
	return "MedicinalProductInteraction", r.Id
}

type OtherMedicinalProductInteraction MedicinalProductInteraction

// MarshalJSON marshals the given MedicinalProductInteraction as JSON into a byte slice
func (r MedicinalProductInteraction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicinalProductInteraction
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicinalProductInteraction: OtherMedicinalProductInteraction(r),
		ResourceType:                     "MedicinalProductInteraction",
	})
}

// UnmarshalMedicinalProductInteraction unmarshals a MedicinalProductInteraction.
func UnmarshalMedicinalProductInteraction(b []byte) (MedicinalProductInteraction, error) {
	var medicinalProductInteraction MedicinalProductInteraction
	if err := json.Unmarshal(b, &medicinalProductInteraction); err != nil {
		return medicinalProductInteraction, err
	}
	return medicinalProductInteraction, nil
}
