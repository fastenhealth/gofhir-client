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

// Binary is documented here http://hl7.org/fhir/StructureDefinition/Binary
// A resource that represents the data of a single raw artifact as digital content accessible in its native format.  A Binary resource can contain any content, whether text, image, pdf, zip archive, etc.
type Binary struct {
	Id              *string    `bson:"id,omitempty" json:"id,omitempty"`
	Meta            *Meta      `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules   *string    `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language        *string    `bson:"language,omitempty" json:"language,omitempty"`
	ContentType     string     `bson:"contentType" json:"contentType"`
	SecurityContext *Reference `bson:"securityContext,omitempty" json:"securityContext,omitempty"`
	Data            *string    `bson:"data,omitempty" json:"data,omitempty"`
}

// This function returns resource reference information
func (r Binary) ResourceRef() (string, *string) {
	return "Binary", r.Id
}

// This function returns resource reference information
func (r Binary) ContainedResources() []json.RawMessage {
	return nil
}

type OtherBinary Binary

// MarshalJSON marshals the given Binary as JSON into a byte slice
func (r Binary) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherBinary
		ResourceType string `json:"resourceType"`
	}{
		OtherBinary:  OtherBinary(r),
		ResourceType: "Binary",
	})
}

// UnmarshalBinary unmarshals a Binary.
func UnmarshalBinary(b []byte) (Binary, error) {
	var binary Binary
	if err := json.Unmarshal(b, &binary); err != nil {
		return binary, err
	}
	return binary, nil
}
