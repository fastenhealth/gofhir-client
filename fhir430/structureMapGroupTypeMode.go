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

import (
	"encoding/json"
	"fmt"
	"strings"
)

// THIS FILE IS GENERATED BY https://github.com/fastenhealth/gofhir-models-gen
// PLEASE DO NOT EDIT BY HAND

// StructureMapGroupTypeMode is documented here http://hl7.org/fhir/ValueSet/map-group-type-mode
type StructureMapGroupTypeMode int

const (
	StructureMapGroupTypeModeNone StructureMapGroupTypeMode = iota
	StructureMapGroupTypeModeTypes
	StructureMapGroupTypeModeTypeAndTypes
)

func (code StructureMapGroupTypeMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *StructureMapGroupTypeMode) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	s = strings.ToLower(s)
	switch s {
	case "none":
		*code = StructureMapGroupTypeModeNone
	case "types":
		*code = StructureMapGroupTypeModeTypes
	case "type-and-types":
		*code = StructureMapGroupTypeModeTypeAndTypes
	default:
		return fmt.Errorf("unknown StructureMapGroupTypeMode code `%s`", s)
	}
	return nil
}
func (code StructureMapGroupTypeMode) String() string {
	return code.Code()
}
func (code StructureMapGroupTypeMode) Code() string {
	switch code {
	case StructureMapGroupTypeModeNone:
		return "none"
	case StructureMapGroupTypeModeTypes:
		return "types"
	case StructureMapGroupTypeModeTypeAndTypes:
		return "type-and-types"
	}
	return "<unknown>"
}
func (code StructureMapGroupTypeMode) Display() string {
	switch code {
	case StructureMapGroupTypeModeNone:
		return "Not a Default"
	case StructureMapGroupTypeModeTypes:
		return "Default for Type Combination"
	case StructureMapGroupTypeModeTypeAndTypes:
		return "Default for type + combination"
	}
	return "<unknown>"
}
func (code StructureMapGroupTypeMode) Definition() string {
	switch code {
	case StructureMapGroupTypeModeNone:
		return "This group is not a default group for the types."
	case StructureMapGroupTypeModeTypes:
		return "This group is a default mapping group for the specified types and for the primary source type."
	case StructureMapGroupTypeModeTypeAndTypes:
		return "This group is a default mapping group for the specified types."
	}
	return "<unknown>"
}
