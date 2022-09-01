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

// SearchParameter is documented here http://hl7.org/fhir/StructureDefinition/SearchParameter
// A search parameter that defines a named search item that can be used to search/filter on a resource.
type SearchParameter struct {
	Id                *string                    `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                      `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                    `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                    `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                 `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension                `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url               string                     `bson:"url" json:"url"`
	Version           *string                    `bson:"version,omitempty" json:"version,omitempty"`
	Name              string                     `bson:"name" json:"name"`
	DerivedFrom       *string                    `bson:"derivedFrom,omitempty" json:"derivedFrom,omitempty"`
	Status            PublicationStatus          `bson:"status" json:"status"`
	Experimental      *bool                      `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date              *string                    `bson:"date,omitempty" json:"date,omitempty"`
	Publisher         *string                    `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact           []ContactDetail            `bson:"contact,omitempty" json:"contact,omitempty"`
	Description       string                     `bson:"description" json:"description"`
	UseContext        []UsageContext             `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept          `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Purpose           *string                    `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Code              string                     `bson:"code" json:"code"`
	Base              []ResourceType             `bson:"base" json:"base"`
	Type              SearchParamType            `bson:"type" json:"type"`
	Expression        *string                    `bson:"expression,omitempty" json:"expression,omitempty"`
	Xpath             *string                    `bson:"xpath,omitempty" json:"xpath,omitempty"`
	XpathUsage        *XPathUsageType            `bson:"xpathUsage,omitempty" json:"xpathUsage,omitempty"`
	Target            []ResourceType             `bson:"target,omitempty" json:"target,omitempty"`
	MultipleOr        *bool                      `bson:"multipleOr,omitempty" json:"multipleOr,omitempty"`
	MultipleAnd       *bool                      `bson:"multipleAnd,omitempty" json:"multipleAnd,omitempty"`
	Comparator        []SearchComparator         `bson:"comparator,omitempty" json:"comparator,omitempty"`
	Modifier          []SearchModifierCode       `bson:"modifier,omitempty" json:"modifier,omitempty"`
	Chain             []string                   `bson:"chain,omitempty" json:"chain,omitempty"`
	Component         []SearchParameterComponent `bson:"component,omitempty" json:"component,omitempty"`
}

// Used to define the parts of a composite search parameter.
type SearchParameterComponent struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Definition        string      `bson:"definition" json:"definition"`
	Expression        string      `bson:"expression" json:"expression"`
}

// This function returns resource reference information
func (r SearchParameter) ResourceRef() (string, *string) {
	return "SearchParameter", r.Id
}

type OtherSearchParameter SearchParameter

// MarshalJSON marshals the given SearchParameter as JSON into a byte slice
func (r SearchParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSearchParameter
		ResourceType string `json:"resourceType"`
	}{
		OtherSearchParameter: OtherSearchParameter(r),
		ResourceType:         "SearchParameter",
	})
}

// UnmarshalSearchParameter unmarshals a SearchParameter.
func UnmarshalSearchParameter(b []byte) (SearchParameter, error) {
	var searchParameter SearchParameter
	if err := json.Unmarshal(b, &searchParameter); err != nil {
		return searchParameter, err
	}
	return searchParameter, nil
}
