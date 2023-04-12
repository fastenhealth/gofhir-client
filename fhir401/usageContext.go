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

// THIS FILE IS GENERATED BY https://github.com/fastenhealth/gofhir-models-gen
// PLEASE DO NOT EDIT BY HAND

// UsageContext is documented here http://hl7.org/fhir/StructureDefinition/UsageContext
// Base StructureDefinition for UsageContext Type: Specifies clinical/business/etc. metadata that can be used to retrieve, index and/or categorize an artifact. This metadata can either be specific to the applicable population (e.g., age category, DRG) or the specific context of care (e.g., venue, care setting, provider of care).
type UsageContext struct {
	Id                   *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension            []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	Code                 Coding          `bson:"code" json:"code"`
	ValueCodeableConcept CodeableConcept `bson:"valueCodeableConcept" json:"valueCodeableConcept"`
	ValueQuantity        Quantity        `bson:"valueQuantity" json:"valueQuantity"`
	ValueRange           Range           `bson:"valueRange" json:"valueRange"`
	ValueReference       Reference       `bson:"valueReference" json:"valueReference"`
}
