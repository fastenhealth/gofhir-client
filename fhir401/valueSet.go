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

// ValueSet is documented here http://hl7.org/fhir/StructureDefinition/ValueSet
// A ValueSet resource instance specifies a set of codes drawn from one or more code systems, intended for use in a particular context. Value sets link between [[[CodeSystem]]] definitions and their use in [coded elements](terminologies.html).
type ValueSet struct {
	Id                *string            `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta              `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string            `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string            `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative         `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension        `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension        `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url               *string            `bson:"url,omitempty" json:"url,omitempty"`
	Identifier        []Identifier       `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version           *string            `bson:"version,omitempty" json:"version,omitempty"`
	Name              *string            `bson:"name,omitempty" json:"name,omitempty"`
	Title             *string            `bson:"title,omitempty" json:"title,omitempty"`
	Status            PublicationStatus  `bson:"status" json:"status"`
	Experimental      *bool              `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date              *string            `bson:"date,omitempty" json:"date,omitempty"`
	Publisher         *string            `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact           []ContactDetail    `bson:"contact,omitempty" json:"contact,omitempty"`
	Description       *string            `bson:"description,omitempty" json:"description,omitempty"`
	UseContext        []UsageContext     `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept  `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Immutable         *bool              `bson:"immutable,omitempty" json:"immutable,omitempty"`
	Purpose           *string            `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Copyright         *string            `bson:"copyright,omitempty" json:"copyright,omitempty"`
	Compose           *ValueSetCompose   `bson:"compose,omitempty" json:"compose,omitempty"`
	Expansion         *ValueSetExpansion `bson:"expansion,omitempty" json:"expansion,omitempty"`
}

// A set of criteria that define the contents of the value set by including or excluding codes selected from the specified code system(s) that the value set draws from. This is also known as the Content Logical Definition (CLD).
type ValueSetCompose struct {
	Id                *string                  `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension              `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension              `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	LockedDate        *string                  `bson:"lockedDate,omitempty" json:"lockedDate,omitempty"`
	Inactive          *bool                    `bson:"inactive,omitempty" json:"inactive,omitempty"`
	Include           []ValueSetComposeInclude `bson:"include" json:"include"`
	Exclude           []ValueSetComposeInclude `bson:"exclude,omitempty" json:"exclude,omitempty"`
}

// Include one or more codes from a code system or other value set(s).
// All the conditions in an include must be true. If a system is listed, all the codes from the system are listed. If one or more filters are listed, all of the filters must apply. If one or more value sets are listed, the codes must be in all the value sets. E.g. each include is 'include all the codes that meet all these conditions'.
type ValueSetComposeInclude struct {
	Id                *string                         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	System            *string                         `bson:"system,omitempty" json:"system,omitempty"`
	Version           *string                         `bson:"version,omitempty" json:"version,omitempty"`
	Concept           []ValueSetComposeIncludeConcept `bson:"concept,omitempty" json:"concept,omitempty"`
	Filter            []ValueSetComposeIncludeFilter  `bson:"filter,omitempty" json:"filter,omitempty"`
	ValueSet          []string                        `bson:"valueSet,omitempty" json:"valueSet,omitempty"`
}

// Specifies a concept to be included or excluded.
// The list of concepts is considered ordered, though the order might not have any particular significance. Typically, the order of an expansion follows that defined in the compose element.
type ValueSetComposeIncludeConcept struct {
	Id                *string                                    `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                                `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                                `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              string                                     `bson:"code" json:"code"`
	Display           *string                                    `bson:"display,omitempty" json:"display,omitempty"`
	Designation       []ValueSetComposeIncludeConceptDesignation `bson:"designation,omitempty" json:"designation,omitempty"`
}

// Additional representations for this concept when used in this value set - other languages, aliases, specialized purposes, used for particular purposes, etc.
// Concepts have both a ```display``` and an array of ```designation```. The display is equivalent to a special designation with an implied ```designation.use``` of "primary code" and a language equal to the [Resource Language](resource.html#language).
type ValueSetComposeIncludeConceptDesignation struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Language          *string     `bson:"language,omitempty" json:"language,omitempty"`
	Use               *Coding     `bson:"use,omitempty" json:"use,omitempty"`
	Value             string      `bson:"value" json:"value"`
}

// Select concepts by specify a matching criterion based on the properties (including relationships) defined by the system, or on filters defined by the system. If multiple filters are specified, they SHALL all be true.
// Selecting codes by specifying filters based on properties is only possible where the underlying code system defines appropriate properties. Note that in some cases, the underlying code system defines the logical concepts but not the literal codes for the concepts. In such cases, the literal definitions may be provided by a third party.
type ValueSetComposeIncludeFilter struct {
	Id                *string        `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Property          string         `bson:"property" json:"property"`
	Op                FilterOperator `bson:"op" json:"op"`
	Value             string         `bson:"value" json:"value"`
}

// A value set can also be "expanded", where the value set is turned into a simple collection of enumerated codes. This element holds the expansion, if it has been performed.
/*
Expansion is performed to produce a collection of codes that are ready to use for data entry or validation. Value set expansions are always considered to be stateless - they are a record of the set of codes in the value set at a point in time under a given set of conditions, and are not subject to ongoing maintenance.

Expansion.parameter is  a simplified list of parameters - a subset of the features of the [Parameters](parameters.html) resource.
*/
type ValueSetExpansion struct {
	Id                *string                      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        *string                      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Timestamp         string                       `bson:"timestamp" json:"timestamp"`
	Total             *int                         `bson:"total,omitempty" json:"total,omitempty"`
	Offset            *int                         `bson:"offset,omitempty" json:"offset,omitempty"`
	Parameter         []ValueSetExpansionParameter `bson:"parameter,omitempty" json:"parameter,omitempty"`
	Contains          []ValueSetExpansionContains  `bson:"contains,omitempty" json:"contains,omitempty"`
}

// A parameter that controlled the expansion process. These parameters may be used by users of expanded value sets to check whether the expansion is suitable for a particular purpose, or to pick the correct expansion.
// The server decides which parameters to include here, but at a minimum, the list SHOULD include all of the parameters that affect the $expand operation. If the expansion will be persisted all of these parameters SHALL be included. If the codeSystem on the server has a specified version then this version SHALL be provided as a parameter in the expansion (note that not all code systems have a version).
type ValueSetExpansionParameter struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              string       `bson:"name" json:"name"`
	ValueString       *string      `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueBoolean      *bool        `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty"`
	ValueInteger      *int         `bson:"valueInteger,omitempty" json:"valueInteger,omitempty"`
	ValueDecimal      *json.Number `bson:"valueDecimal,omitempty" json:"valueDecimal,omitempty"`
	ValueUri          *string      `bson:"valueUri,omitempty" json:"valueUri,omitempty"`
	ValueCode         *string      `bson:"valueCode,omitempty" json:"valueCode,omitempty"`
	ValueDateTime     *string      `bson:"valueDateTime,omitempty" json:"valueDateTime,omitempty"`
}

// The codes that are contained in the value set expansion.
type ValueSetExpansionContains struct {
	Id                *string                                    `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                                `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                                `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	System            *string                                    `bson:"system,omitempty" json:"system,omitempty"`
	Abstract          *bool                                      `bson:"abstract,omitempty" json:"abstract,omitempty"`
	Inactive          *bool                                      `bson:"inactive,omitempty" json:"inactive,omitempty"`
	Version           *string                                    `bson:"version,omitempty" json:"version,omitempty"`
	Code              *string                                    `bson:"code,omitempty" json:"code,omitempty"`
	Display           *string                                    `bson:"display,omitempty" json:"display,omitempty"`
	Designation       []ValueSetComposeIncludeConceptDesignation `bson:"designation,omitempty" json:"designation,omitempty"`
	Contains          []ValueSetExpansionContains                `bson:"contains,omitempty" json:"contains,omitempty"`
}

// This function returns resource reference information
func (r ValueSet) ResourceRef() (string, *string) {
	return "ValueSet", r.Id
}

type OtherValueSet ValueSet

// MarshalJSON marshals the given ValueSet as JSON into a byte slice
func (r ValueSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherValueSet
		ResourceType string `json:"resourceType"`
	}{
		OtherValueSet: OtherValueSet(r),
		ResourceType:  "ValueSet",
	})
}

// UnmarshalValueSet unmarshals a ValueSet.
func UnmarshalValueSet(b []byte) (ValueSet, error) {
	var valueSet ValueSet
	if err := json.Unmarshal(b, &valueSet); err != nil {
		return valueSet, err
	}
	return valueSet, nil
}
