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

// VisionPrescription is documented here http://hl7.org/fhir/StructureDefinition/VisionPrescription
// An authorization for the provision of glasses and/or contact lenses to a patient.
type VisionPrescription struct {
	Id                *string                               `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                                 `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                               `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                               `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                            `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension                           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier                          `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            FinancialResourceStatusCodes          `bson:"status" json:"status"`
	Created           string                                `bson:"created" json:"created"`
	Patient           Reference                             `bson:"patient" json:"patient"`
	Encounter         *Reference                            `bson:"encounter,omitempty" json:"encounter,omitempty"`
	DateWritten       string                                `bson:"dateWritten" json:"dateWritten"`
	Prescriber        Reference                             `bson:"prescriber" json:"prescriber"`
	LensSpecification []VisionPrescriptionLensSpecification `bson:"lensSpecification" json:"lensSpecification"`
}

// Contain the details of  the individual lens specifications and serves as the authorization for the fullfillment by certified professionals.
type VisionPrescriptionLensSpecification struct {
	Id                *string                                    `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                                `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                                `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Product           CodeableConcept                            `bson:"product" json:"product"`
	Eye               VisionEyes                                 `bson:"eye" json:"eye"`
	Sphere            *json.Number                               `bson:"sphere,omitempty" json:"sphere,omitempty"`
	Cylinder          *json.Number                               `bson:"cylinder,omitempty" json:"cylinder,omitempty"`
	Axis              *int                                       `bson:"axis,omitempty" json:"axis,omitempty"`
	Prism             []VisionPrescriptionLensSpecificationPrism `bson:"prism,omitempty" json:"prism,omitempty"`
	Add               *json.Number                               `bson:"add,omitempty" json:"add,omitempty"`
	Power             *json.Number                               `bson:"power,omitempty" json:"power,omitempty"`
	BackCurve         *json.Number                               `bson:"backCurve,omitempty" json:"backCurve,omitempty"`
	Diameter          *json.Number                               `bson:"diameter,omitempty" json:"diameter,omitempty"`
	Duration          *Quantity                                  `bson:"duration,omitempty" json:"duration,omitempty"`
	Color             *string                                    `bson:"color,omitempty" json:"color,omitempty"`
	Brand             *string                                    `bson:"brand,omitempty" json:"brand,omitempty"`
	Note              []Annotation                               `bson:"note,omitempty" json:"note,omitempty"`
}

// Allows for adjustment on two axis.
type VisionPrescriptionLensSpecificationPrism struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Amount            json.Number `bson:"amount" json:"amount"`
	Base              VisionBase  `bson:"base" json:"base"`
}

// This function returns resource reference information
func (r VisionPrescription) ResourceRef() (string, *string) {
	return "VisionPrescription", r.Id
}

type OtherVisionPrescription VisionPrescription

// MarshalJSON marshals the given VisionPrescription as JSON into a byte slice
func (r VisionPrescription) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherVisionPrescription
		ResourceType string `json:"resourceType"`
	}{
		OtherVisionPrescription: OtherVisionPrescription(r),
		ResourceType:            "VisionPrescription",
	})
}

// UnmarshalVisionPrescription unmarshals a VisionPrescription.
func UnmarshalVisionPrescription(b []byte) (VisionPrescription, error) {
	var visionPrescription VisionPrescription
	if err := json.Unmarshal(b, &visionPrescription); err != nil {
		return visionPrescription, err
	}
	return visionPrescription, nil
}
