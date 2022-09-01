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

// Resource is documented here http://hl7.org/fhir/StructureDefinition/Resource
// This is the base resource type for everything.
type Resource struct {
	Id            *string `bson:"id,omitempty" json:"id,omitempty"`
	Meta          *Meta   `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules *string `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language      *string `bson:"language,omitempty" json:"language,omitempty"`
}

// This function returns resource reference information
func (r Resource) ResourceRef() (string, *string) {
	return "Resource", r.Id
}

type OtherResource Resource

// MarshalJSON marshals the given Resource as JSON into a byte slice
func (r Resource) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherResource
		ResourceType string `json:"resourceType"`
	}{
		OtherResource: OtherResource(r),
		ResourceType:  "Resource",
	})
}

// UnmarshalResource unmarshals a Resource.
func UnmarshalResource(b []byte) (Resource, error) {
	var resource Resource
	if err := json.Unmarshal(b, &resource); err != nil {
		return resource, err
	}
	return resource, nil
}
