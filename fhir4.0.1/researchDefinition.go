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

// ResearchDefinition is documented here http://hl7.org/fhir/StructureDefinition/ResearchDefinition
type ResearchDefinition struct {
	Id                  *string           `bson:"id,omitempty" json:"id,omitempty"`
	Meta                *Meta             `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules       *string           `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language            *string           `bson:"language,omitempty" json:"language,omitempty"`
	Text                *Narrative        `bson:"text,omitempty" json:"text,omitempty"`
	Extension           []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url                 *string           `bson:"url,omitempty" json:"url,omitempty"`
	Identifier          []Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version             *string           `bson:"version,omitempty" json:"version,omitempty"`
	Name                *string           `bson:"name,omitempty" json:"name,omitempty"`
	Title               *string           `bson:"title,omitempty" json:"title,omitempty"`
	ShortTitle          *string           `bson:"shortTitle,omitempty" json:"shortTitle,omitempty"`
	Subtitle            *string           `bson:"subtitle,omitempty" json:"subtitle,omitempty"`
	Status              PublicationStatus `bson:"status" json:"status"`
	Experimental        *bool             `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date                *string           `bson:"date,omitempty" json:"date,omitempty"`
	Publisher           *string           `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact             []ContactDetail   `bson:"contact,omitempty" json:"contact,omitempty"`
	Description         *string           `bson:"description,omitempty" json:"description,omitempty"`
	Comment             []string          `bson:"comment,omitempty" json:"comment,omitempty"`
	UseContext          []UsageContext    `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction        []CodeableConcept `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Purpose             *string           `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Usage               *string           `bson:"usage,omitempty" json:"usage,omitempty"`
	Copyright           *string           `bson:"copyright,omitempty" json:"copyright,omitempty"`
	ApprovalDate        *string           `bson:"approvalDate,omitempty" json:"approvalDate,omitempty"`
	LastReviewDate      *string           `bson:"lastReviewDate,omitempty" json:"lastReviewDate,omitempty"`
	EffectivePeriod     *Period           `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	Topic               []CodeableConcept `bson:"topic,omitempty" json:"topic,omitempty"`
	Author              []ContactDetail   `bson:"author,omitempty" json:"author,omitempty"`
	Editor              []ContactDetail   `bson:"editor,omitempty" json:"editor,omitempty"`
	Reviewer            []ContactDetail   `bson:"reviewer,omitempty" json:"reviewer,omitempty"`
	Endorser            []ContactDetail   `bson:"endorser,omitempty" json:"endorser,omitempty"`
	RelatedArtifact     []RelatedArtifact `bson:"relatedArtifact,omitempty" json:"relatedArtifact,omitempty"`
	Library             []string          `bson:"library,omitempty" json:"library,omitempty"`
	Population          Reference         `bson:"population" json:"population"`
	Exposure            *Reference        `bson:"exposure,omitempty" json:"exposure,omitempty"`
	ExposureAlternative *Reference        `bson:"exposureAlternative,omitempty" json:"exposureAlternative,omitempty"`
	Outcome             *Reference        `bson:"outcome,omitempty" json:"outcome,omitempty"`
}
type OtherResearchDefinition ResearchDefinition

// MarshalJSON marshals the given ResearchDefinition as JSON into a byte slice
func (r ResearchDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherResearchDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherResearchDefinition: OtherResearchDefinition(r),
		ResourceType:            "ResearchDefinition",
	})
}

// UnmarshalResearchDefinition unmarshals a ResearchDefinition.
func UnmarshalResearchDefinition(b []byte) (ResearchDefinition, error) {
	var researchDefinition ResearchDefinition
	if err := json.Unmarshal(b, &researchDefinition); err != nil {
		return researchDefinition, err
	}
	return researchDefinition, nil
}
