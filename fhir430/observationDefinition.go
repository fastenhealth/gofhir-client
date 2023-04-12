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

// ObservationDefinition is documented here http://hl7.org/fhir/StructureDefinition/ObservationDefinition
// Set of definitional characteristics for a kind of observation or measurement produced or consumed by an orderable health care service.
type ObservationDefinition struct {
	Id                     *string                                   `bson:"id,omitempty" json:"id,omitempty"`
	Meta                   *Meta                                     `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules          *string                                   `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language               *string                                   `bson:"language,omitempty" json:"language,omitempty"`
	Text                   *Narrative                                `bson:"text,omitempty" json:"text,omitempty"`
	Extension              []Extension                               `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension      []Extension                               `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Category               []CodeableConcept                         `bson:"category,omitempty" json:"category,omitempty"`
	Code                   CodeableConcept                           `bson:"code" json:"code"`
	Identifier             []Identifier                              `bson:"identifier,omitempty" json:"identifier,omitempty"`
	PermittedDataType      []ObservationDataType                     `bson:"permittedDataType,omitempty" json:"permittedDataType,omitempty"`
	MultipleResultsAllowed *bool                                     `bson:"multipleResultsAllowed,omitempty" json:"multipleResultsAllowed,omitempty"`
	Method                 *CodeableConcept                          `bson:"method,omitempty" json:"method,omitempty"`
	PreferredReportName    *string                                   `bson:"preferredReportName,omitempty" json:"preferredReportName,omitempty"`
	QuantitativeDetails    *ObservationDefinitionQuantitativeDetails `bson:"quantitativeDetails,omitempty" json:"quantitativeDetails,omitempty"`
	QualifiedInterval      []ObservationDefinitionQualifiedInterval  `bson:"qualifiedInterval,omitempty" json:"qualifiedInterval,omitempty"`
	ValidCodedValueSet     *Reference                                `bson:"validCodedValueSet,omitempty" json:"validCodedValueSet,omitempty"`
	NormalCodedValueSet    *Reference                                `bson:"normalCodedValueSet,omitempty" json:"normalCodedValueSet,omitempty"`
	AbnormalCodedValueSet  *Reference                                `bson:"abnormalCodedValueSet,omitempty" json:"abnormalCodedValueSet,omitempty"`
	CriticalCodedValueSet  *Reference                                `bson:"criticalCodedValueSet,omitempty" json:"criticalCodedValueSet,omitempty"`
}

// Characteristics for quantitative results of this observation.
type ObservationDefinitionQuantitativeDetails struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	CustomaryUnit     *CodeableConcept `bson:"customaryUnit,omitempty" json:"customaryUnit,omitempty"`
	Unit              *CodeableConcept `bson:"unit,omitempty" json:"unit,omitempty"`
	ConversionFactor  *json.Number     `bson:"conversionFactor,omitempty" json:"conversionFactor,omitempty"`
	DecimalPrecision  *int             `bson:"decimalPrecision,omitempty" json:"decimalPrecision,omitempty"`
}

// Multiple  ranges of results qualified by different contexts for ordinal or continuous observations conforming to this ObservationDefinition.
type ObservationDefinitionQualifiedInterval struct {
	Id                *string                   `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension               `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension               `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Category          *ObservationRangeCategory `bson:"category,omitempty" json:"category,omitempty"`
	Range             *Range                    `bson:"range,omitempty" json:"range,omitempty"`
	Context           *CodeableConcept          `bson:"context,omitempty" json:"context,omitempty"`
	AppliesTo         []CodeableConcept         `bson:"appliesTo,omitempty" json:"appliesTo,omitempty"`
	Gender            *AdministrativeGender     `bson:"gender,omitempty" json:"gender,omitempty"`
	Age               *Range                    `bson:"age,omitempty" json:"age,omitempty"`
	GestationalAge    *Range                    `bson:"gestationalAge,omitempty" json:"gestationalAge,omitempty"`
	Condition         *string                   `bson:"condition,omitempty" json:"condition,omitempty"`
}

// This function returns resource reference information
func (r ObservationDefinition) ResourceRef() (string, *string) {
	return "ObservationDefinition", r.Id
}

type OtherObservationDefinition ObservationDefinition

// MarshalJSON marshals the given ObservationDefinition as JSON into a byte slice
func (r ObservationDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherObservationDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherObservationDefinition: OtherObservationDefinition(r),
		ResourceType:               "ObservationDefinition",
	})
}

// UnmarshalObservationDefinition unmarshals a ObservationDefinition.
func UnmarshalObservationDefinition(b []byte) (ObservationDefinition, error) {
	var observationDefinition ObservationDefinition
	if err := json.Unmarshal(b, &observationDefinition); err != nil {
		return observationDefinition, err
	}
	return observationDefinition, nil
}
