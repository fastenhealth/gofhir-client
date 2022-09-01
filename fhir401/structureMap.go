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

// StructureMap is documented here http://hl7.org/fhir/StructureDefinition/StructureMap
// A Map of relationships between 2 structures that can be used to transform data.
type StructureMap struct {
	Id                *string                 `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                   `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                 `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                 `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative              `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension             `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension             `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url               string                  `bson:"url" json:"url"`
	Identifier        []Identifier            `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version           *string                 `bson:"version,omitempty" json:"version,omitempty"`
	Name              string                  `bson:"name" json:"name"`
	Title             *string                 `bson:"title,omitempty" json:"title,omitempty"`
	Status            PublicationStatus       `bson:"status" json:"status"`
	Experimental      *bool                   `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date              *string                 `bson:"date,omitempty" json:"date,omitempty"`
	Publisher         *string                 `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact           []ContactDetail         `bson:"contact,omitempty" json:"contact,omitempty"`
	Description       *string                 `bson:"description,omitempty" json:"description,omitempty"`
	UseContext        []UsageContext          `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept       `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Purpose           *string                 `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Copyright         *string                 `bson:"copyright,omitempty" json:"copyright,omitempty"`
	Structure         []StructureMapStructure `bson:"structure,omitempty" json:"structure,omitempty"`
	Import            []string                `bson:"import,omitempty" json:"import,omitempty"`
	Group             []StructureMapGroup     `bson:"group" json:"group"`
}

// A structure definition used by this map. The structure definition may describe instances that are converted, or the instances that are produced.
// It is not necessary for a structure map to identify any dependent structures, though not listing them may restrict its usefulness.
type StructureMapStructure struct {
	Id                *string               `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url               string                `bson:"url" json:"url"`
	Mode              StructureMapModelMode `bson:"mode" json:"mode"`
	Alias             *string               `bson:"alias,omitempty" json:"alias,omitempty"`
	Documentation     *string               `bson:"documentation,omitempty" json:"documentation,omitempty"`
}

// Organizes the mapping into manageable chunks for human review/ease of maintenance.
type StructureMapGroup struct {
	Id                *string                   `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension               `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension               `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              string                    `bson:"name" json:"name"`
	Extends           *string                   `bson:"extends,omitempty" json:"extends,omitempty"`
	TypeMode          StructureMapGroupTypeMode `bson:"typeMode" json:"typeMode"`
	Documentation     *string                   `bson:"documentation,omitempty" json:"documentation,omitempty"`
	Input             []StructureMapGroupInput  `bson:"input" json:"input"`
	Rule              []StructureMapGroupRule   `bson:"rule" json:"rule"`
}

// A name assigned to an instance of data. The instance must be provided when the mapping is invoked.
// If no inputs are named, then the entry mappings are type based.
type StructureMapGroupInput struct {
	Id                *string               `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              string                `bson:"name" json:"name"`
	Type              *string               `bson:"type,omitempty" json:"type,omitempty"`
	Mode              StructureMapInputMode `bson:"mode" json:"mode"`
	Documentation     *string               `bson:"documentation,omitempty" json:"documentation,omitempty"`
}

// Transform Rule from source to target.
type StructureMapGroupRule struct {
	Id                *string                          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              string                           `bson:"name" json:"name"`
	Source            []StructureMapGroupRuleSource    `bson:"source" json:"source"`
	Target            []StructureMapGroupRuleTarget    `bson:"target,omitempty" json:"target,omitempty"`
	Rule              []StructureMapGroupRule          `bson:"rule,omitempty" json:"rule,omitempty"`
	Dependent         []StructureMapGroupRuleDependent `bson:"dependent,omitempty" json:"dependent,omitempty"`
	Documentation     *string                          `bson:"documentation,omitempty" json:"documentation,omitempty"`
}

// Source inputs to the mapping.
type StructureMapGroupRuleSource struct {
	Id                *string                     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                 `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                 `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Context           string                      `bson:"context" json:"context"`
	Min               *int                        `bson:"min,omitempty" json:"min,omitempty"`
	Max               *string                     `bson:"max,omitempty" json:"max,omitempty"`
	Type              *string                     `bson:"type,omitempty" json:"type,omitempty"`
	Element           *string                     `bson:"element,omitempty" json:"element,omitempty"`
	ListMode          *StructureMapSourceListMode `bson:"listMode,omitempty" json:"listMode,omitempty"`
	Variable          *string                     `bson:"variable,omitempty" json:"variable,omitempty"`
	Condition         *string                     `bson:"condition,omitempty" json:"condition,omitempty"`
	Check             *string                     `bson:"check,omitempty" json:"check,omitempty"`
	LogMessage        *string                     `bson:"logMessage,omitempty" json:"logMessage,omitempty"`
}

// Content to create because of this mapping rule.
type StructureMapGroupRuleTarget struct {
	Id                *string                                `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Context           *string                                `bson:"context,omitempty" json:"context,omitempty"`
	ContextType       *StructureMapContextType               `bson:"contextType,omitempty" json:"contextType,omitempty"`
	Element           *string                                `bson:"element,omitempty" json:"element,omitempty"`
	Variable          *string                                `bson:"variable,omitempty" json:"variable,omitempty"`
	ListMode          []StructureMapTargetListMode           `bson:"listMode,omitempty" json:"listMode,omitempty"`
	ListRuleId        *string                                `bson:"listRuleId,omitempty" json:"listRuleId,omitempty"`
	Transform         *StructureMapTransform                 `bson:"transform,omitempty" json:"transform,omitempty"`
	Parameter         []StructureMapGroupRuleTargetParameter `bson:"parameter,omitempty" json:"parameter,omitempty"`
}

// Parameters to the transform.
type StructureMapGroupRuleTargetParameter struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
}

// Which other rules to apply in the context of this rule.
type StructureMapGroupRuleDependent struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              string      `bson:"name" json:"name"`
	Variable          []string    `bson:"variable" json:"variable"`
}

// This function returns resource reference information
func (r StructureMap) ResourceRef() (string, *string) {
	return "StructureMap", r.Id
}

type OtherStructureMap StructureMap

// MarshalJSON marshals the given StructureMap as JSON into a byte slice
func (r StructureMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherStructureMap
		ResourceType string `json:"resourceType"`
	}{
		OtherStructureMap: OtherStructureMap(r),
		ResourceType:      "StructureMap",
	})
}

// UnmarshalStructureMap unmarshals a StructureMap.
func UnmarshalStructureMap(b []byte) (StructureMap, error) {
	var structureMap StructureMap
	if err := json.Unmarshal(b, &structureMap); err != nil {
		return structureMap, err
	}
	return structureMap, nil
}
