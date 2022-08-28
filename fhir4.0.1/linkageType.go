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

package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// THIS FILE IS GENERATED BY https://github.com/fastenhealth/gofhir-models-gen
// PLEASE DO NOT EDIT BY HAND

// LinkageType is documented here http://hl7.org/fhir/ValueSet/linkage-type
type LinkageType int

const (
	LinkageTypeSource LinkageType = iota
	LinkageTypeAlternate
	LinkageTypeHistorical
)

func (code LinkageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *LinkageType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "source":
		*code = LinkageTypeSource
	case "alternate":
		*code = LinkageTypeAlternate
	case "historical":
		*code = LinkageTypeHistorical
	default:
		return fmt.Errorf("unknown LinkageType code `%s`", s)
	}
	return nil
}
func (code LinkageType) String() string {
	return code.Code()
}
func (code LinkageType) Code() string {
	switch code {
	case LinkageTypeSource:
		return "source"
	case LinkageTypeAlternate:
		return "alternate"
	case LinkageTypeHistorical:
		return "historical"
	}
	return "<unknown>"
}
func (code LinkageType) Display() string {
	switch code {
	case LinkageTypeSource:
		return "Source of Truth"
	case LinkageTypeAlternate:
		return "Alternate Record"
	case LinkageTypeHistorical:
		return "Historical/Obsolete Record"
	}
	return "<unknown>"
}
func (code LinkageType) Definition() string {
	switch code {
	case LinkageTypeSource:
		return "The resource represents the \"source of truth\" (from the perspective of this Linkage resource) for the underlying event/condition/etc."
	case LinkageTypeAlternate:
		return "The resource represents an alternative view of the underlying event/condition/etc.  The resource may still be actively maintained, even though it is not considered to be the source of truth."
	case LinkageTypeHistorical:
		return "The resource represents an obsolete record of the underlying event/condition/etc.  It is not expected to be actively maintained."
	}
	return "<unknown>"
}
