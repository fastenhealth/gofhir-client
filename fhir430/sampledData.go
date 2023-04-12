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

// SampledData is documented here http://hl7.org/fhir/StructureDefinition/SampledData
// Base StructureDefinition for SampledData Type: A series of measurements taken by a device, with upper and lower limits. There may be more than one dimension in the data.
type SampledData struct {
	Id         *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension  []Extension  `bson:"extension,omitempty" json:"extension,omitempty"`
	Origin     Quantity     `bson:"origin" json:"origin"`
	Period     json.Number  `bson:"period" json:"period"`
	Factor     *json.Number `bson:"factor,omitempty" json:"factor,omitempty"`
	LowerLimit *json.Number `bson:"lowerLimit,omitempty" json:"lowerLimit,omitempty"`
	UpperLimit *json.Number `bson:"upperLimit,omitempty" json:"upperLimit,omitempty"`
	Dimensions int          `bson:"dimensions" json:"dimensions"`
	Data       *string      `bson:"data,omitempty" json:"data,omitempty"`
}
